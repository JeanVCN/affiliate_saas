# Affiliate SaaS

AI-assisted affiliate commerce platform for creators, affiliates, agencies, and operators who manage products, affiliate links, campaign assets, publishing tasks, and performance across multiple affiliate programs.

## Current Direction

The MVP is not a generic video clipping tool. It is an affiliate commerce operating system:

`product -> affiliate link -> AI campaign -> channel package -> click tracking -> conversion import -> insight`

Initial playbooks:

- TikTok Shop Affiliate
- Shopee Affiliates
- Mercado Livre
- Amazon Associates/Creators

## Documentation

Start with:

- [AGENTS.md](AGENTS.md)
- [docs/README.md](docs/README.md)
- [docs/INDEX.md](docs/INDEX.md)
- [docs/_meta/engineering-documentation-blueprint.md](docs/_meta/engineering-documentation-blueprint.md)
- [docs/product/business-plan-affiliate-multimarketplace.md](docs/product/business-plan-affiliate-multimarketplace.md)

The documentation system is intentionally designed for AI-assisted development with compact context, clear sources of truth, handoffs, memory, and future MCP/code-index support.

## Stack Direction

- Backend: Go
- Database: PostgreSQL
- Queue: Redis Streams
- Storage: S3-compatible
- Frontend: Next.js + React
- Deploy: Docker first
