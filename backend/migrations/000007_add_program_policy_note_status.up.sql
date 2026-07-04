ALTER TABLE program_policy_notes
  ADD COLUMN status TEXT NOT NULL DEFAULT 'active',
  ADD CONSTRAINT program_policy_notes_status_check CHECK (status IN ('active', 'archived'));

CREATE INDEX program_policy_notes_program_status_idx
  ON program_policy_notes (program_id, status);
