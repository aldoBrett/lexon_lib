package engine

import (
	"context"

	"amox/lex_engine_lib/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EngineParams struct {
	Ctx  context.Context
	Pool *pgxpool.Pool
}

type EngineHandler struct {
	repos          *repository.Repositories
	legalProcedure *LegalProcedure
}

func NewEngineHandler(params EngineParams) *EngineHandler {
	return &EngineHandler{
		repos: repository.NewRepositories(params.Ctx, params.Pool),
	}
}
