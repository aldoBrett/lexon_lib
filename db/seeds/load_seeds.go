package seeds

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Load runs every available seed loader against pool.
//
// A subset of MachineStateTransitions is commented out in its seed data:
// those rows use the "*" wildcard or pseudo-state exit markers (see
// MachineStateTransition's doc comment), which violate machine_state_transitions'
// foreign keys as currently defined in migration 003, pending a decision on
// how to represent them.
func Load(ctx context.Context, pool *pgxpool.Pool) error {
	if err := SeedMachineStateStages(ctx, pool); err != nil {
		return fmt.Errorf("seeds: load machine state stages: %w", err)
	}

	if err := SeedMachineStates(ctx, pool); err != nil {
		return fmt.Errorf("seeds: load machine states: %w", err)
	}

	if err := SeedMachineEvents(ctx, pool); err != nil {
		return fmt.Errorf("seeds: load machine events: %w", err)
	}

	if err := SeedMachineStateTransitions(ctx, pool); err != nil {
		return fmt.Errorf("seeds: load machine state transitions: %w", err)
	}

	if err := SeedLegalActions(ctx, pool); err != nil {
		return fmt.Errorf("seeds: load legal actions: %w", err)
	}

	return nil
}
