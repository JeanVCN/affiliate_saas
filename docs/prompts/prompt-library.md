---
title: Prompt Library
status: active
owner: ai-engineer
last_verified_at: 2026-07-03
source_of_truth: true
depends_on:
  - ../domains/campaign/README.md
  - ../domains/compliance/README.md
  - ../quality/evals.md
---

# Prompt Library

This library defines prompt families and governance for AI-assisted campaign generation. It does not yet define final production prompts.

## Prompt Versioning

Each production prompt must have:

- stable ID;
- version;
- owner;
- target channel or workflow;
- input contract;
- output contract;
- compliance requirements;
- eval cases;
- changelog.

Suggested ID format:

```text
campaign.tiktok_shop.ugc_review.v1
campaign.shopee.achadinhos.v1
compliance.affiliate_copy_check.v1
```

## Prompt Families

| Family | Status | Purpose |
|---|---|---|
| `campaign.tiktok_shop.ugc_review` | planned | TikTok Shop review-style campaign draft. |
| `campaign.tiktok_shop.achadinhos` | planned | Short-form discovery/deal angle. |
| `campaign.product_comparison` | planned | Product comparison script/caption. |
| `campaign.demonstration` | planned | Demonstration or how-to angle. |
| `campaign.flash_offer` | planned | Offer-focused copy with strict price/availability guardrails. |
| `compliance.affiliate_copy_check` | planned | Check disclosure, unsupported claims, and policy-sensitive text. |

## Required Inputs

Campaign prompts must receive structured inputs from product/link/campaign domains:

- product name;
- marketplace/program context;
- user-provided benefits and cautions;
- destination channel;
- affiliate disclosure requirement;
- known offer details if user-provided;
- prohibited claims or missing data notes.

Prompts must not infer:

- price;
- discount;
- availability;
- guarantee;
- health result;
- income result;
- official endorsement;
- marketplace approval.

## Required Output Shape

Campaign output should be structured enough for review:

- `title`
- `hook`
- `script_or_caption`
- `cta`
- `affiliate_disclosure`
- `claims_used`
- `assumptions`
- `compliance_notes`

## Safety Rules

- Generated copy is draft content, not final approved content.
- User review is required before publishing.
- Compliance checks run before content is marked ready.
- Do not generate instructions for scraping, artificial engagement, browser automation, or policy evasion.
- Do not use copyrighted assets or third-party creative unless user authorization is recorded.

## Eval Link

Prompt changes must add or update eval cases in `docs/quality/evals.md`.
