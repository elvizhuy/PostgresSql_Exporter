package metrics

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
)

type Metric interface {
	Scrape(*sql.DB) error
	Collect(chan<- prometheus.Metric)
	Describe(chan<- *prometheus.Desc)
}
