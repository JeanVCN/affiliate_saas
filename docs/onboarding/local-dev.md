---
title: Local Development
status: active
owner: infrastructure-engineer
last_verified_at: 2026-07-06
source_of_truth: true
depends_on:
  - ../deployment/docker.md
  - ../quality/quality-gates.md
---

# Local Development

The repository now includes a Go backend and an initial Next.js frontend. These commands define the current verification path and the intended local workflow as implementation grows.

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

## Local API Container

Build and start the API with PostgreSQL:

```bash
docker compose up -d postgres api
```

Run migrations against the Compose PostgreSQL service:

```bash
docker compose run --rm migrate
```

The API is exposed on:

```text
http://localhost:18080
```

Health check:

```bash
curl http://localhost:18080/healthz
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
signup -> session cookie -> workspace RBAC -> marketplace program -> product -> campaign -> channel package -> offer -> affiliate link -> short redirect -> click event -> analytics query
```

Do not point `AFFILIATE_TEST_DATABASE_URL` at a production or shared database.

## Frontend Workflow

The Next.js app lives in `frontend/`. It proxies `/backend/*` to the local API container so browser requests can use the backend session cookie without adding CORS requirements to the Go API during the first frontend slice.

```bash
cd frontend
npm install
npm run dev
npm run lint
npm run build
```

Default local values:

```bash
NEXT_PUBLIC_API_BASE_URL=/backend
API_PROXY_ORIGIN=http://localhost:18080
```

The frontend dev server runs on:

```text
http://localhost:13000
```

Start the backend stack first from the repository root:

```bash
docker compose up -d postgres api
docker compose run --rm migrate
```

Current frontend slice:

```text
signup/login -> session restore -> marketplace program -> product -> offer -> affiliate link -> short link -> analytics overview
```

The frontend normalizes empty program lists because the current backend can return `null` for an empty list response.

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

Full local stack rebuild:

```bash
docker compose up --build
```

Migrations should run as a separate documented command or one-shot service.

## First Implementation Slice

The active frontend implementation target is:

```text
workspace -> marketplace program -> product -> offer -> affiliate link -> short link -> analytics overview
```

Do not begin with AI generation, marketplace integrations, OAuth, or browser automation.
