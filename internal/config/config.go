package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AxiomConfig struct {
	AxiomURL   string
	AxiomToken string
}

type AppConfig struct {
	Axiom *AxiomConfig
}

func LoadAxiomConfig() *AxiomConfig {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(".env not found")
	}

	return &AxiomConfig{
		AxiomURL:   os.Getenv("AXIOM_URL"),
		AxiomToken: os.Getenv("AXIOM_TOKEN"),
	}
}

func LoadAppConfig() *AppConfig {
	return &AppConfig{
		Axiom: LoadAxiomConfig(),
	}
}
