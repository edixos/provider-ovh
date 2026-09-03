# Release Notes

> IMPORTANT: Update this file for each tagged release before (or as part of) CI so it accurately reflects changes delivered in that version. The CI workflow will append this file along with other extension assets to the published package.

## Unreleased

## v2.19.1 - 2026-09-03

> [!WARNING]
> Tracks the same upstream OVHcloud Terraform provider as v2.19.0. Despite the
> patch version, this release contains breaking CRD schema changes. Read the
> section below before upgrading.

### ⚠️ Changed (breaking)
- SDK-backed resources now reconcile in-process instead of shelling out to the
  Terraform CLI, and the CLI and the bundled native provider are no longer
  present in the image. Reconciliation no longer writes a workspace to disk or
  runs `terraform init`, so the provider needs neither a writable working
  directory nor registry network access at runtime.
- Integer-valued fields are now declared as `type: integer` in the CRD schemas
  rather than `type: number`, across 506 properties. `terraform providers
  schema -json` renders Terraform's `TypeInt` as `number`, which is what the
  generated schemas previously carried; none of the affected fields are ever
  fractional, so this is a correction rather than a behaviour change. The real
  break is in the generated Go types, where the corresponding fields move from
  `*float64` to `*int64`. ⚠️ **Action required** only if you import
  `apis/{cluster,namespaced}/...` directly; manifests and stored objects are
  unaffected.
- `x-kubernetes-map-type: granular` is no longer set on the string maps
  `User.openstackRc` (`cloud`, in `forProvider`, `initProvider` and
  `atProvider`) and `PrivateNetwork.regionsOpenstackIds` (`network`), in both
  the cluster-scoped and namespaced API groups — 8 paths in total. Those maps
  are now atomic for server-side apply: a client that applies the map takes
  ownership of it whole, so two appliers can no longer each own individual keys
  and a partial apply replaces the map rather than merging into it.
- `vrack`: `CloudProject` now carries a CEL validation rule making
  `spec.forProvider.projectId` required whenever `managementPolicies` includes
  `Create`, `Update` or `*`. ⚠️ **Action required:** existing objects that do
  not set `projectId` (or `spec.initProvider.projectId`) and have either policy
  will be rejected on their next update until the field is filled in. Find
  affected objects before upgrading:

  ```bash
  # cluster-scoped
  kubectl get cloudprojects.vrack.ovh.edixos.io -o json \
    | jq -r '.items[]
        | select((.spec.forProvider.projectId // .spec.initProvider.projectId) == null)
        | .metadata.name'

  # namespaced
  kubectl get cloudprojects.vrack.ovh.m.edixos.io -A -o json \
    | jq -r '.items[]
        | select((.spec.forProvider.projectId // .spec.initProvider.projectId) == null)
        | "\(.metadata.namespace)/\(.metadata.name)"'
  ```

### Fixed
- SDKv2 resources now receive a configured provider meta. upjet passes
  `terraform.Setup`'s `Meta` straight through to the resource CRUD functions on
  the SDKv2 path and never populates it itself, unlike the Framework path where
  it configures the provider from `Setup.Configuration`. Every OVH SDKv2
  resource begins with `meta.(*Config)`, so without this all SDK-backed
  resources fail their first `Observe`. Configured metas are cached per
  ProviderConfig and effective configuration, because configuring the provider
  runs `Config.loadAndValidate` and calls `GET /auth/details`; caching keeps
  that to one call rather than one per `Connect`. A rotated OAuth access token
  produces a new cache key, so a cached meta is never reused with credentials
  that have since changed.

### Security
- Remediated the vulnerabilities Trivy reported against the v2.19.0 image. Go
  moves from 1.25.6 to 1.25.13, the Alpine runtime from 3.17.1 (end of life) to
  3.24.1, and the affected Go dependencies to patched versions. A Trivy scan of
  the resulting image reports no HIGH or CRITICAL findings, and `govulncheck`
  reports the provider code as affected by none.

  `golang.org/x/crypto/openpgp` is still flagged module-wide as GO-2026-5932.
  There is no fixed version, and the affected package is neither imported nor
  called by this provider.

  The build-time Terraform pin moves from 1.8.1 to 1.15.9. It is used only to
  generate `config/schema.json` and is not shipped; regenerating with 1.15.9
  produces byte-identical output, so no CRD changes follow from it.

