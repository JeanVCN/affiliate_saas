---
title: Secrets Policy
status: active
owner: security-engineer
last_verified_at: 2026-07-03
source_of_truth: true
depends_on:
  - ../decisions/adr/004-auth-session-strategy.md
  - ../domains/identity/README.md
---

# Secrets Policy

Affiliate SaaS must never commit secrets, credentials, customer data, marketplace reports, OAuth tokens, or private exports.

## Secret Types

- Database URLs and passwords.
- Cookie/session signing keys.
- Marketplace credentials and OAuth tokens.
- AI provider API keys.
- S3-compatible storage keys.
- SMTP credentials.
- Customer CSV reports or private commission exports.

## Storage

- Local development uses `.env.local` or shell environment variables.
- Production uses the deployment platform secret manager.
- Tests use disposable local credentials only.
- Documentation uses placeholder values only.

## Git Rules

Never commit:

- `.env`;
- `.env.local`;
- marketplace reports;
- customer exports;
- OAuth token dumps;
- production database snapshots;
- private keys or certificates.

Use `.env.example` for required variable names and safe sample values once implementation starts.

## Rotation

Rotate immediately if a secret is exposed in:

- Git history;
- logs;
- screenshots;
- chat;
- issue tracker;
- CI output.

## Application Requirements

- Session cookies must be signed or backed by unguessable session IDs.
- Password hashes use Argon2id or another approved memory-hard algorithm.
- Marketplace tokens are not part of MVP storage unless an official integration is approved.
- AI provider keys remain server-side only.

## Agent Rules

- Agents must not request broad write access to secret stores.
- MCP/tool access is deny-by-default and read-only unless explicitly approved.
- If a task needs a real secret, ask the user to configure it outside Git.
