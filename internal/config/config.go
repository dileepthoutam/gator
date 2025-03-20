package config

import (
	"encoding/json"
	"log"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
    DbUrl string `json:"db_url"`
    CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(username string) {
    c.CurrentUserName = username
    cfg, err := json.Marshal(c)
    if err != nil {
        log.Fatal(err)
    }
    cfgFile, err := getConfigFilePath()
    if err != nil {
        log.Fatal("Error getting config file path.")
    }
    err = os.WriteFile(cfgFile, cfg, 0644)
    if err != nil {
        log.Fatal("Error writing to the file.")
    }
}

func Read() *Config {
    cfg, err := getConfigFilePath()
    file, err := os.ReadFile(cfg)
    if err != nil {
        log.Fatal("Error reading config file.")
    }

    var config Config
    err = json.Unmarshal(file, &config)
    if err != nil {
        log.Fatal("Error parsing the config file")
    }

    return &config
}

func getConfigFilePath() (string, error) {
    home, err := os.UserHomeDir()
    if err != nil {
        return "", err
    }
    return home + "/" + configFileName, nil
}

