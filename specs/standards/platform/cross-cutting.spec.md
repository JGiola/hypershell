# Cross-Cutting Platform Conventions

Conventions that apply across all HyperShell components.

## Image references must match across the stack

Every image name and tag used in manifests, env vars, kustomization overlays, kind load commands, and Makefile targets must resolve to the same artifact.

After changing an image name or tag, grep all overlays, Makefiles, and deployment manifests. Mismatches cause silent deployment failures.

## Reconcile, don't create-or-skip

Operator and backend code that creates K8s resources must use update-or-create (reconcile) patterns, not create-and-ignore-`AlreadyExists`.

**Pattern to avoid**:
```go
err := client.Create(ctx, obj)
if apierrors.IsAlreadyExists(err) {
    return nil // BAD: silently skips spec updates
}
```

**Correct pattern**:
```go
existing := &v1.Resource{}
err := client.Get(ctx, key, existing)
if apierrors.IsNotFound(err) {
    return client.Create(ctx, obj)
}
existing.Spec = obj.Spec
return client.Update(ctx, existing)
```

## Never silently swallow partial failures

Every error path must propagate or explicitly log the failure. If a step can fail independently, collect errors and return them together.

```go
var errs []error
for _, item := range items {
    if err := reconcile(item); err != nil {
        errs = append(errs, fmt.Errorf("reconcile %s: %w", item.Name, err))
    }
}
return errors.Join(errs...)
```

## Restricted SecurityContext on all containers

All containers in manifests must set:
- `runAsNonRoot: true`
- `capabilities.drop: ["ALL"]`
- `readOnlyRootFilesystem: true` (unless a specific write path is required)

```yaml
securityContext:
  runAsNonRoot: true
  allowPrivilegeEscalation: false
  capabilities:
    drop: ["ALL"]
  readOnlyRootFilesystem: true
  seccompProfile:
    type: RuntimeDefault
```
