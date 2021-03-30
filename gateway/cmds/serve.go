package cmds

import (
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"vwood/app/gateway/application"

	"github.com/gorilla/handlers"
	"github.com/urfave/cli"
)

func serve(c *cli.Context) {
	configPath := c.Parent().String("config")
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	cfg := new(application.Config)
	err = json.Unmarshal(data, cfg)
	if err != nil {
		log.Fatal(err)
	}

	application.ApplicationContext = application.NewContext(cfg)

	switch {
	case c.Bool("80"):
		println("start server on :80")
		router := wrapRouter(application.ApplicationContext.Router80())
		err := http.ListenAndServe(":80", router)
		if err != nil {
			panic(err)
		}
		break
	case c.Bool("8443"):
		println("start server on :443")
		router := wrapRouter(application.ApplicationContext.Router443())
		// 暂时不用key
		err := http.ListenAndServeTLS(":443",
			cfg.SSLChainCrtPath,
			cfg.SSLKeyPath,
			router)
		err = http.ListenAndServe(":443", router)
		if err != nil {
			panic(err)
		}
	default:
		panic("open a port at a time")
	}
}

type contextMiddleware struct {
	next http.Handler
}

func (h *contextMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	h.next.ServeHTTP(w, r)
}

func wrapRouter(sr http.Handler) http.Handler {
	r := handlers.CompressHandlerLevel(sr, gzip.DefaultCompression)
	r = handlers.CombinedLoggingHandler(os.Stdout, r)
	r = &contextMiddleware{
		next: r,
	}

	r = handlers.CORS(
		handlers.AllowedHeaders([]string{"authorization", "content-type", "token"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{http.MethodPost, http.MethodPut,
			http.MethodDelete, http.MethodGet, http.MethodPatch, http.MethodOptions}),
	)(r)

	return r
}
