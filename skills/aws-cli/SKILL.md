---
name: aws-cli
description: "AWS CLI for S3, EC2, Lambda, CloudWatch, RDS, and ECS"
metadata: {"openclaw":{"always":false,"emoji":"☁️"}}
---
# AWS CLI

Manage AWS services via the aws CLI.

## Setup

1. **Check if installed:**
   ```bash
   command -v aws && aws --version
   ```

2. **Install:**
   ```bash
   # macOS
   brew install awscli

   # Official installer (Linux/macOS)
   curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
   unzip awscliv2.zip && sudo ./aws/install && rm -rf aws awscliv2.zip
   ```

3. **Credentials:** Use the vault for secrets. Stored keys are auto-injected as env vars (UPPERCASE).
   ```bash
   # Save to vault (key names lowercase)
   vault_save aws_access_key_id "AKIA..."
   vault_save aws_secret_access_key "secret..."
   vault_save aws_default_region "us-east-1"

   # Or interactive: aws configure
   ```

## S3

```bash
# List buckets
aws s3 ls

# List objects
aws s3 ls s3://<bucket>/<prefix>/

# Upload/Download
aws s3 cp <file> s3://<bucket>/<key>
aws s3 cp s3://<bucket>/<key> <file>

# Sync
aws s3 sync <dir> s3://<bucket>/<prefix>/
aws s3 sync s3://<bucket>/<prefix>/ <dir>

# Remove
aws s3 rm s3://<bucket>/<key>
aws s3 rm s3://<bucket>/<prefix>/ --recursive
```

## EC2

```bash
# List instances
aws ec2 describe-instances --query 'Reservations[].Instances[].[InstanceId,State.Name,InstanceType,PublicIpAddress,Tags[?Key==`Name`].Value|[0]]' --output table

# Start/Stop
aws ec2 start-instances --instance-ids <id>
aws ec2 stop-instances --instance-ids <id>

# Security Groups
aws ec2 describe-security-groups --group-ids <sg-id>
```

## Lambda

```bash
# List functions
aws lambda list-functions --query 'Functions[].[FunctionName,Runtime,LastModified]' --output table

# Invoke
aws lambda invoke --function-name <name> --payload '{"key":"value"}' output.json

# Logs
aws logs filter-log-events --log-group-name /aws/lambda/<name> --limit 20
```

## CloudWatch

```bash
# List log groups
aws logs describe-log-groups --query 'logGroups[].logGroupName'

# Search logs
aws logs filter-log-events --log-group-name <group> --filter-pattern "ERROR" --limit 20

# Metrics
aws cloudwatch get-metric-statistics --namespace AWS/EC2 --metric-name CPUUtilization --dimensions Name=InstanceId,Value=<id> --start-time <iso> --end-time <iso> --period 300 --statistics Average
```

## RDS

```bash
# List instances
aws rds describe-db-instances --query 'DBInstances[].[DBInstanceIdentifier,DBInstanceStatus,Engine,Endpoint.Address]' --output table
```

## ECS

```bash
# Clusters
aws ecs list-clusters

# Services
aws ecs list-services --cluster <cluster>
aws ecs describe-services --cluster <cluster> --services <service>

# Tasks
aws ecs list-tasks --cluster <cluster> --service-name <service>
aws ecs describe-tasks --cluster <cluster> --tasks <task-arn>
```

## Tips

- Use `--output table` for readable output
- Use `--query` (JMESPath) to filter fields
- Use `--profile <name>` for multiple accounts
- Use `--region <region>` when needed
- Configure with `aws configure` or env vars `AWS_ACCESS_KEY_ID` + `AWS_SECRET_ACCESS_KEY`
