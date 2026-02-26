CREATE TABLE users (
  id TEXT PRIMARY KEY,
  org_id TEXT NOT NULL,
  email TEXT NOT NULL,
  name TEXT NOT NULL,
  role TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
);

CREATE UNIQUE INDEX users_org_id_email_uq ON users (org_id, email);
CREATE INDEX users_org_id_idx ON users (org_id);

CREATE TABLE services (
  id TEXT PRIMARY KEY,
  org_id TEXT NOT NULL,
  owner_id TEXT NOT NULL,
  name TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  CONSTRAINT services_owner_id_fk FOREIGN KEY (owner_id) REFERENCES users (id)
);

CREATE UNIQUE INDEX services_org_id_name_uq ON services (org_id, name);
CREATE INDEX services_org_id_idx ON services (org_id);
CREATE INDEX services_owner_id_idx ON services (owner_id);

CREATE TABLE routes (
  id TEXT PRIMARY KEY,
  org_id TEXT NOT NULL,
  service_id TEXT NOT NULL,
  method TEXT NOT NULL,
  path_pattern TEXT NOT NULL,
  is_shadow_enabled BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  CONSTRAINT routes_service_id_fk FOREIGN KEY (service_id) REFERENCES services (id)
);

CREATE UNIQUE INDEX routes_service_method_path_uq ON routes (service_id, method, path_pattern);
CREATE INDEX routes_org_id_idx ON routes (org_id);
CREATE INDEX routes_service_id_idx ON routes (service_id);

CREATE TABLE shadow_targets (
  id TEXT PRIMARY KEY,
  org_id TEXT NOT NULL,
  route_id TEXT NOT NULL,
  target_service TEXT NOT NULL,
  weight INTEGER NOT NULL DEFAULT 100,
  timeout_ms INTEGER NOT NULL DEFAULT 1000,
  retries INTEGER NOT NULL DEFAULT 0,
  enabled BOOLEAN NOT NULL DEFAULT TRUE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  CONSTRAINT shadow_targets_route_id_fk FOREIGN KEY (route_id) REFERENCES routes (id) ON DELETE CASCADE,
  CONSTRAINT shadow_targets_weight_check CHECK (weight >= 0),
  CONSTRAINT shadow_targets_timeout_ms_check CHECK (timeout_ms >= 0),
  CONSTRAINT shadow_targets_retries_check CHECK (retries >= 0)
);

CREATE INDEX shadow_targets_org_id_idx ON shadow_targets (org_id);
CREATE INDEX shadow_targets_route_id_idx ON shadow_targets (route_id);

CREATE TABLE experiments (
  id TEXT PRIMARY KEY,
  org_id TEXT NOT NULL,
  route_id TEXT NOT NULL,
  name TEXT NOT NULL,
  status TEXT NOT NULL,
  started_at TIMESTAMPTZ,
  ended_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  CONSTRAINT experiments_route_id_fk FOREIGN KEY (route_id) REFERENCES routes (id) ON DELETE CASCADE
);

CREATE INDEX experiments_org_id_idx ON experiments (org_id);
CREATE INDEX experiments_route_id_idx ON experiments (route_id);
CREATE INDEX experiments_status_idx ON experiments (status);

CREATE TABLE policies (
  id TEXT PRIMARY KEY,
  org_id TEXT NOT NULL,
  name TEXT NOT NULL,
  mismatch_max DOUBLE PRECISION NOT NULL DEFAULT 0,
  latency_max_ms INTEGER NOT NULL DEFAULT 0,
  enabled BOOLEAN NOT NULL DEFAULT TRUE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
);

CREATE UNIQUE INDEX policies_org_id_name_uq ON policies (org_id, name);
CREATE INDEX policies_org_id_idx ON policies (org_id);

CREATE TABLE diff_results (
  id TEXT PRIMARY KEY,
  org_id TEXT NOT NULL,
  experiment_id TEXT NOT NULL,
  route_id TEXT NOT NULL,
  request_id TEXT NOT NULL,
  verdict TEXT NOT NULL,
  field TEXT NOT NULL,
  primary_value TEXT NOT NULL,
  shadow_value TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  CONSTRAINT diff_results_experiment_id_fk FOREIGN KEY (experiment_id) REFERENCES experiments (id),
  CONSTRAINT diff_results_route_id_fk FOREIGN KEY (route_id) REFERENCES routes (id)
);

CREATE INDEX diff_results_org_id_idx ON diff_results (org_id);
CREATE INDEX diff_results_experiment_id_idx ON diff_results (experiment_id);
CREATE INDEX diff_results_route_id_idx ON diff_results (route_id);
CREATE INDEX diff_results_request_id_idx ON diff_results (request_id);
CREATE INDEX diff_results_verdict_idx ON diff_results (verdict);

CREATE TABLE audit_logs (
  id TEXT PRIMARY KEY,
  org_id TEXT NOT NULL,
  actor_id TEXT NOT NULL,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  meta TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  CONSTRAINT audit_logs_actor_id_fk FOREIGN KEY (actor_id) REFERENCES users (id)
);

CREATE INDEX audit_logs_org_id_idx ON audit_logs (org_id);
CREATE INDEX audit_logs_actor_id_idx ON audit_logs (actor_id);
CREATE INDEX audit_logs_created_at_idx ON audit_logs (created_at);
