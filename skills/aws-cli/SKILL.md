---
name: aws-cli
description: "AWS CLI para S3, EC2, Lambda, CloudWatch, RDS e ECS"
metadata: {"openclaw":{"always":false,"emoji":"☁️"}}
---
# AWS CLI

Interface com serviços AWS via aws cli.

## S3

```bash
# Listar buckets
aws s3 ls

# Listar objetos
aws s3 ls s3://<bucket>/<prefix>/

# Upload/Download
aws s3 cp <file> s3://<bucket>/<key>
aws s3 cp s3://<bucket>/<key> <file>

# Sync
aws s3 sync <dir> s3://<bucket>/<prefix>/
aws s3 sync s3://<bucket>/<prefix>/ <dir>

# Remover
aws s3 rm s3://<bucket>/<key>
aws s3 rm s3://<bucket>/<prefix>/ --recursive
```

## EC2

```bash
# Listar instâncias
aws ec2 describe-instances --query 'Reservations[].Instances[].[InstanceId,State.Name,InstanceType,PublicIpAddress,Tags[?Key==`Name`].Value|[0]]' --output table

# Start/Stop
aws ec2 start-instances --instance-ids <id>
aws ec2 stop-instances --instance-ids <id>

# Security Groups
aws ec2 describe-security-groups --group-ids <sg-id>
```

## Lambda

```bash
# Listar funções
aws lambda list-functions --query 'Functions[].[FunctionName,Runtime,LastModified]' --output table

# Invocar
aws lambda invoke --function-name <name> --payload '{"key":"value"}' output.json

# Logs
aws logs filter-log-events --log-group-name /aws/lambda/<name> --limit 20
```

## CloudWatch

```bash
# Listar log groups
aws logs describe-log-groups --query 'logGroups[].logGroupName'

# Buscar logs
aws logs filter-log-events --log-group-name <group> --filter-pattern "ERROR" --limit 20

# Métricas
aws cloudwatch get-metric-statistics --namespace AWS/EC2 --metric-name CPUUtilization --dimensions Name=InstanceId,Value=<id> --start-time <iso> --end-time <iso> --period 300 --statistics Average
```

## RDS

```bash
# Listar instâncias
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

- Use `--output table` para output legível
- Use `--query` (JMESPath) para filtrar campos
- Use `--profile <name>` para múltiplas contas
- Use `--region <region>` quando necessário
- Configure com `aws configure` ou env vars `AWS_ACCESS_KEY_ID` + `AWS_SECRET_ACCESS_KEY`
