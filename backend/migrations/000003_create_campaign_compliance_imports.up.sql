CREATE TABLE campaigns (
  id TEXT PRIMARY KEY,
  workspace_id TEXT NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
  product_id TEXT REFERENCES products(id) ON DELETE SET NULL,
  name TEXT NOT NULL,
  status TEXT NOT NULL DEFAULT 'draft',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  archived_at TIMESTAMPTZ,
  CONSTRAINT campaigns_status_check CHECK (status IN ('draft', 'ready', 'published', 'archived'))
);

CREATE INDEX campaigns_workspace_product_idx ON campaigns (workspace_id, product_id);

ALTER TABLE link_variants
  ADD CONSTRAINT link_variants_campaign_id_fkey
  FOREIGN KEY (campaign_id) REFERENCES campaigns(id) ON DELETE SET NULL;

CREATE TABLE channel_packages (
  id TEXT PRIMARY KEY,
  workspace_id TEXT NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
  campaign_id TEXT NOT NULL REFERENCES campaigns(id) ON DELETE CASCADE,
  channel TEXT NOT NULL,
  title TEXT,
  body TEXT,
  status TEXT NOT NULL DEFAULT 'draft',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  CONSTRAINT channel_packages_status_check CHECK (status IN ('draft', 'ready', 'archived'))
);

CREATE INDEX channel_packages_workspace_campaign_idx ON channel_packages (workspace_id, campaign_id);

CREATE TABLE compliance_checks (
  id TEXT PRIMARY KEY,
  workspace_id TEXT NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
  campaign_id TEXT REFERENCES campaigns(id) ON DELETE CASCADE,
  status TEXT NOT NULL DEFAULT 'completed',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  CONSTRAINT compliance_checks_status_check CHECK (status IN ('completed', 'failed'))
);

CREATE TABLE compliance_findings (
  id TEXT PRIMARY KEY,
  workspace_id TEXT NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
  compliance_check_id TEXT NOT NULL REFERENCES compliance_checks(id) ON DELETE CASCADE,
  severity TEXT NOT NULL,
  code TEXT NOT NULL,
  message TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  CONSTRAINT compliance_findings_severity_check CHECK (severity IN ('info', 'warning', 'blocker'))
);

CREATE INDEX compliance_findings_workspace_check_idx
  ON compliance_findings (workspace_id, compliance_check_id);

CREATE TABLE conversion_imports (
  id TEXT PRIMARY KEY,
  workspace_id TEXT NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
  source TEXT NOT NULL DEFAULT 'manual',
  status TEXT NOT NULL DEFAULT 'draft',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  CONSTRAINT conversion_imports_status_check CHECK (status IN ('draft', 'processed', 'failed', 'archived'))
);

CREATE TABLE conversion_import_rows (
  id TEXT PRIMARY KEY,
  workspace_id TEXT NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
  conversion_import_id TEXT NOT NULL REFERENCES conversion_imports(id) ON DELETE CASCADE,
  product_id TEXT REFERENCES products(id) ON DELETE SET NULL,
  affiliate_link_id TEXT REFERENCES affiliate_links(id) ON DELETE SET NULL,
  occurred_at TIMESTAMPTZ,
  order_reference TEXT,
  gross_amount_cents INTEGER,
  commission_cents INTEGER,
  currency TEXT,
  raw_payload JSONB NOT NULL DEFAULT '{}'::jsonb,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  CONSTRAINT conversion_import_rows_gross_amount_check CHECK (gross_amount_cents IS NULL OR gross_amount_cents >= 0),
  CONSTRAINT conversion_import_rows_commission_check CHECK (commission_cents IS NULL OR commission_cents >= 0)
);

CREATE INDEX conversion_import_rows_workspace_import_idx
  ON conversion_import_rows (workspace_id, conversion_import_id);
