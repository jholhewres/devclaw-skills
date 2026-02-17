---
name: terraform
description: "Infraestrutura como código com Terraform"
metadata: {"openclaw":{"always":false,"emoji":"🏗️"}}
---
# Terraform

Gerenciamento de infraestrutura como código.

## Workflow Básico

```bash
# Inicializar (baixar providers)
terraform init

# Planejar (preview de mudanças)
terraform plan
terraform plan -out=plan.tfplan   # salvar plano

# Aplicar
terraform apply
terraform apply plan.tfplan       # aplicar plano salvo
terraform apply -auto-approve     # sem confirmação

# Destruir
terraform destroy
terraform destroy -target=aws_instance.web   # recurso específico
```

## State

```bash
# Listar recursos no state
terraform state list

# Detalhes de um recurso
terraform state show <resource>

# Mover recurso (refactoring)
terraform state mv <old> <new>

# Remover do state (sem destruir)
terraform state rm <resource>

# Pull/Push remote state
terraform state pull > state.json
terraform state push state.json
```

## Output e Variáveis

```bash
# Ver outputs
terraform output
terraform output -json
terraform output <name>

# Validar configuração
terraform validate

# Formatar código
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
# Importar recurso existente para o state
terraform import <resource_type>.<name> <id>
```

## Tips

- Sempre rode `terraform plan` antes de `apply`
- Use `-target` para aplicar recursos específicos
- Use `terraform fmt` para manter código padronizado
- State é sensível — nunca commite em repos públicos
- Use backend remoto (S3, GCS) para trabalho em equipe
