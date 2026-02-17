---
name: gcloud
description: "Google Cloud CLI para Compute, Cloud Run, GCS e BigQuery"
metadata: {"openclaw":{"always":false,"emoji":"🌐"}}
---
# Google Cloud CLI

Interface com serviços GCP via gcloud e gsutil.

## Compute Engine

```bash
# Listar instâncias
gcloud compute instances list

# Detalhes
gcloud compute instances describe <name> --zone=<zone>

# SSH
gcloud compute ssh <name> --zone=<zone>
gcloud compute ssh <name> --zone=<zone> --command="<cmd>"

# Start/Stop
gcloud compute instances start <name> --zone=<zone>
gcloud compute instances stop <name> --zone=<zone>

# SCP (transferir arquivos)
gcloud compute scp <local> <name>:<remote> --zone=<zone>
gcloud compute scp <name>:<remote> <local> --zone=<zone>

# Serial port (debug boot)
gcloud compute instances get-serial-port-output <name> --zone=<zone>
```

## Cloud Run

```bash
# Listar services
gcloud run services list

# Deploy
gcloud run deploy <service> --image=<image> --region=<region> --allow-unauthenticated
gcloud run deploy <service> --source=. --region=<region>

# Logs
gcloud run services logs read <service> --region=<region> --limit=50

# Describe
gcloud run services describe <service> --region=<region>
```

## Cloud Storage (gsutil)

```bash
# Listar buckets
gsutil ls

# Listar objetos
gsutil ls gs://<bucket>/<prefix>/

# Upload/Download
gsutil cp <file> gs://<bucket>/<path>
gsutil cp gs://<bucket>/<path> <file>

# Sync
gsutil -m rsync -r <dir> gs://<bucket>/<prefix>/

# Remover
gsutil rm gs://<bucket>/<path>
gsutil rm -r gs://<bucket>/<prefix>/
```

## BigQuery

```bash
# Listar datasets
bq ls

# Query
bq query --use_legacy_sql=false 'SELECT * FROM `project.dataset.table` LIMIT 10'

# Listar tabelas
bq ls <dataset>

# Schema
bq show --schema <dataset>.<table>
```

## Configuração

```bash
# Projeto ativo
gcloud config get-value project
gcloud config set project <project-id>

# Conta ativa
gcloud auth list
gcloud config set account <email>

# Região/Zona padrão
gcloud config set compute/region <region>
gcloud config set compute/zone <zone>
```

## Tips

- Use `--format=json` ou `--format=table` para output estruturado
- Use `--project=<id>` para operar em outro projeto
- Use `--quiet` para suprimir confirmações
- Para auth de service account: `gcloud auth activate-service-account --key-file=<key.json>`
