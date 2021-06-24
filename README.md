## Alertmanager Dingtalk Webhook

Webhook service support send Prometheus alert messages to Dingtalk.



### Build image

```
docker build -t <NAME>/dingtalk-hook:<TAG> .
docker image prune -f
```



### deploy.yaml

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: dingtalk-hook
  namespace: monitoring
spec:
  selector:
    matchLabels:
      app: dingtalk-hook
  template:
    metadata:
      labels:
        app: dingtalk-hook
    spec:
      containers:
        - name: dingtalk-hook
          image: <NAME>/dingtalk-hook:<TAG> # use your image
          imagePullPolicy: IfNotPresent
          ports:
            - containerPort: 5000
              name: http
          env:
            - name: TZ
              value: Asia/Shanghai
            - name: PROME_URL
              value: # your prometheus url 
            - name: LOG_LEVEL
              value: debug
            - name: ROBOT_TOKEN
              valueFrom:
                secretKeyRef:
                  name: dingtalk-secret
                  key: token
            - name: ROBOT_SECRET
              valueFrom:
                secretKeyRef:
                  name: dingtalk-secret
                  key: secret
          resources:
            requests:
              cpu: 50m
              memory: 100Mi
            limits:
              cpu: 50m
              memory: 100Mi

---
apiVersion: v1
kind: Service
metadata:
  name: dingtalk-hook
  namespace: monitoring
spec:
  selector:
    app: dingtalk-hook
  ports:
    - name: hook
      port: 5000
      targetPort: http
  type: NodePort
```

