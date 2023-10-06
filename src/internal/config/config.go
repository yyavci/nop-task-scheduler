package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type AppConfig struct {
	StoreUrl string
}

func ReadConfiguration() (*AppConfig, error) {

	fmt.Println("reading configuration...")

	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Printf("Error occured opening config file! Err:%s\n", err)
		return nil, err
	}

	fmt.Println("opened config.json")

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("Error occured reading json file! Err:%s\n", err)
		return nil, err
	}

	var appConfig *AppConfig
	err = json.Unmarshal(byteValue, &appConfig)
	if err != nil {
		fmt.Printf("Error occured parsing json file! Err:%s\n", err)
		return nil, err
	}

	if len(appConfig.StoreUrl) == 0 {
		return nil, errors.New("store url is empty")
	}

	return appConfig, nil
}
