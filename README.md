# 制作helm包

```bash
helm create fms
helm lint --strict fms
helm package fms
helm install fms fms-{version}.tgz
```

# 更新helm包
```bash
helm upgrade --install fms /path/to/fms
```

# 本地测试
```bash
kubectl port-forward {pod-name} 8080:8080
```
