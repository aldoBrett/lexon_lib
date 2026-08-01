package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	LegalProcedure  LegalProcedureRepository
	MachineInstance MachineInstanceRepository
}

func NewRepositories(ctx context.Context, pool *pgxpool.Pool) *Repositories {
	return &Repositories{
		LegalProcedure:  NewLegalProcedureRepositoryHandler(ctx, pool),
		MachineInstance: NewMachineInstanceRepositoryHandler(ctx, pool),
	}
}
