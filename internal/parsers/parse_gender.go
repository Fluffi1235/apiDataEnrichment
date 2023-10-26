package parsers

import (
	"Effective_Mobile/internal/model"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type userGender struct {
	Gender string `json:"gender"`
}

func GetUserGender(name string, cfg *model.Config) (string, error) {
	client := http.DefaultClient
	url := cfg.GenderApi + name
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

	user := &userGender{}
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return "", err
	}

	return user.Gender, nil
}
