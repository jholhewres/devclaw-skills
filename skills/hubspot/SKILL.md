---
name: hubspot
version: 0.1.0
author: devclaw
description: "Manage HubSpot CRM — contacts, companies, deals, tickets"
category: integrations
tags: [hubspot, crm, contacts, deals, companies, marketing, sales]
requires:
  env: [HUBSPOT_ACCESS_TOKEN]
---
# HubSpot

Manage HubSpot CRM via the REST API v3.

## Setup

1. Go to HubSpot → Settings → Integrations → Private Apps
2. Create a private app with scopes: `crm.objects.contacts.read`, `crm.objects.contacts.write`, `crm.objects.companies.read`, `crm.objects.companies.write`, `crm.objects.deals.read`, `crm.objects.deals.write`, `content`
3. Copy the access token
4. Store in vault:
   ```bash
   copilot config vault-set HUBSPOT_ACCESS_TOKEN=pat-na1-xxxxx
   ```

## List Contacts

```bash
curl -s "https://api.hubapi.com/crm/v3/objects/contacts?limit=10&properties=firstname,lastname,email,phone,company" \
  -H "Authorization: Bearer $HUBSPOT_ACCESS_TOKEN" | jq '.results[] | {id, properties}'
```

## Search Contacts

```bash
curl -s -X POST "https://api.hubapi.com/crm/v3/objects/contacts/search" \
  -H "Authorization: Bearer $HUBSPOT_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "filterGroups": [{"filters": [{"propertyName": "email", "operator": "CONTAINS_TOKEN", "value": "example.com"}]}],
    "properties": ["firstname", "lastname", "email", "phone"],
    "limit": 10
  }' | jq '.results[] | {id, properties}'
```

## Create Contact

```bash
curl -s -X POST "https://api.hubapi.com/crm/v3/objects/contacts" \
  -H "Authorization: Bearer $HUBSPOT_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "properties": {
      "email": "contact@example.com",
      "firstname": "John",
      "lastname": "Doe",
      "phone": "+5511999999999",
      "company": "Company Name"
    }
  }' | jq '{id, properties}'
```

## Update Contact

```bash
curl -s -X PATCH "https://api.hubapi.com/crm/v3/objects/contacts/CONTACT_ID" \
  -H "Authorization: Bearer $HUBSPOT_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"properties": {"phone": "+5511888888888"}}' | jq '{id, properties}'
```

## List Companies

```bash
curl -s "https://api.hubapi.com/crm/v3/objects/companies?limit=10&properties=name,domain,industry,city,phone" \
  -H "Authorization: Bearer $HUBSPOT_ACCESS_TOKEN" | jq '.results[] | {id, properties}'
```

## Create Company

```bash
curl -s -X POST "https://api.hubapi.com/crm/v3/objects/companies" \
  -H "Authorization: Bearer $HUBSPOT_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "properties": {
      "name": "Company Name",
      "domain": "company.com",
      "industry": "TECHNOLOGY",
      "city": "São Paulo",
      "phone": "+551130000000"
    }
  }' | jq '{id, properties}'
```

## List Deals

```bash
curl -s "https://api.hubapi.com/crm/v3/objects/deals?limit=10&properties=dealname,amount,dealstage,closedate,pipeline" \
  -H "Authorization: Bearer $HUBSPOT_ACCESS_TOKEN" | jq '.results[] | {id, properties}'
```

## Create Deal

```bash
curl -s -X POST "https://api.hubapi.com/crm/v3/objects/deals" \
  -H "Authorization: Bearer $HUBSPOT_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "properties": {
      "dealname": "Deal Name",
      "amount": "10000",
      "dealstage": "qualifiedtobuy",
      "pipeline": "default",
      "closedate": "2026-03-31"
    }
  }' | jq '{id, properties}'
```

## List Tickets

```bash
curl -s "https://api.hubapi.com/crm/v3/objects/tickets?limit=10&properties=subject,content,hs_pipeline_stage,hs_ticket_priority" \
  -H "Authorization: Bearer $HUBSPOT_ACCESS_TOKEN" | jq '.results[] | {id, properties}'
```

## Associate Objects

```bash
# Associate contact with company
curl -s -X PUT "https://api.hubapi.com/crm/v3/objects/contacts/CONTACT_ID/associations/companies/COMPANY_ID/contact_to_company" \
  -H "Authorization: Bearer $HUBSPOT_ACCESS_TOKEN"
```

## Deal Pipelines

```bash
curl -s "https://api.hubapi.com/crm/v3/pipelines/deals" \
  -H "Authorization: Bearer $HUBSPOT_ACCESS_TOKEN" | jq '.results[] | {id, label, stages: [.stages[] | {id, label}]}'
```

## Tips

- API rate limit: 100 requests/10s for private apps
- Use search endpoint for filtered queries (up to 3 filter groups)
- Properties must be explicitly requested via `?properties=` parameter
- Deal stages vary by pipeline — list pipelines first
- Associations link contacts, companies, deals, and tickets
- Use batch endpoints for bulk operations (up to 100 records)
- Confirm before deleting any CRM records

## Triggers

hubspot, crm, contato, contact, empresa, company, negócio, deal, ticket, pipeline, lead
