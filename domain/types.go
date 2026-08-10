package domain

import "time"

type LegalProcedure struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

type LegalRecord struct {
	ID               string  `json:"id"`
	LegalProcedureID *string `json:"legal_procedure_id"`
	TrialKind        *string `json:"trial_kind"`
	RecordNumber     *string `json:"record_number"`
	//? Should it be also an table in the database or just a string?
	Actor            string  `json:"actor"`
	Defendant        string  `json:"defendant"`
	DefendantAddress *string `json:"defendant_address"`
	// Id of the legal action associated with this legal record, if any. This field is optional and can be null.
	ActionID *string `json:"action_id"`
}

type LegalClaim struct {
	ID            string `json:"id"`
	LegalRecordID string `json:"legal_record_id"`
	Description   string `json:"description"`
}

type LegalAction struct {
	ID          string `json:"id"`
	Category    string `json:"category"`
	SubCategory string `json:"sub_category"`
	ActionName  string `json:"action_name"`
	Via         string `json:"via"`
}

type DocumentData struct {
	Kind string `json:"kind"`
}

type MachineStateStage struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Content     string `json:"content"`
	StageOrder  int    `json:"stage_order"`
}

type MachineState struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	StageID     string `json:"stage_id"`
	Kind        string `json:"kind"`
}

type MachineStateTransition struct {
	ID string `json:"id"`
	// Code          string `json:"code"`
	SourceStateID string    `json:"source_state_id"`
	TargetStateID string    `json:"target_state_id"`
	EventID       string    `json:"event_id"`
	Condition     string    `json:"condition"`
	Actions       string    `json:"actions"`
	Risk          string    `json:"risk"`
	Note          string    `json:"note"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	// SourceStateID string `json:"source_state_id"`

	// SourceID string `json:"source"`
	// EventID  string `json:"event"`
	// TargetID string `json:"target"`
}

type MachineStateTransitionHistory struct {
	ID                string    `json:"id"`
	MachineInstanceID string    `json:"machine_instance_id"`
	TransitionID      string    `json:"transition_id"`
	CreatedAt         time.Time `json:"created_at"`
}

type MachineEvent struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Kind        string `json:"kind"`
}

type MachineInstance struct {
	ID               string    `json:"id"`
	CurrentStateID   *string   `json:"current_state_id,omitempty"`
	LegalProcedureID string    `json:"legal_procedure_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// It represents a signal that can come from different origins,
// this signal will be processed by the engine machine.
type MachineSignal struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	// The possible values for the origin are:
	//   - document: the signal comes from a digitalized document
	//   - ui: the signal comes from the user interface
	//   - system: the signal comes from the system (e.g., a scheduled task)
	Origin string `json:"origin"`
	// Kind string `json:"kind"`
	// If the signal comes from a document this are the possible values:
	//   - demanda
	//   - prevencion
	DocumentType string  `json:"document_type"`
	EventID      *string `json:"event_id,omitempty"`
}
