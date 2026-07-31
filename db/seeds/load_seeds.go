package seeds

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Load runs every available seed loader against pool.
//
// MachineStateTransitions is not loaded here: its source/target values
// include the "*" wildcard and pseudo-states (see MachineStateTransition's
// doc comment), which violate machine_state_transitions' foreign keys as
// currently defined in migration 003.
func Load(ctx context.Context, pool *pgxpool.Pool) error {
	if err := SeedMachineStates(ctx, pool); err != nil {
		return fmt.Errorf("seeds: load machine states: %w", err)
	}

	if err := SeedMachineEvents(ctx, pool); err != nil {
		return fmt.Errorf("seeds: load machine events: %w", err)
	}

	return nil
}
