ALTER TABLE conversion_import_rows
  ADD COLUMN reconciliation_status TEXT NOT NULL DEFAULT 'pending',
  ADD COLUMN reconciliation_note TEXT,
  ADD COLUMN reconciled_at TIMESTAMPTZ,
  ADD CONSTRAINT conversion_import_rows_reconciliation_status_check
    CHECK (reconciliation_status IN ('pending', 'matched', 'unmatched', 'ignored'));

UPDATE conversion_import_rows
SET reconciliation_status = CASE
  WHEN product_id IS NOT NULL OR affiliate_link_id IS NOT NULL THEN 'matched'
  ELSE 'pending'
END;

CREATE INDEX conversion_import_rows_workspace_reconciliation_idx
  ON conversion_import_rows (workspace_id, reconciliation_status);
