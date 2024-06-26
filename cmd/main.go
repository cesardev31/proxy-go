package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config representa la estructura del archivo de configuración YAML
type Config struct {
	Servers []ServerConfig `yaml:"servers"`
}

// ServerConfig representa la configuración de cada servidor
type ServerConfig struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

func main() {
	// Cargar configuración desde el archivo YAML
	config, err := loadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Imprimir la configuración cargada (solo como ejemplo)
	fmt.Println("Servers:")
	for _, server := range config.Servers {
		fmt.Printf("Name: %s, URL: %s\n", server.Name, server.URL)
	}

	// Aquí continuarías con la lógica de tu aplicación, por ejemplo, utilizar esta configuración para configurar el proxy reverso, etc.
}

// loadConfig carga el archivo de configuración YAML
func loadConfig(filename string) (*Config, error) {
	// Leer el archivo YAML
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	// Decodificar el archivo YAML en la estructura Config
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error parsing config file: %v", err)
	}

	return &config, nil
}
