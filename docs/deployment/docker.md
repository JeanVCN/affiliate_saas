---
title: Docker Deployment Baseline
status: active
owner: infrastructure-engineer
last_verified_at: 2026-07-06
source_of_truth: true
depends_on:
  - ../architecture/system-overview.md
  - ../database/migrations.md
---

# Docker Deployment Baseline

Affiliate SaaS starts Docker-first. Kubernetes is deferred until operational scale requires it.

## Target Local Services

Initial local stack once code exists:

- Go API.
- PostgreSQL.
- Redis, only when async jobs are introduced.
- Next.js frontend.
- Optional object storage emulator when product assets are implemented.

## Expected Compose Shape

Current `docker-compose.yml` provides:

- `postgres`
- `api`
- `migrate` as a one-shot migration command

Future local Compose growth should add:

- `frontend`
- `redis` when needed

## Environment

Use safe local defaults in `.env.example` once code exists:

- `DATABASE_URL`
- `AFFILIATE_TEST_DATABASE_URL`
- `SESSION_SECRET`
- `APP_ENV`
- `API_BASE_URL`
- `FRONTEND_BASE_URL`
- `REDIS_URL` when needed

No real secrets belong in Compose files.

## Migration Policy

Migrations run outside API startup:

```bash
migrate -path backend/migrations -database "$DATABASE_URL" up
```

Local Compose may expose this as a one-shot service or documented command.

Current local one-shot command:

```bash
docker compose run --rm migrate
```

Current local API startup:

```bash
docker compose up -d postgres api
```

## Production Direction

Start with:

- one API container;
- one frontend deployment;
- managed or containerized PostgreSQL depending on host;
- Redis only after first async job;
- S3-compatible storage after product assets are in scope.

## Deployment Gate

Before production:

- health endpoint exists;
- migrations run from zero;
- secrets are configured outside Git;
- logs avoid credentials and customer data;
- HTTPS terminates before browser-facing traffic;
- session cookies use secure production settings.
