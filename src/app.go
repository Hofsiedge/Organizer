package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Hofsiedge/Organizer/src/handlers"
	"github.com/Hofsiedge/Organizer/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

func main() {
	app := fiber.New()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	dbPool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		err = fmt.Errorf("Unable to connect to database: %w\n", err)
		sugar.Error(err)
		panic(err)
	}
	defer dbPool.Close()

	appContext := utils.AppContext{
		DbPool: dbPool,
		Logger: sugar,
	}

	characterAPI := app.Group("/task")
	characterAPI.Get(
		"/:id",
		withAppCtx(&appContext, handlers.GetTask),
	)
	characterAPI.Post(
		"/",
		withAppCtx(&appContext, handlers.CreateTask),
	)

	if err := app.Listen(":80"); err != nil {
		panic(err)
	}
}

func withAppCtx(ac *utils.AppContext, handler func(*utils.AppContext, *fiber.Ctx) error) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return handler(ac, c)
	}
}
