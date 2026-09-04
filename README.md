# astmn

A lightweight asset manager designed for tracking and extracting binary files via Git-friendly manifests.

> [!WARNING]
> This project is in **early development**. Features and CLI interfaces are unstable and may not work.

## Supported Formats

| Archive Type | Status      |
| :----------- | :---------- |
| `.zip`       | Implemented |
| `.7z`        | Implemented |
| `.tar.gz`    | Implemented |
| `.gzip`      | Planned     |
| `.tar`       | Planned     |
| `.rar`       | Planned     |

## How It Works

### Current Mode

1. Reads a YAML manifest tracking binary asset versions and target paths.
2. Downloads full archives from remote sources (e.g., Google Drive, direct HTTP).
3. Verifies integrity using SHA-256 checksums and safe path extraction.
4. Overwrites local assets when the manifest in the registry changes.

### Roadmap

- File-level Hashing: Track individual binary files instead of full archives to download only updated files, drastically reducing bandwidth usage.
- Dedicated Hosting/Server: Custom server implementation to handle delta syncs and chunked binary storage.

## Quickstart

```bash
# Build the binary
make build

# View the package contents
./astmn view templates/Example_UE_Asset_Pack.yml

# Install assets from a manifest
./astmn install templates/Example_UE_Asset_Pack.yml

# Interactively generate a new package manifest
./astmn pack
```
