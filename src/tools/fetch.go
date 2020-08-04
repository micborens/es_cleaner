package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Indice struct {
	Indice string                  `json:"index"`
}

type Indices []Indice

func (i *Indices) GetIndicesInfos(endpoint string) (err error) {
	url := fmt.Sprintf("%s/_cat/indices?format=json", endpoint)
	var req *http.Response
	req, err = http.Get(url)
	if err != nil {
		return
	}
	defer req.Body.Close()
	err = json.NewDecoder(req.Body).Decode(i)
	return
}

func NewIndicesList() *Indices {
	return new(Indices)
}
