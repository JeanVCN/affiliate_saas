---
title: ADR-005 Short Link Tracking Strategy
status: active
owner: system-architect
last_verified_at: 2026-07-03
source_of_truth: true
---

# ADR-005: Short Link Tracking Strategy

Status: accepted

Date: 2026-07-03

## Context

The MVP thesis depends on connecting products, affiliate links, campaigns, clicks, imported conversions, and insights. The product must work without closed marketplace APIs, so first-party click tracking is a core capability.

Short links must preserve affiliate destination URLs, record useful click events, and avoid behaviors that violate platform terms or user trust.

## Decision

Use first-party short links handled by the Go API.

Initial behavior:

- Generate collision-resistant base62 slugs for affiliate links or campaign-specific link variants.
- Serve redirects from a dedicated route such as `/r/{slug}`.
- Record a click event before redirecting when storage is available.
- Redirect to the configured destination URL without scraping, cloaking, or browser automation.
- Capture minimal analytics fields needed for MVP insights: timestamp, link ID, campaign ID when present, workspace ID, referrer, user agent, IP-derived coarse metadata when allowed, and UTM parameters.
- Store raw IP only if a later privacy/security decision explicitly allows it; prefer hashed or truncated values for analytics.
- Treat failed click recording as non-blocking for the redirect after logging/metrics are available.

## Alternatives Considered

- Third-party short-link provider: faster initially, but weakens ownership of analytics and can become costly or limiting.
- Marketplace-native tracking only: insufficient because the MVP must work across programs and with manual conversion imports.
- Client-side tracking pixel only: misses many click events and does not solve channel-friendly short links.
- Per-channel static UTM links without redirect service: useful supplementary data, but not enough for product-level attribution.

## Consequences

Positive:

- Makes click tracking a first-party product capability.
- Works before marketplace APIs are approved.
- Enables dashboard queries by product, channel, campaign, and link.
- Keeps redirect behavior transparent and policy-aligned.

Negative:

- Redirect endpoint becomes latency-sensitive.
- Bot filtering and fraud heuristics will need later iteration.
- Privacy policy and retention rules must be defined before beta.
- Attribution remains approximate until conversion imports and marketplace reports are normalized.

## Verification

- A short-link request redirects to the configured destination URL.
- A click event is created with the expected link/campaign/workspace relationship.
- Redirect still succeeds when click recording fails in a recoverable way.
- Analytics docs and database docs define retention, bot filtering assumptions, and conversion import relationships before implementation broadens.

## Links

- Related docs:
  - `docs/product/vision.md`
  - `docs/product/roadmap.md`
  - `docs/architecture/context-map.md`
- Related issues:
  - None yet.
