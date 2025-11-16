# Community OVH Cloud Provider

`provider-ovh` is a community-maintained Crossplane infrastructure provider for [OVH Cloud](https://ovh.com/), built with [Upjet](https://github.com/crossplane/upjet).

> **⚠️ Community Effort**
> This provider is **not an official OVH offering**. It’s maintained by volunteers and contributors in the open-source community. OVH has not (yet) published an official Crossplane provider; this project fills that gap so users can manage OVH resources declaratively. Features, coverage, and support are best‑effort—please report issues and consider contributing improvements.

## Features
- Implements a growing subset of OVH Cloud resources as Crossplane Managed Resources.
- SafeStart capability enabled.
- Rich API reference: https://doc.crds.dev/github.com/edixos/provider-ovh

## Installation
```
up ctp provider install edixos/provider-ovh
```
Or declaratively:
```
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-ovh
spec:
  package: edixos/provider-ovh
```

## Configuration
Create a `ProviderConfig` with your OVH credentials (application key, application secret, consumer key, and endpoint). See repository examples under `examples/providerconfig/`.

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
