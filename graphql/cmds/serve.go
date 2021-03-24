package cmds

import (
	"vwood/app/graphql/application"

	"github.com/gorilla/handlers"
	"github.com/urfave/cli"

	"log"
	"net/http"
	"os"
	"strconv"
)

type contentTypeMiddleware struct {
	next http.Handler
}

func (h *contentTypeMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	log.Printf("method:%s, url:%s\n", r.Method, r.URL.String())
	h.next.ServeHTTP(w, r)
}

func Serve(c *cli.Context) {
	cpath := c.Parent().String("config")
	var err error
	application.ApplicationContext, err = application.NewContext(cpath)
	if err != nil {
		panic(err)
	}

	routeHandler := handlers.CombinedLoggingHandler(os.Stdout, application.NewRouter())
	routeHandler = &contentTypeMiddleware{
		next: routeHandler,
	}

	port := int(application.ApplicationContext.GetConfig().Port)
	if port == 0 {
		port = 3064
	}
	log.Printf("start comment service on %d\n", port)
	err = http.ListenAndServe(":"+strconv.Itoa(port), routeHandler)
	if err != nil {
		log.Fatal(err)
	}
}
