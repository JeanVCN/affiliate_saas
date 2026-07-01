---
title: Documentation Governance
status: draft
owner: technical-writer
last_verified_at: 2026-07-01
source_of_truth: true
---

# Documentation Governance

## Lifecycle

Documents move through:

`draft -> active -> stale -> superseded -> archived`

## Required Metadata

Durable docs should include:

```yaml
title:
status:
owner:
last_verified_at:
source_of_truth:
depends_on:
supersedes:
superseded_by:
```

## Split Criteria

Split a document when it exceeds roughly 250-400 lines, mixes owners, changes at different cadences, or forces agents to read unrelated context.

## Archive Criteria

Archive when a document is superseded, rejected, historically useful but inactive, or no longer has an owner.

## Pull Request Checklist

- [ ] API changed: API docs or generated contract updated.
- [ ] Database changed: schema docs or migration notes updated.
- [ ] Architecture changed: ADR created or updated.
- [ ] Domain behavior changed: domain docs updated.
- [ ] AI prompt/provider behavior changed: prompt/eval docs updated.
- [ ] Deployment changed: deployment docs updated.

