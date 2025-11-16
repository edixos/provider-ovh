# Provider OVH Marketplace Extension Readme

`provider-ovh` is the Crossplane infrastructure provider for [OVH Cloud](https://ovh.com/), built with [Upjet](https://github.com/crossplane/upjet).

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

## Contributing
We welcome PRs! See `build/CONTRIBUTING.md` for guidelines.

## License
Apache-2.0
