package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"gitlab.com/zxtang/at"
	atjsonapi "gitlab.com/zxtang/at/jsonapi"
)

var (
	ErrUserID = at.NewError("at:gateway:application/ErrUserID", "user id is required")
)

//包装其他服返回的错误信息，实现Error接口
type ServerErrors struct {
	Errors []atjsonapi.JsonapiError `json:"errors"`
}

func (e *ServerErrors) Error() string {
	str := ""
	for _, v := range e.Errors {
		str += fmt.Sprintf("code:%d detail:%s meta:%v", v.Code, v.Detail, v.Meta)
	}
	return str
}

type authVrifyReq struct {
	Action   string
	Resource string
	Token    string
}

type UserInfo struct {
	ID         string
	Role       string
	OperatorID string
}

func authVerify(authURL string, req authVrifyReq) (bool, *UserInfo, error) {
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return false, nil, fmt.Errorf("authVrifyReq json.Marshal error")
	}

	resp, err := http.Post(authURL, "", bytes.NewReader(reqBytes))
	if err != nil {
		return false, nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, nil, err
	}

	jsonErrs := &atjsonapi.JsonapiDocument{}
	if err := json.Unmarshal(respBytes, jsonErrs); err != nil {
		return false, nil, fmt.Errorf("authVerify respStruct json.Unmarshal error")
	}

	if len(jsonErrs.Errors) > 0 {
		//@todo 此处应该解析出iam返回的错误进行分类，而不是直接返回，然后用同一个错误
		errs := &ServerErrors{
			Errors: jsonErrs.Errors,
		}

		return false, nil, errs
	}

	respStruct := struct {
		IsValid bool
		User    *UserInfo
	}{}
	err = json.Unmarshal(respBytes, &respStruct)
	if err != nil {
		return false, nil, fmt.Errorf("authVerify respStruct json.Unmarshal error")
	}

	return respStruct.IsValid, respStruct.User, nil
}

func authMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		action := r.Method
		resource := r.URL.Path

		token := r.Header.Get("token")
		reference := r.Referer()
		domain, url, err := getURLInfo(reference)
		if err != nil {
			handleError(w, fmt.Errorf("gateway get url info failed:%v", err))
		}
		tmp := strings.Split(resource, "/")
		if len(tmp) < 3 {
			handleError(w, fmt.Errorf("invalid url:%s, url should be /v{number}/{apiservename}", r.URL.Path))
			return
		}
		api := tmp[2]

		// 文章获取不需要认证
		// file获取不要认证
		r.Header.Set("reference", reference)
		r.Header.Set("domain", domain)
		r.Header.Set("url", url)

		if ApplicationContext.NoVerify(api, resource, strings.ToLower(action)) == false {

			if token == "" {
				handleError(w, fmt.Errorf("NOSESSIONTOKEN"))
				return
			}

			turl, err := ApplicationContext.APIProxyURL(APIUsers)
			if err != nil {
				handleError(w, err)
				return
			}
			ok, user, err := authVerify(fmt.Sprintf("%s://%s", turl.Scheme, turl.Host+"/v1/users/verify"), authVrifyReq{
				Action:   action,
				Resource: resource,
				Token:    token,
			})
			if err != nil {
				log.Printf("call authVerify error, %v", err)
				handleError(w, err)
			} else if !ok {
				handleError(w, fmt.Errorf("access was denied by auth, resource:%v", resource))
			} else {

				if user.ID == "" {
					handleError(w, ErrUserID)
					return
				}

				r.Header.Set("userID", user.ID)
				r.Header.Set("operatorID", user.OperatorID)

				next.ServeHTTP(w, r)
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func getURLInfo(refer string) (string, string, error) {
	u, err := url.Parse(refer)
	if err != nil {
		return "", "", nil
	}
	domain := u.Host
	url := domain + u.RequestURI()
	return domain, url, nil
}
