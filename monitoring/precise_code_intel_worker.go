package main

func PreciseCodeIntelWorker() *Container {
	return &Container{
		Name:        "precise-code-intel-worker",
		Title:       "Precise Code Intel Worker",
		Description: "Handles conversion of uploaded precise code intelligence bundles.",
		Groups: []Group{
			{
				Title: "General",
				Rows: []Row{
					{
						{
							Name:            "queue_size",
							Description:     "queue size",
							Query:           `max(src_precise_code_intel_worker_queue_size)`,
							DataMayNotExist: true,
							Warning:         Alert{LessOrEqual: 25},
							Critical:        Alert{LessOrEqual: 15},
							PanelOptions:    PanelOptions().LegendFormat("{{instance}}"),
							PossibleSolutions: `
								TODO - update this
								- **Provision more disk space:** Sourcegraph will begin deleting least-used repository clones at 10% disk space remaining which may result in decreased performance, users having to wait for repositories to clone, etc.
							`,
						},
						{
							Name:              "job_total",
							Description:       "job total every 5m",
							Query:             `sum(increase(src_precise_code_intel_worker_jobs_total[5m]))`,
							DataMayNotExist:   true,
							Warning:           Alert{GreaterOrEqual: 5},
							Critical:          Alert{GreaterOrEqual: 20},
							PanelOptions:      PanelOptions().LegendFormat("errors"),
							PossibleSolutions: "none",
						},
						{
							Name:              "job_errors",
							Description:       "job errors every 5m",
							Query:             `sum(increase(src_precise_code_intel_worker_jobs_errors_total[5m]))`,
							DataMayNotExist:   true,
							Warning:           Alert{GreaterOrEqual: 5},
							Critical:          Alert{GreaterOrEqual: 20},
							PanelOptions:      PanelOptions().LegendFormat("errors"),
							PossibleSolutions: "none",
						},
					},
					// TODO - bundle manager stuff
					// TODO - gitserver
					{
						{
							Name:        "99th_percentile_db_duration",
							Description: "99th percentile successful db query duration over 5m",
							// TODO - by op?
							Query:             `histogram_quantile(0.99, sum by (le)(rate(src_precise_code_intel_worker_db_duration_seconds_bucket[5m])))`,
							DataMayNotExist:   true,
							Warning:           Alert{GreaterOrEqual: 20},
							PanelOptions:      PanelOptions().LegendFormat("duration").Unit(Seconds),
							PossibleSolutions: "none",
						},
						{
							Name:              "db_errors",
							Description:       "db errors every 5m",
							Query:             `sum(sum by (op)(increase(src_precise_code_intel_worker_db_errors_total[5m])))`,
							DataMayNotExist:   true,
							Warning:           Alert{GreaterOrEqual: 5},
							Critical:          Alert{GreaterOrEqual: 20},
							PanelOptions:      PanelOptions().LegendFormat("errors"),
							PossibleSolutions: "none",
						},
					},
					{
						sharedFrontendInternalAPIErrorResponses("precise-code-intel-worker"),
					},
				},
			},
			{
				Title:  "Container monitoring (not available on k8s or server)",
				Hidden: true,
				Rows: []Row{
					{
						sharedContainerRestarts("precise-code-intel-worker"),
						sharedContainerMemoryUsage("precise-code-intel-worker"),
						sharedContainerCPUUsage("precise-code-intel-worker"),
					},
				},
			},
		},
	}
}
