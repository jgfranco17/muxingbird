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

## Tooling

We use [`syft`](https://github.com/anchore/syft) to generate SBOMs in multiple standard formats,
and [`cyclonedx`](https://github.com/CycloneDX/cyclonedx-cli) to validate the entries.

## Generating a New SBOM File

From the project root, the Justfile provides basic utilities to create and validate SBOM files:

```bash
just generate-sbom
just validate-sbom
```

## Notes

- The SBOM reflects the state of dependencies at the time of the latest build.
- If you change project dependencies, you must regenerate the SBOM.
