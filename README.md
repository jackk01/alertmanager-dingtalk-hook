## Alertmanager Dingtalk Webhook

Webhook service support send Prometheus alert messages to Dingtalk.

## Install
```
kubectl apply -f deploy.yaml
```

### Build image

```
docker build -t <NAME>/dingtalk-hook:<TAG> .
docker image prune -f
```

