package types

import (
	"context"
	"time"
)

// Exchange defines the standard interface that all exchanges must implement
type Exchange interface {
	// Basic Information
	GetName() string
	GetVersion() string
	
	// Public API
	LoadMarkets(ctx context.Context, reload bool) ([]Market, error)
	FetchMarkets(ctx context.Context) ([]Market, error)
	FetchCurrencies(ctx context.Context) (map[string]Currency, error)
	FetchTicker(ctx context.Context, symbol string) (*Ticker, error)
	FetchTickers(ctx context.Context, symbols []string) (map[string]*Ticker, error)
	FetchOrderBook(ctx context.Context, symbol string, limit int) (*OrderBook, error)
	FetchOHLCV(ctx context.Context, symbol string, timeframe string, since time.Time, limit int) ([]OHLCV, error)
	FetchStatus(ctx context.Context) (*ExchangeStatus, error)
	FetchTrades(ctx context.Context, symbol string, since time.Time, limit int) ([]Trade, error)
	
	// Private API
	FetchBalance(ctx context.Context) (*Balance, error)
	CreateOrder(ctx context.Context, symbol string, orderType OrderType, side OrderSide, amount float64, price float64) (*Order, error)
	CancelOrder(ctx context.Context, id string, symbol string) error
	FetchOrder(ctx context.Context, id string, symbol string) (*Order, error)
	FetchOrders(ctx context.Context, symbol string, since time.Time, limit int) ([]Order, error)
	FetchOpenOrders(ctx context.Context, symbol string) ([]Order, error)
	FetchClosedOrders(ctx context.Context, symbol string) ([]Order, error)
	FetchMyTrades(ctx context.Context, symbol string, since time.Time, limit int) ([]Trade, error)
	
	// Deposit & Withdrawal
	FetchDepositAddress(ctx context.Context, currency string) (*DepositAddress, error)
	CreateDepositAddress(ctx context.Context, currency string) (*DepositAddress, error)
	Withdraw(ctx context.Context, currency string, amount float64, address string, tag string) (*WithdrawalResponse, error)
	
	// Authentication
	SetCredentials(apiKey string, apiSecret string)
}

// Ticker represents current market status
type Ticker struct {
	Symbol        string
	Last         float64
	Bid          float64
	Ask          float64
	High         float64
	Low          float64
	Volume       float64
	Timestamp    time.Time
}

// OrderBook represents market depth
type OrderBook struct {
	Symbol     string
	Timestamp  time.Time
	Bids      [][2]float64 // [price, amount]
	Asks      [][2]float64 // [price, amount]
}

// Trade represents a single trade
type Trade struct {
	ID        string
	Symbol    string
	Side      OrderSide
	Price     float64
	Amount    float64
	Timestamp time.Time
}

// OHLCV represents candlestick data
type OHLCV struct {
	Timestamp time.Time
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    float64
}

// Balance represents account balance
type Balance struct {
	Total     map[string]float64
	Free      map[string]float64
	Used      map[string]float64
	Timestamp time.Time
}

// Currency represents a cryptocurrency or fiat currency
type Currency struct {
	ID          string
	Code        string
	Name        string
	Active      bool
	Precision   int
	Limits      CurrencyLimits
	Networks    map[string]NetworkInfo
}

// ExchangeStatus represents the current status of an exchange
type ExchangeStatus struct {
	Status      string
	Updated     time.Time
	Eta         time.Time
	URL         string
	Message     string
}

// DepositAddress represents a cryptocurrency deposit address
type DepositAddress struct {
	Currency    string
	Address     string
	Tag         string
	Network     string
}

// WithdrawalResponse represents the response from a withdrawal request
type WithdrawalResponse struct {
	ID          string
	Currency    string
	Amount      float64
	Address     string
	Tag         string
	Status      string
	Timestamp   time.Time
	TxID        string
}

// NetworkInfo represents blockchain network information for a currency
type NetworkInfo struct {
	ID              string
	Network         string
	Active          bool
	DepositFee      float64
	WithdrawalFee   float64
	Limits          CurrencyLimits
}

// CurrencyLimits defines minimum and maximum values for currency operations
type CurrencyLimits struct {
	Deposit     MinMax
	Withdrawal  MinMax
}

// MinMax represents minimum and maximum values
type MinMax struct {
	Min float64
	Max float64
} 