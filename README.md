# MintCocoa Pipeline Demo

`pipeline-demo` is a small release-probe service used to verify the MintCocoa
staging-first GitOps promotion flow.

- `GET /healthz` returns `ok`
- `GET /` returns the service name, image version, Kubernetes environment,
  hostname, and current UTC time

The application repository builds and pushes the container image. The
`mintcocoa-ops` repository declares which image tag runs in staging and
production.
