CREATE TABLE marketplaces (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  slug TEXT NOT NULL,
  status TEXT NOT NULL DEFAULT 'active',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  CONSTRAINT marketplaces_status_check CHECK (status IN ('active', 'archived'))
);

CREATE UNIQUE INDEX marketplaces_slug_unique ON marketplaces (slug);

CREATE TABLE programs (
  id TEXT PRIMARY KEY,
  marketplace_id TEXT NOT NULL REFERENCES marketplaces(id) ON DELETE RESTRICT,
  name TEXT NOT NULL,
  slug TEXT NOT NULL,
  program_type TEXT NOT NULL DEFAULT 'affiliate',
  status TEXT NOT NULL DEFAULT 'active',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  CONSTRAINT programs_status_check CHECK (status IN ('active', 'archived'))
);

CREATE UNIQUE INDEX programs_marketplace_slug_unique ON programs (marketplace_id, slug);

CREATE TABLE workspace_programs (
  id TEXT PRIMARY KEY,
  workspace_id TEXT NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
  program_id TEXT NOT NULL REFERENCES programs(id) ON DELETE RESTRICT,
  external_account_label TEXT,
  status TEXT NOT NULL DEFAULT 'active',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  archived_at TIMESTAMPTZ,
  CONSTRAINT workspace_programs_status_check CHECK (status IN ('active', 'paused', 'archived'))
);

CREATE UNIQUE INDEX workspace_programs_workspace_program_unique
  ON workspace_programs (workspace_id, program_id);

CREATE TABLE products (
  id TEXT PRIMARY KEY,
  workspace_id TEXT NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
  name TEXT NOT NULL,
  category TEXT,
  description TEXT,
  status TEXT NOT NULL DEFAULT 'active',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  archived_at TIMESTAMPTZ,
  CONSTRAINT products_status_check CHECK (status IN ('active', 'paused', 'archived'))
);

CREATE INDEX products_workspace_id_idx ON products (workspace_id);

CREATE TABLE offers (
  id TEXT PRIMARY KEY,
  workspace_id TEXT NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
  product_id TEXT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
  workspace_program_id TEXT NOT NULL REFERENCES workspace_programs(id) ON DELETE RESTRICT,
  title TEXT,
  price_cents INTEGER,
  currency TEXT,
  status TEXT NOT NULL DEFAULT 'active',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  archived_at TIMESTAMPTZ,
  CONSTRAINT offers_status_check CHECK (status IN ('active', 'paused', 'archived')),
  CONSTRAINT offers_price_cents_check CHECK (price_cents IS NULL OR price_cents >= 0)
);

CREATE INDEX offers_workspace_product_idx ON offers (workspace_id, product_id);

CREATE TABLE affiliate_links (
  id TEXT PRIMARY KEY,
  workspace_id TEXT NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
  product_id TEXT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
  offer_id TEXT REFERENCES offers(id) ON DELETE SET NULL,
  destination_url TEXT NOT NULL,
  label TEXT,
  status TEXT NOT NULL DEFAULT 'active',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  archived_at TIMESTAMPTZ,
  CONSTRAINT affiliate_links_status_check CHECK (status IN ('active', 'paused', 'archived'))
);

CREATE INDEX affiliate_links_workspace_product_idx ON affiliate_links (workspace_id, product_id);

CREATE TABLE link_variants (
  id TEXT PRIMARY KEY,
  workspace_id TEXT NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
  affiliate_link_id TEXT NOT NULL REFERENCES affiliate_links(id) ON DELETE CASCADE,
  campaign_id TEXT,
  channel TEXT,
  utm_source TEXT,
  utm_medium TEXT,
  utm_campaign TEXT,
  status TEXT NOT NULL DEFAULT 'active',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  archived_at TIMESTAMPTZ,
  CONSTRAINT link_variants_status_check CHECK (status IN ('active', 'paused', 'archived'))
);

CREATE INDEX link_variants_workspace_link_idx ON link_variants (workspace_id, affiliate_link_id);

CREATE TABLE short_links (
  id TEXT PRIMARY KEY,
  workspace_id TEXT NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
  affiliate_link_id TEXT NOT NULL REFERENCES affiliate_links(id) ON DELETE CASCADE,
  link_variant_id TEXT REFERENCES link_variants(id) ON DELETE SET NULL,
  slug TEXT NOT NULL,
  status TEXT NOT NULL DEFAULT 'active',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  archived_at TIMESTAMPTZ,
  CONSTRAINT short_links_status_check CHECK (status IN ('active', 'paused', 'archived'))
);

CREATE UNIQUE INDEX short_links_slug_unique ON short_links (slug);
CREATE INDEX short_links_workspace_link_idx ON short_links (workspace_id, affiliate_link_id);

CREATE TABLE click_events (
  id TEXT PRIMARY KEY,
  workspace_id TEXT NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
  short_link_id TEXT NOT NULL REFERENCES short_links(id) ON DELETE CASCADE,
  affiliate_link_id TEXT NOT NULL REFERENCES affiliate_links(id) ON DELETE CASCADE,
  product_id TEXT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
  occurred_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  referrer TEXT,
  user_agent TEXT,
  ip_hash TEXT,
  utm_source TEXT,
  utm_medium TEXT,
  utm_campaign TEXT
);

CREATE INDEX click_events_workspace_occurred_idx ON click_events (workspace_id, occurred_at DESC);
CREATE INDEX click_events_workspace_product_idx ON click_events (workspace_id, product_id);
CREATE INDEX click_events_workspace_link_idx ON click_events (workspace_id, affiliate_link_id);
