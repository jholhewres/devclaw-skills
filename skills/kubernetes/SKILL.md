---
name: kubernetes
description: "Gerenciar clusters Kubernetes via kubectl"
metadata: {"openclaw":{"always":false,"emoji":"☸️"}}
---
# Kubernetes

Gerenciamento de clusters Kubernetes via kubectl.

## Recursos Básicos

```bash
# Pods
kubectl get pods -A                          # todos namespaces
kubectl get pods -n <ns> -o wide             # com IPs e nodes
kubectl describe pod <pod> -n <ns>
kubectl logs <pod> -n <ns> --tail=100
kubectl logs <pod> -n <ns> -c <container>    # multi-container
kubectl exec -it <pod> -n <ns> -- sh

# Deployments
kubectl get deployments -n <ns>
kubectl describe deployment <name> -n <ns>
kubectl scale deployment <name> --replicas=3 -n <ns>
kubectl rollout status deployment <name> -n <ns>
kubectl rollout restart deployment <name> -n <ns>
kubectl rollout undo deployment <name> -n <ns>

# Services
kubectl get svc -n <ns>
kubectl describe svc <name> -n <ns>
kubectl port-forward svc/<name> 8080:80 -n <ns>

# ConfigMaps e Secrets
kubectl get configmap -n <ns>
kubectl get secret -n <ns>
kubectl get secret <name> -n <ns> -o jsonpath='{.data}'
```

## Diagnóstico

```bash
# Eventos (útil para debug)
kubectl get events -n <ns> --sort-by='.lastTimestamp'

# Top (métricas)
kubectl top pods -n <ns>
kubectl top nodes

# Nodes
kubectl get nodes -o wide
kubectl describe node <name>
kubectl cordon <node>    # marcar unschedulable
kubectl drain <node> --ignore-daemonsets --delete-emptydir-data
```

## Apply e Delete

```bash
# Aplicar manifesto
kubectl apply -f <file.yaml>
kubectl apply -f <directory>/

# Dry run
kubectl apply -f <file.yaml> --dry-run=client

# Delete
kubectl delete -f <file.yaml>
kubectl delete pod <pod> -n <ns>
```

## Contextos

```bash
# Listar contextos
kubectl config get-contexts

# Trocar contexto
kubectl config use-context <name>

# Ver contexto atual
kubectl config current-context
```

## Tips

- Use `-o yaml` ou `-o json` para output detalhado
- Use `--watch` para monitorar mudanças em tempo real
- Use `-l app=myapp` para filtrar por labels
- Sempre especifique `-n <namespace>` para evitar surpresas
