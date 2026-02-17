---
name: terraform
description: "Infrastructure as Code with Terraform"
metadata: {"openclaw":{"always":false,"emoji":"🏗️"}}
---
# Terraform

Infrastructure as Code management.

## Basic Workflow

```bash
# Initialize (download providers)
terraform init

# Plan (preview changes)
terraform plan
terraform plan -out=plan.tfplan   # save plan

# Apply
terraform apply
terraform apply plan.tfplan       # apply saved plan
terraform apply -auto-approve     # no confirmation

# Destroy
terraform destroy
terraform destroy -target=aws_instance.web   # specific resource
```

## State

```bash
# List resources in state
terraform state list

# Resource details
terraform state show <resource>

# Move resource (refactoring)
terraform state mv <old> <new>

# Remove from state (without destroying)
terraform state rm <resource>

# Pull/Push remote state
terraform state pull > state.json
terraform state push state.json
```

## Output & Variables

```bash
# View outputs
terraform output
terraform output -json
terraform output <name>

# Validate configuration
terraform validate

# Format code
terraform fmt
terraform fmt -recursive
```

## Workspaces

```bash
terraform workspace list
terraform workspace new <name>
terraform workspace select <name>
terraform workspace delete <name>
```

## Import

```bash
# Import existing resource into state
terraform import <resource_type>.<name> <id>
```

## Tips

- Always run `terraform plan` before `apply`
- Use `-target` to apply specific resources
- Use `terraform fmt` to keep code standardized
- State is sensitive — never commit to public repos
- Use remote backend (S3, GCS) for team work
