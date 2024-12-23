package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"yp-webinar/pkg/logging"
)

type Controller struct {
	l *logging.ZapLogger
	r chi.Router
}

func New(logger *logging.ZapLogger) *Controller {
	return &Controller{
		l: logger,
		r: chi.NewRouter(),
	}
}

func (c *Controller) recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				c.l.PanicCtx(r.Context(), "recovered from panic", zap.Any("panic", rec))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (c *Controller) Routes() *Controller {
	c.r.Use(c.recover)

	c.r.Get("/hello", c.hello)
	return c
}

func (c *Controller) hello(w http.ResponseWriter, r *http.Request) {
	c.l.InfoCtx(r.Context(), "hello world")
}

func (c *Controller) ServeHTTP() error {
	s := http.Server{
		Addr:     ":8080",
		ErrorLog: c.l.Std(),
		Handler:  c.r,
	}

	return s.ListenAndServe()
}
