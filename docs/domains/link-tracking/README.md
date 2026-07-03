---
title: Link Tracking Domain
status: active
owner: system-architect
last_verified_at: 2026-07-03
source_of_truth: true
depends_on:
  - ../../product/vision.md
  - ../../product/roadmap.md
  - ../../architecture/context-map.md
  - ../../decisions/adr/005-short-link-tracking-strategy.md
---

# Link Tracking Domain

## Purpose

Link Tracking owns first-party short links, redirect behavior, and click events. It connects campaign distribution to measurable product interest.

## MVP Responsibilities

- Generate and resolve short-link slugs.
- Redirect users to affiliate destinations.
- Record click events with workspace, product, link, and campaign context.
- Preserve redirect reliability even when analytics recording has a recoverable failure.
- Provide click data to analytics.

## Entities

- Short Link: first-party redirect URL with a unique slug.
- Click Event: recorded redirect attempt with timestamp and attribution context.
- Visitor Fingerprint Metadata: privacy-preserving request metadata for deduplication and bot heuristics.
- UTM Context: channel and campaign parameters used for reporting.

## Invariants

- Slugs are unique and non-guessable enough for MVP use.
- Redirect destination comes from an active affiliate link.
- Click recording must not block redirect for recoverable storage failures.
- Raw tracking data belongs to the workspace and is read by analytics.
- Privacy-sensitive fields must follow retention and minimization rules before beta.

## Lifecycle

1. User creates or activates an affiliate link.
2. System creates a short link for a link or campaign variant.
3. Visitor opens the short link.
4. System records a click event and redirects to the destination.
5. Analytics aggregates clicks by product, link, campaign, and channel.

## First Vertical Slice

The first slice needs `/r/{slug}` behavior, a click event record, and a dashboard-ready query for clicks by product or link.

## Risks

- Redirect latency hurting user experience and conversion.
- Bot traffic distorting early analytics.
- Collecting more personal data than needed.
- Broken slugs or inactive links leading to dead campaigns.

## Open Questions

- What is the minimum bot filtering for beta?
- What retention period applies to raw click events?
- Should short links have custom slugs in MVP or only generated slugs?

## Links

- `docs/product/vision.md`
- `docs/product/roadmap.md`
- `docs/architecture/context-map.md`
- `docs/decisions/adr/005-short-link-tracking-strategy.md`
