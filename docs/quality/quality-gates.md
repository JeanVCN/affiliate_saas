---
title: Quality Gates
status: active
owner: qa-engineer
last_verified_at: 2026-07-03
source_of_truth: true
depends_on:
  - testing-strategy.md
  - ../workflows/development/documentation-phases.md
---

# Quality Gates

Quality gates define the minimum checks before merging or presenting work.

## Verification Budget

During active implementation, prefer focused checks that match the files or package being changed. Avoid repeating full test suites, broad `git diff`, or detailed `git status` output after every small edit.

Run complete gates when a stage is ready for review, before committing, before presenting work as complete, or earlier when the change is high risk.

High-risk changes include auth, RBAC, sessions, migrations, billing, data deletion, external integrations, secrets handling, and tracking/privacy behavior.

Git usage:

- use targeted diffs while implementing, such as a specific file or package;
- use broad `git status` and full diffs at checkpoints, review, or commit time;
- never include secrets, OAuth tokens, customer data, or private exports in diffs or commits;
- do not commit until the relevant gate has passed or any skipped check is explicitly called out.

## Documentation-Only Gate

Run:

```bash
node -e "for (const f of ['.docs-index.json','.context-manifest.json']) JSON.parse(require('fs').readFileSync(f,'utf8'));"
```

Check:

- changed docs are linked from `docs/INDEX.md` when durable;
- source-of-truth changes update `docs/_meta/source-of-truth-map.md` when needed;
- ADRs are created or updated for architecture decisions;
- stale docs are updated or explicitly left as historical.

## Backend Gate

Active now that `backend/` exists:

```bash
cd backend
GOCACHE=/tmp/affiliate-saas-go-cache go test ./...
```

Required:

- tests pass;
- handlers keep Gin at the delivery boundary;
- domain modules do not import `gin.Context`;
- migrations are reviewed with code that depends on them;
- auth and workspace authorization are tested for protected endpoints.

## Database Gate

Active once migrations exist:

```bash
migrate -path backend/migrations -database "$DATABASE_URL" up
```

Required:

- fresh database migrates from zero;
- migration filenames are append-only once merged;
- schema docs are updated for entity or relationship changes.

## Frontend Gate

Active once frontend exists:

```bash
npm test
npm run build
```

Required:

- user-facing errors are visible;
- protected routes handle unauthenticated state;
- core flows work on desktop and mobile viewport sizes;
- API contract changes are reflected in UI data handling.

## Security Gate

Required for any code or config change:

- no secrets in Git;
- no marketplace credentials, OAuth tokens, customer reports, or private exports committed;
- auth/session changes update security docs;
- tracking changes consider data minimization and retention.
