package main

import (
	"context"
	"log"

	"go.etcd.io/bbolt"
	"go.uber.org/zap"
	"yp-webinar/internal/service"
	"yp-webinar/pkg/logging"
)

func main() {
	ctx := context.Background()
	l, err := logging.NewZapLogger(zap.InfoLevel)
	if err != nil {
		log.Fatal(err)
	}

	backup, err := bbolt.Open("backup.db", 0600, nil)
	if err != nil {
		l.PanicCtx(ctx, "failed to open backup database", zap.Error(err))
	}

	s := service.New(l, backup)
	defer s.CloseBackup()

	data, err := s.Get(ctx, []byte("key"))
	if err != nil {
		l.ErrorCtx(ctx, "failed to get value", zap.Error(err))
	}

	l.InfoCtx(ctx, "value", zap.String("data", string(data)))

}
