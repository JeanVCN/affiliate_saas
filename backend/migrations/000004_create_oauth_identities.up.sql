CREATE TABLE oauth_identities (
  id TEXT PRIMARY KEY,
  user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  provider TEXT NOT NULL,
  provider_subject TEXT NOT NULL,
  email TEXT,
  display_name TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  CONSTRAINT oauth_identities_provider_check CHECK (provider IN ('google', 'tiktok', 'amazon'))
);

CREATE UNIQUE INDEX oauth_identities_provider_subject_unique
  ON oauth_identities (provider, provider_subject);

CREATE INDEX oauth_identities_user_id_idx ON oauth_identities (user_id);

CREATE TABLE oauth_states (
  id TEXT PRIMARY KEY,
  provider TEXT NOT NULL,
  state_hash TEXT NOT NULL,
  redirect_url TEXT,
  expires_at TIMESTAMPTZ NOT NULL,
  consumed_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  CONSTRAINT oauth_states_provider_check CHECK (provider IN ('google', 'tiktok', 'amazon'))
);

CREATE UNIQUE INDEX oauth_states_state_hash_unique ON oauth_states (state_hash);
CREATE INDEX oauth_states_expires_at_idx ON oauth_states (expires_at);
