# GitHub Exporter

An exporter for [Prometheus](https://prometheus.io/) that collects metrics from [GitHub Enterprise](https://github.service.anz).


## Build




## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). This project requires Go >= v1.11.

```bash
git clone https://github.service.anz/ecp/github-enterprise-exporter.git ghe-exporter
cd ghe-exporter

make build

./bin/github-enterprise-exporter -h
```
