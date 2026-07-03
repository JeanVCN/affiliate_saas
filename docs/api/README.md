---
title: API Overview
status: active
owner: system-architect
last_verified_at: 2026-07-03
source_of_truth: true
depends_on:
  - ../domains/README.md
  - ../decisions/adr/002-http-router-framework.md
---

# API Overview

Affiliate SaaS starts with a JSON REST API served by the Go backend using Gin at the HTTP delivery boundary.

## Goals

- Support the first vertical slice:

```text
workspace -> marketplace program -> product -> affiliate link -> short redirect -> click event -> dashboard query
```

- Keep contracts understandable before code exists.
- Keep `gin.Context` out of domain/application code.
- Make workspace scoping explicit on every protected resource.
- Avoid marketplace API dependency in MVP contracts.

## Route Shape

- API routes use `/api/v1`.
- Public redirect route uses `/r/{slug}` outside the API prefix.
- Auth/session routes live under `/api/v1/auth`.
- Workspace-scoped resources live under `/api/v1/workspaces/{workspace_id}`.

## Common Conventions

- Request and response bodies are JSON unless the endpoint is a redirect.
- IDs are opaque strings at the API boundary.
- Timestamps use RFC 3339 UTC strings.
- Pagination uses `limit` and `cursor` where list size can grow.
- Mutations return the created or updated resource representation where practical.
- Soft-deleted or archived resources are hidden from default list endpoints.

## Error Shape

```json
{
  "error": {
    "code": "validation_failed",
    "message": "Human-readable summary.",
    "fields": {
      "name": "Name is required."
    }
  }
}
```

Initial error codes:

- `unauthenticated`
- `forbidden`
- `not_found`
- `validation_failed`
- `conflict`
- `rate_limited`
- `internal_error`

## Auth And Authorization

- MVP auth uses secure cookie sessions.
- Protected endpoints require an authenticated session.
- Workspace endpoints require membership in the referenced workspace.
- Authorization decisions belong to the backend, not the frontend.

## API Docs

- MVP endpoint inventory: `docs/api/rest/mvp-endpoints.md`
- Domain rules: `docs/domains/`
- Database model: `docs/database/`
