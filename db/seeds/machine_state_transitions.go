package seeds

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// MachineStateTransition mirrors a row of lexon.machine_state_transitions.
//
// A handful of transitions from the source material use "*" as a wildcard
// source (any state) or pseudo-state exit markers (FUERA_DE_ALCANCE,
// FUJO_RECURSO, ALERTA_CADUCIDAD, ESTADO_PRINCIPAL_PREVIO, ESTADO_PREVIO,
// ESTADO_ACTUAL) that are not rows in machine_states. Those are commented
// out below pending a decision on how to represent them; the rest reference
// only real machine_states IDs and are safe to load as-is.
type MachineStateTransition struct {
	ID          string
	SourceState string
	TargetState string
	EventID     string
	Condition   string
	Actions     string
	Risk        string
	Note        string
}

var MachineStateTransitions = []MachineStateTransition{
	{
		ID:          "T001",
		SourceState: "CIV.ORD.S00",
		EventID:     "CIV.ORD.E001",
		Condition:   "",
		TargetState: "CIV.ORD.S00",
		Actions:     "",
		Risk:        "",
		Note:        "",
	},
	{
		ID:          "T002",
		SourceState: "CIV.ORD.S00",
		EventID:     "CIV.ORD.E002",
		Condition:   "VIA = JUICIO ORDINARIO CIVIL",
		TargetState: "CIV.ORD.S01",
		Actions:     "",
		Risk:        "",
		Note:        "",
	},
	{
		ID:          "T003",
		SourceState: "CIV.ORD.S01",
		EventID:     "CIV.ORD.E003",
		Condition:   "",
		TargetState: "CIV.ORD.S02",
		Actions:     "",
		Risk:        "",
		Note:        "",
	},
	{
		ID:          "T004",
		SourceState: "CIV.ORD.S02",
		EventID:     "CIV.ORD.E004",
		Condition:   "",
		TargetState: "CIV.ORD.S03",
		Actions:     "",
		Risk:        "",
		Note:        "",
	},
	{
		ID:          "T005",
		SourceState: "CIV.ORD.S03",
		EventID:     "CIV.ORD.E005",
		Condition:   "",
		TargetState: "CIV.ORD.S04",
		Actions:     "",
		Risk:        "",
		Note:        "",
	},
	{
		ID:          "T006",
		SourceState: "CIV.ORD.S04",
		EventID:     "CIV.ORD.E010",
		Condition:   "",
		TargetState: "CIV.ORD.S08",
		Actions:     "",
		Risk:        "",
		Note:        "",
	},
	{
		ID:          "T007",
		SourceState: "CIV.ORD.S04",
		EventID:     "CIV.ORD.E006",
		Condition:   "",
		TargetState: "CIV.ORD.S05",
		Actions:     "",
		Risk:        "",
		Note:        "",
	},
	{
		ID:          "T008",
		SourceState: "CIV.ORD.S05",
		EventID:     "CIV.ORD.E007",
		Condition:   "",
		TargetState: "CIV.ORD.S07",
		Actions:     "",
		Risk:        "",
		Note:        "",
	},
	{
		ID:          "T009",
		SourceState: "CIV.ORD.S07",
		EventID:     "CIV.ORD.E008",
		Condition:   "",
		TargetState: "CIV.ORD.S08",
		Actions:     "",
		Risk:        "",
		Note:        "",
	},
	{
		ID:          "T010",
		SourceState: "CIV.ORD.S07",
		EventID:     "CIV.ORD.E009",
		Condition:   "",
		TargetState: "CIV.ORD.S34",
		Actions:     "",
		Risk:        "",
		Note:        "",
	},
}

// SeedMachineStateTransitions upserts MachineStateTransitions into lexon.machine_state_transitions.
func SeedMachineStateTransitions(ctx context.Context, pool *pgxpool.Pool) error {
	batch := &pgx.Batch{}
	for _, t := range MachineStateTransitions {
		batch.Queue(
			`INSERT INTO lexon.machine_state_transitions (id, source_state_id, target_state_id, event_id, condition, actions, risk, note)
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			 ON CONFLICT (id) DO UPDATE SET
			   source_state_id = EXCLUDED.source_state_id,
			   target_state_id = EXCLUDED.target_state_id,
			   event_id = EXCLUDED.event_id,
			   condition = EXCLUDED.condition,
			   actions = EXCLUDED.actions,
			   risk = EXCLUDED.risk,
			   note = EXCLUDED.note`,
			t.ID, t.SourceState, t.TargetState, t.EventID, t.Condition, t.Actions, t.Risk, t.Note,
		)
	}

	br := pool.SendBatch(ctx, batch)
	defer br.Close()

	for range MachineStateTransitions {
		if _, err := br.Exec(); err != nil {
			return fmt.Errorf("seeds: insert machine state transition: %w", err)
		}
	}

	return nil
}
