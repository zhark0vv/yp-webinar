package main

import (
	"context"
	"log"
	"time"

	"go.uber.org/zap"
	"yp-webinar/internal/service"
	"yp-webinar/pkg/logging"
	"yp-webinar/pkg/rand"
)

func main() {
	ctx := context.Background()
	l, err := logging.NewZapLogger(zap.InfoLevel)

	if err != nil {
		log.Panic(err)
	}

	ctx = l.WithContextFields(ctx,
		zap.String("app", "logging"),
		zap.String("service", "main"))

	defer l.Sync()

	s := service.New(l, nil)

	for i := 0; i < 10; i++ {

		user := rand.Name()

		ctx = l.WithContextFields(ctx,
			zap.String("user", user),
			zap.String("password", "qwerty"))

		s.DoSomeJob(ctx)
		time.Sleep(1 * time.Second)
	}
}
