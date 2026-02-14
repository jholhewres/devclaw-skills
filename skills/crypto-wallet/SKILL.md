---
name: crypto-wallet
version: 0.1.0
author: goclaw
description: "Cryptocurrency wallet — check balances, addresses, and transactions"
category: finance
tags: [crypto, bitcoin, ethereum, wallet, blockchain]
requires:
  bins: [curl, jq]
---
# Crypto Wallet

Check cryptocurrency balances, addresses, and transactions.

## Bitcoin

```bash
# Check address balance (blockchain.com)
curl -s "https://blockchain.info/q/addressbalance/BITCOIN_ADDRESS" | jq '.'

# Address info
curl -s "https://blockchain.info/rawaddr/BITCOIN_ADDRESS" | jq '{
  balance: .final_balance,
  total_received: .total_received,
  n_tx: .n_tx
}'

# Current price
curl -s "https://blockchain.info/ticker" | jq '.USD.last'

# Transaction info
curl -s "https://blockchain.info/rawtx/TX_HASH" | jq '{hash, size, block_height}'
```

## Ethereum

```bash
# Check balance (etherscan.io)
export ETHERSCAN_API_KEY="your_key"
curl -s "https://api.etherscan.io/api?module=account&action=balance&address=ADDRESS&tag=latest&apikey=$ETHERSCAN_API_KEY" | jq '.result'

# Get ETH price
curl -s "https://api.etherscan.io/api?module=stats&action=ethprice&apikey=$ETHERSCAN_API_KEY" | jq '.result'

# Transaction list
curl -s "https://api.etherscan.io/api?module=account&action=txlist&address=ADDRESS&apikey=$ETHERSCAN_API_KEY" | jq '.result[:5]'

# ERC-20 token balance
curl -s "https://api.etherscan.io/api?module=account&action=tokenbalance&contractaddress=TOKEN_ADDR&address=WALLET_ADDR&apikey=$ETHERSCAN_API_KEY" | jq '.result'
```

## Multi-chain (BlockCypher)

```bash
# BTC, LTC, DOGE supported
curl -s "https://api.blockcypher.com/v1/btc/main/addrs/ADDRESS/balance" | jq '.'

# Ethereum
curl -s "https://api.blockcypher.com/v1/eth/main/addrs/ADDRESS/balance" | jq '.'

# Litecoin
curl -s "https://api.blockcypher.com/v1/ltc/main/addrs/ADDRESS/balance" | jq '.'
```

## Prices (CoinGecko - Free)

```bash
# Current prices
curl -s "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum,solana&vs_currencies=usd" | jq '.'

# Coin info
curl -s "https://api.coingecko.com/api/v3/coins/bitcoin" | jq '{
  name, symbol,
  price: .market_data.current_price.usd,
  change_24h: .market_data.price_change_percentage_24h
}'

# Top coins
curl -s "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=10" | jq '.[] | {name, current_price, market_cap}'
```

## Prices (CoinCap - Free)

```bash
# Get asset price
curl -s "https://api.coincap.io/v2/assets/bitcoin" | jq '.data.priceUsd'

# Get multiple assets
curl -s "https://api.coincap.io/v2/assets?ids=bitcoin,ethereum" | jq '.data[]'

# Historical prices
curl -s "https://api.coincap.io/v2/assets/bitcoin/history?interval=d1" | jq '.data[-7:]'
```

## Solana

```bash
# Balance (mainnet)
curl -s -X POST "https://api.mainnet-beta.solana.com" \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","id":1,"method":"getBalance","params":["WALLET_ADDRESS"]}' | jq '.result.value'
```

## Transaction Tracking

```bash
# Bitcoin transaction
curl -s "https://blockchain.info/rawtx/TX_HASH" | jq '{
  hash, block_height, inputs: .inputs | length, outputs: .out | length
}'

# Ethereum transaction
curl -s "https://api.etherscan.io/api?module=proxy&action=eth_getTransactionByHash&txhash=TX_HASH&apikey=$ETHERSCAN_API_KEY" | jq '.result'
```

## Generate Addresses (offline)

```bash
# Generate random private key (for testing only!)
openssl rand -hex 32

# Use bitcoinlib or similar for proper key management
pip install bitcoinlib
python -c "from bitcoinlib.wallets import Wallet; w = Wallet.create('test'); print(w.get_key().address)"
```

## Tips

- Never share private keys
- Use API keys for higher rate limits
- Etherscan: 5 calls/second free tier
- CoinGecko: 10-30 calls/minute free tier
- Always verify addresses before sending

## Triggers

crypto, bitcoin, ethereum, wallet balance, crypto price,
check crypto, blockchain, cryptocurrency
