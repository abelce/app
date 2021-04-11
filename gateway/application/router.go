package application

import (
	normal_context "abelce/app/gateway/application/context"
	"abelce/app/gateway/application/middleware"
	"abelce/at"
	atjsonapi "abelce/at/jsonapi"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"runtime"
	"strings"

	"github.com/gorilla/mux"
)

func handleOthers(w http.ResponseWriter, r *http.Request) {
	tmp := strings.Split(r.URL.Path, "/")
	if len(tmp) < 3 {
		handleError(w, fmt.Errorf("invalid url:%s, url should be /v{number}/{apiservename}", r.URL.Path))
		return
	}
	api := tmp[2]

	turl, err := ApplicationContext.APIProxyURL(api)
	if at.Ensure(&err) {
		handleError(w, err)
		return
	}

	// 设置反向代理
	proxy := httputil.NewSingleHostReverseProxy(turl)
	proxy.ServeHTTP(w, r)
}

func NewRouter(port int) *mux.Router {
	r := mux.NewRouter()
	r = r.PathPrefix(ApplicationContext.Config().Version).Subrouter()
	// r.Host("www.example.com")
	// 添加中间件
	AddMiddleware(port, r)

	// r.HandleFunc("/registry", handleAuth).Methods(http.MethodPost)
	// r.HandleFunc("/auth", handleAuth).Methods(http.MethodPost)

	r.PathPrefix("/").HandlerFunc(handleOthers)

	return r
}

// 添加中间件
func AddMiddleware(port int, r *mux.Router) {
	// 创建中间建的上下文
	ctx := normal_context.NewNormalContext(port, ApplicationContext.Config())
	middleware.GetMiddlewares(ctx)
}


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

func handleError(w http.ResponseWriter, e error) {

	if e != nil {
		switch tErr := e.(type) {
		case *ServerErrors:
			w.Header().Set("Content-Type", "application/json")
			if tErr.Errors[0].Detail == "UNAUTHORIZED" {
				w.WriteHeader(401)
			} else {
				w.WriteHeader(400)
			}

			json, e := errs2doc(tErr.Errors)

			if e != nil {
				w.WriteHeader(500)
				fmt.Fprintln(w, e.Error())
			}
			fmt.Fprintln(w, json)
			return

		default:
			errs := []atjsonapi.JsonapiError{}
			at.Ensure(&tErr)
			_, f, l, _ := runtime.Caller(1)
			errs = append(errs, atjsonapi.JsonapiError{
				Code:   400,
				Detail: tErr.Error(),
				Meta: map[string]interface{}{
					"file": f,
					"line": l,
				},
			})
			w.Header().Set("Content-Type", "application/json")
			if strings.Contains(errs[0].Detail, "NOSESSIONTOKEN") {
				w.WriteHeader(401)
			} else {
				w.WriteHeader(400)
			}
			json, e := errs2doc(errs)
			if e != nil {
				w.WriteHeader(500)
				fmt.Fprintln(w, e.Error())
			}
			fmt.Fprintln(w, json)
			return
		}
	}
}

func errs2doc(errs []atjsonapi.JsonapiError) (string, error) {
	doc := atjsonapi.JsonapiDocument{
		Errors: errs,
	}
	b, e := json.Marshal(doc)
	if e != nil {
		return "", errors.New("can not encode errors object")
	}
	return string(b), nil
}

