# Affiliate SaaS

Affiliate SaaS is an AI-assisted workspace for affiliate commerce.

It helps creators, affiliates, agencies, and small commerce operators organize products, affiliate links, campaign ideas, publishing tasks, click tracking, conversion imports, and performance insights across multiple affiliate programs.

## Product Direction

The product is built around a practical affiliate workflow:

```text
product -> affiliate link -> campaign package -> tracking -> imported conversions -> insight
```

The first version focuses on helping operators answer:

- Which product should I promote next?
- Which angle or channel is working?
- Which links and campaigns generated clicks?
- Which imported sales or commissions should inform the next campaign?

## MVP Focus

Initial marketplace playbooks:

- TikTok Shop Affiliate
- Shopee Affiliates
- Mercado Livre
- Amazon Associates/Creators

The MVP is designed to work without depending on closed marketplace APIs. It favors manual setup, user-authorized imports, clear tracking, and safe AI-assisted campaign drafting.

## Principles

- Keep the workflow commerce-first.
- Avoid scraping, browser automation, and artificial engagement.
- Treat affiliate disclosures, product claims, pricing, and asset rights as first-class concerns.
- Prefer official APIs and user-authorized imports when integrations become available.

## Project Status

The product development documentation base is complete. Backend implementation has started with the Go/Gin scaffold, health endpoint, config loading, PostgreSQL connection layer, and initial SQL migrations.

The active implementation slice is:

```text
workspace -> marketplace program -> product -> affiliate link -> short redirect -> click event -> dashboard query
```

## Documentation

Start here:

- [AGENTS.md](AGENTS.md)
- [docs/README.md](docs/README.md)
- [docs/INDEX.md](docs/INDEX.md)
- [docs/workflows/development/project-resume-brief.md](docs/workflows/development/project-resume-brief.md)

## Stack Direction

- Backend: Go + Gin
- Database: PostgreSQL
- Queue: Redis Streams when async jobs are needed
- Frontend: Next.js + React
- Deploy: Docker first
