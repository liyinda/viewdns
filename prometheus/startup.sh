#!/bin/bash

docker run -d --name=grafana -e "GF_SECURITY_ADMIN_PASSWORD=123456" -p 3000:3000 grafana/grafana

docker run -d -p 9090:9090 -v /home/gopath/src/github.com/liyinda/viewdns/prometheus/prometheus.yaml:/etc/prometheus/prometheus.yml  prom/prometheus
