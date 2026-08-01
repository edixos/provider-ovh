---
name: upgrading-upstream-provider
description: Use when provider-ovh must track a newer upstream ovh/terraform-provider-ovh release — a requested version bump, a new upstream tag to pick up, an `ovh_*` resource that exists upstream but has no CRD here, or a question about upjet CRD regeneration, crddiff, or schema-version-diff tied to the pinned provider version.
---

# Upgrading the Upstream Terraform Provider

## Overview

**Core principle:** Read the upstream diff → bump every pin → regenerate → let the CI gates classify the breakage → tag to mirror upstream.

**Announce at start:** "I'm using the upgrading-upstream-provider skill to bump to upstream v\<X.Y.Z\>."

Nearly every file this touches is generated. The judgement is in three places: which resources to wire in, how to classify the CRD breakage, and what the PR body tells users to do about it.

## Quick Reference

Every artifact carrying the upstream version. All must agree.

| Artifact | Location | Exact form |
|---|---|---|
| Schema + docs pin | `Makefile`, `export TERRAFORM_PROVIDER_VERSION ?=` | `<X.Y.Z>` — **no** leading `v` |
| Native provider binary | `Makefile`, `export TERRAFORM_NATIVE_PROVIDER_BINARY ?=` | `terraform-provider-ovh_v<X.Y.Z>` — **leading `v`** |
| Go dependency | `go.mod`, `github.com/ovh/terraform-provider-ovh/v2` | `v<X.Y.Z>` — leading `v` |
| Branch | git | `upgrade/terraform-provider-ovh-v<X.Y.Z>` |
| Release tag | git | `v<X.Y.Z>` — mirrors upstream exactly |

---

## Step 1 — Resolve the target version

```bash
gh api repos/ovh/terraform-provider-ovh/releases --jq '.[0].tag_name'
grep -n 'TERRAFORM_PROVIDER_VERSION\|TERRAFORM_NATIVE_PROVIDER_BINARY' Makefile
```

The second command gives the current pin. Confirm the target if none was named.

---

## Step 2 — Read the upstream changelog BEFORE regenerating

```bash
gh api repos/ovh/terraform-provider-ovh/releases/tags/v<NEW> --jq .body
gh api repos/ovh/terraform-provider-ovh/compare/v<OLD>...v<NEW> --jq '.files[].filename'
```

Classify what you find into four buckets:

| Bucket | Consequence here |
|---|---|
| New `ovh_*` resources | Candidates for Step 4 — no CRD until wired in |
| Removed / renamed resources | CRD disappears; user migration required |
| Removed / renamed fields | Trips `crddiff`; the breaking class |
| State `SchemaVersion` bumps | Trips `schema-version-diff` |

Do this first. The regenerated diff is thousands of lines of `zz_*` churn, unreadable unless you already know what to expect in it.

---

## Step 3 — Bump the three pins

Edit both `Makefile` variables, then:

```bash
go get github.com/ovh/terraform-provider-ovh/v2@v<X.Y.Z>
go mod tidy
go mod vendor
```

**`go mod vendor` is not optional.** This repo vendors, and `vendor/` is gitignored. Skipping it makes every later `go build` fail with `inconsistent vendoring` — which reads like a code error and is not one.

Then confirm the release actually ships the binary the provider image will download, for **every** platform in `PLATFORMS`:

```bash
for a in amd64 arm64; do
  curl -o /dev/null -sIL -w "$a %{http_code}\n" \
    "https://github.com/ovh/terraform-provider-ovh/releases/download/v<X.Y.Z>/terraform-provider-ovh_<X.Y.Z>_linux_$a.zip"
done
```

A 404 here surfaces much later and far less legibly — as a `failed to load cache key: invalid response status 404` in the `local-deploy` job, after lint and tests have already gone green.

---

## Step 4 — Wire new resources

Skip only if Step 2 found no new resources worth exposing.

A new upstream resource produces **no CRD, no types, no controller** until its Terraform name appears in a map in `config/external_name.go` — those maps double as the include list.

1. **Pick the map** by how upstream registers it: `ResourcesMap` in `ovh/provider.go` → `TerraformPluginSDKExternalNameConfigs`; `Resources()` in `ovh/provider_new.go` → `TerraformPluginFrameworkExternalNameConfigs`.
2. **Choose the external name.** OVH IDs are usually composite (`serviceName/id`), needing a custom `config.ExternalName` with a `GetIDFn` plus `DisableNameInitializer: true`. Reuse a helper from the top of `config/external_name.go` when the shape matches.
3. **Add the `Configure` entry to BOTH trees** — `config/cluster/<group>/config.go` and `config/namespaced/<group>/config.go`, byte-identical.
4. **If the group is new**, register it in both lists in `config/provider.go` (`GetProvider` and `GetProviderNamespaced`).

---

## Step 5 — Regenerate and verify

```bash
rm -rf .work/terraform .work/ovh   # MUST come first — see below
make generate                      # needs network: terraform download, schema fetch, docs sparse-clone
make reviewable                    # generate + lint + test
```

**Purge `.work/` or you will regenerate the old version.** `.work/terraform` holds a `.terraform.lock.hcl` pinned to the previous provider, and the `pull-docs` target is guarded by `if [ ! -d ]` so it never re-clones. Both silently produce a "successful" regeneration of the version you were trying to leave.

**Never hand-edit a `zz_*` file.** Change `config/`, regenerate.

**When the release removes a resource**, delete that group's `zz_generated.managed.go` and `zz_generated.managedlist.go` before regenerating. `apis/generate.go` deliberately preserves those two files, so a stale entry survives and `angryjet` then fails with `undefined: <Kind>`.

