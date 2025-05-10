# Software Bill of Materials (SBOM)

This directory contains the **SBOM (Software Bill of Materials)** for the `muxingbird` project.
An SBOM is a detailed inventory of all the dependencies and components used to build this
software, including their versions, licenses, and origins.

## Why SBOM Matters

Including an SBOM provides critical benefits:

- **Transparency**: Reveals all components included in the binary or container image.
- **Security**: Enables automated vulnerability scanning and impact analysis for CVEs.
- **Compliance**: Assists with license auditing and regulatory standards (e.g. Executive Order
  14028, NIST SSDF).
- **Reproducibility**: Helps others verify or reproduce builds accurately.

## Tooling: [Syft](https://github.com/anchore/syft)

We use [`syft`](https://github.com/anchore/syft) by Anchore to generate SBOMs in multiple
standard formats.

### 🔧 Installation

If you do not have `syft` installed, you can install it with this shell script:

```bash
#!/usr/bin/env bash
if ! command -v syft &> /dev/null; then
  echo "Syft not found. Installing..."
  curl -sSfL https://raw.githubusercontent.com/anchore/syft/main/install.sh | sh -s -- -b "$HOME/.local/bin"
else
  echo "Syft already installed."
fi
```

## Generating a New SBOM File

From the project root, you can regenerate the SBOM as follows:

```bash
just sbom
```

## Notes

- The SBOM reflects the state of dependencies at the time of the latest build.
- If you change project dependencies, you must regenerate the SBOM.
