package common

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type ExchangeCredentials struct {
	APIKey    string
	APISecret string
}

type Config struct {
	Binance ExchangeCredentials
	Bybit   ExchangeCredentials
}

var (
	config *Config
	once   sync.Once
)

// GetConfig returns the singleton config instance
func GetConfig() *Config {
	once.Do(func() {
		// Load .env file
		godotenv.Load()

		config = &Config{
			Binance: ExchangeCredentials{
				APIKey:    os.Getenv("BINANCE_API_KEY"),
				APISecret: os.Getenv("BINANCE_API_SECRET"),
			},
			Bybit: ExchangeCredentials{
				APIKey:    os.Getenv("BYBIT_API_KEY"),
				APISecret: os.Getenv("BYBIT_API_SECRET"),
			},
		}
	})
	return config
} 