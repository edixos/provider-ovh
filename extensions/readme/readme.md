# Community OVHcloud Provider

`provider-ovh` is a community-maintained Crossplane infrastructure provider for [OVHcloud](https://www.ovhcloud.com/), built with [Upjet](https://github.com/crossplane/upjet).

> **⚠️ Community Effort**
> This provider is **not an official OVHcloud offering**. It’s maintained by volunteers and contributors in the open-source community. OVHcloud has not (yet) published an official Crossplane provider; this project fills that gap so users can manage OVHcloud resources declaratively. Features, coverage, and support are best‑effort—please report issues and consider contributing improvements.

## Features
- Implements a growing subset of OVH Cloud resources as Crossplane Managed Resources.
- SafeStart capability enabled.
- Rich API reference: https://doc.crds.dev/github.com/edixos/provider-ovh

## Installation
```bash
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-ovh
spec:
  package: xpkg.upbound.io/edixos/provider-ovh:latest
EOF
```
Or using the CLI:
```bash
up ctp provider install edixos/provider-ovh
```

## Configuration
Create a `ProviderConfig` with your OVH credentials (application key, application secret, consumer key, and endpoint). See examples in our repository under [`examples/cluster/providerconfig/`](https://github.com/edixos/provider-ovh/tree/main/examples/cluster/providerconfig) for Cluster-scopped resource, and [`examples/namespaced/providerconfig/`](https://github.com/edixos/provider-ovh/tree/main/examples/namespaced/providerconfig) for Namespace-scopped resource.

## Support & Feedback
File issues and feature requests: https://github.com/edixos/provider-ovh/issues

When opening an issue, include:
- Provider version (`kubectl get providers.pkg.crossplane.io provider-ovh -o jsonpath='{.status.atProvider.version}'` if installed)
- Relevant CR YAML (redact secrets)
- Logs from the provider pod

Community discussions and roadmap proposals are welcome via issues labeled `enhancement` or `discussion`.

## Contributing
We welcome PRs! See root `CONTRIBUTING.md` and `build/CONTRIBUTING.md` for detailed guidance (DCO, release process, marketplace assets).

## License
Apache-2.0
