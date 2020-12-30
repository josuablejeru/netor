# 📈 NETOR

## 📖 Overwiew

NETOR is the Inter**NE**t Moni**TOR**. It shows you the current UP and DOWN speed of your Internet connection.

## 🚀 Run Exporter

The metrix will be available on the default prometheus Path `/metrics`

```bash
docker run --rm -p 8080:8080 josuablejeru/netor:latest
```

If you want to test it localy use docker-compose. This will provide you Grafana on localhost with the port `:3000`.

Prometheus (on `:9090`) and netor (on `:8080`) will be allready setup

```bash
docker-compose up
```
