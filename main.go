package main

import (
	"context"
	"log"

	serviceProvider "github.com/Artenso/FloodControl/internal/app/service_provider"
)

func main() {
	ctx := context.Background()

	app, err := serviceProvider.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to create app: %s", err.Error())
	}

	_ = app

}

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}
