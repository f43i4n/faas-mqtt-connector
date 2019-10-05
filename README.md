MQTT Connector for OpenFaaS
===========================

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


