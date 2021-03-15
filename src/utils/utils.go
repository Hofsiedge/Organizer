package utils

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type AppContext struct {
	DbPool *pgxpool.Pool
	Logger *zap.SugaredLogger
}
