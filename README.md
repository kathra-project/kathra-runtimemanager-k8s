# Kathra Runtime Manager K8S

This service implements Kathra Runtime Manager 1.1.0.
Using helm-client and kubctl, it can do :

* Manage runtime environment (k8s namespace)
* Deploy application (using helm chart)
* Get information application (loadbalancer, services, containers, logs)


## Configuration

| Env var                         | Description                          | Default                                   |
| --------------------------------- | ------------------------------------ | ----------------------------------------- |
| `PORT`            | Port to expose           | `random`                 |
| `KATHRA_REPO_NAME`            | Chart repository name for Helm           | `kathra-local`                 |
| `KATHRA_REPO_URL`             | Chart repository URL                     |                                |
| `KATHRA_REPO_CREDENTIAL_ID`   | Chart repository Credential ID           |                                |
| `KATHRA_REPO_SECRET`          | Chart repository Credential secret       |                                |
| `HELM_UPDATE_INTERVAL`            | Cron settings for helm update            | `* * * * *`                    |


## How to run

```
go run cmd/kathra-runtime-environment-manager-server/main.go
```
## Open Api Spec

https://gitlab.com/kathra/kathra/kathra-services/kathra-runtimemanager/kathra-runtimemanager-api/raw/master/swagger.yaml