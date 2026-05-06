# Discordo MTG Fork &middot; [![ci](https://github.com/MTG-Thomas/discordo/actions/workflows/ci.yml/badge.svg)](https://github.com/MTG-Thomas/discordo/actions/workflows/ci.yml) [![license](https://img.shields.io/github/license/MTG-Thomas/discordo?logo=github)](https://github.com/MTG-Thomas/discordo/blob/main/LICENSE)

This is the Midtown Technology Group fork of [`ayn2op/discordo`](https://github.com/ayn2op/discordo), a lightweight Discord terminal client.

The fork keeps the runtime executable name `discordo` so existing config, cache, and keyring paths continue to work. Fork builds identify themselves in `discordo -version` with `distribution=mtg-fork`, and release assets are named `discordo-mtg_*`.

Fork-local changes focus on operator ergonomics:

- nano-friendly key aliases and pane focus cycling with `Tab` / `Shift+Tab`
- faster mouse wheel scrolling
- a softer terminal theme
- wrapped guild/channel tree navigation
- MTG fork release metadata and branded binary artifacts

For upstream project details, see [`ayn2op/discordo`](https://github.com/ayn2op/discordo). Heavily work-in-progress, expect breaking changes.

![Preview](.github/preview.png)

## Installation

### Prebuilt binaries

MTG fork binaries are published on the [MTG-Thomas/discordo releases page](https://github.com/MTG-Thomas/discordo/releases). Each release includes matching source at the release tag, platform archives, checksums, and GitHub artifact attestations.

Upstream prebuilt binaries are available from the upstream project.

### Package managers

- Arch Linux: `yay -S discordo-git`
- Gentoo (available on the guru repos as a live ebuild): `emerge net-im/discordo`
- FreeBSD: `pkg install discordo` or via the ports system `make -C /usr/ports/net-im/discordo install clean`.
- Nix: Add `pkgs.discordo` to `environment.systemPackages` or `home.packages`.

- Windows (Scoop):

```sh
scoop bucket add vvxrtues https://github.com/vvirtues/bucket
scoop install discordo
```

### Building from source

```bash
git clone https://github.com/MTG-Thomas/discordo
cd discordo
go build -trimpath -ldflags="-s" .
```

### Wayland clipboard support

`wl-clipboard` is required for clipboard support.

## Usage

### Token (UI, recommended)

1. Run the `discordo` executable with no arguments.

2. Enter your token and click on the "Login" button to save it.

### Token (environment variable)

Set the value of the `DISCORDO_TOKEN` environment variable to the authentication token to log in with.

```sh
DISCORDO_TOKEN="OTI2MDU5NTQxNDE2Nzc5ODA2.Yc2KKA.2iZ-5JxgxG-9Ub8GHzBSn-NJjNg" discordo
```

### QR (UI)

1. Run the `discordo` executable with no arguments.

2. Click on the "Login with QR" button.

3. Follow the instructions in the QR Login screen.

## Configuration

The configuration file allows you to configure and customize the behavior, keybindings, and theme of the application.

- Unix: `$XDG_CONFIG_HOME/discordo/config.toml` or `$HOME/.config/discordo/config.toml`
- Darwin: `$HOME/Library/Application Support/discordo/config.toml`
- Windows: `%AppData%/discordo/config.toml`

Discordo uses the default configuration if a configuration file is not found in the aforementioned path; however, the default configuration file is not written to the path. [The default configuration can be found here](./internal/config/config.toml).

> [!IMPORTANT]
> Automated user accounts or "self-bots" are against Discord's Terms of Service. I am not responsible for any loss caused by using "self-bots" or Discordo.

## License

Copyright (C) 2025-present ayn2op

This fork and the upstream project are licensed under the GNU General Public License v3.0 (GPL-3.0).
See the [LICENSE](./LICENSE) file for the full license text.

When distributing MTG fork binaries, provide the matching source for the exact release tag. The release workflow publishes binaries from tags on `main` so the corresponding GPL source is available in this repository.
