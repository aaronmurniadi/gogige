---
name: commit-changes
description: Commit all working-tree changes as brief, related-topic commits; update CHANGELOG.md with detail; mark completed ROADMAP.md items; bump the Go package version only when code changed. Use only when explicitly invoked (e.g. /commit-changes or attach this skill).
disable-model-invocation: true
---

# Commit Changes

When explicitly invoked: commit **all** pending changes, split by topic, keep messages brief, put detail in `CHANGELOG.md`, sync `ROADMAP.md` if a point is done, and bump the Go package version when code changed.

## Workflow

1. **Inspect** (parallel): `git status`, `git diff` (staged + unstaged), `git log -8 --oneline`, latest version (`git describe --tags --abbrev=0` and any `version.go` / `VERSION`), read `ROADMAP.md` sections touched by the diff.
2. **Group** changes into the fewest coherent commits (same concern = one commit). Do not dump unrelated files into one commit.
3. **CHANGELOG.md**: append a dated entry under the **new** version heading with the detail that does **not** belong in the commit subject. Create the file if missing. Prefer Keep a Changelog style:

```markdown
## [X.Y.Z] - YYYY-MM-DD

### Added | Changed | Fixed | Removed
- Concrete bullet of what changed and why it matters
```

4. **ROADMAP.md**: if the change set fully completes a tracked item, flip `[ ]` / `[~]` → `[x]` and adjust Notes if needed. Partial progress → `[~]` only when clearly warranted. If nothing matches, leave ROADMAP alone.
5. **Commit** each group (include matching CHANGELOG/ROADMAP hunks in the same commit when they document that group; leftover changelog/roadmap-only edits get their own brief docs commit).
6. **Bump Go package version** (after content commits). Skip entirely — no bump, no tag, no new CHANGELOG heading — when the change set contains **no code** (e.g. README/docs/.githooks/.gitattributes-only changes).
   - Semver from change set: breaking → major, feature → minor, fix/chore/docs → patch. No prior tag → start at `v0.1.0` (or next patch if already above).
   - Update `version.go` `Version` const if present; otherwise add root `version.go` with `package gogige` and `const Version = "X.Y.Z"`.
   - Record the version in the CHANGELOG heading for this release.
   - Final commit: `chore: bump version to X.Y.Z` (version file + changelog heading if not already committed).
   - Annotated tag: `git tag -a vX.Y.Z -m "vX.Y.Z"`. Do not push tags unless asked.
7. **Verify**: `git status` clean (or only intentional leftovers). Do not push unless asked.

## Commit messages

- Brief subject only (≈50 chars). Match repo tone (`feat:`, `fix:`, `chore:`, …).
- No body unless a safety note is required (secrets, breaking behavior).
- Detail lives in `CHANGELOG.md`, not the commit message.

Example:

```text
feat: discover cameras per interface
```

not a multi-paragraph body describing discovery.

## Git safety

Follow the user's committing-changes-with-git rules: no config changes, no amend/force/hooks-skip unless explicitly requested, no secrets in commits, HEREDOC for `-m`, new commit after hook failure (never amend a failed commit).

## Checklist

```
- [ ] All changes grouped by related topic
- [ ] Brief message per commit
- [ ] CHANGELOG.md updated with detail under new version (skipped for docs-only changes)
- [ ] ROADMAP.md checked; completed items marked [x]
- [ ] Version bumped (version.go + annotated tag vX.Y.Z) — only when code changed
- [ ] git status clean (or leftovers explained)
```
