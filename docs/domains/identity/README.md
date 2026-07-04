---
title: Identity Domain
status: active
owner: system-architect
last_verified_at: 2026-07-04
source_of_truth: true
depends_on:
  - ../../product/roadmap.md
  - ../../architecture/context-map.md
  - ../../decisions/adr/004-auth-session-strategy.md
---

# Identity Domain

## Purpose

Identity owns users, authentication, workspaces, membership, and authorization boundaries. It answers who is using the product and which workspace data they can access.

## MVP Responsibilities

- Register and authenticate users.
- Create the initial workspace for a user.
- Store workspace memberships and roles.
- Authorize product, marketplace, link, campaign, import, and analytics operations by workspace.
- Support server-side session revocation.
- Hash passwords with Argon2id.
- Prepare OAuth identity/state storage without storing provider tokens before integration approval.

## Entities

- User: human account with email, password hash, profile fields, and lifecycle status.
- Workspace: tenant boundary for products, affiliate programs, links, campaigns, imports, and analytics.
- Membership: relation between user and workspace.
- Role: permission tier inside a workspace.
- Session: server-side login session tied to a user and revocable by the backend.
- OAuth Identity: optional provider identity linked to a user after provider login is configured.
- OAuth State: short-lived authorization state for OAuth CSRF protection.

## Invariants

- Every business object belongs to exactly one workspace.
- Protected operations require an authenticated user and an authorized workspace membership.
- A session identifies a user, not a workspace; workspace context is selected or inferred per request.
- Password hashes are never reversible and never leave the backend.
- Workspace membership checks happen server-side, not only in the frontend.
- Workspace roles are ordered as owner, admin, and member.
- OAuth provider access/refresh tokens are not stored until a provider-specific token policy exists.

## Lifecycle

1. User signs up or is invited.
2. User verifies access through login.
3. User creates or joins a workspace.
4. User performs workspace-scoped operations.
5. User logs out, is removed, or has sessions revoked.

## First Vertical Slice

The first slice can start with one user and one workspace per account, but the data model must not block future team memberships.

## Risks

- Leaking data across workspaces through missing authorization checks.
- Treating frontend route protection as sufficient security.
- Adding marketplace OAuth before the core workspace model is stable.
- Under-specifying password reset and session revocation before beta.

## Open Questions

- Which roles are needed before beta: owner only, owner/member, or owner/admin/member?
- Will early private beta use invite-only signup?
- When should email verification become mandatory?

## Links

- `docs/product/roadmap.md`
- `docs/architecture/context-map.md`
- `docs/decisions/adr/004-auth-session-strategy.md`
