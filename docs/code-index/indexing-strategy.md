---
title: Code Indexing Strategy
status: active
owner: context-engineering-specialist
last_verified_at: 2026-07-03
source_of_truth: true
depends_on:
  - ../context/context-engineering-guide.md
  - ../architecture/system-overview.md
  - ../api/rest/mvp-endpoints.md
---

# Code Indexing Strategy

The indexing strategy keeps agent context small and accurate as the repository grows from documentation into implementation.

## Search Order

1. Exact file lookup from `docs/INDEX.md` or `.context-manifest.json`.
2. Exact text search with `rg`.
3. Symbol/package search once code exists.
4. Generated inventories such as route maps or migration maps.
5. Semantic search only when exact search cannot identify the right files.

## What To Index

Current documentation index:

- source-of-truth docs;
- ADRs;
- domains;
- API contracts;
- database docs;
- quality/security/deployment/onboarding docs.

Future code index:

- Go packages under `backend/internal/modules/*`;
- Gin route registrations;
- migrations under `backend/migrations`;
- repository interfaces and implementations;
- frontend routes and components;
- tests by module/domain.

## Index Freshness

- Generated indexes must include generation timestamp and command.
- Generated indexes must not be treated as source of truth over code, migrations, tests, or accepted ADRs.
- Update indexes when routes, migrations, module boundaries, or frontend routes change.

## Exact Search Roles

Use exact search for:

- endpoint paths;
- table names;
- migration filenames;
- ADR numbers;
- domain names;
- error codes;
- environment variable names;
- Gin handler or route registration names.

## Symbol Search Roles

Use symbol search for:

- Go functions, structs, interfaces, methods;
- route registration functions;
- repository methods;
- service/use-case methods;
- frontend components/hooks.

## Semantic Search Roles

Use semantic search only for:

- broad product/domain questions;
- finding conceptually related docs when exact terms are unknown;
- summarizing older design intent.

Semantic search must not override executable source, migrations, tests, or accepted ADRs.

## Planned Generated Indexes

Once implementation starts:

- `docs/code-index/generated/routes.md`
- `docs/code-index/generated/backend-modules.md`
- `docs/code-index/generated/database-entities.md`
- `docs/code-index/generated/frontend-routes.md`
- `docs/code-index/generated/test-inventory.md`

Generated indexes should be rebuilt by documented commands before being trusted.
