---
title: ADR-004 Auth Session Strategy
status: active
owner: system-architect
last_verified_at: 2026-07-03
source_of_truth: true
---

# ADR-004: Auth Session Strategy

Status: accepted

Date: 2026-07-03

## Context

Affiliate SaaS is a workspace-based web SaaS. The MVP needs users, workspace membership, and protected product/link/campaign/dashboard data. It does not need enterprise SSO, marketplace OAuth, public API tokens, or mobile clients at the start.

The frontend is expected to be Next.js and the backend Go API will own application data and authorization decisions.

## Decision

Use first-party email/password authentication with secure cookie sessions for the MVP.

Guidelines:

- Store users, password hashes, sessions, and workspace memberships in PostgreSQL.
- Hash passwords with a memory-hard algorithm such as Argon2id.
- Set session cookies as `HttpOnly`, `Secure` in non-local environments, and `SameSite=Lax` by default.
- Authorize every protected operation through workspace membership and role checks.
- Keep OAuth, magic links, SSO, and API tokens out of the MVP unless a concrete customer requirement appears.
- Design session storage so sessions can be revoked server-side.

## Alternatives Considered

- Stateless JWT auth: useful for multi-client APIs, but revocation and browser security are more awkward for this web-first MVP.
- Third-party auth provider from day one: fast to ship, but adds external dependency and may complicate local/demo environments.
- OAuth-only login: good later for creator workflows, but not necessary before marketplace integrations are approved.
- Magic links only: reduces password handling, but depends on email deliverability and is slower for local MVP iteration.

## Consequences

Positive:

- Simple web-first security model.
- Server-side revocation is straightforward.
- Keeps authorization close to workspace data.
- Avoids premature dependency on external identity providers.

Negative:

- The project must implement password reset and credential hardening carefully before real users.
- Cookie/CORS behavior must be tested with the Next.js app and Go API deployment topology.
- Mobile/API-token use cases will need additional auth patterns later.

## Verification

- Protected endpoints reject unauthenticated requests.
- Users can only access resources through authorized workspace membership.
- Session revocation invalidates future requests.
- Security docs define password reset, cookie settings, CSRF posture, and secret handling before beta.

## Links

- Related docs:
  - `docs/architecture/context-map.md`
  - `docs/product/roadmap.md`
- Related issues:
  - None yet.
