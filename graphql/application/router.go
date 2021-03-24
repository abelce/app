package application

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"vwood/app/graphql/queryType"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
)

const (
	version    = "/v1"
	successStr = `{"responseStatus":{"success":true}}`
)

type QueryParams struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

var Schema graphql.Schema

func init() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType.RootQueryType,
		// Mutation: MutationType,
	})
	if err != nil {
		log.Fatal(err)
	}
	Schema = schema
}

func graphqlHander(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var queryPrams QueryParams
	err = json.Unmarshal(data, &queryPrams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := graphql.Do(graphql.Params{
		Schema:         Schema,
		RequestString:  queryPrams.Query,
		VariableValues: queryPrams.Variables,
		OperationName:  queryPrams.Operation,
	})

	json.NewEncoder(w).Encode(result)
}

func NewRouter() *mux.Router {
	r := mux.NewRouter().PathPrefix("/" + version).Subrouter()
	// r.HandleFunc("/graphql/list", graphqlHander).Methods(http.MethodPost)
	r.HandleFunc("/graphql", graphqlHander).Methods(http.MethodPost)

	return r
}
