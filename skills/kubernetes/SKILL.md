---
name: kubernetes
description: "Manage Kubernetes clusters via kubectl"
metadata: {"openclaw":{"always":false,"emoji":"☸️"}}
---
# Kubernetes

Manage Kubernetes clusters via kubectl.

## Setup

1. **Check if installed:**
   ```bash
   command -v kubectl && kubectl version --client
   ```

2. **Install:**
   ```bash
   # macOS
   brew install kubectl

   # Ubuntu / Debian
   curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
   chmod +x kubectl && sudo mv kubectl /usr/local/bin/

   # Or via apt (see https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/)
   sudo apt update && sudo apt install -y kubectl
   ```

## Basics

```bash
# Pods
kubectl get pods -A                          # all namespaces
kubectl get pods -n <ns> -o wide             # with IPs and nodes
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

# ConfigMaps and Secrets
kubectl get configmap -n <ns>
kubectl get secret -n <ns>
kubectl get secret <name> -n <ns> -o jsonpath='{.data}'
```

## Diagnostics

```bash
# Events (useful for debug)
kubectl get events -n <ns> --sort-by='.lastTimestamp'

# Top (metrics)
kubectl top pods -n <ns>
kubectl top nodes

# Nodes
kubectl get nodes -o wide
kubectl describe node <name>
kubectl cordon <node>    # mark unschedulable
kubectl drain <node> --ignore-daemonsets --delete-emptydir-data
```

## Apply & Delete

```bash
# Apply manifest
kubectl apply -f <file.yaml>
kubectl apply -f <directory>/

# Dry run
kubectl apply -f <file.yaml> --dry-run=client

# Delete
kubectl delete -f <file.yaml>
kubectl delete pod <pod> -n <ns>
```

## Contexts

```bash
# List contexts
kubectl config get-contexts

# Switch context
kubectl config use-context <name>

# View current context
kubectl config current-context
```

## Tips

- Use `-o yaml` or `-o json` for detailed output
- Use `--watch` to monitor changes in real time
- Use `-l app=myapp` to filter by labels
- Always specify `-n <namespace>` to avoid surprises
