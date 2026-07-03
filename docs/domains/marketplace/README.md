---
title: Marketplace Domain
status: active
owner: system-architect
last_verified_at: 2026-07-03
source_of_truth: true
depends_on:
  - ../../product/vision.md
  - ../../product/roadmap.md
  - ../../architecture/context-map.md
---

# Marketplace Domain

## Purpose

Marketplace owns the internal representation of affiliate programs and commerce channels such as TikTok Shop, Shopee, Mercado Livre, Amazon, and future approved integrations.

For the MVP, it must work without closed marketplace APIs by supporting manual program setup and user-authorized imports.

## MVP Responsibilities

- Maintain supported marketplace/program definitions.
- Let a workspace configure marketplace affiliate program participation manually.
- Normalize marketplace names, program types, regions, and link rules enough for products and campaigns.
- Provide marketplace context to product, affiliate link, campaign, compliance, and analytics domains.

## Entities

- Marketplace: commerce platform or affiliate network family.
- Program: affiliate program or creator program within a marketplace.
- Workspace Program: workspace-specific enrollment/configuration.
- Program Policy Note: human-readable constraints for links, disclosures, claims, and imports.

## Invariants

- Marketplace records describe external programs; they do not own product truth.
- Workspace program configuration belongs to one workspace.
- MVP setup cannot require API approval from a marketplace.
- Program policy notes guide compliance but do not replace legal review or platform terms.

## Lifecycle

1. System defines an initial marketplace/program catalog.
2. Workspace enables a program manually.
3. Products and affiliate links reference the enabled program.
4. Campaign generation uses marketplace context and policy notes.
5. Imports and reports map back to configured programs when available.

## First Vertical Slice

Start with manually configured programs for TikTok Shop, Shopee, Mercado Livre, and Amazon. Each product and affiliate link should reference the workspace program it belongs to.

## Risks

- Accidentally implying official API integration before approval exists.
- Encoding platform-specific assumptions too deeply in product or campaign modules.
- Letting marketplace policy notes become stale.
- Mixing marketplace catalog data with workspace-owned product data.

## Open Questions

- Which marketplace fields are mandatory for the first private beta?
- Do we need region/country-specific program variants immediately?
- Where will platform policy research be versioned before beta?

## Links

- `docs/product/vision.md`
- `docs/product/roadmap.md`
- `docs/architecture/context-map.md`
