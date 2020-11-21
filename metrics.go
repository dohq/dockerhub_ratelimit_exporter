package main

import (
	"context"
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

type Collector struct {
	username string
	password string
}

var (
	RateLimitGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "limit",
		Help:      "DockerHub RateLimit-Limit gauge",
	})
	RemainingGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "remaining",
		Help:      "DockerHub RateLimit-Remaining gauge",
	})
)

func (c Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- RateLimitGauge.Desc()
	ch <- RemainingGauge.Desc()
}

func (c Collector) Collect(ch chan<- prometheus.Metric) {
	ctx := context.Background()
	l, err := checkLimit(ctx, c.username, c.password)
	if err != nil {
		log.Printf("could not get limit status: %v", err)
	}

	ch <- prometheus.MustNewConstMetric(
		RateLimitGauge.Desc(),
		prometheus.GaugeValue,
		float64(l.Limit),
	)
	ch <- prometheus.MustNewConstMetric(
		RemainingGauge.Desc(),
		prometheus.GaugeValue,
		float64(l.Remaining),
	)
}
