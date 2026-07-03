---
title: Affiliate Domain
status: active
owner: system-architect
last_verified_at: 2026-07-03
source_of_truth: true
depends_on:
  - ../../product/vision.md
  - ../../product/roadmap.md
  - ../../architecture/context-map.md
---

# Affiliate Domain

## Purpose

Affiliate owns the workspace's affiliate link registry and the relationship between products, marketplace programs, campaigns, and destination URLs.

## MVP Responsibilities

- Store affiliate destination URLs provided by users.
- Associate affiliate links with workspace programs and products.
- Support campaign-specific link variants.
- Preserve destination URLs without scraping, cloaking, or artificial engagement.
- Provide link metadata to short-link tracking and analytics.

## Entities

- Affiliate Link: user-provided destination URL for a product/program.
- Link Variant: campaign, channel, or angle-specific variant.
- Tracking Parameters: UTM and internal attribution fields appended or recorded for analysis.
- Link Status: draft, active, paused, archived, or invalid.

## Invariants

- Every affiliate link belongs to one workspace.
- Every active affiliate link references a product and workspace program.
- The system must not alter links in ways that break affiliate attribution or mislead users.
- Short links redirect to the configured destination URL.
- Link validation must not require automated marketplace browsing.

## Lifecycle

1. User creates an affiliate link for a product and program.
2. System validates basic URL shape and stores metadata.
3. User creates campaign/channel variants when needed.
4. Link variants are used by short links and campaign packages.
5. User pauses or archives outdated links.

## First Vertical Slice

The first slice needs one active affiliate link per product, one generated short-link slug, and enough metadata for clicks to be grouped by product and campaign.

## Risks

- Breaking affiliate tracking by rewriting URLs incorrectly.
- Allowing unsafe redirect destinations.
- Confusing raw affiliate links with first-party short links.
- Depending on marketplace APIs before MVP validation.

## Open Questions

- Which URL validation and allowlist rules are required before beta?
- Should users be allowed to store multiple affiliate links for the same product/channel?
- How should expired or unavailable offers be represented before official integrations exist?

## Links

- `docs/product/vision.md`
- `docs/product/roadmap.md`
- `docs/architecture/context-map.md`
- `docs/decisions/adr/005-short-link-tracking-strategy.md`
