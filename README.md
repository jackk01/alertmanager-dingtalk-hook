## Alertmanager Dingtalk Webhook

Webhook service support send Prometheus alert messages to Dingtalk.

## Secret
```
kubectl create secret generic dingtalk-secret \
--from-literal=token=<your_token> \
--from-literal=secret=<your_secret> -n monitoring
```

## Install
```
kubectl apply -f deploy.yaml
```

### Build image

```
docker build -t <NAME>/dingtalk-hook:<TAG> .
docker image prune -f
```

