# Contributing to provider-ovh

Thanks for your interest in contributing! This guide gives a quick, provider-focused overview. For the full build system contribution details (DCO, workflow, etc.), also read `build/CONTRIBUTING.md`.

## Table of Contents
1. Code of Conduct
2. Getting Set Up
3. Development Workflow
4. Testing & Linting
5. Adding / Updating Resources
6. Release Preparation
7. Marketplace Extension Assets
8. Support & Questions

---
## 1. Code of Conduct
We follow the project-wide Code of Conduct in `CODE_OF_CONDUCT.md`. Please adhere to it in all interactions.

## 2. Getting Set Up
Prerequisites:
- Go (version defined in CI/env `GO_VERSION` – currently 1.24).
- Docker (for image builds) and optional Kubernetes cluster for live testing.
- `up` CLI (>= v0.39.0) for provider install and marketplace operations.

Install dependencies:
```bash
make vendor
```
Generate code (Upjet):
```bash
go run cmd/generator/main.go "$PWD"
```
Run locally (controller only):
```bash
make run
```
Build everything:
```bash
make all
```

## 3. Development Workflow
1. Branch from `main` (or an appropriate `release-` branch):
   ```bash
   git checkout -b feat/my-change
   ```
2. Make focused commits with clear messages; sign off with `git commit -s`.
3. Keep PRs small and focused; large changes benefit from an issue first.
4. Open a Pull Request; ensure CI checks pass.

## 4. Testing & Linting
Run unit tests:
```bash
make test
```
Run lint:
```bash
make lint
```
(Uses golangci-lint; configuration inherited from build tooling.)

## 5. Adding / Updating Resources
- Resource definitions come from the Upjet generator and provider schema.
- Adjust external name logic or references in `config/` subdirectories (e.g. `config/cluster/kube/`).
- Re-run the generator after changes: `go run cmd/generator/main.go "$PWD"`.
- Verify API changes: CI step will flag breaking CRD schema changes.

## 6. Release Preparation
Follow this checklist when publishing a new version:
1. Ensure tests & lint pass locally.
2. Update `extensions/release-notes/release_notes.md` moving items from `Unreleased` to a new section:
   ```
   ## vX.Y.Z - YYYY-MM-DD
   ```
3. (Optional) Refresh icon or marketplace readme under `extensions/`.
4. Commit & sign off: `git commit -s -m "release: prepare vX.Y.Z"`.
5. Create annotated tag: `git tag -a vX.Y.Z -m "vX.Y.Z"` and push: `git push origin vX.Y.Z`.
6. CI `publish-artifacts` job builds & publishes and appends extension assets.
7. Verify Marketplace listing at: https://marketplace.upbound.io/providers/edixos/provider-ovh
8. Start new `## Unreleased` section for ongoing work.
(See detailed flow in `build/CONTRIBUTING.md`).

## 7. Marketplace Extension Assets
Structure:
```
extensions/
  icons/icon.svg
  readme/readme.md
  release-notes/release_notes.md
  sboms/sbom.json
```
CI regenerates SBOM and runs:
```bash
up alpha xpkg append --extensions-root=./extensions xpkg.upbound.io/edixos/provider-ovh:<version>
```
Manual recovery if needed:
```bash
docker login xpkg.upbound.io
up alpha xpkg append --extensions-root=./extensions xpkg.upbound.io/edixos/provider-ovh:vX.Y.Z
```

## 8. Support & Questions
- Issues: https://github.com/edixos/provider-ovh/issues
- Security, breaking changes, or roadmap questions can start as an issue labeled accordingly.

## License & DCO
- Apache 2.0 license (`LICENSE`).
- All commits must be signed off (`Signed-off-by:`) to satisfy the DCO.

Happy hacking! 🚀
