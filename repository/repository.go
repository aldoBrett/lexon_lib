package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	LegalProcedure  LegalProcedureRepository
	MachineInstance MachineInstanceRepository
	LegalRecord     LegalRecordRepository
	LegalClaims     LegalClaimsRepository
}

func NewRepositories(ctx context.Context, pool *pgxpool.Pool) *Repositories {
	return &Repositories{
		LegalProcedure:  NewLegalProcedureRepositoryHandler(ctx, pool),
		MachineInstance: NewMachineInstanceRepositoryHandler(ctx, pool),
		LegalRecord:     NewLegalRecordRepositoryHandler(ctx, pool),
		LegalClaims:     NewLegalClaimsRepositoryHandler(ctx, pool),
	}
}
