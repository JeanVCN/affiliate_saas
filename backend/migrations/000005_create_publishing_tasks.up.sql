CREATE TABLE publishing_tasks (
  id TEXT PRIMARY KEY,
  workspace_id TEXT NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
  campaign_id TEXT NOT NULL REFERENCES campaigns(id) ON DELETE CASCADE,
  channel_package_id TEXT REFERENCES channel_packages(id) ON DELETE SET NULL,
  channel TEXT NOT NULL,
  title TEXT NOT NULL,
  notes TEXT,
  scheduled_for TIMESTAMPTZ,
  status TEXT NOT NULL DEFAULT 'todo',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  completed_at TIMESTAMPTZ,
  CONSTRAINT publishing_tasks_status_check CHECK (status IN ('todo', 'scheduled', 'done', 'canceled'))
);

CREATE INDEX publishing_tasks_workspace_campaign_idx
  ON publishing_tasks (workspace_id, campaign_id);
