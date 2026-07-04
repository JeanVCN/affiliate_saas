---
title: Analytics Domain
status: active
owner: system-architect
last_verified_at: 2026-07-04
source_of_truth: true
depends_on:
  - ../../product/vision.md
  - ../../product/roadmap.md
  - ../../architecture/context-map.md
---

# Analytics Domain

## Purpose

Analytics owns read models and performance insights built from click events, imported conversions, commissions, products, links, and campaigns.

## MVP Responsibilities

- Aggregate clicks by product, link, campaign, channel, and time range.
- Prepare dashboard queries for top products and links.
- Ingest manual or CSV conversion and commission records after import contracts exist.
- Compare click data with imported sales when attribution fields are available.
- Feed insight back to campaign planning.

## Entities

- Click Aggregate: summarized click metrics.
- Conversion Import Record: user-provided sale or commission row.
- Attribution Match: relationship between imported conversion data and internal product/link/campaign context.
- Dashboard Metric: user-facing summary value.

## Invariants

- Analytics reads event and import data; it does not mutate campaign content or product truth.
- Raw click events remain owned by Link Tracking.
- Imported conversions are user-provided unless an official integration is approved.
- Metrics must be labeled honestly when attribution is approximate.
- Workspace boundaries apply to every query.

## Lifecycle

1. Link Tracking records click events.
2. Analytics aggregates events for dashboards.
3. User imports conversion or commission data.
4. System maps imports to products, links, or campaigns where possible.
5. Dashboard surfaces actionable insights for the next campaign.

## First Vertical Slice

Implemented first:

- click metrics grouped by product or link;
- analytics overview with clicks, imported conversions, gross amount, and commission;
- top products by clicks and imported conversion performance;
- manual conversion import batches;
- manual conversion import rows with optional product/link attribution.

CSV upload, attribution reconciliation, richer dashboard filters, and materialized read models remain later slices.

## Risks

- Presenting approximate attribution as exact.
- Mixing workspaces in aggregate queries.
- Overbuilding reporting before event contracts are stable.
- Letting bot traffic dominate early performance signals.

## Open Questions

- Which beta metrics are essential: clicks, sales, commission, CTR, EPC, conversion rate?
- How should unmatched conversion import rows be surfaced?
- When do we need materialized summaries versus direct SQL queries?

## Links

- `docs/product/vision.md`
- `docs/product/roadmap.md`
- `docs/architecture/context-map.md`
