# TODO

Outstanding work tracked in [`ROADMAP.md`](ROADMAP.md). Items below are the not-yet-done `[ ]` / partial `[~]` rows.

## Phase 3 — GenApi 2.1.1

- `Category` / `StructReg` as first-class node types (parsed/skipped today) — `[ ]`, `ROADMAP.md:154`
- `pIsImplemented` / `pIsAvailable` / `pIsLocked` / `pInvalidator` cache invalidation — `[ ]`, `ROADMAP.md:156`
- ManifestTable (`0x01D0`) path, preferred over FirstURL when present — `[ ]`, `ROADMAP.md:157`
- SwissKnife ops (`+ - * / % ** & | ^ << >> && || ?:`) — subset only — `[~]`, `ROADMAP.md:158`
- SwissKnife funcs (`SQRT`, `FLOOR`, `CEIL`, `ABS`) — verify against § formula grammar — `[~]`, `ROADMAP.md:159`
- SFNC-required features for streaming — formal `Gev*` coverage — `[~]`, `ROADMAP.md:161`

## Phase 4 — GenTL 1.6 (optional, pure-Go primary)

- `gentl/cti.go` `.cti` loader (`dlopen` / CGO) — `[ ]`, `ROADMAP.md:51,179`
- GenTL module ladder (`TLOpen`, `IFOpenDevice`, `DevOpenDataStream`, buffer announce/queue, `DSStartAcquisition` + `EVENT_NEW_BUFFER`, payload-type mapping, SFNC feature names) — `[ ]`, `ROADMAP.md:180-186`
- Optional `gentl/` CGO producer/consumer — `[ ]`, `ROADMAP.md:170`
- `cmd/` CLIs — discover done, stream done; extend as needed — `[~]`, `ROADMAP.md:171`

## Layout debt

- Root extras (`device.go`, `grab.go`, `live.go`, `alias.go`, …) — converge on Phase 4 surface — `[~]`, `ROADMAP.md:64`