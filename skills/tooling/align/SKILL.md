---
name: align
description: >
  Run a convention alignment check across the codebase to measure adherence to
  documented standards. Use when you want to check health of the codebase,
  verify convention compliance, get a scored report, or find violations.
  Triggers on: "check conventions", "alignment scan", "codebase health",
  "convention violations", "quality check".
---

# Convention Alignment Check

Measure codebase adherence to documented conventions. Produces a scored report.

## Usage

```text
/align                # Full codebase scan
/align api-server     # API server checks only
/align control-plane  # Control plane checks only
/align security       # Security checks only
```

## User Input

```text
$ARGUMENTS
```

## How It Works

1. **Parse scope** from `$ARGUMENTS` (default: full scan)
2. **Load convention docs** from `specs/standards/`
3. **Run checks** via grep/glob/bash
4. **Produce scored report**

## Categories and Checks

| Category | Checks | Weight | Key concerns |
|----------|--------|--------|-------------|
| API Server | 8 | 40% | panic, error handling, generated code edits |
| Control Plane | 7 | 30% | SecurityContext, reconciliation, context propagation |
| Security | 5 | 30% | secrets in logs, input validation, fleet isolation |

## Report Format

- Overall weighted score (0-100%)
- Per-category scores
- Pass/fail per check with file:line references
- Failures grouped by severity
- Top 3 recommendations

## Interpreting Results

- **90-100%**: Excellent alignment. Ship with confidence.
- **70-89%**: Good alignment. Address blockers before merge.
- **50-69%**: Moderate alignment. Technical debt accumulating.
- **Below 50%**: Significant drift. Prioritize convention adherence.
