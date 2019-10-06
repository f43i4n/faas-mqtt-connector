MQTT Connector for OpenFaaS
===========================

[![Go Report Card](https://goreportcard.com/badge/github.com/f43i4n/faas-mqtt-connector)](https://goreportcard.com/badge/github.com/f43i4n/faas-mqtt-connector)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![OpenFaaS](https://img.shields.io/badge/openfaas-serverless-blue.svg)](https://www.openfaas.com)

Inspired by the [Kafka Connector](https://github.com/openfaas-incubator/kafka-connector).
Tested on Kubernetes cluster running on Raspberry Pi 3s.

Deployment
----------

To deploy to Kubernetes cluster change the MQTT broker URL in connector.yaml.
An then run

```
kubectl apply -f deployment/kubernetes/connector.yaml
```

Backlog of open issues
----------------------

- [ ] QoS is currently fixed to 2
- [ ] No possibility to configure MQTT credentials
- [ ] Error handling and retry mechanisms.
- [ ] Content type of request is not set.


