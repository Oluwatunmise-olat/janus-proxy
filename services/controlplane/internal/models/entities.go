package models

import "time"

type User struct {
	ID        string    `json:"id"`
	OrgID     string    `json:"org_id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Service struct {
	ID        string    `json:"id"`
	OrgID     string    `json:"org_id"`
	OwnerID   string    `json:"owner_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Route struct {
	ID              string    `json:"id"`
	OrgID           string    `json:"org_id"`
	ServiceID       string    `json:"service_id"`
	Method          string    `json:"method"`
	PathPattern     string    `json:"path_pattern"`
	IsShadowEnabled bool      `json:"is_shadow_enabled"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type ShadowTarget struct {
	ID              string    `json:"id"`
	OrgID           string    `json:"org_id"`
	RouteID         string    `json:"route_id"`
	TargetServiceID string    `json:"target_service_id"`
	Weight          int       `json:"weight"`
	TimeoutMs       int       `json:"timeout_ms"`
	Retries         int       `json:"retries"`
	Enabled         bool      `json:"enabled"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Experiment struct {
	ID        string    `json:"id"`
	OrgID     string    `json:"org_id"`
	RouteID   string    `json:"route_id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Policy struct {
	ID           string    `json:"id"`
	OrgID        string    `json:"org_id"`
	Name         string    `json:"name"`
	MismatchMax  float64   `json:"mismatch_max"`
	LatencyMaxMs int       `json:"latency_max_ms"`
	Enabled      bool      `json:"enabled"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type DiffResult struct {
	ID           string    `json:"id"`
	OrgID        string    `json:"org_id"`
	ExperimentID string    `json:"experiment_id"`
	RouteID      string    `json:"route_id"`
	RequestID    string    `json:"request_id"`
	Verdict      string    `json:"verdict"`
	Field        string    `json:"field"`
	PrimaryValue string    `json:"primary_value"`
	ShadowValue  string    `json:"shadow_value"`
	CreatedAt    time.Time `json:"created_at"`
}

type AuditLog struct {
	ID          string    `json:"id"`
	OrgID       string    `json:"org_id"`
	ActorID     string    `json:"actor_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Meta        string    `json:"meta"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
