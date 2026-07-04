---
title: Database Entity Catalog
status: active
owner: system-architect
last_verified_at: 2026-07-04
source_of_truth: true
depends_on:
  - schema-overview.md
  - ../domains/README.md
---

# Database Entity Catalog

This catalog names the initial PostgreSQL entities before migrations are implemented.

## Identity

| Entity | Table | Purpose |
|---|---|---|
| User | `users` | Human account. |
| Workspace | `workspaces` | Tenant boundary for business data. |
| Membership | `workspace_memberships` | User access to workspace. |
| Session | `sessions` | Server-side cookie session record. |
| OAuth Identity | `oauth_identities` | Optional external identity linked to a user. |
| OAuth State | `oauth_states` | Short-lived OAuth state for provider authorization flow protection. |

## Marketplace

| Entity | Table | Purpose |
|---|---|---|
| Marketplace | `marketplaces` | Supported platform/network family. |
| Program | `programs` | Affiliate or creator program definition. |
| Workspace Program | `workspace_programs` | Workspace-enabled program configuration. |
| Program Policy Note | `program_policy_notes` | Dated policy/compliance notes. |

## Product

| Entity | Table | Purpose |
|---|---|---|
| Product | `products` | Workspace product catalog item. |
| Offer | `offers` | Marketplace/program-specific offer. |
| Product Angle | `product_angles` | Promotional framing for campaign generation. |
| Product Asset | `product_assets` | Optional user-provided or authorized asset metadata. |

## Affiliate And Link Tracking

| Entity | Table | Purpose |
|---|---|---|
| Affiliate Link | `affiliate_links` | User-provided destination URL. |
| Link Variant | `link_variants` | Campaign/channel/angle-specific attribution variant. |
| Short Link | `short_links` | First-party redirect slug. |
| Click Event | `click_events` | Redirect/click tracking event. |

## Campaign

| Entity | Table | Purpose |
|---|---|---|
| Campaign | `campaigns` | Promotional initiative. |
| Channel Package | `channel_packages` | Channel-specific copy package. |
| Publishing Task | `publishing_tasks` | Manual publishing/calendar task. |
| AI Generation Request | `ai_generation_requests` | Deferred until AI provider docs exist. |

## Compliance

| Entity | Table | Purpose |
|---|---|---|
| Compliance Check | `compliance_checks` | Check run against setup or campaign content. |
| Compliance Finding | `compliance_findings` | Warning/blocker/info result. |

## Analytics And Imports

| Entity | Table | Purpose |
|---|---|---|
| Conversion Import | `conversion_imports` | Manual or CSV import batch. |
| Conversion Import Row | `conversion_import_rows` | Imported sale/commission row. |

## Views Or Read Models Later

Likely read models after raw entities exist:

- `analytics_clicks_by_product`
- `analytics_clicks_by_link`
- `analytics_top_products`
- `analytics_campaign_performance`

Do not create these until query needs are proven by the first dashboard implementation.
