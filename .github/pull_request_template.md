## Summary

<!--
The smallest view that makes the point: pseudocode, a call tree, a file
tree, a Mermaid diagram or a diff sketch, each beside one or two lines of
why. If it settles something about Herdr's API, say what was probed.
-->

Closes #

## Evidence

<!--
A before and an after: a failing then passing test, a probe's output, a
tab bar from a live session. `make check` is the gate; show what you looked
at beyond it.
-->

- **Before:**
  **After:**

## Merge Danger

<!--
A two-way door is cheap to walk back; a one-way door is not — a label
written onto the user's tabs, a token reported to a workspace, a setting
whose meaning changes.
-->

**Door:** <!-- one-way or two-way -->

**Blast Radius:** <!-- one word: what could break, and for whom -->

## Before merging

- [ ] `make check` is green locally.
- [ ] One logical change per commit, Conventional Commits, no co-author
      trailer. **PRs are rebased, never squashed or merged**, so the commit
      messages are what lands in the history and in the changelog.
- [ ] Tests land with the behaviour they cover.
- [ ] `CHANGELOG.md`, the tags and the version in `herdr-plugin.toml` are
      untouched — they belong to release-please.
- [ ] A probe that taught something new is recorded in `AGENTS.md` and in
      `docs/architecture/herdr-socket-api.md`.
- [ ] This pull request is one thing.
