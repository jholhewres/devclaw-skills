---
name: azure-cli
description: "Azure CLI for VMs, App Service, Blob Storage, Azure SQL"
---
# Azure CLI

Use the **bash** tool with the az CLI for Azure cloud operations.

## Auth & Account
```bash
az login
az account show
az account list --output table
az account set --subscription "NAME"
```

## Virtual Machines
```bash
az vm list --output table
az vm start --resource-group RG --name VM
az vm stop --resource-group RG --name VM
az vm show --resource-group RG --name VM --output json
az vm create --resource-group RG --name VM --image Ubuntu2204 --size Standard_B2s
```

## App Service
```bash
az webapp list --output table
az webapp log tail --resource-group RG --name APP
az webapp deployment source config-zip --resource-group RG --name APP --src app.zip
az webapp config appsettings set --resource-group RG --name APP --settings KEY=VALUE
```

## Blob Storage
```bash
az storage blob list --container-name CONTAINER --account-name ACCOUNT --output table
az storage blob upload --file FILE --container-name CONTAINER --name BLOB --account-name ACCOUNT
az storage blob download --container-name CONTAINER --name BLOB --file OUTPUT --account-name ACCOUNT
```

## Azure SQL
```bash
az sql server list --output table
az sql db list --server SERVER --resource-group RG --output table
```

## Tips
- Use --output table for readable output, --output json for parsing
- Use vault_get to retrieve Azure credentials
- Use az configure --defaults to set default resource group
