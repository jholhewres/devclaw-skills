---
name: payments
version: 0.1.0
author: goclaw
description: "Payment processing — Stripe, PayPal, Lightning Network payments"
category: finance
tags: [payments, stripe, paypal, bitcoin, lightning, invoice]
requires:
  bins: [curl, jq]
---
# Payments

Process payments using various payment providers.

## Stripe

```bash
# Setup
export STRIPE_SECRET_KEY="sk_test_xxx"

# Create payment intent
curl -s -X POST "https://api.stripe.com/v1/payment_intents" \
  -u "$STRIPE_SECRET_KEY:" \
  -d "amount=2000" \
  -d "currency=usd" \
  -d "payment_method_types[]=card" | jq '.'

# Retrieve payment intent
curl -s "https://api.stripe.com/v1/payment_intents/pi_xxx" \
  -u "$STRIPE_SECRET_KEY:" | jq '.'

# Create customer
curl -s -X POST "https://api.stripe.com/v1/customers" \
  -u "$STRIPE_SECRET_KEY:" \
  -d "email=customer@example.com" \
  -d "name=John Doe" | jq '.'

# Create invoice
curl -s -X POST "https://api.stripe.com/v1/invoices" \
  -u "$STRIPE_SECRET_KEY:" \
  -d "customer=cus_xxx" | jq '.'

# List charges
curl -s "https://api.stripe.com/v1/charges?limit=10" \
  -u "$STRIPE_SECRET_KEY:" | jq '.data[]'

# Create checkout session
curl -s -X POST "https://api.stripe.com/v1/checkout/sessions" \
  -u "$STRIPE_SECRET_KEY:" \
  -d "mode=payment" \
  -d "success_url=https://example.com/success" \
  -d "line_items[0][price]=price_xxx" \
  -d "line_items[0][quantity]=1" | jq '.url'
```

## PayPal

```bash
# Setup
export PAYPAL_CLIENT_ID="xxx"
export PAYPAL_SECRET="xxx"
export PAYPAL_API="https://api-m.sandbox.paypal.com"  # or api-m.paypal.com for live

# Get access token
ACCESS_TOKEN=$(curl -s -X POST "$PAYPAL_API/v1/oauth2/token" \
  -u "$PAYPAL_CLIENT_ID:$PAYPAL_SECRET" \
  -d "grant_type=client_credentials" | jq -r '.access_token')

# Create order
curl -s -X POST "$PAYPAL_API/v2/checkout/orders" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -d '{
    "intent": "CAPTURE",
    "purchase_units": [{
      "amount": {"currency_code": "USD", "value": "100.00"}
    }]
  }' | jq '.'

# Capture order
curl -s -X POST "$PAYPAL_API/v2/checkout/orders/ORDER_ID/capture" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ACCESS_TOKEN" | jq '.'
```

## Lightning Network (Alby)

```bash
# Setup
export ALBY_ACCESS_TOKEN="xxx"

# Get balance
curl -s "https://api.getalby.com/balance" \
  -H "Authorization: Bearer $ALBY_ACCESS_TOKEN" | jq '.'

# Create invoice
curl -s -X POST "https://api.getalby.com/invoices" \
  -H "Authorization: Bearer $ALBY_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"amount": 1000, "description": "Payment for service"}' | jq '.'

# Pay invoice
curl -s -X POST "https://api.getalby.com/payments/bolt11" \
  -H "Authorization: Bearer $ALBY_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"invoice": "lnbc..."}' | jq '.'

# Check invoice status
curl -s "https://api.getalby.com/invoices/PAYMENT_HASH" \
  -H "Authorization: Bearer $ALBY_ACCESS_TOKEN" | jq '.'
```

## OpenNode (Bitcoin)

```bash
# Setup
export OPENNODE_KEY="xxx"

# Create invoice
curl -s -X POST "https://api.opennode.com/v1/charges" \
  -H "Authorization: $OPENNODE_KEY" \
  -H "Content-Type: application/json" \
  -d '{"amount": 1000, "currency": "USD", "description": "Order #123"}' | jq '.'

# Check charge status
curl -s "https://api.opennode.com/v1/charge/CHARGE_ID" \
  -H "Authorization: $OPENNODE_KEY" | jq '.'
```

## LNURL (Lightning)

```bash
# Generate LNURL-pay
# Requires lnurl library
pip install lnurl

python3 -c "
from lnurl import LnurlPayResponse
import json

# Example LNURL-pay metadata
data = {
    'callback': 'https://your-domain.com/callback',
    'maxSendable': 1000000,
    'minSendable': 1000,
    'metadata': '[["text/plain", "Payment for service"]]',
    'tag': 'payRequest'
}
print(json.dumps(data))
"
```

## Tips

- Stripe: Test with `sk_test_` keys first
- PayPal: Use sandbox for testing
- Lightning: Alby offers easy API for Bitcoin
- Always use HTTPS for payment APIs
- Store payment IDs for reconciliation

## Triggers

payment, stripe, paypal, bitcoin, lightning, invoice,
create payment, process payment, payment api
