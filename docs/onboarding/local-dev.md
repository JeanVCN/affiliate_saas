---
title: Local Development
status: active
owner: infrastructure-engineer
last_verified_at: 2026-07-03
source_of_truth: true
depends_on:
  - ../deployment/docker.md
  - ../quality/quality-gates.md
---

# Local Development

The repository is currently documentation-only. These commands define the current verification path and the intended local workflow once implementation begins.

## Current Setup

From the repository root:

```bash
node -e "for (const f of ['.docs-index.json','.context-manifest.json']) JSON.parse(require('fs').readFileSync(f,'utf8'));"
```

This is the active local verification command until application code exists.

## Future Backend Workflow

Expected once `backend/` exists:

```bash
cd backend
go test ./...
```

Expected services:

- PostgreSQL for persistence.
- Redis only after async jobs are introduced.

## Future Frontend Workflow

Expected once the Next.js app exists:

```bash
npm install
npm run dev
npm test
npm run build
```

Exact package manager and app directory will be documented when the frontend is scaffolded.

## Future Docker Workflow

Expected once Compose exists:

```bash
docker compose up --build
```

Migrations should run as a separate documented command or one-shot service.

## First Implementation Slice

When coding starts, implement and verify:

```text
workspace -> marketplace program -> product -> affiliate link -> short redirect -> click event -> dashboard query
```

Do not begin with AI generation, marketplace integrations, OAuth, or browser automation.
