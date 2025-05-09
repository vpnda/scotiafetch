# ScotiaFetch - Go API Client for Scotiabank

A clean, modern Go client for interacting with Scotiabank's API.

## Overview

**ScotiaFetch** is a Go client library that provides seamless integration with Scotiabank's banking services API. This client was generated using [OpenAPI Generator](https://openapi-generator.tech) based on Scotiabank's official API specifications.

With ScotiaFetch, you can access account information, transactions, credit card details, rewards, and more from your Go applications.

## Features

- Account summaries and balances
- Credit card information and transactions
- Banking transaction history
- Reward program details
- Campaign information

## Installation

```sh
go get github.com/vpnda/scotiafetch
```

### Dependencies

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

## Quick Start

```go
import (
	"context"
	openapi "github.com/vpnda/scotiafetch"
)

func main() {
	client := openapi.NewAPIClient(openapi.NewConfiguration())
	ctx := context.Background()
	
	// Get account summary
	summary, resp, err := client.DefaultAPI.ApiAccountsSummaryGet(ctx).Execute()
	if err != nil {
		// Handle error
	}
	
	// Use summary data...
}
```

## Proxy Configuration

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Server Configuration

### Custom Server Selection

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Server Variables

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

## Available API Endpoints

All URIs are relative to *https://secure.scotiabank.com*

| Endpoint | Description |
|----------|-------------|
| `GET /api/accounts/summary` | Get accounts summary |
| `GET /api/campaigns` | Get marketing campaigns |
| `GET /api/cards/{cardId}` | Get card details |
| `GET /api/credit/{creditId}` | Get credit account details |
| `GET /api/credit/{creditId}/transactions` | Get credit account transactions |
| `GET /api/mpsa-accounts/{accountId}` | Get MPSA account details |
| `GET /api/mpsa-accounts/{accountId}/transactions` | Get MPSA account transactions |
| `GET /api/rewards/{rewardId}` | Get rewards information |
| `GET /api/transactions/deposit-accounts/{depositAccountId}` | Get deposit account transactions |

## Documentation

For detailed documentation of all models and endpoints, see the [docs](docs/) directory.

## Author

Victor Pineda
