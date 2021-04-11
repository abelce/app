package middleware

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	normal_context "abelce/app/gateway/application/context"
	"abelce/at"
	atjsonapi "abelce/at/jsonapi"
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

type AuthMiddleware struct {
}

func NewAuthMiddleware() AuthMiddleware {
	return AuthMiddleware{}
}

// 处理请求认证的中间件
func (m AuthMiddleware) Register(ctx *normal_context.Context) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
}

// func authMW(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		action := r.Method
// 		resource := r.URL.Path

// 		token := r.Header.Get("token")
// 		reference := r.Referer()
// 		domain, url, err := getURLInfo(reference)
// 		if err != nil {
// 			handleError(w, fmt.Errorf("gateway get url info failed:%v", err))
// 		}
// 		tmp := strings.Split(resource, "/")
// 		if len(tmp) < 3 {
// 			handleError(w, fmt.Errorf("invalid url:%s, url should be /v{number}/{apiservename}", r.URL.Path))
// 			return
// 		}
// 		api := tmp[2]

// 		// 文章获取不需要认证
// 		// file获取不要认证
// 		r.Header.Set("reference", reference)
// 		r.Header.Set("domain", domain)
// 		r.Header.Set("url", url)

// 		if NoVerify(api, resource, strings.ToLower(action)) == false {

// 			if token == "" {
// 				handleError(w, fmt.Errorf("NOSESSIONTOKEN"))
// 				return
// 			}

// 			turl, err := APIProxyURL(APIUsers)
// 			if err != nil {
// 				handleError(w, err)
// 				return
// 			}
// 			ok, user, err := authVerify(fmt.Sprintf("%s://%s", turl.Scheme, turl.Host+"/v1/users/verify"), authVrifyReq{
// 				Action:   action,
// 				Resource: resource,
// 				Token:    token,
// 			})
// 			if err != nil {
// 				log.Printf("call authVerify error, %v", err)
// 				handleError(w, err)
// 			} else if !ok {
// 				handleError(w, fmt.Errorf("access was denied by auth, resource:%v", resource))
// 			} else {

// 				if user.ID == "" {
// 					handleError(w, ErrUserID)
// 					return
// 				}

// 				r.Header.Set("userID", user.ID)
// 				r.Header.Set("operatorID", user.OperatorID)

// 				next.ServeHTTP(w, r)
// 			}
// 		} else {
// 			next.ServeHTTP(w, r)
// 		}
// 	})
// }

// func getURLInfo(refer string) (string, string, error) {
// 	u, err := url.Parse(refer)
// 	if err != nil {
// 		return "", "", nil
// 	}
// 	domain := u.Host
// 	url := domain + u.RequestURI()
// 	return domain, url, nil
// }

// func authVerify(authURL string, req authVrifyReq) (bool, *UserInfo, error) {
// 	reqBytes, err := json.Marshal(req)
// 	if err != nil {
// 		return false, nil, fmt.Errorf("authVrifyReq json.Marshal error")
// 	}

// 	resp, err := http.Post(authURL, "", bytes.NewReader(reqBytes))
// 	if err != nil {
// 		return false, nil, err
// 	}
// 	defer resp.Body.Close()

// 	respBytes, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return false, nil, err
// 	}

// 	jsonErrs := &atjsonapi.JsonapiDocument{}
// 	if err := json.Unmarshal(respBytes, jsonErrs); err != nil {
// 		return false, nil, fmt.Errorf("authVerify respStruct json.Unmarshal error")
// 	}

// 	if len(jsonErrs.Errors) > 0 {
// 		//@todo 此处应该解析出iam返回的错误进行分类，而不是直接返回，然后用同一个错误
// 		errs := &ServerErrors{
// 			Errors: jsonErrs.Errors,
// 		}

// 		return false, nil, errs
// 	}

// 	respStruct := struct {
// 		IsValid bool
// 		User    *UserInfo
// 	}{}
// 	err = json.Unmarshal(respBytes, &respStruct)
// 	if err != nil {
// 		return false, nil, fmt.Errorf("authVerify respStruct json.Unmarshal error")
// 	}

// 	return respStruct.IsValid, respStruct.User, nil
// }

// func handleError(w http.ResponseWriter, e error) {

// 	if e != nil {
// 		switch tErr := e.(type) {
// 		case *ServerErrors:
// 			w.Header().Set("Content-Type", "application/json")
// 			if tErr.Errors[0].Detail == "UNAUTHORIZED" {
// 				w.WriteHeader(401)
// 			} else {
// 				w.WriteHeader(400)
// 			}

// 			json, e := errs2doc(tErr.Errors)

// 			if e != nil {
// 				w.WriteHeader(500)
// 				fmt.Fprintln(w, e.Error())
// 			}
// 			fmt.Fprintln(w, json)
// 			return

// 		default:
// 			errs := []atjsonapi.JsonapiError{}
// 			at.Ensure(&tErr)
// 			_, f, l, _ := runtime.Caller(1)
// 			errs = append(errs, atjsonapi.JsonapiError{
// 				Code:   400,
// 				Detail: tErr.Error(),
// 				Meta: map[string]interface{}{
// 					"file": f,
// 					"line": l,
// 				},
// 			})
// 			w.Header().Set("Content-Type", "application/json")
// 			if strings.Contains(errs[0].Detail, "NOSESSIONTOKEN") {
// 				w.WriteHeader(401)
// 			} else {
// 				w.WriteHeader(400)
// 			}
// 			json, e := errs2doc(errs)
// 			if e != nil {
// 				w.WriteHeader(500)
// 				fmt.Fprintln(w, e.Error())
// 			}
// 			fmt.Fprintln(w, json)
// 			return
// 		}
// 	}
// } /*
//  */

// //key: 服务, url: url， action: 方法(get, post, delete, put)
// func (c *Context) NoVerify(key, url, action string) bool {
// 	fmt.Println("#v", c.Config().NoVerify)
// 	verifyMethod := c.Config().NoVerify[key]
// 	if verifyMethod == nil {
// 		return false
// 	}

// 	// all表示所以接口都不验证
// 	if verifyMethod.All == true {
// 		return true
// 	}

// 	method := verifyMethod.Methods[action]
// 	if method == nil {
// 		return false
// 	}

// 	if method.All == true {
// 		return true
// 	}

// 	if len(method.Urls) > 0 {
// 		for _, v := range method.Urls {
// 			if v == url {
// 				return true
// 			}
// 		}
// 	}

// 	if len(method.Regs) > 0 {
// 		for _, v := range method.Regs {
// 			ok, err := regexp.MatchString(v, url)
// 			if err != nil {
// 				return false
// 			}
// 			if ok {
// 				return true
// 			}
// 		}
// 	}

// 	return false
// } /*
//  */
