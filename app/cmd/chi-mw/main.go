package main

import (
	"context"
	"log"

	"go.uber.org/zap"
	"yp-webinar/internal/controller"
	"yp-webinar/pkg/logging"
)

func main() {
	ctx := context.Background()
	l, err := logging.NewZapLogger(zap.InfoLevel)

	if err != nil {
		log.Panic(err)
	}

	c := controller.New(l)
	err = c.Routes().ServeHTTP()
	if err != nil {
		l.PanicCtx(ctx, "failed to start server", zap.Error(err))
	}
}
