---
title: Product Domain
status: active
owner: system-architect
last_verified_at: 2026-07-03
source_of_truth: true
depends_on:
  - ../../product/vision.md
  - ../../product/roadmap.md
  - ../../architecture/context-map.md
---

# Product Domain

## Purpose

Product owns the workspace catalog of products and offers that affiliates want to promote. It is the commercial center of the MVP workflow.

## MVP Responsibilities

- Register products manually.
- Store marketplace/program association.
- Store product positioning fields used by campaign generation.
- Track offer metadata needed for user decisions.
- Provide product context to affiliate links, campaigns, compliance, and analytics.

## Entities

- Product: item or service the workspace may promote.
- Offer: marketplace-specific commercial offer for a product.
- Product Asset: approved image, note, or reference material owned or authorized by the user.
- Product Angle: positioning idea such as comparison, demonstration, pain point, review, or flash offer.

## Invariants

- Product truth belongs to the workspace, not to generated campaigns.
- A product can have multiple offers across marketplaces or programs.
- Price, discount, and availability are user-provided or imported unless an official integration is approved.
- Campaign copy must not invent product claims that are absent from product data or user input.
- Assets must be user-provided, authorized, or otherwise safe to use.

## Lifecycle

1. User creates a product with basic name, category, and marketplace context.
2. User adds one or more offers and affiliate links.
3. User enriches the product with benefits, cautions, audience, and content angles.
4. Campaigns are generated from product data and selected channel templates.
5. Analytics report clicks and imported conversions back to the product.

## First Vertical Slice

The first slice needs enough product data to create an affiliate link, generate a campaign draft, redirect a short link, and query clicks by product.

## Risks

- Treating marketplace pages as scrape targets.
- Using AI to fabricate pricing, guarantees, health claims, or performance claims.
- Storing product assets without rights metadata.
- Letting product data become duplicated inside campaign records.

## Open Questions

- Which product fields are mandatory for campaign generation v1?
- Should product assets be in MVP or deferred until storage docs exist?
- Do we represent variants/SKUs before the first beta?

## Links

- `docs/product/vision.md`
- `docs/product/roadmap.md`
- `docs/architecture/context-map.md`
