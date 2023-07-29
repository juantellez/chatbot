// config.go

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config representa la configuración del sistema
type Config struct {
	DatabasePath  string `json:"database_path"`
	LNbitsAPIKey  string `json:"lnbits_api_key"`
	EmailSMTPHost string `json:"email_smtp_host"`
	EmailSMTPPort int    `json:"email_smtp_port"`
	EmailSMTPUser string `json:"email_smtp_user"`
	EmailSMTPPass string `json:"email_smtp_pass"`
}

// LoadConfig carga la configuración del sistema desde un archivo JSON
func LoadConfig(filePath string) (*Config, error) {
	// Leer el contenido del archivo de configuración
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	// Parsear el contenido del archivo JSON en la estructura Config
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return &config, nil
}
