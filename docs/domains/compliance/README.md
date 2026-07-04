---
title: Compliance Domain
status: active
owner: system-architect
last_verified_at: 2026-07-04
source_of_truth: true
depends_on:
  - ../../product/vision.md
  - ../../product/roadmap.md
  - ../../architecture/context-map.md
---

# Compliance Domain

## Purpose

Compliance owns policy-sensitive checks for affiliate disclosures, marketplace constraints, claim safety, rights to assets, and risky AI-generated commercial copy.

It is a guardrail domain, not a copy generation domain.

## MVP Responsibilities

- Provide a basic checklist for campaign drafts.
- Flag missing affiliate disclosures.
- Flag unsupported claims about price, discount, availability, health, income, guarantees, or performance.
- Surface marketplace/program policy notes where available.
- Block or warn on obvious risky outputs before a campaign package is considered ready.

## Entities

- Compliance Check: rule or checklist item applied to content or setup.
- Compliance Finding: warning, blocker, or informational note.
- Disclosure Requirement: affiliate disclosure guidance by channel or program.
- Policy Note: marketplace/program-specific rule summary.

## Invariants

- Compliance checks annotate or block; they do not generate campaign copy.
- The system must not claim legal approval.
- AI-generated copy is treated as untrusted until checked and reviewed.
- Product claims must be grounded in user-provided product data or approved source material.
- MVP workflows avoid scraping, browser automation, artificial engagement, and copyright-risk reuse.

## Lifecycle

1. Campaign draft or product setup requests a compliance pass.
2. Compliance evaluates checklist/rules against available context.
3. Findings are attached to the draft or setup item.
4. User resolves blockers or accepts warnings where allowed.
5. Analytics and future reports can identify repeated compliance friction.

## First Vertical Slice

Implemented first:

- run a basic checklist for campaign content;
- flag missing affiliate disclosure as a blocker;
- flag unsupported absolute claims as blockers;
- warn on price or availability claims that must be checked against product truth;
- flag prohibited automation or artificial engagement references.
- attach manual marketplace/program policy notes as findings when campaign products map to configured programs.
- support review and archive workflow for manual policy notes.

This is an MVP guardrail, not legal approval.

## Risks

- Users treating checklist output as legal advice.
- Marketplace policy notes becoming outdated.
- Blocking too aggressively and slowing early users.
- Being too permissive with AI-generated commercial claims.

## Open Questions

- Which compliance findings are blockers versus warnings for beta?
- How will platform policy notes be reviewed and dated?
- Do we need channel-specific disclosure templates before AI generation starts?

## Links

- `docs/product/vision.md`
- `docs/product/roadmap.md`
- `docs/architecture/context-map.md`
