package config

import (
	"Effective_Mobile/internal/model"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
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
	clientTimeOutStr := os.Getenv("clientTimeOut")
	contextTimeOutStr := os.Getenv("contextTimeOut")

	clientTimeOut, err := strconv.Atoi(clientTimeOutStr)
	if err != nil {
		return nil, err
	}

	contextTimeOut, err := strconv.Atoi(contextTimeOutStr)
	if err != nil {
		return nil, err
	}

	cfg := &model.Config{
		DBUrl:          databaseURL,
		AgeApi:         ageApi,
		GenderApi:      genderApi,
		CountryApi:     countryApi,
		ClientTimeOut:  clientTimeOut,
		ContextTimeOut: contextTimeOut,
	}

	return cfg, nil
}
