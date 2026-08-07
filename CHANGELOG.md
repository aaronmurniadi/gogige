# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2026-08-08

### Changed
- `examples/dump-xml` default filename is `<vendor>-<serial>-genicam.xml` (path-safe tokens from discovery ABRM) instead of an IP-based `*-genapi.xml` name

### Added
- GigE Vision local references: `GigE_Vision_for_Realtime_MV_11052010.pdf`, `GigE_Features_Reference.pdf`
- Root `.gitignore` for example camera dumps (`*.xml` / `*.jpg` / `*.log` under `examples/`)
- Cursor skill `.cursor/skills/commit-changes` for topic-split commits, changelog, roadmap, and version bumps
