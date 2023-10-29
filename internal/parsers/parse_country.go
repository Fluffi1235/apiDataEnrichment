package parsers

import (
	"Effective_Mobile/internal/model"
	"context"
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"time"
)

type userCountry struct {
	Country []country `json:"country"`
}

type country struct {
	CountryId string `json:"country_id"`
}

func GetUserCountry(name string, cfg *model.Config) (string, error) {
	client := http.DefaultClient
	url := cfg.CountryApi + name
	client.Timeout = time.Duration(cfg.ClientTimeOut) * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ContextTimeOut)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	countries := &userCountry{}
	err = json.NewDecoder(resp.Body).Decode(&countries)
	if err != nil {
		return "", err
	}
	if len(countries.Country) == 0 {
		return "", errors.New("Country error")
	}
	index := rand.Intn(len(countries.Country) - 1)

	return countries.Country[index].CountryId, nil
}
