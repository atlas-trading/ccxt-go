package types

// Market represents trading market information
type Market struct {
	ID           string
	Symbol       string
	Base        string
	Quote       string
	Active      bool
	Precision   MarketPrecision
	Limits      MarketLimits
	MakerFee    float64
	TakerFee    float64
}

type MarketPrecision struct {
	Amount int
	Price  int
}

type MarketLimits struct {
	Amount MarketMinMax
	Price  MarketMinMax
	Cost   MarketMinMax
}

type MarketMinMax struct {
	Min float64
	Max float64
} 