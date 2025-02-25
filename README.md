# ccxt-go

# Recommended Structure

```bash
your-project/
│
├── cmd/                    # Executable applications (main function entry point)
│   └── your-project/
│       └── main.go
│
├── internal/                # Internal packages (inaccessible from outside)
│   ├── exchanges/           # Exchange-related logic
│   │   ├── binance.go
│   │   └── bybit.go
│   │
│   ├── trading/             # Trading strategies and execution logic
│   │   ├── arbitrage.go
│   │   ├── risk_management.go
│   │   └── market_making.go
│   │
│   ├── utils/               # Common utility functions
│   │   ├── logger.go
│   │   └── config.go
│   │
│   └── models/              # Data model definitions
│       ├── order.go
│       └── ticker.go
│
├── config/                  # Configuration files (e.g., API keys, logging settings)
│   └── config.yaml
│
├── scripts/                 # Automation scripts (deployment, testing, etc.)
│   └── deploy.sh
│
├── test/                    # Test code
│   ├── exchanges_test.go
│   ├── trading_test.go
│   └── utils_test.go
│
├── go.mod                   # Go module configuration file
├── go.sum                   # Dependency hash file
├── .env                     # Environment variables file (API keys, etc., should be added to .gitignore)
├── .gitignore
└── README.md

```