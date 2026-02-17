---
name: stock-prices
version: 0.1.0
author: devclaw
description: "Stock prices and market data — quotes, charts, and financial info"
category: finance
tags: [stocks, finance, market, trading, quotes]
requires:
  bins: [curl, jq]
---
# Stock Prices

Get stock prices, market data, and financial information.

## Yahoo Finance (yfinance - Free)

```bash
# Get quote (requires yfinance)
pip install yfinance

# Python script for quotes
python3 -c "
import yfinance as yf
import json

ticker = yf.Ticker('AAPL')
info = ticker.info
print(json.dumps({
    'symbol': info.get('symbol'),
    'price': info.get('currentPrice'),
    'change': info.get('regularMarketChange'),
    'change_percent': info.get('regularMarketChangePercent'),
    'volume': info.get('regularMarketVolume'),
    'market_cap': info.get('marketCap')
}, indent=2))
"

# Get multiple quotes
python3 -c "
import yfinance as yf
import json

tickers = yf.Tickers('AAPL GOOGL MSFT')
for symbol, ticker in tickers.tickers.items():
    print(json.dumps({
        'symbol': symbol,
        'price': ticker.info.get('currentPrice')
    }))
"
```

## Alpha Vantage (Free API)

```bash
# Setup
export ALPHAVANTAGE_KEY="your_key"

# Get quote
curl -s "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=AAPL&apikey=$ALPHAVANTAGE_KEY" | jq '."Global Quote"'

# Intraday data
curl -s "https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=AAPL&interval=5min&apikey=$ALPHAVANTAGE_KEY" | jq '.["Time Series (5min)"] | to_entries | .[0:5]'

# Daily data
curl -s "https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=AAPL&apikey=$ALPHAVANTAGE_KEY" | jq '.["Time Series (Daily)"] | to_entries | .[0:5]'

# Company overview
curl -s "https://www.alphavantage.co/query?function=OVERVIEW&symbol=AAPL&apikey=$ALPHAVANTAGE_KEY" | jq '.'
```

## Finnhub (Free API)

```bash
# Setup
export FINNHUB_KEY="your_key"

# Get quote
curl -s "https://finnhub.io/api/v1/quote?symbol=AAPL&token=$FINNHUB_KEY" | jq '.'

# Company profile
curl -s "https://finnhub.io/api/v1/stock/profile2?symbol=AAPL&token=$FINNHUB_KEY" | jq '.'

# News
curl -s "https://finnhub.io/api/v1/company-news?symbol=AAPL&from=2025-01-01&to=2025-01-15&token=$FINNHUB_KEY" | jq '.[0:5]'

# Basic financials
curl -s "https://finnhub.io/api/v1/stock/metric?symbol=AAPL&metric=all&token=$FINNHUB_KEY" | jq '.metric'
```

## Twelve Data (Free API)

```bash
# Setup
export TWELVEDATA_KEY="your_key"

# Get quote
curl -s "https://api.twelvedata.com/quote?symbol=AAPL&apikey=$TWELVEDATA_KEY" | jq '.'

# Time series
curl -s "https://api.twelvedata.com/time_series?symbol=AAPL&interval=1day&outputsize=5&apikey=$TWELVEDATA_KEY" | jq '.values'

# Price
curl -s "https://api.twelvedata.com/price?symbol=AAPL&apikey=$TWELVEDATA_KEY" | jq '.'
```

## Polygon.io

```bash
# Setup
export POLYGON_KEY="your_key"

# Previous close
curl -s "https://api.polygon.io/v2/aggs/ticker/AAPL/prev?adjusted=true&apiKey=$POLYGON_KEY" | jq '.results[0]'

# Daily open/close
curl -s "https://api.polygon.io/v1/open-close/AAPL/2025-01-15?apiKey=$POLYGON_KEY" | jq '.'

# Company details
curl -s "https://api.polygon.io/v3/reference/tickers/AAPL?apiKey=$POLYGON_KEY" | jq '.results'
```

## Market Indices

```bash
# S&P 500, Dow Jones, NASDAQ
python3 -c "
import yfinance as yf
import json

indices = {'SPY': 'S&P 500', 'DIA': 'Dow Jones', 'QQQ': 'NASDAQ'}
for symbol, name in indices.items():
    ticker = yf.Ticker(symbol)
    print(json.dumps({
        'index': name,
        'symbol': symbol,
        'price': ticker.info.get('currentPrice'),
        'change': ticker.info.get('regularMarketChangePercent')
    }))
"
```

## Cryptocurrency Prices

```bash
# Using CoinGecko
curl -s "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum&vs_currencies=usd" | jq '.'
```

## Tips

- Alpha Vantage: 5 calls/minute, 500/day free
- Finnhub: 60 calls/minute free
- Twelve Data: 8 calls/minute, 800/day free
- Use yfinance for unlimited Python-based queries
- Cache results to avoid rate limits

## Triggers

stock, stock price, stock quote, market data, stock market,
stock api, finance, trading, nasdaq, nyse
