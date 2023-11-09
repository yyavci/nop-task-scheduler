package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type AppConfig struct {
	ConnectionString string
}

func ReadConfiguration(file string) (*AppConfig, error) {

	fmt.Println("reading configuration...")

	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error occured opening config file! Err:%+v\n", err)
		return nil, err
	}

	fmt.Println("opened config.json")

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("Error occured reading json file! Err:%+v\n", err)
		return nil, err
	}
	var appConfig *AppConfig
	err = json.Unmarshal(byteValue, &appConfig)
	if err != nil {
		fmt.Printf("Error occured parsing json file! Err:%+v\n", err)
		return nil, err
	}

	return appConfig, nil
}