### Thanks
Both of the changes in this release came from the community.

- 🎉 **@VeSeWe** made their **first contribution** in
  [#63](https://github.com/edixos/provider-ovh/pull/63), the security
  remediation — Go, Alpine and dependency updates that take the image to zero
  HIGH and CRITICAL Trivy findings. Welcome, and thank you.
- **@ekarlso** contributed [#65](https://github.com/edixos/provider-ovh/pull/65),
  the move to in-process reconciliation and the removal of the Terraform CLI
  from the image, and followed up with the SDKv2 provider meta fix and its
  cache. Thanks as always.

## v2.19.0 - 2026-08-08
### Added
- `cloud`: `Instance` — new managed resource in both the cluster-scoped
  (`cloud.ovh.edixos.io`) and namespaced (`cloud.ovh.m.edixos.io`) API groups,
  wrapping the new upstream `ovh_cloud_instance` resource. It covers `flavorId`
  (resizes in place), `imageId`, `powerState` (`ACTIVE`, `SHUTOFF`, `SHELVED`),
  `networks` (public Ext-Net auto-assign, or a private `networkId` + `subnetId`),
  `securityGroupIds`, `volumeIds` and `shares`. `groupId` resolves by reference
  or selector to an `InstanceGroup`, and `sshKeyName` to an `SSHKey`.

  Note that changing `imageId` rebuilds the instance and **wipes the root disk**,
  and that `availabilityZone`, `groupId` and `sshKeyName` are immutable.
  Omitting `securityGroupIds` applies the project's `default` security group;
  an explicit empty list applies none, so the instance accepts no inbound
  traffic.
- `cloud`: `InstanceGroup` — new managed resource in both API groups, wrapping
  `ovh_cloud_instance_group`. A placement group with an `AFFINITY` or
  `ANTI_AFFINITY` `policy`. The group is immutable: changing `name`, `region` or
  `policy` replaces it. Membership is set only through an `Instance`'s `groupId`,
  never from this resource.

This is the first of the two OVHcloud instance APIs to be exposed here as a
distinct CRD: the pre-existing `ProjectInstance` (`ovh_cloud_project_instance`)
is untouched and keeps its own schema.

### Changed
- Upgrade OVH Terraform provider from 2.18.0 to 2.19.0.
- `KeyManagerSecret`: upstream now normalizes `secretType`, `algorithm` and
  `mode` to upper case to match the API, so lower-case values no longer produce
  a permanent diff. CRD change is documentation-only.
- `ProjectStorage` accepts `DEEP_ARCHIVE`, `GLACIER`, `GLACIER_IR`,
  `INTELLIGENT_TIERING` and `ONEZONE_IA` as a replication rule `storageClass`,
  and `ProjectStorageLifecycleConfiguration` accepts `DEEP_ARCHIVE` and
  `GLACIER_IR` for a noncurrent version transition. Both are validated upstream,
  not in the CRD schema, so no CRD change accompanies this.
- `CloudGateway` reports the failing task reason when a gateway ends in `ERROR`,
  instead of a generic unexpected-state message.
- `BlockVolume` keeps unset `createFrom` attributes null rather than `""`, which
  removes the inconsistent-result errors seen on create.
- Several generated Go types in the `cloud` API group were renamed by upjet to
  disambiguate them from the new `Instance`/`InstanceGroup` types — for example
  `MembersObservation` → `LoadbalancerMembersObservation` and
  `FlavorParameters` → `ProjectInstanceFlavorParameters`. The JSON field names
  are unchanged, so the CRD schemas and existing manifests are unaffected; only
  Go code importing `apis/{cluster,namespaced}/cloud/v1alpha1` directly needs
  updating.

### Removed (breaking)
None. No upstream resource or field was removed in 2.19.0.

## v2.18.0 - 2026-08-01
### Added
- `storage`: `FileShareACL` — new managed resource in both the cluster-scoped
  (`storage.ovh.edixos.io`) and namespaced (`storage.ovh.m.edixos.io`) API groups,
  wrapping the new upstream `ovh_cloud_storage_file_share_acl` resource. It grants
  one IP or CIDR (`accessTo`) either `READ_WRITE` or `READ_ONLY` (`accessLevel`)
  access to a file share. `shareId` resolves by reference or selector to a
  `FileShare`. Every configurable field forces replacement, so an ACL is updated
  by recreating it.
- `BlockVolume` gains an optional `availabilityZone` field, plus
  `status.atProvider.currentState.location.availabilityZone`.

### Changed
- Upgrade OVH Terraform provider from 2.17.0 to 2.18.0.
- `BlockVolume` no longer replaces the volume when `encryption` is left unchanged —
  upstream now keeps the value from state instead of treating it as a diff.

### Removed (breaking)
Upstream removed the inline `access_rules` attribute from
`ovh_cloud_storage_file_share` in 2.18.0 and replaced it with the standalone
`ovh_cloud_storage_file_share_acl` resource. Accordingly, `FileShare` loses the
following fields in both API groups:

- `spec.forProvider.accessRules`
- `spec.initProvider.accessRules`
- `status.atProvider.accessRules`
- `status.atProvider.currentState.accessRules`

**Migration.** Declare one `FileShareACL` per rule that was previously inline.
Note that `accessRules` is not rejected after the CRDs are upgraded — it is
silently pruned by the API server, so a `FileShare` manifest left untouched will
appear to apply cleanly while its rules quietly stop being managed. Rewrite
manifests before upgrading:

```yaml
# before (v2.17.0)
spec:
  forProvider:
    accessRules:
      - accessTo: 10.0.0.0/24
        accessLevel: READ_WRITE

# after (v2.18.0)
apiVersion: storage.ovh.edixos.io/v1alpha1
kind: FileShareACL
metadata:
  name: example-acl
spec:
  forProvider:
    serviceName: <service-name>
    shareIdRef:
      name: example-share
    accessTo: 10.0.0.0/24
    accessLevel: READ_WRITE
```

Existing access rules already provisioned in OVHcloud are not deleted by this
upgrade; import them by setting `crossplane.io/external-name` to
`<serviceName>/<shareId>/<aclId>` on the new `FileShareACL`.

## v2.17.0 - 2026-07-25
### Added
Wired the remaining 31 OVHcloud Terraform resources, bringing managed-resource
coverage to the provider's full 166-resource schema (134 → 165 generated
resources per API scope). New CRDs, in both the cluster-scoped and namespaced
API groups:

- `cloud`: `FloatingIP`, `Quota`, `SecurityGroup`, `SSHKey`
- `gateway`: `CloudGateway`
- `network`: `PrivateVrackNetwork`, `PrivateVrackSubnet`
- `kms`: `KeyManagerContainer`, `KeyManagerContainerConsumer`, `KeyManagerSecret`, `KeyManagerSecretConsumer`
- `storage`: `EFS`, `BlockVolume`, `BlockVolumeBackup`, `BlockVolumeSnapshot`, `FileShare`, `FileShareNetwork`, `FileShareSnapshot`, `ProjectFileStorageShare`, `ProjectFileStorageShareNetwork`, `ProjectStorageLifecycleConfiguration`, `ProjectStorageReplicationJob`
- `databases`: `ProjectDatabaseClickhouseUser`, `ProjectDatabaseLogSubscription`
- `kube`: `LogSubscription`
- `logs`: `LogsEncryptionKey`, `LogsOutputGraylogStream`
- `me`: `IdentityUserToken`
- `email` (new group): `DomainAccount`
- `vrack`: `PublicRoutingPriority`, `VrackServicesOrder`

`IdentityUserToken` writes its `token` to the connection secret rather than
status, as the upstream attribute is marked sensitive.

### Changed
- Upgrade OVH Terraform provider from 2.13.1 to 2.17.0.

### Removed (breaking)
Upstream removed the deprecated `cassandra`, `m3db`, `m3aggregator` and `redis`
database engines in 2.17.0, along with the `ovh_cloud_project_database_ip_restriction`
resource. The following managed resources and their CRDs are therefore gone from both
the cluster-scoped (`databases.ovh.edixos.io`) and namespaced (`databases.m.ovh.edixos.io`)
API groups:
- `ProjectDatabaseIPRestriction` — declare IP restrictions directly on `ProjectDatabase` via `ipRestrictions` instead.
- `ProjectDatabaseM3DbNamespace`
- `ProjectDatabaseM3DbUser`
- `ProjectDatabaseRedisUser` — use `ProjectDatabaseValkeyUser` with a `valkey` engine cluster instead.

`ProjectDatabase` now rejects the `cassandra`, `m3db`, `m3aggregator` and `redis`
engines client-side; supported engines are `clickhouse`, `grafana`, `kafka`,
`kafkaConnect`, `kafkaMirrorMaker`, `mongodb`, `mysql`, `opensearch`, `postgresql`
and `valkey`.

### Infrastructure
- Initial setup of Marketplace extension assets (icon, readme, release notes, SBOM).
- Added CI step to run `up alpha xpkg append`.
- Link to Edixos from the marketplace readme and `meta.crossplane.io/readme`.

## v2.9.1 - 2025-11-16
### Changed
- Make up CLI install step resilient (fallback path handling).
- Ensure publish-artifacts job installs up CLI into PATH.
- Test Marketplace extension alpha append feature.

## v2.9.0 - 2025-11-15
### Changed
- Upgrade OVH Terraform provider to 2.9.0.
- Upgrade to Upjet v2 for Crossplane v2 compatibility.
- Update Terraform documentation scraper options.
- Include tags in release branch filter so v* tags trigger publish.
### Fixed
- Tag fetch failure fix in CI.
- Tag-related CI issue (#34).
- Private network reference fix.

## v1.1.0 - 2024-12-09
### Added
- New OVH resources introduced with provider upgrade.
### Changed
- Upgrade OVH Terraform provider to v1.1.0.

## v0.49.1 - 2024-09-25
### Added
- Support for using OAuth credentials.
- References to DB resources.

## v0.49.0 - 2024-09-24
### Changed
- Dependency updates.

## v0.40.1 - 2024-09-24
### Fixed
- External secret stores panic hotfix.

## v0.40.0 - 2024-03-21
### Changed
- Upgrade Terraform provider to v0.40.0.
- (Carry-over) upgrade actions from prior 0.39.0 work.

## v0.39.0 - 2024-03-16
### Changed
- Upgrade OVH Terraform provider to v0.39.0.
- Bump changed-files GitHub Action.
- Bump protobuf dependency.

## v0.37.0 - 2024-03-04
### Changed
- Upgrade OVH Terraform provider to v0.37.0.

## v0.36.1 - 2024-02-01
### Added
- Configuration example.
### Changed
- Improve README documentation.

## v0.1.4 - 2024-01-18
### Changed
- Upgrade Terraform module to version 0.36.1.
- Update private network kinds and references.

## v0.1.3 - 2024-01-09
### Added
- Rework project and user cloud resources.
### Changed
- Rename default kube resource for better naming convention.

## v0.1.2 - 2024-01-07
### Added
- Adapt container registry dependencies.

## v0.1.1 - 2024-01-06
### Added
- Update externalName for nodepools and ip restrictions.
### Changed
- Update external name for kube resource.

## v0.1.0 - 2024-01-05
### Added
- Initial codeowners configuration.

## Template For Future Releases
```
## vX.Y.Z - YYYY-MM-DD
### Added
- ...
### Changed
- ...
### Fixed
- ...
### Deprecated / Removed
- ...
```
