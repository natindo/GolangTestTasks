package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	TasksCreatedTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "tasks_created_total",
		Help: "Общее количество созданных задач",
	})

	TaskCreationDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "task_creation_duration_seconds",
		Help:    "Гистограмма времени обработки запроса на создание задачи",
		Buckets: prometheus.DefBuckets,
	})
)

func InitMetrics() {
	prometheus.MustRegister(TasksCreatedTotal, TaskCreationDuration)
}
