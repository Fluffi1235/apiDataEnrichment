package config

import (
	"Effective_Mobile/internal/model"
	"github.com/joho/godotenv"
	"os"
)

func ReadConfigEnv() (*model.Config, error) {
	path := "./config/config.env"
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}
	databaseURL := os.Getenv("connectDb")
	ageApi := os.Getenv("ageApi")
	genderApi := os.Getenv("genderApi")
	countryApi := os.Getenv("countryApi")
	cfg := &model.Config{
		DBUrl:      databaseURL,
		AgeApi:     ageApi,
		GenderApi:  genderApi,
		CountryApi: countryApi,
	}
	return cfg, nil
}