Commit every generated artifact — `apis/**/zz_*.go`, `internal/controller/**/zz_*.go`, `package/crds/*.yaml`, `examples-generated/`, `config/schema.json`. Note that the `check-diff` job is currently commented out in `ci.yml`, so **CI does not verify this for you** — an uncommitted regeneration will merge silently and only surface as drift on the next bump.

**Stop-gate — do not proceed to a PR while `make reviewable` is red.** Report instead:

```
make reviewable failed after regenerating for v<X.Y.Z>. Not opening a PR.
<failing target and output>
```

---

## Step 6 — The two repo-specific CI gates

```bash
make crddiff              # breaking CRD OpenAPI v3 schema changes
make schema-version-diff  # native state schema version changes
```

**Neither gate can fail your PR. Run them locally and read the output yourself.**

- `crddiff` runs in CI's `report-breaking-changes` job when `package/crds/**` changed, but the target loops and prints `Breaking change detected!` without ever exiting non-zero. A green step does **not** mean no breaking change.
- `schema-version-diff` never runs in CI at all: its step is gated on `inputs.upjet-based-provider`, and `ci.yml` has no `workflow_call` trigger, so `inputs` is always empty. Locally it is also broken two ways — it seds the base `Makefile` for `TERRAFORM_PROVIDER_VERSION[[:space:]]*:=` while the file declares `?=` (so `PREV_PROVIDER_VERSION` resolves empty), and it reads `config/generated.lst`, which does not exist in the repo.

So a removed upstream field is something **you** must catch from Step 2 and confirm in the generated CRD diff — not something CI will hand you. Document it in the PR body with a migration note for users.

---

## Step 7 — Branch, commit, PR

Branch off `main` as `upgrade/terraform-provider-ovh-v<X.Y.Z>`. Every commit must be DCO signed off:

```bash
git commit -s
```

The PR body must carry all four:

- Upstream release URL
- New and removed resource summary from Step 2
- Breaking changes with a user migration note (or an explicit "none")
- Confirmation that `make reviewable` passed locally

---

## Step 8 — Drive CI green, then merge

```bash
gh pr checks <pr>
gh run view --log-failed   # on failure
```

**Stop-gate — never merge on red CI.** Only `lint`, `unit-tests`, `local-deploy`, and `publish-artifacts` can actually fail the PR. `report-breaking-changes` is informational (Step 6) and `check-diff` is disabled — a fully green run does not mean the regeneration was committed or the schema is compatible.

---

## Step 9 — Tag and publish

**Stop-gate — do not tag until the merge commit is on `main`.**

First, `extensions/release-notes/release_notes.md` must carry a `## v<X.Y.Z> - YYYY-MM-DD` section — CI appends this file to the published package, so it is user-facing. Either include it in the upgrade PR or land a separate `release: prepare v<X.Y.Z>` commit (that is the only file the previous release commit touched — there is no version field anywhere else; the package version is purely the git tag). Leave a fresh `## Unreleased` heading above it.

The provider tag mirrors the upstream version exactly. Confirm that against history rather than assuming:

```bash
git tag --sort=-v:refname | head
```

Tag `v<X.Y.Z>` on the merge commit — push the tag, or dispatch the `Tag` workflow (`.github/workflows/tag.yaml`). CI triggers on `v*` tag pushes; its `publish-artifacts` job builds the xpkg, generates the SBOM, and pushes to `xpkg.upbound.io`. Create the GitHub Release, then confirm that job succeeded.

---

## Common Mistakes

| Mistake | Reality |
|---|---|
| "`go mod tidy` is enough" | This repo vendors. Without `go mod vendor`, every build fails with `inconsistent vendoring`. |
| "I'll just fix the field in the `zz_*` file" | Generated. The next `make generate` silently overwrites it. Change `config/`. |
| "I bumped the version in the Makefile" | There are two Makefile variables. One takes a bare `<X.Y.Z>`, the other `terraform-provider-ovh_v<X.Y.Z>`. |
| "I added the configurator to `config/cluster/`" | `config/cluster/**` and `config/namespaced/**` are byte-identical. Both, or the namespaced API surface silently lacks the resource. |
| "The resource is in upstream, so it will generate" | Absent from both include-list maps in `config/external_name.go` = no CRD, no types, no controller. |
| "crddiff printed a breaking change — I'll regenerate until it stops" | There is nothing to suppress; it never blocks. Its output *is* the finding. Document the migration in the PR body. |
| "The breaking-changes job is green, so nothing broke" | `crddiff` never exits non-zero and `schema-version-diff` never runs. Green means nothing. |
| "Generated files bloat the diff, I'll skip them" | `check-diff` is commented out in CI, so nothing catches it — it becomes someone else's mystery diff. Commit all generated output. |
| "`make generate` succeeded, so it used the new version" | Check the log line `generating provider schema for ovh/ovh <X.Y.Z>`. A stale `.work/` regenerates the old one just as successfully. |
| "The changelog can wait until after regeneration" | Then you are reading a thousand-file diff with no hypothesis. Read it first. |
| "lint and tests are green, so the bump is sound" | Neither builds the provider image. A missing release asset only fails in `local-deploy`, minutes later. |

## Notes

- No version is hardcoded here. The target comes from Step 1, the current pin from the `Makefile`.
- Targets come from the `Makefile`; if one is missing, run `make help` rather than inventing it.
- `make` failing on missing includes means the `build/` submodule is uninitialised — `make submodules`.
- Repo-wide conventions live in `AGENTS.md`; this skill covers only the upgrade path.
