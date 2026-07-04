---
title: ADR-006 Token-Efficient Assisted Development
status: active
owner: system-architect
last_verified_at: 2026-07-04
source_of_truth: true
---

# ADR-006: Token-Efficient Assisted Development

Status: accepted

Date: 2026-07-04

## Context

Affiliate SaaS is being developed with AI assistance as a first-class workflow. As the backend grows, every always-loaded instruction, broad documentation read, and cross-layer change increases context cost. The project still needs clear domain ownership, security discipline, and an evolution path beyond a throwaway MVP.

Recent vendor guidance for coding agents points to the same practical pattern: keep repository instructions concise, load context on demand, scope work tightly, and ask agents to show verification evidence. Emerging empirical research on repository-level agent files also suggests that unnecessary requirements can increase exploration and inference cost.

## Decision

Use a token-efficient modular monolith:

- Keep domain ownership under `backend/internal/modules/*`.
- Prefer vertical, domain-scoped changes over broad layer-by-layer edits.
- Keep always-loaded instructions minimal. `AGENTS.md` is a routing guide, not a knowledge base.
- Use `docs/INDEX.md` only as a routing map when the task needs project context.
- Keep detailed architecture, domain, API, database, security, and workflow knowledge under focused docs.
- Avoid forced file fan-out. A module may start with the files it actually needs.
- Keep handler/service/repository separation as the default shape for real behavior, but split by responsibility when size or mixed concerns justify it.
- Keep Go files at or below 400 lines unless there is an explicit split plan.
- Route registration maps URLs to named handler methods or functions; business logic does not live in inline lambdas.
- Use a verification budget: run focused checks while implementing. Run full validation only when the user asks to validate, review, close a stage, commit, or says it is the last step before commit.
- Avoid repeated broad `git diff`/`git status` output during active implementation; use targeted Git inspection unless a checkpoint, review, or commit is being prepared.
- End meaningful changes with focused verification evidence.

This means `handler.go`, `service.go`, `repository.go`, and `models.go` remain familiar anchors, but they are not mandatory ceremony for every small module. When a domain grows, prefer responsibility-based splits such as `session_repository.go`, `oauth_repository.go`, `workspace_handler.go`, or `click_service.go`.

## Alternatives Considered

- Strict layer files for every module from day one: predictable, but it creates more files per change and raises token cost for small features.
- Single package or flat MVP store: cheap at the start, but makes domain boundaries unclear and creates a later restructuring cost.
- Full Clean Architecture per module: strong separation, but too much boilerplate and context overhead before domain behavior stabilizes.
- Large always-on agent manual: helpful for rare tasks, but wastes context on common tasks and reduces instruction adherence.

## Consequences

Positive:

- Reduces routine context loading and file reads.
- Reduces repeated command/log output during active implementation.
- Keeps the codebase navigable for humans and agents.
- Preserves a path from MVP to durable product architecture.
- Encourages smaller, easier-to-review backend slices.

Negative:

- Requires discipline to avoid underspecifying module boundaries.
- Some modules will not look perfectly symmetrical.
- Agents must decide when to load deeper docs instead of relying on always-on context.

## Verification

- `AGENTS.md` remains concise and points to docs instead of duplicating them.
- New tasks start with targeted search or the smallest relevant doc set.
- Architecture changes update ADRs instead of expanding always-on instructions.
- Backend changes do not require touching many files only to satisfy ceremony.
- During implementation, agents use package/file-scoped checks when sufficient.
- At user-declared checkpoints and before commits, agents run the relevant complete gate and summarize evidence.
- No regular Go source file exceeds 400 lines without an intentional split plan.

## Research Basis

- Anthropic Claude Code memory guidance: keep repository memory concise and use scoped rules to save context.
- Anthropic Claude Code best practices: explore, plan, implement, commit, and provide verification evidence.
- GitHub Copilot repository instructions: repository instructions are automatically added to requests, so stacked instructions consume context.
- "Evaluating AGENTS.md: Are Repository-Level Context Files Helpful for Coding Agents?" notes that unnecessary repository-level requirements can increase cost and reduce success in some tasks.

## Links

- Related docs:
  - `AGENTS.md`
  - `docs/architecture/system-overview.md`
  - `docs/decisions/adr/001-backend-module-layout.md`
- External references:
  - `https://code.claude.com/docs/en/memory`
  - `https://docs.anthropic.com/en/docs/claude-code/best-practices`
  - `https://docs.github.com/en/copilot/how-tos/configure-custom-instructions/add-repository-instructions`
  - `https://arxiv.org/abs/2602.11988`
