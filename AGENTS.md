# AGENTS.md

This file is the always-on operating guide for AI agents working in this repository. Keep it short; load detailed docs only when the task needs them.

## Project

Affiliate SaaS is an AI-assisted affiliate commerce platform for creators and operators who manage products, links, campaigns, publishing tasks, and performance across TikTok Shop, Shopee, Mercado Livre, Amazon, and other affiliate programs.

Core MVP thesis:

- Product + affiliate link + AI campaign generation + click tracking + CSV/manual conversion imports.
- Do not depend on closed marketplace APIs for the MVP.
- Avoid scraping, browser automation, artificial engagement, and copyright-risk workflows.
- Prefer official APIs, user-authorized imports, and manual-package workflows until integrations are approved.

## Source Of Truth

Use `docs/INDEX.md` as the routing map when project context is needed. Do not open broad docs by default for narrow code or command tasks.

Precedence when sources conflict:

1. Executable code, schemas, migrations, and tests.
2. Accepted ADRs in `docs/decisions/adr/`.
3. Source-of-truth documents listed in `docs/_meta/source-of-truth-map.md`.
4. Derived docs.
5. Agent memory/session summaries.
6. Old conversations.

## Development Workflow

- Explore before editing when a task touches multiple files, architecture, product scope, data model, or integrations.
- Make small, scoped changes and keep unrelated churn out.
- Use `rg`/symbol search before broad file reads.
- Prefer docs, indexes, and focused snippets over opening large files.
- Update docs when changing architecture, domain behavior, API contracts, database schema, prompts, deployment, or workflows.
- Record architecturally significant decisions as ADRs.
- End work with verification evidence when possible.

## Context Budget

- Keep always-on context minimal; durable detail belongs under `docs/` and is loaded on demand.
- Prefer one bounded vertical slice per turn, then checkpoint the next step.
- For multi-file work, identify the small file set first, edit, verify, and summarize evidence.
- Favor domain modules with low file fan-out. Split files by responsibility when they approach 400 lines or mix concerns; do not create layers only for ceremony.
- Use a verification budget: run focused checks while implementing; save full test suites, broad diffs, and detailed Git status for checkpoints and pre-commit validation.
- Keep route handlers as named functions or methods, not inline lambdas.

## Stack Preference

- Backend: Go.
- Database: PostgreSQL.
- Queue: Redis Streams for MVP; RabbitMQ/Temporal/River only if workflow complexity justifies it.
- Storage: S3-compatible object storage.
- Frontend: Next.js + React.
- Deploy: Docker first; Kubernetes only when scale requires it.

## Safety

- Never commit secrets, tokens, marketplace credentials, OAuth tokens, customer data, or private reports.
- Avoid scraping or automation that violates platform terms.
- Treat AI-generated commercial copy as policy-sensitive: check claims, affiliate disclosure, price/offer accuracy, and rights to images/assets.
- MCP/tool access must be deny-by-default and read-only unless explicitly approved.

## Documentation Rules

- Keep always-on instructions concise.
- Do not duplicate source-of-truth content across docs; link to it.
- Add `status`, `owner`, and `last_verified_at` frontmatter to durable docs when practical.
- Archive stale docs instead of silently deleting context.
