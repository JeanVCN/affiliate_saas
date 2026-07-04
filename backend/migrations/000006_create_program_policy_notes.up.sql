CREATE TABLE program_policy_notes (
  id TEXT PRIMARY KEY,
  program_id TEXT NOT NULL REFERENCES programs(id) ON DELETE CASCADE,
  title TEXT NOT NULL,
  body TEXT NOT NULL,
  severity TEXT NOT NULL DEFAULT 'info',
  source_url TEXT,
  effective_at TIMESTAMPTZ,
  reviewed_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  CONSTRAINT program_policy_notes_severity_check CHECK (severity IN ('info', 'warning', 'blocker'))
);

CREATE INDEX program_policy_notes_program_idx ON program_policy_notes (program_id);
