DROP INDEX IF EXISTS program_policy_notes_program_status_idx;

ALTER TABLE program_policy_notes
  DROP CONSTRAINT IF EXISTS program_policy_notes_status_check,
  DROP COLUMN IF EXISTS status;
