---
title: Local Development
status: active
owner: infrastructure-engineer
last_verified_at: 2026-07-04
source_of_truth: true
depends_on:
  - ../deployment/docker.md
  - ../quality/quality-gates.md
---

# Local Development

The repository now includes the initial Go backend scaffold. These commands define the current verification path and the intended local workflow as implementation grows.

## Current Setup

From the repository root:

```bash
node -e "for (const f of ['.docs-index.json','.context-manifest.json']) JSON.parse(require('fs').readFileSync(f,'utf8'));"
```

This remains the documentation verification command.

## Backend Workflow

From the repository root:

```bash
cd backend
GOCACHE=/tmp/affiliate-saas-go-cache go test ./...
```

The normal test suite is expected to pass without local services. PostgreSQL integration tests are skipped unless `AFFILIATE_TEST_DATABASE_URL` is set.

Expected services:

- PostgreSQL for persistence once repository endpoints are implemented.
- Redis only after async jobs are introduced.

## Local PostgreSQL

From the repository root:

```bash
docker compose up -d postgres
```

The local PostgreSQL URL exposed to host-run backend commands is:

```bash
postgres://affiliate:affiliate@localhost:55432/affiliate_saas?sslmode=disable
```

## Backend Integration Test

Use a disposable local PostgreSQL database and set:

```bash
export AFFILIATE_TEST_DATABASE_URL='postgres://affiliate:affiliate@localhost:55432/affiliate_saas?sslmode=disable'
cd backend
GOCACHE=/tmp/affiliate-saas-go-cache go test ./tests/integration
```

The integration test creates an isolated PostgreSQL schema, applies all `backend/migrations/*.up.sql` files, and exercises the first product slice through HTTP:

```text
workspace -> marketplace program -> product -> offer -> affiliate link -> short redirect -> click event -> analytics query
```

Do not point `AFFILIATE_TEST_DATABASE_URL` at a production or shared database.

## Future Frontend Workflow

Expected once the Next.js app exists:

```bash
npm install
npm run dev
npm test
npm run build
```

Exact package manager and app directory will be documented when the frontend is scaffolded.

## Migrations

SQL migrations live in:

```text
backend/migrations/
```

Expected command once a local database is available:

```bash
migrate -path backend/migrations -database "$DATABASE_URL" up
```

## Future Docker Workflow

Expected once Compose exists:

```bash
docker compose up --build
```

Migrations should run as a separate documented command or one-shot service.

## First Implementation Slice

The active implementation target is:

```text
workspace -> marketplace program -> product -> affiliate link -> short redirect -> click event -> dashboard query
```

Do not begin with AI generation, marketplace integrations, OAuth, or browser automation.
