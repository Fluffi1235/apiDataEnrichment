package config

import (
	"Effective_Mobile/internal/model"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func ReadConfigEnv() (*model.Config, error) {
	var path string
	flag.StringVar(&path, "path", "./config/config.env", "The path to the file")

	flag.Parse()

	fmt.Printf("Config path: %s\n", path)

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
