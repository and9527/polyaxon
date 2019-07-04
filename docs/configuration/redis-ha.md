---
title: "PostgreSQL HA"
sub_link: "postgresql-ha"
meta_title: "High availability of postgresql database in Polyaxon - Configuration"
meta_description: "Keeping database records of your users, projects, experiments, and jobs is very important. Polyaxon offers a couple of ways to set a high available database."
tags:
    - configuration
    - polyaxon
    - postgresql
    - scaling
    - high-availability
    - kubernetes
    - docker-compose
sidebar: "configuration"
---

Polyaxon ships with a default redis based on the stable [Helm chart](https://github.com/helm/charts/tree/master/stable/redis).

You can check the chart values to extend it's configuration.

## External Redis

If you prefer to have Redis managed by you or hosted outside of Kubernetes, 
you need to disable the in-cluster redis, and provide the information needed to establish a connection to the external one, e.g.:


```yaml
redis:
  enabled: false

externalServices:
  redis:
    password: polyaxon
    host: 35.262.163.88
```

## External Redis with password

If your redis instance requires a password, you need to provide it as well:


```yaml
redis:
  enabled: false

externalServices:
  redis:
    usePassword: true
    password: polyaxon
    host: 35.262.163.88
    port: 1234
```


### Memorystore for Redis

You can use [Cloud MemoryStore for Redis](https://cloud.google.com/memorystore/) if you are running Polyaxon on GKE, 
please follow this [integration guide](/integrations/redis/).