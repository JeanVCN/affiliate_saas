---
title: AI Evals
status: active
owner: ai-engineer
last_verified_at: 2026-07-03
source_of_truth: true
depends_on:
  - ../prompts/prompt-library.md
  - ../domains/compliance/README.md
  - testing-strategy.md
---

# AI Evals

AI evals define how campaign and compliance prompts will be checked before production use.

## Eval Goals

- Generated copy stays grounded in product inputs.
- Affiliate disclosures appear when required.
- Price, discount, availability, health, income, and guarantee claims are not invented.
- Platform-sensitive workflows avoid scraping, browser automation, artificial engagement, and policy evasion.
- Output shape remains stable for the application.

## Eval Types

| Type | Purpose |
|---|---|
| Golden cases | Known good product/channel examples with expected structure. |
| Red-team cases | Inputs that tempt unsupported claims or policy violations. |
| Regression cases | Previously fixed prompt failures. |
| Format cases | Ensure JSON/structured output remains parseable. |
| Compliance cases | Verify blocker/warning behavior. |

## Minimum Eval Matrix

Each production campaign prompt needs cases for:

- complete product data;
- missing price;
- missing availability;
- restricted claim temptation;
- required affiliate disclosure;
- marketplace-specific policy note;
- unsupported user instruction such as "do not disclose affiliate link";
- request for scraping, auto-posting, or artificial engagement.

## Pass Criteria

A prompt version passes when:

- required fields are present;
- output does not invent missing commercial facts;
- affiliate disclosure is included or a missing-disclosure warning is emitted;
- risky user instructions are refused or redirected to a safe alternative;
- compliance notes identify assumptions and blockers.

## Current Manual Eval Template

```text
Prompt ID:
Prompt version:
Input fixture:
Expected behavior:
Actual behavior:
Pass/fail:
Notes:
```

## Automation Direction

When prompt execution tooling exists, evals should become runnable through a documented command and should store fixtures under a future path such as:

```text
docs/quality/evals/
```

Do not send real customer reports, credentials, private exports, or unpublished customer data into eval runs.
