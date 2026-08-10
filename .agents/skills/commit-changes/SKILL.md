---
name: commit-changes
description: Commit all working-tree changes as brief, related-topic commits; update CHANGELOG.md with detail; mark completed ROADMAP.md items; bump the Go package version only when code changed. Use only when explicitly invoked (e.g. /commit-changes or attach this skill).
disable-model-invocation: true
---

# Commit Changes

When explicitly invoked: commit **all** pending changes, split by topic, keep messages brief, put detail in `CHANGELOG.md`, sync `ROADMAP.md` if a point is done, and bump the Go package version when code changed.

## Workflow

1. **Inspect** (parallel): `git status`, `git diff` (staged + unstaged), `git log -8 --oneline`, latest version (`git describe --tags --abbrev=0` and any `version.go` / `VERSION`), read the `ROADMAP.md` sections touched by the diff. If there's nothing to commit, say so and stop here.

2. **Classify the change set** before writing anything: does it contain code, or only docs/config (README, `.githooks`, `.gitattributes`, etc.)? This decides whether a version bump and CHANGELOG heading happen at all, so settle it now rather than guessing later. If code changed, also decide the semver bump: breaking → major, feature → minor, fix/chore/docs-alongside-code → patch. No prior tag → start at `v0.1.0` (or the next patch if already above that). Compute the target version number now — you'll use it in step 3, but the actual file bump/tag/commit still happens last, in step 6.

3. **Group** changes into the fewest coherent commits. A group is everything that serves one concern — typically the files under one feature, directory, or fix, not "all Go files" or "everything touched today." Don't dump unrelated files into one commit just because they were edited in the same session.

4. **CHANGELOG.md** (skip entirely for docs/config-only change sets — see step 2): append a dated entry under the target version heading from step 2, with the detail that doesn't belong in the commit subject. Create the file if missing. Prefer Keep a Changelog style:

```markdown
## [X.Y.Z] - YYYY-MM-DD

### Added | Changed | Fixed | Removed
- Concrete bullet of what changed and why it matters
```

5. **ROADMAP.md**: if the change set fully completes a tracked item, flip `[ ]` / `[~]` → `[x]` and adjust Notes if needed. Partial progress → `[~]` only when clearly warranted. If nothing matches, leave it alone.

6. **Commit** each group from step 3 (include the matching CHANGELOG/ROADMAP hunks in the same commit when they document that group; leftover changelog/roadmap-only edits get their own brief docs commit). Then, only if step 2 found code changes:
   - Update `version.go`'s `Version` const if present; otherwise add a root `version.go` with `package gogige` and `const Version = "X.Y.Z"`, using the version computed in step 2.
   - Final commit: `chore: bump version to X.Y.Z` (version file + CHANGELOG heading if not already committed).
   - Annotated tag: `git tag -a vX.Y.Z -m "vX.Y.Z"`. Don't push tags unless asked.

7. **Verify**: `git status` is clean (or only intentional leftovers remain). Don't push unless asked.

## Commit messages

- Brief subject only (≈50 chars). Match the repo's existing tone (`feat:`, `fix:`, `chore:`, …).
- No body unless a safety note is required (secrets, breaking behavior) — the detail belongs in `CHANGELOG.md`, not the commit message.

Example:

```text
feat: discover cameras per interface
```

not a multi-paragraph body describing the discovery logic.

## Git safety

Follow the user's committing-changes-with-git rules: no config changes, no amend/force/hooks-skip unless explicitly requested, no secrets in commits, HEREDOC for `-m`, and a fresh commit after a hook failure (never amend a failed commit).

## Checklist

```
- [ ] Nothing pending, or everything grouped by related topic
- [ ] Brief message per commit
- [ ] CHANGELOG.md updated under the target version (skipped for docs/config-only changes)
- [ ] ROADMAP.md checked; completed items marked [x]
- [ ] Version bumped (version.go + annotated tag vX.Y.Z) — only when code changed
- [ ] git status clean (or leftovers explained)
```