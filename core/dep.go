package core

import (
	"strings"
)

const (
	// MetricLimit is the name of the metric for current limit
	MetricLimit = "limit"
	// MetricDropped is the name of the metric for dropped counts
	MetricDropped = "dropped"
	// MetricInFlight is the name of the metric for current in flight count
	MetricInFlight = "inflight"
	// MetricPartitionLimit is the name of the metric for a current partition's limit
	MetricPartitionLimit = "limit.partition"
	// MetricMinRTT is the name of the metric for the Minimum Round Trip Time
	MetricMinRTT = "min_rtt"
	// MetricWindowMinRTT is the name of the metric for the Window's Minimum Round Trip Time
	MetricWindowMinRTT = "window.min_rtt"
	// MetricWindowQueueSize represents the name of hte metric for the Window's Queue Size
	MetricWindowQueueSize = "window.queue_size"
)

// PrefixMetricWithName will prefix a given name with the metric name in the form "<name>.<metric>"
func PrefixMetricWithName(metric, name string) string {
	if strings.HasSuffix(name, ".") {
		return name + metric
	}
	return name + "." + metric
}
