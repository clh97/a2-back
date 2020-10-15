package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"main/types"
	"net/http"
)

// FetchDataByCep fetches data from specific CEP endpoint and returns as types.Address structure
func FetchDataByCep(cep string) (*types.Address, error) {
	res, err := http.Get(fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%s: %s", "Could not fetch data", res.Status)
	}

	buf, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, errors.New("Could not read response data")
	}

	result := types.Address{}

	err = json.Unmarshal(buf, &result)

	if err != nil {
		return nil, errors.New("Could not unmarshal data")
	}

	return &result, nil
}
