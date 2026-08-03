---
name: reconcile
description: >
  Top-level autonomous orchestrator that reconciles all specs against the codebase.
  Reads skills/RECONCILE.md for checkpoint state, runs field-level gap analysis,
  plans waves, and executes the full-stack-pipeline per wave. Idempotent: safe to
  run repeatedly. Use when: "reconcile everything", "run full reconciliation",
  "what's the gap", "autonomous build", "build everything", "spec coverage".
---

# Reconcile

Autonomous code reconciliation against the spec corpus. Orchestrates all other
skills into a single convergence loop.

## User Input

```text
$ARGUMENTS
```

Supported arguments:
- *(empty)* -- full reconciliation across all specs
- `--dry-run` -- gap analysis only, no code changes
- `--domain <platform|standards>` -- scope to a single spec domain
- `--spec <path>` -- scope to a single spec file
- `--wave <N>` -- start from a specific wave number

## Checkpoint: skills/RECONCILE.md

**Read `skills/RECONCILE.md` first.** This is the checkpoint file.

### Idempotency Rules

1. If `RECONCILE.md` has a gap table and `Last analyzed` matches the current
   HEAD commit, skip Phases 1-4. Jump to Phase 5 or Phase 6.
2. If specs or code changed since `Last analyzed`, re-run gap analysis for
   affected specs only.
3. After each wave, update `RECONCILE.md` in-place.
4. Commit `RECONCILE.md` alongside code changes.

## Phases

### Phase 1-2: Spec Discovery & Dependency Graph

Read the Spec Registry in `specs/index.spec.md` and the dependency order in
`RECONCILE.md`.

### Phase 3: Gap Analysis

For each spec in topological order, check every requirement at field level:

| Layer | What to check | Where |
|-------|---------------|-------|
| API | Routes, schemas, `required[]` | `components/api-server/openapi/` |
| BE | Models, DAOs, handlers, migrations | `components/api-server/plugins/` |
| gRPC | Proto definitions, handlers, presenters | `components/api-server/proto/`, `plugins/*/grpc_*` |
| CP | Watcher, reconciler logic | `components/control-plane/` |

### Phase 4: Update RECONCILE.md

Merge gap tables. Update coverage summary. If `--dry-run`, stop here.

### Phase 5: Wave Planning

Waves follow `/full-stack-pipeline` dependency order:
API → gRPC → BE → CP → Integration

### Phase 6: Execution Loop

For each wave:
1. Present wave plan
2. Dispatch work following `/full-stack-pipeline` patterns
3. Verify with lint and build gates
4. Run `/align` for affected scope
5. Update `RECONCILE.md`

## Constraints

- Never modify specs during reconciliation (code changes only)
- One wave at a time; downstream waits for upstream verification
- Max 3 retries per wave before human escalation
- Always update `skills/RECONCILE.md` after state changes
