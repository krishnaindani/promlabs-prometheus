## Grafana Labs Loki Workshop

## Current challenges 
- Hard to operate at scale 
- Expensive to scale(hardware, licenses)
- Doesn't correlate well outside vendor stack. Navigation between logs and metrics and traces.

## Loki 
- Cost-effective at scale. Highly scalable and smart indexing
- Easy to operate and scale to maintain at large scale. (Uses object storage)
- Seamless integration with prometheus.(label correlation, alerting, service discovery) (correlate metrics, traces and logs)
- Logs as metrics, alerting and predictions. 
- Format agnostic accept all logs formats(json, regex, logfmt) and log structures.
- Loki accepts logs from all places
- Promtail agent sends logs to loki. Loki stores in objest storage. 
  - Grafana 
  - REST API 
  - log cli 
  - Alertmanager 
- Grafana agent has promtail wrapped in it 
- Attach labels to log lines 
- Advanced pipeline mechanism to filter logs and parsing, transforming 
- Build and expose custom metrics from your logs data. 
- Loki Architecture 
  - Simple Binary 
  - Simple scalable deployment 
  - Microservices
- Loki is line prometheus but not for logs. 
- Prometheus model 
  - timestamp, metric_name, label/selectors, metric_value - prometheus model 
  - timestamp, label/selectors, content log line - loki logs model 
  - Timestamp and label selectors are indexed. (10TB for logs data needs 200MB index)
- Log stream is a stream of log entries with the same labels 
- What makes a good label?(at ingest time) same line prometheus. 
- Query processing 
  - Log any and all formats 
  - Smaller indexes 
  - Cheaper to run 
  - Fast queries 
  - Cut and slice your logs in dynamic ways 
- Ad hoc metrics from logs 
  - Metrics from non indexed fields 
  - Recording rules for loki to store the metric as prometheus metric
  - 