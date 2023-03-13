package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"log"
	"net/http"
)

func main() {

	registry := prometheus.NewRegistry()

	registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	registry.MustRegister(collectors.NewGoCollector())

	//Gauges
	temp := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "home_temperature_celsius",
		Help: "The current temperature in degrees Celsius.",
	})

	registry.MustRegister(temp)

	temp.Set(42)
	temp.Inc()
	temp.Dec()
	temp.Add(10)
	temp.Sub(10)

	//total requests
	//counter
	totalRequests := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "The total number of handled HTTP requests.",
	})

	totalRequests.Inc()
	totalRequests.Add(10)

	//histograms
	requestDurations := prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "A histogram of the HTTP request durations in seconds.",
		// Bucket configuration: the first bucket includes all requests finishing in 0.05 seconds, the last one includes all requests finishing in 10 seconds.
		Buckets: []float64{0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
	})

	requestDurations.Observe(0.42)

	//prometheus go client for tracking durations for histogram
	// Start a timer.
	timer := prometheus.NewTimer(requestDurations)

	// [...handle the request in your application here...]

	// Stop the timer and observe its duration into the "requestDurations" histogram metric.
	timer.ObserveDuration()

	//Summaries

	requestDurationsSummaries := prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "http_request_duration_seconds",
		Help: "A summary of the HTTP request durations in seconds.",
		Objectives: map[float64]float64{
			0.5:  0.05,  // 50th percentile with a max. absolute error of 0.05.
			0.9:  0.01,  // 90th percentile with a max. absolute error of 0.01.
			0.99: 0.001, // 99th percentile with a max. absolute error of 0.001.
		},
	},
	)

	requestDurationsSummaries.Observe(0.42)

	//adding labels

	tempWithLabels := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "home_temperature_celsius",
			Help: "The current temperature in degrees Celsius.",
		},
		// The two label names by which to split up the metric.
		[]string{"house", "room"},
	)

	// Set the metric value for home="julius" and room="living-room".
	tempWithLabels.WithLabelValues("julius", "living-room").Set(23.5)

	//to add explicitly
	tempWithLabels.With(prometheus.Labels{"house": "julius", "room": "living-room"}).Set(90)

	//Note that both methods panic during runtime if you pass an inconsistent set of parameters to
	//them (incorrect number of labels or incorrect label names)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
