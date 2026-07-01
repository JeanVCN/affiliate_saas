# AGENTS.md

This file is the always-on operating guide for AI agents working in this repository. Keep it short. Put detailed process, architecture, and domain knowledge under `docs/` and retrieve it on demand.

## Project

Affiliate SaaS is an AI-assisted affiliate commerce platform for creators and operators who manage products, links, campaigns, publishing tasks, and performance across TikTok Shop, Shopee, Mercado Livre, Amazon, and other affiliate programs.

Core MVP thesis:

- Product + affiliate link + AI campaign generation + click tracking + CSV/manual conversion imports.
- Do not depend on closed marketplace APIs for the MVP.
- Avoid scraping, browser automation, artificial engagement, and copyright-risk workflows.
- Prefer official APIs, user-authorized imports, and manual-package workflows until integrations are approved.

## Source Of Truth

Read `docs/INDEX.md` first when looking for project context.

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

