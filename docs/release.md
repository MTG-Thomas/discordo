# MTG Fork Release Runbook

This fork publishes GPL-3.0 binaries from `MTG-Thomas/discordo` while keeping the runtime executable name `discordo`.

## Policy

- Release from `main` only.
- Use fork release tags such as `v0.1.0-mtg.3`.
- Do not tag a pull request branch for release.
- Keep `LICENSE` and upstream notices intact.
- Release assets must be named `discordo-mtg_*`.
- `discordo -version` must include `distribution=mtg-fork`.

## Before Tagging

1. Merge the release PR into `main`.
2. Update the local checkout:

   ```sh
   git switch main
   git pull --ff-only origin main
   ```

3. Run checks:

   ```sh
   pwsh ./scripts/check.ps1
   ```

## Cut a Release

Create and push an annotated tag from `main`:

```sh
git tag -a v0.1.0-mtg.3 -m "MTG fork release v0.1.0-mtg.3"
git push origin v0.1.0-mtg.3
```

The release workflow verifies the tag commit is reachable from `main`, creates or updates the GitHub release, builds platform archives, uploads checksums, and creates artifact attestations.

## Verify the Release

Check the release page:

```sh
gh release view v0.1.0-mtg.3 --repo MTG-Thomas/discordo
```

Expected assets:

- `discordo-mtg_Linux_ARM64.tar.gz`
- `discordo-mtg_Linux_X64.tar.gz`
- `discordo-mtg_macOS_ARM64.tar.gz`
- `discordo-mtg_macOS_X64.tar.gz`
- `discordo-mtg_Windows_ARM64.zip`
- `discordo-mtg_Windows_X64.zip`
- one checksum file per platform

Download one archive and confirm:

```sh
discordo -version
```

The output should include:

```text
distribution=mtg-fork
```

## GPL Source Access

For each distributed binary, the matching source is the repository content at the same release tag. Do not replace release assets with binaries built from a different commit unless the tag is also moved and the release workflow is rerun from that commit.
