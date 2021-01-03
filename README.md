# 📈 NETOR

## 📖 Overwiew

NETOR is the Inter**NE**t Moni**TOR**. It shows you the current speed of your local network.

## 🚀 Getting Started

### ⚙️ Prerequisites and installation

> **_NOTE:_** Make shure to have Docker and Docker-Compose installed on you Machine.

Pull the Image from the official Docker Regristry.

```bash
$ docker pull josuablejeru/netor:latest
```

### Usage | standalone

Run and bind the port `:8080`.

```bash
$ docker run --rm -p 8080:8080 josuablejeru/netor:latest
```

The metrix will be available on the default Prometheus path `/metrics`.

### Usage | with Grafana and Prometheus

Run Netor with prometheus and Grafana localy with docker compose in the Repo root.

```bash
$ docker-compose up
```

Prometheus and Grafana will be available on port `:9090` and `:3000`.

## 📝 License

Distributed under the MIT License. See LICENSE for more information.

## Acknowledgements

- [speedtest-go](https://github.com/showwin/speedtest-go)
