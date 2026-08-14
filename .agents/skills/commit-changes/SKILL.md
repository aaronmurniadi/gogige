---
name: commit-changes
description: Commit all working-tree changes as brief, related-topic commits; update CHANGELOG.md with detail; mark completed ROADMAP.md items; bump the Go package version only when code changed. Use only when explicitly invoked (e.g. /commit-changes or attach this skill).
disable-model-invocation: true
---

# Commit Changes

When explicitly invoked (e.g. `/commit-changes`): commit **all** pending changes, split by topic, keep messages brief, put detail in `CHANGELOG.md`, sync `ROADMAP.md` if a point is done, and bump the Go library version when code changed. This is a project-scoped Ante skill at `.agents/skills/`.

## Workflow

1. **Inspect** (parallel): `git status`, `git diff` (staged + unstaged), `git log -8 --oneline`, latest version (`git describe --tags --abbrev=0`), read the `ROADMAP.md` sections touched by the diff. If there's nothing to commit, say so and stop here.

2. **Classify the change set** before writing anything: does it contain code, or only docs/config (README, `.githooks`, `.gitattributes`, AGENTS.md, etc.)? This decides whether a version bump and CHANGELOG heading happen at all, so settle it now rather than guessing later. If code changed, also decide the semver bump: breaking → major, feature → minor, fix/chore/docs-alongside-code → patch. No prior tag → start at `v0.1.0` (or the next patch if already above that). Compute the target version number now — you'll use it in step 3, but the actual file bump/tag/commit still happens last, in step 6.

3. **Group** changes into the fewest coherent commits. A group is everything that serves one concern — typically the files under one feature, directory, or fix, not "all Go files" or "everything touched today." Don't dump unrelated files into one commit just because they were edited in the same session.

4. **CHANGELOG.md** (skip entirely for docs/config-only change sets — see step 2): append a dated entry under the target version heading from step 2, with the detail that doesn't belong in the commit subject. Match this repo's existing style exactly: `## [X.Y.Z] - YYYY-MM-DD`, then `### Added | Changed | Fixed | Removed` section headings with concrete, backticked symbols, plus a `### Tests` heading naming the new/exercised tests. Reuse an existing heading if a version entry is already being written from an earlier change in this same release; otherwise create it:

```markdown
## [X.Y.Z] - YYYY-MM-DD

### Added

- `symbol` — concrete change and why it matters.

### Changed

- `symbol` — breaking renames / behavior change.

### Fixed

- `symbol` — bug fixed and effect.

### Tests

- `TestFoo`, `TestBar` names, or the behavior assertions added.
```

5. **ROADMAP.md**: if the change set fully completes a tracked item, flip `[ ]` / `[~]` → `[x]` and adjust Notes if needed. Partial progress → `[~]` only when clearly warranted; restate open items as-is. Also keep the **Outstanding work** summary section (`.md:~7`) consistent with any row you flip. If nothing matches, leave it alone.

6. **Commit** each group from step 3 (include the matching CHANGELOG/ROADMAP/TODO hunks in the same commit when they document that group; leftover changelog/roadmap-only edits get their own brief docs commit). Then, only if step 2 found code changes:
   - Bump the `Version` const — it lives in `options.go` (root `package gogige`), not a separate `version.go`. Update `const Version = "X.Y.Z"` to the number computed in step 2. There is no `version.go` in this repo; don't create one.
   - Final commit: `chore: bump version to X.Y.Z` (options.go + CHANGELOG heading if not already committed).
   - Annotated tag: `git tag -a vX.Y.Z -m "vX.Y.Z"`. Don't push tags unless asked.

7. **Verify**: `git status` is clean (or only intentional leftovers remain). The `.githooks/pre-commit` hook runs `gofmt` + `go test ./...` on every commit — don't skip it. Don't push unless asked.

## Commit messages

- Brief subject only (≈50 chars). Match the repo's existing tone (`feat:`, `fix:`, `chore:`, `refactor:`, `docs:`).
- No body unless a safety note is required (secrets, breaking behavior) — the detail belongs in `CHANGELOG.md`, not the commit message.

Example:

```text
feat: discover cameras per interface
```

not a multi-paragraph body describing the discovery logic.

## Git safety

- Never commit secrets; never amend a hook-failed commit — the pre-commit hook may reformat files, so after a hook failure inspect what changed and create a fresh commit.
- No git config changes, no `--no-verify`/`--no-gpg-sign`, no amend/force, unless the user explicitly asks.
- Use HEREDOC (`git commit -m "..."`) for messages so subjects with quotes/special chars commit cleanly.

## Checklist

```
- [ ] Nothing pending, or everything grouped by related topic
- [ ] Brief message per commit
- [ ] CHANGELOG.md updated under the target version, with Tests heading (skipped for docs/config-only changes)
- [ ] ROADMAP.md checked; completed items marked [x], Outstanding summary synced
- [ ] Version bumped in options.go (const Version) + annotated tag vX.Y.Z — only when code changed
- [ ] git status clean (or leftovers explained)
```
