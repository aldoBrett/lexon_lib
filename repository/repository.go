package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	LegalProcedure            LegalProcedureRepository
	LegalRecord               LegalRecordRepository
	LegalClaims               LegalClaimsRepository
	MachineInstances          MachineInstancesRepository
	MachineTransitions        MachineTransitionsRepository
	MachineTransitionsHistory MachineStateTransitionsHistoryRepository
	MachineStates             MachineStatesRepository
	MachineStateStages        MachineStateStagesRepository
}

func NewRepositories(ctx context.Context, pool *pgxpool.Pool) *Repositories {
	return &Repositories{
		LegalProcedure:            NewLegalProcedureRepositoryHandler(ctx, pool),
		LegalRecord:               NewLegalRecordRepositoryHandler(ctx, pool),
		LegalClaims:               NewLegalClaimsRepositoryHandler(ctx, pool),
		MachineInstances:          NewMachineInstancesRepositoryHandler(ctx, pool),
		MachineTransitions:        NewMachineTransitionsRepositoryHandler(ctx, pool),
		MachineTransitionsHistory: NewMachineStateTransitionsHistoryRepositoryHandler(ctx, pool),
		MachineStates:             NewMachineStatesRepositoryHandler(ctx, pool),
		MachineStateStages:        NewMachineStateStagesRepositoryHandler(ctx, pool),
	}
}
