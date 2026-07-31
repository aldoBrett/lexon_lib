package domain

import (
	"encoding/json"
	"time"
)

type MachineName string
type StateID string
type EventType string
type InstanceID string
type TransitionID string

type TruthValue string

const (
	TruthTrue    TruthValue = "TRUE"
	TruthFalse   TruthValue = "FALSE"
	TruthUnknown TruthValue = "UNKNOWN"
)

type TransitionKind string

const (
	TransitionExternal TransitionKind = "EXTERNAL"
	TransitionInternal TransitionKind = "INTERNAL"
)

type TargetMode string

const (
	TargetExplicit   TargetMode = "EXPLICIT"
	TargetCurrent    TargetMode = "CURRENT"
	TargetPrevious   TargetMode = "PREVIOUS"
	TargetStartChild TargetMode = "START_CHILD_MACHINE"
	TargetOutOfScope TargetMode = "OUT_OF_SCOPE"
)

type InstanceStatus string

const (
	InstanceActive    InstanceStatus = "ACTIVE"
	InstanceCompleted InstanceStatus = "COMPLETED"
	InstanceCancelled InstanceStatus = "CANCELLED"
)

type RelationType string

const (
	RelationIncident         RelationType = "INCIDENT"
	RelationDefendantService RelationType = "DEFENDANT_SERVICE"
	RelationEvidence         RelationType = "EVIDENCE"
	RelationAppeal           RelationType = "APPEAL"
	RelationExecution        RelationType = "EXECUTION"
	RelationSuspension       RelationType = "SUSPENSION"
)

type RiskLevel string

const (
	RiskLow      RiskLevel = "LOW"
	RiskMedium   RiskLevel = "MEDIUM"
	RiskHigh     RiskLevel = "HIGH"
	RiskCritical RiskLevel = "CRITICAL"
)

type State struct {
	ID       StateID         `json:"id"`
	Name     string          `json:"name"`
	ParentID *StateID        `json:"parent_id,omitempty"`
	Initial  *StateID        `json:"initial,omitempty"`
	Terminal bool            `json:"terminal,omitempty"`
	Metadata json.RawMessage `json:"metadata,omitempty"`
}

type EventDefinition struct {
	Type     EventType       `json:"type"`
	Name     string          `json:"name"`
	Category string          `json:"category,omitempty"`
	Emitter  string          `json:"emitter,omitempty"`
	Metadata json.RawMessage `json:"metadata,omitempty"`
}

type Actor struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

type Event struct {
	ID            string          `json:"id"`
	Type          EventType       `json:"type"`
	InstanceID    InstanceID      `json:"instance_id"`
	CorrelationID string          `json:"correlation_id,omitempty"`
	CausationID   string          `json:"causation_id,omitempty"`
	OcurredAt     time.Time       `json:"ocurred_at"`
	Actor         Actor           `json:"actor,omitempty"`
	Payload       json.RawMessage `json:"payload,omitempty"`
	Facts         map[string]any  `json:"facts,omitempty"`
	Metadata      map[string]any  `json:"metadata,omitempty"`
}

type ActionSpect struct {
	Type       string         `json:"type"`
	Parameters map[string]any `json:"parameters,omitempty"`
}

type Transition struct {
	ID         TransitionID   `json:"id"`
	Name       string         `json:"name"`
	Source     StateID        `json:"source"`
	Event      EventType      `json:"event"`
	Target     StateID        `json:"target"`
	TargetMode TargetMode     `json:"target_mode,omitempty"`
	Kind       TransitionKind `json:"kind,omitempty"`
	Priority   int            `json:"priority,omitempty"`
	// Condition Expression   `json:"condition,omitempty"`
	// Actions []ActionSpec `json:"actions,omitempty"`
	Risk       RiskLevel `json:"risk,omitempty"`
	LegalBasis string    `json:"legal_basis,omitempty"`
	// Review *ReviewSpec `json:"review,omitempty"`
}

type ReviewSpec struct {
	Required bool      `json:"required"`
	Role     string    `json:"role,omitempty"`
	Risk     RiskLevel `json:"risk,omitempty"`
}

type Definition struct {
	Name        MachineName                   `json:"name"`
	Version     int                           `json:"version"`
	RootState   StateID                       `json:"root_state"`
	States      map[StateID]State             `json:"states"`
	Events      map[EventType]EventDefinition `json:"events"`
	Transitions []Transition                  `json:"transitions"`
}

type Instance struct {
	ID               InstanceID      `json:"id"`
	MachineName      MachineName     `json:"machine_name"`
	MachineVersion   int             `json:"machine_version"`
	BusinessKey      string          `json:"business_key,omitempty"`
	ParentInstanceID *InstanceID     `json:"parent_instance_id,omitempty"`
	RelationType     *RelationType   `json:"relation_type,omitempty"`
	CurrentState     StateID         `json:"current_state"`
	PreviousState    *StateID        `json:"previous_state,omitempty"`
	Data             json.RawMessage `json:"data,omitempty"`
	Status           InstanceStatus  `json:"status"`
	Version          int64           `json:"version"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
}

type ReviewTask struct {
	ID           string       `json:"id"`
	InstanceID   InstanceID   `json:"instance_id"`
	EventID      string       `json:"event_id"`
	TransitionID TransitionID `json:"transition_id"`
	Reason       string       `json:"reason"`
	Risk         string       `json:"risk"`
	Status       string       `json:"status"`
	CreatedAt    time.Time    `json:"created_at"`
}

type OutboxMessage struct {
	ID         string          `json:"id"`
	InstanceID InstanceID      `json:"instance_id"`
	Topic      string          `json:"topic"`
	EventType  EventType       `json:"event_type"`
	Payload    json.RawMessage `json:"payload"`
	CreatedAt  time.Time       `json:"created_at"`
}
