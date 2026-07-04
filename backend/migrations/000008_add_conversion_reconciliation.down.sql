DROP INDEX IF EXISTS conversion_import_rows_workspace_reconciliation_idx;

ALTER TABLE conversion_import_rows
  DROP CONSTRAINT IF EXISTS conversion_import_rows_reconciliation_status_check,
  DROP COLUMN IF EXISTS reconciled_at,
  DROP COLUMN IF EXISTS reconciliation_note,
  DROP COLUMN IF EXISTS reconciliation_status;
