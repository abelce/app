/*
 * File: router.go
 * Project: application
 * File Created: Wednesday, 24th March 2021 1:51:31 pm
 * Author: zxtang (1061225829@qq.com)
 * -----
 * Last Modified: Wednesday, 24th March 2021 1:51:34 pm
 * Modified By: zxtang (1061225829@qq.com>)
 * -----
 * Copyright 2017 - 2021 Your Company, Your Company
 */
package application

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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

func NewRouter(schema graphql.Schema) *mux.Router {
	Schema = schema
	r := mux.NewRouter().PathPrefix("/v1").Subrouter()
	// r.HandleFunc("/graphql/list", graphqlHander).Methods(http.MethodPost)
	r.HandleFunc("/graphql", graphqlHander).Methods(http.MethodPost)

	return r
}
