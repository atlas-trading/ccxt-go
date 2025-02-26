package base

import (
	"github.com/atlas-trading/ccxt-go/common"
)

// BaseExchange provides common functionality for all exchanges
type BaseExchange struct {
	Name      string
	ApiKey    string
	ApiSecret string
	Client    *common.HTTPClient
}

// NewBaseExchange creates a new base exchange instance
func NewBaseExchange(name string) *BaseExchange {
	return &BaseExchange{
		Name:   name,
		Client: common.NewHTTPClient(),
	}
}

// GetName returns the exchange name
func (e *BaseExchange) GetName() string {
	return e.Name
}

// SetCredentials sets API credentials
func (e *BaseExchange) SetCredentials(apiKey string, apiSecret string) {
	e.ApiKey = apiKey
	e.ApiSecret = apiSecret
} 