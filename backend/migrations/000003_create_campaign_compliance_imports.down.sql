DROP TABLE IF EXISTS conversion_import_rows;
DROP TABLE IF EXISTS conversion_imports;
DROP TABLE IF EXISTS compliance_findings;
DROP TABLE IF EXISTS compliance_checks;
DROP TABLE IF EXISTS channel_packages;

ALTER TABLE IF EXISTS link_variants
  DROP CONSTRAINT IF EXISTS link_variants_campaign_id_fkey;

DROP TABLE IF EXISTS campaigns;
