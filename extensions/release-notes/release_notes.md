# Release Notes

> IMPORTANT: Update this file for each tagged release before (or as part of) CI so it accurately reflects changes delivered in that version. The CI workflow will append this file along with other extension assets to the published package.

## Unreleased
- Initial setup of Marketplace extension assets (icon, readme, release notes, SBOM).
- Added CI step to run `up alpha xpkg append`.

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
