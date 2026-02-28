---
name: shopify
version: 0.1.0
author: devclaw
description: "Manage Shopify store — products, orders, customers, inventory"
category: integrations
tags: [shopify, ecommerce, orders, products, customers, inventory]
requires:
  env: [SHOPIFY_STORE, SHOPIFY_ACCESS_TOKEN]
---
# Shopify

Manage a Shopify store via the Admin REST API.

## Setup

1. Go to your Shopify admin → Settings → Apps and sales channels → Develop apps
2. Create a custom app with Admin API access (scopes: `read_products`, `write_products`, `read_orders`, `write_orders`, `read_customers`, `read_inventory`)
3. Install the app and copy the Admin API access token
4. Store in vault:
   ```bash
   copilot config vault-set SHOPIFY_STORE=your-store.myshopify.com
   copilot config vault-set SHOPIFY_ACCESS_TOKEN=shpat_xxxxx
   ```

## List Products

```bash
curl -s "https://$SHOPIFY_STORE/admin/api/2024-01/products.json?limit=10" \
  -H "X-Shopify-Access-Token: $SHOPIFY_ACCESS_TOKEN" | jq '.products[] | {id, title, status, variants: [.variants[] | {id, price, inventory_quantity}]}'
```

## Get Product Details

```bash
curl -s "https://$SHOPIFY_STORE/admin/api/2024-01/products/PRODUCT_ID.json" \
  -H "X-Shopify-Access-Token: $SHOPIFY_ACCESS_TOKEN" | jq '.product | {id, title, body_html, vendor, product_type, tags, status}'
```

## Create Product

```bash
curl -s -X POST "https://$SHOPIFY_STORE/admin/api/2024-01/products.json" \
  -H "X-Shopify-Access-Token: $SHOPIFY_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "product": {
      "title": "Product Name",
      "body_html": "<p>Description</p>",
      "vendor": "Vendor",
      "product_type": "Type",
      "tags": "tag1, tag2",
      "variants": [{"price": "29.99", "sku": "SKU001"}]
    }
  }' | jq '.product | {id, title}'
```

## Update Product

```bash
curl -s -X PUT "https://$SHOPIFY_STORE/admin/api/2024-01/products/PRODUCT_ID.json" \
  -H "X-Shopify-Access-Token: $SHOPIFY_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"product": {"title": "New Title", "tags": "updated, tags"}}' | jq '.product | {id, title}'
```

## List Orders

```bash
# Recent orders
curl -s "https://$SHOPIFY_STORE/admin/api/2024-01/orders.json?status=any&limit=10" \
  -H "X-Shopify-Access-Token: $SHOPIFY_ACCESS_TOKEN" | jq '.orders[] | {id, name, total_price, financial_status, fulfillment_status, created_at}'

# Unfulfilled orders
curl -s "https://$SHOPIFY_STORE/admin/api/2024-01/orders.json?fulfillment_status=unfulfilled&limit=20" \
  -H "X-Shopify-Access-Token: $SHOPIFY_ACCESS_TOKEN" | jq '.orders[] | {id, name, total_price, customer: .customer.email}'
```

## Get Order Details

```bash
curl -s "https://$SHOPIFY_STORE/admin/api/2024-01/orders/ORDER_ID.json" \
  -H "X-Shopify-Access-Token: $SHOPIFY_ACCESS_TOKEN" | jq '.order | {id, name, total_price, line_items: [.line_items[] | {title, quantity, price}], shipping_address, customer: {name: .customer.first_name, email: .customer.email}}'
```

## List Customers

```bash
curl -s "https://$SHOPIFY_STORE/admin/api/2024-01/customers.json?limit=10" \
  -H "X-Shopify-Access-Token: $SHOPIFY_ACCESS_TOKEN" | jq '.customers[] | {id, first_name, last_name, email, orders_count, total_spent}'
```

## Search Customers

```bash
curl -s "https://$SHOPIFY_STORE/admin/api/2024-01/customers/search.json?query=email:customer@example.com" \
  -H "X-Shopify-Access-Token: $SHOPIFY_ACCESS_TOKEN" | jq '.customers[]'
```

## Inventory

```bash
# Get inventory levels for a location
curl -s "https://$SHOPIFY_STORE/admin/api/2024-01/inventory_levels.json?location_ids=LOCATION_ID&limit=50" \
  -H "X-Shopify-Access-Token: $SHOPIFY_ACCESS_TOKEN" | jq '.inventory_levels[]'

# Adjust inventory
curl -s -X POST "https://$SHOPIFY_STORE/admin/api/2024-01/inventory_levels/adjust.json" \
  -H "X-Shopify-Access-Token: $SHOPIFY_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"location_id": LOCATION_ID, "inventory_item_id": ITEM_ID, "available_adjustment": 5}'
```

## Refunds

```bash
# Calculate refund
curl -s -X POST "https://$SHOPIFY_STORE/admin/api/2024-01/orders/ORDER_ID/refunds/calculate.json" \
  -H "X-Shopify-Access-Token: $SHOPIFY_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"refund": {"refund_line_items": [{"line_item_id": LINE_ITEM_ID, "quantity": 1}]}}' | jq '.refund'
```

## Tips

- Store URL format: `your-store.myshopify.com` (without https://)
- API rate limit: 2 requests/second (REST), bucket-based for GraphQL
- Use `?fields=id,title` to limit response fields
- Pagination: use `page_info` cursor from Link header
- Always confirm before creating orders or processing refunds
- Use `status=any` to include archived orders

## Triggers

shopify, pedido, order, produto, product, estoque, inventory, cliente, customer, loja, store, reembolso, refund
