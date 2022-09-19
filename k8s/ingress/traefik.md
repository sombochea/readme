# Traefik Controller

```shell
helm upgrade --install traefik traefik/traefik --create-namespace -n traefik \
--set "ports.websecure.tls.enabled=true" \
--set "providers.kubernetesIngress.publishedService.enabled=true"
```
