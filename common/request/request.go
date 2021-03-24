package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	sjson "github.com/bitly/go-simplejson"
)

type Request struct {
	Url        string      `json:"url"`
	Method     string      `json:"method"`
	Params     interface{} `json:"params"`
	Payload    interface{} `json:"payload"`
	OperatorID string      `json:"operatorID"`
}

func (r *Request) Do() ([]byte, error) {
	fmt.Println("request:", r.Url)
	js, err := json.Marshal(r.Payload)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(r.Method, r.Url, bytes.NewReader(js))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	if r.Payload != nil {
		request.Header.Set("operatorID", r.OperatorID)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 300 {
		sjError, err := sjson.NewJson(result)
		if err != nil {
			return nil, err
		}
		errStr := sjError.GetPath("errors").GetIndex(0).GetPath("detail").MustString()
		return nil, errors.New(errStr)
	}

	return result, nil
} 