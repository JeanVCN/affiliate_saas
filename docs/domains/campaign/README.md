---
title: Campaign Domain
status: active
owner: system-architect
last_verified_at: 2026-07-04
source_of_truth: true
depends_on:
  - ../../product/vision.md
  - ../../product/roadmap.md
  - ../../architecture/context-map.md
---

# Campaign Domain

## Purpose

Campaign owns planned promotional work: campaign drafts, channel packages, content angles, publishing tasks, and AI-assisted copy generation requests.

## MVP Responsibilities

- Create campaign drafts for products and links.
- Generate channel-specific copy packages for TikTok, Reels, Shorts, WhatsApp, Telegram, and Blog.
- Support TikTok Shop-style templates such as UGC review, achadinhos, comparison, demonstration, and flash offer.
- Track simple publishing tasks or calendar items.
- Send generated copy through compliance checks before it is treated as usable.

## Entities

- Campaign: promotional initiative for one or more products.
- Campaign Draft: editable campaign content before publishing.
- Channel Package: copy/assets tailored to a channel.
- Content Angle: selected promotional framing.
- Publishing Task: planned manual publishing action.
- AI Generation Request: request metadata for generated campaign content.

## Invariants

- Campaigns reference product truth; they do not own or mutate it.
- AI-generated copy must be reviewable and editable by the user.
- Generated copy cannot invent product claims, discounts, or availability.
- Publishing remains manual in the MVP unless an official integration is approved.
- Compliance checks can block or warn on risky outputs.

## Lifecycle

1. User selects product, affiliate link, channel, and angle.
2. System creates a draft manually or through AI generation.
3. Compliance annotates warnings or blockers.
4. User edits and approves the package.
5. User executes publishing manually and marks tasks complete.
6. Analytics later reports performance back to the campaign.

## First Vertical Slice

Implemented first:

- list campaign drafts in a workspace;
- create a manual campaign draft tied to an optional product;
- read campaign detail with channel packages;
- update campaign name or status;
- add draft channel packages manually.
- create, list, schedule, and complete manual publishing tasks.

AI generation remains a later slice. Publishing tasks are manual only; no browser automation or automated posting is allowed in the MVP.

## Risks

- Starting with AI before product/link truth is stable.
- Creating copy that violates affiliate disclosure or platform policy.
- Blurring manual publishing tasks with prohibited browser automation.
- Making campaign records duplicate stale product data.

## Open Questions

- Which channel package should ship first for beta?
- What prompt/version metadata must be stored once AI generation starts?
- Should campaign approval be a formal state before publishing tasks exist?

## Links

- `docs/product/vision.md`
- `docs/product/roadmap.md`
- `docs/architecture/context-map.md`
