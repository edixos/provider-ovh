# Provider OVH

`provider-ovh` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
OVHcloud API.

## Prerequisites

* Install [Crossplane in your cluster](https://docs.crossplane.io/v2.1/get-started/install/)

## Getting Started

Install the provider by using the following command:
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

If you want to specify a version of the provider, add an image tag
to the [latest release](https://marketplace.upbound.io/providers/edixos/provider-ovh).
Example:
`xpkg.upbound.io/edixos/provider-ovh:v2.10.1`

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/edixos/provider-ovh).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/edixos/provider-ovh/issues).

## Marketplace Extension Assets

This repository includes optional Marketplace extension assets under the `extensions/` directory:

```
extensions/
  icons/icon.svg          # Provider icon
  readme/readme.md        # Marketplace readme (concise overview)
  release-notes/release_notes.md  # Human-authored release notes per version
  sboms/sbom.json         # CycloneDX SBOM (regenerated in CI)
```

During CI publish, we run:

```
up alpha xpkg append --extensions-root=./extensions xpkg.upbound.io/edixos/provider-ovh:<version>
```

Where `<version>` is read from `_output/version` produced by the build. You can test locally after logging into the Upbound registry:

```bash
curl -sL https://cli.upbound.io | sh
VERSION=$(cat _output/version)
up alpha xpkg append --extensions-root=./extensions xpkg.upbound.io/edixos/provider-ovh:${VERSION}
```

To update for a new release:
1. Edit `extensions/release-notes/release_notes.md` adding a section for the new `vX.Y.Z`.
2. Optionally refine `extensions/readme/readme.md` or update the icon.
3. Commit changes before tagging the release so CI appends correct assets.

SBOM generation currently uses [Syft](https://github.com/anchore/syft) in CI to regenerate `extensions/sboms/sbom.json`. You can reproduce locally:

```bash
curl -sSfL https://raw.githubusercontent.com/anchore/syft/main/install.sh | sh -s -- -b /usr/local/bin
syft dir:. -o cyclonedx-json > extensions/sboms/sbom.json
```

This is an alpha feature (requires `up` CLI v0.39.0+).



