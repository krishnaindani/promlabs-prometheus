package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {

	// Create a gauge with two label names ("house" and "room").
	temp := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "home_temperature_celsius",
			Help: "The current temperature in degrees Celsius.",
		},
		// The two label names by which to split the metric.
		[]string{"house", "room"},
	)

	// Register the gauge with our metrics registry.
	prometheus.MustRegister(temp)

	// Set the temperature to different values, depending on house and room.
	temp.WithLabelValues("julius", "living-room").Set(23.5)
	temp.WithLabelValues("julius", "bedroom").Set(21.2)
	temp.WithLabelValues("fred", "living-room").Set(23.6)
	temp.WithLabelValues("fred", "bedroom").Set(19.0)

	//NOTE: When using metrics with label dimensions, the time series for any label combination will
	//only appear in the /metrics output once that label combination has been been accessed at least
	//once. This can cause problems in PromQL queries that expect certain series to always be present.
	//When feasible, pre-initialize all important label combinations to default values when the program
	//first starts.

	// Expose our custom registry over HTTP on /metrics.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
