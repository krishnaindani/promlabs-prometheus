package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {

	registry := prometheus.NewRegistry()

	registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	registry.MustRegister(collectors.NewGoCollector())

	temp := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "home_temperature_celsius",
		Help: "The current temperature in degrees Celsius.",
	})

	registry.MustRegister(temp)

	temp.Set(42)

	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{
		Registry: registry,
	}))

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
