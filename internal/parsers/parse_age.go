package parsers

import (
	"Effective_Mobile/internal/model"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type userAge struct {
	Age int `json:"age"`
}

func GetUserAge(name string, cfg *model.Config) (int, error) {
	client := http.DefaultClient
	url := cfg.AgeApi + name
	client.Timeout = 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return -1, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()
	user := &userAge{}
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return -1, err
	}
	return user.Age, nil
}
