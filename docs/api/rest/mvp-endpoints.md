---
title: MVP REST Endpoints
status: active
owner: system-architect
last_verified_at: 2026-07-04
source_of_truth: true
depends_on:
  - ../README.md
  - ../../domains/README.md
---

# MVP REST Endpoints

This document lists the initial endpoint contract inventory. It is not a full OpenAPI spec yet.

## Public

| Method | Path | Domain | Purpose |
|---|---|---|---|
| `GET` | `/healthz` | Platform | Process health check. |
| `GET` | `/r/{slug}` | Link Tracking | Resolve short link, record click event, and redirect to destination. |

## Auth

| Method | Path | Domain | Purpose |
|---|---|---|---|
| `POST` | `/api/v1/auth/signup` | Identity | Create user, initial workspace, and session. |
| `POST` | `/api/v1/auth/login` | Identity | Create session for an existing user. |
| `POST` | `/api/v1/auth/logout` | Identity | Revoke current session. |
| `GET` | `/api/v1/auth/me` | Identity | Return current user and workspace memberships. |
| `GET` | `/api/v1/auth/oauth/{provider}/start` | Identity | OAuth boundary endpoint; returns unavailable until provider configuration is approved. |

Auth rules:

- `signup`, `login`, public redirect `/r/{slug}`, and OAuth start are public.
- `logout` and `me` require the session cookie.
- Workspace-scoped routes require an authenticated user with active workspace membership.
- Passwords are hashed with Argon2id.
- Session cookies are `HttpOnly`, `SameSite=Lax`, and `Secure` outside local/test environments.

## Workspaces

| Method | Path | Domain | Purpose |
|---|---|---|---|
| `GET` | `/api/v1/workspaces` | Identity | List workspaces visible to current user. |
| `POST` | `/api/v1/workspaces` | Identity | Create workspace. |
| `GET` | `/api/v1/workspaces/{workspace_id}` | Identity | Read workspace summary. |

## Marketplace Programs

| Method | Path | Domain | Purpose |
|---|---|---|---|
| `GET` | `/api/v1/marketplaces` | Marketplace | List supported marketplace/program definitions. |
| `GET` | `/api/v1/workspaces/{workspace_id}/programs` | Marketplace | List workspace-enabled programs. |
| `POST` | `/api/v1/workspaces/{workspace_id}/programs` | Marketplace | Enable/configure a manual affiliate program. |
| `PATCH` | `/api/v1/workspaces/{workspace_id}/programs/{program_id}` | Marketplace | Update workspace program settings or status. |

## Products And Offers

| Method | Path | Domain | Purpose |
|---|---|---|---|
| `GET` | `/api/v1/workspaces/{workspace_id}/products` | Product | List products. |
| `POST` | `/api/v1/workspaces/{workspace_id}/products` | Product | Create product. |
| `GET` | `/api/v1/workspaces/{workspace_id}/products/{product_id}` | Product | Read product detail. |
| `PATCH` | `/api/v1/workspaces/{workspace_id}/products/{product_id}` | Product | Update product fields. |
| `POST` | `/api/v1/workspaces/{workspace_id}/products/{product_id}/offers` | Product | Create marketplace/program offer for a product. |
| `PATCH` | `/api/v1/workspaces/{workspace_id}/products/{product_id}/offers/{offer_id}` | Product | Update offer metadata or status. |

## Affiliate Links

| Method | Path | Domain | Purpose |
|---|---|---|---|
| `GET` | `/api/v1/workspaces/{workspace_id}/links` | Affiliate | List affiliate links. |
| `POST` | `/api/v1/workspaces/{workspace_id}/links` | Affiliate | Create affiliate destination link for a product/offer. |
| `GET` | `/api/v1/workspaces/{workspace_id}/links/{link_id}` | Affiliate | Read link detail and short-link variants. |
| `PATCH` | `/api/v1/workspaces/{workspace_id}/links/{link_id}` | Affiliate | Update link metadata or status. |
| `POST` | `/api/v1/workspaces/{workspace_id}/links/{link_id}/short-links` | Link Tracking | Create a short-link slug for a link or campaign variant. |

## Campaigns

| Method | Path | Domain | Purpose |
|---|---|---|---|
| `GET` | `/api/v1/workspaces/{workspace_id}/campaigns` | Campaign | List campaigns. |
| `POST` | `/api/v1/workspaces/{workspace_id}/campaigns` | Campaign | Create campaign draft manually. |
| `GET` | `/api/v1/workspaces/{workspace_id}/campaigns/{campaign_id}` | Campaign | Read campaign detail. |
| `PATCH` | `/api/v1/workspaces/{workspace_id}/campaigns/{campaign_id}` | Campaign | Update campaign draft/status. |
| `POST` | `/api/v1/workspaces/{workspace_id}/campaigns/{campaign_id}/channel-packages` | Campaign | Add channel-specific package. |
| `GET` | `/api/v1/workspaces/{workspace_id}/campaigns/{campaign_id}/publishing-tasks` | Campaign | List manual publishing tasks. |
| `POST` | `/api/v1/workspaces/{workspace_id}/campaigns/{campaign_id}/publishing-tasks` | Campaign | Create a manual publishing task. |
| `PATCH` | `/api/v1/workspaces/{workspace_id}/campaigns/{campaign_id}/publishing-tasks/{task_id}` | Campaign | Update publishing task status or schedule. |
| `POST` | `/api/v1/workspaces/{workspace_id}/campaigns/{campaign_id}/compliance-checks` | Compliance | Run compliance checks for campaign content. |

## Analytics

| Method | Path | Domain | Purpose |
|---|---|---|---|
| `GET` | `/api/v1/workspaces/{workspace_id}/analytics/overview` | Analytics | Dashboard summary for clicks, sales imports, and commissions when available. |
| `GET` | `/api/v1/workspaces/{workspace_id}/analytics/clicks` | Analytics | Click metrics grouped by product, link, campaign, or channel. |
| `GET` | `/api/v1/workspaces/{workspace_id}/analytics/top-products` | Analytics | Top products by clicks and imported conversions. |

## Conversion Imports

| Method | Path | Domain | Purpose |
|---|---|---|---|
| `POST` | `/api/v1/workspaces/{workspace_id}/conversion-imports` | Analytics | Create a manual import batch. |
| `POST` | `/api/v1/workspaces/{workspace_id}/conversion-imports/{import_id}/rows` | Analytics | Add manual conversion rows before CSV upload exists. |
| `GET` | `/api/v1/workspaces/{workspace_id}/conversion-imports/{import_id}` | Analytics | Read import status and row summary. |

## Deferred From MVP Contract

- Automated marketplace scraping.
- Browser automation for publishing or engagement.
- Official marketplace API integration before approval.
- Public API tokens.
- OAuth marketplace connections.
- AI generation endpoint details before prompt/provider docs exist.
