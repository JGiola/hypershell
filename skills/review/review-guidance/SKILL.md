---
name: review-guidance
description: >
  HyperShell-specific review standards for PRs. Loads project conventions,
  security requirements, and component-specific checklists. Use when reviewing
  PRs in this repository.
---

# HyperShell Review Guidance

Apply these standards when reviewing PRs in this repository.

## Mandatory Context Files

Before analyzing any PR, load:

1. `CLAUDE.md`
2. `specs/standards/security/security.spec.md`
3. `specs/standards/control-plane/conventions.spec.md`

## Review Checklists

### Go Backend (API Server)

- [ ] No `panic()` in production code -- return `fmt.Errorf` with context
- [ ] Errors wrapped: `fmt.Errorf("context: %w", err)`
- [ ] `errors.IsNotFound` handled for 404 scenarios
- [ ] No secrets in logs or error messages
- [ ] Input validated (K8s DNS labels, URL parsing)
- [ ] Log injection prevented
- [ ] OpenAPI client not manually edited (`make generate` only)

### Go Control Plane

- [ ] SecurityContext on all pod specs
- [ ] Resource limits/requests on containers
- [ ] Status updated on error paths
- [ ] No `panic()` in non-test code
- [ ] Proper context propagation (no `context.TODO()`)
- [ ] `gofmt -w .` applied
- [ ] `go vet ./...` passes
- [ ] Reconcile pattern used (not create-or-skip)

### General

- [ ] No `panic()` in production Go code
- [ ] PostgreSQL for persistent storage
- [ ] Image references consistent across manifests
- [ ] Conventional commit message

## Severity Classification

- **Blocker** -- Must fix. Security vulnerabilities, data loss, secret leaks
- **Critical** -- Should fix. Missing error handling, `panic()` in handlers
- **Major** -- Important. Architecture violations, missing tests
- **Minor** -- Nice-to-have. Style, docs gaps
