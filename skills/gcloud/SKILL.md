---
name: gcloud
description: "Google Cloud CLI for Compute, Cloud Run, GCS, and BigQuery"
metadata: {"openclaw":{"always":false,"emoji":"🌐"}}
---
# Google Cloud CLI

Manage Google Cloud services via gcloud and gsutil.

## Setup

1. **Check if installed:**
   ```bash
   command -v gcloud && gcloud --version
   ```

2. **Install:**
   ```bash
   # macOS
   brew install google-cloud-sdk

   # Official installer (Linux)
   curl https://sdk.cloud.google.com | bash
   exec -l $SHELL
   ```

3. **Auth:**
   ```bash
   # Interactive (browser)
   gcloud auth login

   # Non-interactive (service account)
   gcloud auth activate-service-account --key-file=<key.json>
   ```

## Compute Engine

```bash
# List instances
gcloud compute instances list

# Details
gcloud compute instances describe <name> --zone=<zone>

# SSH
gcloud compute ssh <name> --zone=<zone>
gcloud compute ssh <name> --zone=<zone> --command="<cmd>"

# Start/Stop
gcloud compute instances start <name> --zone=<zone>
gcloud compute instances stop <name> --zone=<zone>

# SCP (transfer files)
gcloud compute scp <local> <name>:<remote> --zone=<zone>
gcloud compute scp <name>:<remote> <local> --zone=<zone>

# Serial port (debug boot)
gcloud compute instances get-serial-port-output <name> --zone=<zone>
```

## Cloud Run

```bash
# List services
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
# List buckets
gsutil ls

# List objects
gsutil ls gs://<bucket>/<prefix>/

# Upload/Download
gsutil cp <file> gs://<bucket>/<path>
gsutil cp gs://<bucket>/<path> <file>

# Sync
gsutil -m rsync -r <dir> gs://<bucket>/<prefix>/

# Remove
gsutil rm gs://<bucket>/<path>
gsutil rm -r gs://<bucket>/<prefix>/
```

## BigQuery

```bash
# List datasets
bq ls

# Query
bq query --use_legacy_sql=false 'SELECT * FROM `project.dataset.table` LIMIT 10'

# List tables
bq ls <dataset>

# Schema
bq show --schema <dataset>.<table>
```

## Configuration

```bash
# Active project
gcloud config get-value project
gcloud config set project <project-id>

# Active account
gcloud auth list
gcloud config set account <email>

# Default region/zone
gcloud config set compute/region <region>
gcloud config set compute/zone <zone>
```

## Tips

- Use `--format=json` or `--format=table` for structured output
- Use `--project=<id>` to operate on another project
- Use `--quiet` to suppress confirmations
- For service account auth: `gcloud auth activate-service-account --key-file=<key.json>`
