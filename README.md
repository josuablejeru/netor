# 📈 NETOR

## 📖 Overwiew

NETOR is the Inter**NE**t Moni**TOR**. It shows you the current speed of your local network.

## 🚀 Getting Started

### ⚙️ Prerequisites and installation

> **_NOTE:_** Make shure to have Docker and Docker-Compose installed on you Machine.

Pull the Image from the official Docker Registry.

```bash
$ docker pull josuablejeru/netor:latest
```

### Usage | standalone

Run and bind the port `:8080`.

```bash
$ docker run --rm -p 8080:8080 josuablejeru/netor:latest
```

The metric will be available on the default Prometheus path `/metrics`.

### Usage | with Grafana and Prometheus

Run Netor with prometheus and Grafana locally with docker-compose .

```bash
$ docker-compose up
```

Prometheus and Grafana will be available on port `:9090` and `:3000`.

## ✌️ Get in touch with me

<a href="https://into-the-code.com" target="_blank"><img alt="Personal Website" src="https://img.shields.io/badge/Personal%20Website-%2312100E.svg?&style=for-the-badge&logoColor=white" /></a>
<a href="https://twitter.com/josuablejeru" target="_blank"><img alt="Twitter" src="https://img.shields.io/badge/twitter-%231DA1F2.svg?&style=for-the-badge&logo=twitter&logoColor=white" /></a>
<a href="https://www.linkedin.com/in/josua-blejeru-a2871a164" target="_blank"><img alt="LinkedIn" src="https://img.shields.io/badge/linkedin-%230077B5.svg?&style=for-the-badge&logo=linkedin&logoColor=white" /></a>

## 📝 License

Distributed under the MIT License. See LICENSE for more information.

## Acknowledgements

- [speedtest-go](https://github.com/showwin/speedtest-go)
