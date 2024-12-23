package service

import (
	"context"
	"fmt"

	"go.etcd.io/bbolt"
	"go.uber.org/zap"
	"yp-webinar/pkg/logging"
	"yp-webinar/pkg/rand"
)

type Service struct {
	l      *logging.ZapLogger
	backup *bbolt.DB
}

func New(logger *logging.ZapLogger,
	backup *bbolt.DB,
) *Service {
	return &Service{
		l:      logger,
		backup: backup,
	}
}

func (s *Service) DoSomeJob(ctx context.Context) {
	s.l.InfoCtx(ctx,
		"this wonderful service is working with our user...",
		// А вот этот аргумент - модификатор одного вызова логгера, в ошибке почты не будет
		zap.String("email", fmt.Sprintf("%s@yandex.kz", rand.Name())))

	s.l.ErrorCtx(ctx, "error occurred...")
}

func (s *Service) CloseBackup() {
	if s.backup != nil {
		s.backup.Close()
	}
}

func (s *Service) Add(ctx context.Context, k, v []byte) error {
	s.l.InfoCtx(ctx, "adding...")

	return s.backup.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("users"))
		if b == nil {
			var err error
			b, err = tx.CreateBucket([]byte("users"))
			if err != nil {
				return err
			}
		}
		return b.Put(k, v)
	})
}

func (s *Service) Get(ctx context.Context, k []byte) ([]byte, error) {
	s.l.InfoCtx(ctx, "getting...")

	var (
		v   []byte
		err error
	)
	err = s.backup.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("users"))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}

		v = b.Get(k)
		if v == nil {
			return fmt.Errorf("not found")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return v, nil
}
