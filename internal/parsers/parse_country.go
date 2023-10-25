package parsers

import (
	"Effective_Mobile/internal/model"
	"context"
	"encoding/json"
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
	client.Timeout = 3 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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
	index := rand.Intn(len(countries.Country) - 1)

	return countries.Country[index].CountryId, nil
}
