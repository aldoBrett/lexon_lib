package engine

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EngineParams struct {
	Ctx  context.Context
	Pool *pgxpool.Pool
}

type EngineHandler struct {
	ctx  context.Context
	pool *pgxpool.Pool
}

func NewEngineHandler(params EngineParams) *EngineHandler {
	return &EngineHandler{
		ctx:  params.Ctx,
		pool: params.Pool,
	}
}
