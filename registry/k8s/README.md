## Note

### Defines the key name of specific fields

GoMicro needs to cooperate with the following fields to run properly on kubernetes:

- `gomicro-service-id`: define the ID of the service
- `gomicro-service-app`: define the name of the service
- `gomicro-service-version`: define the version of the service
- `gomicro-service-metadata`: define the metadata of the service
- `gomicro-service-protocols`: define the protocols of the service

### Example Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: shop_api
  labels:
    app: shop_api
spec:
  replicas: 2
  selector:
    matchLabels:
      app: shop_api
  template:
    metadata:
      labels:
        app: shop_api
        gomicro-service-id: "56991810-c77f-4a95-8190-393efa9c1a61"
        gomicro-service-app: "shop_api"
        gomicro-service-version: "v1.0.0"
      annotations:
        gomicro-service-protocols: |
          {"8080": "http"}
        gomicro-service-metadata: |
          {"region": "sh", "zone": "sh001", "cluster": "pd"}
    spec:
      containers:
        - name: shop_api
          image: shop_api:1.0.0
          ports:
            - containerPort: 8080
```

