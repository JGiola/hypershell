---
name: memory
description: >
  Manage the auto-memory system for this project. Search, audit, prune, and
  create memories with proper frontmatter. Triggers on: "check memory",
  "what do we remember about", "find the memory about", "clean up memories",
  "audit memories", "add to memory".
---

# Memory Management

Manage the auto-memory system for the HyperShell project.

## Usage

```text
/memory                    # Show summary of all memories
/memory search <query>     # Search for a topic
/memory audit              # Check for stale/duplicate memories
/memory prune              # Remove stale memories (with confirmation)
/memory add <topic>        # Create a new memory
```

## User Input

```text
$ARGUMENTS
```

## Memory Location

```text
$HOME/.claude/projects/<project-slug>/memory/
```

## Subcommands

### `/memory` -- Summary

Read `MEMORY.md`, count memories by type, list recently modified.

### `/memory search <query>`

Grep through memory files, show matching excerpts with frontmatter.

### `/memory audit`

Check for: stale memories (>3 months), duplicates, missing frontmatter, orphaned files, broken links, oversized index.

### `/memory prune`

Run audit, present findings, confirm deletions, update `MEMORY.md`.

### `/memory add <topic>`

Create new memory file with frontmatter:

```markdown
---
name: <descriptive name>
description: <one-line description>
type: <user|feedback|project|reference>
---

<memory content>
```

## What NOT to Store

- Code patterns derivable from reading current code
- Git history
- Debugging solutions (the fix is in the code)
- Anything in CLAUDE.md
- Ephemeral task details
- Secrets or credentials
