package middleware

import (
	"MyTodo/engine/v1/starter"

	"github.com/prometheus/client_golang/prometheus"
)

func PromCount(opts prometheus.CounterOpts) starter.TodoHandler {
	counter := prometheus.NewCounter(opts)
	prometheus.MustRegister(counter)
	return func(c starter.TodoContext) {
		counter.Inc()
		c.Context().Next()
	}
}
