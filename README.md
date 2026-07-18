# go-stix

A modern, idiomatic Go library for STIX 2.1 cyber threat intelligence objects.

The goal of this project is to provide a lightweight Go SDK for building,
managing, and exchanging STIX objects with a focus on:

- Cloud-native security
- OpenTelemetry integration
- Cyber threat intelligence automation
- TAXII interoperability
- Security observability

## Project Status

Early development.

Current milestone:

**Milestone 1 - STIX Core Foundation**

Planned:

- STIX object model
- STIX bundles
- Cyber Observable Objects (SCO)
- STIX Domain Objects (SDO)
- Deterministic object identifiers
- OpenTelemetry mappings
- TAXII support

## Design Principles

- Idiomatic Go
- Minimal dependencies
- Strong typing
- Test-driven development
- Extensible architecture

## Development

Requirements:

- Go 1.24+

Run tests:

```bash
make test