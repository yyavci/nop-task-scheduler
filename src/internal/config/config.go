package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type AppConfig struct {
	ConnectionString string
}

var errOpeningConfig = errors.New("error opening config file")
var errReadingConfig = errors.New("error reading config file")
var errParsingConfig = errors.New("error parsing config file")

func ReadConfiguration(file string) (*AppConfig, error) {

	fmt.Println("reading configuration...")

	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Printf("error: error occured opening config file! err:%+v\n %+v\n", errOpeningConfig, err)
		return nil, errOpeningConfig
	}

	fmt.Println("opened config.json")

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("error: error occured reading json file! Err:%+v\n %+v\n", errReadingConfig, err)
		return nil, errReadingConfig
	}
	var appConfig *AppConfig
	err = json.Unmarshal(byteValue, &appConfig)
	if err != nil {
		fmt.Printf("error: error occured parsing json file! Err:%+v\n %+v\n", errParsingConfig, err)
		return nil, errParsingConfig
	}

	return appConfig, nil
}
