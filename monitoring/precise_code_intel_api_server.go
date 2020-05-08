package main

func PreciseCodeIntelAPIServer() *Container {
	return &Container{
		Name:        "precise-code-intel-api-server",
		Title:       "Precise Code Intel API Server",
		Description: "Serves precise code intelligence requests.",
		Groups: []Group{
			{
				Title: "General",
				Rows: []Row{
					{
						{
							Name:        "99th_percentile_code_intel_api_duration",
							Description: "99th percentile successful code intel api query duration over 5m",
							// TODO - by op?
							// TODO - where no error
							Query:             `histogram_quantile(0.99, sum by (le)(rate(src_precise_code_intel_api_server_code_intel_api_duration_seconds_bucket[5m])))`,
							DataMayNotExist:   true,
							Warning:           Alert{GreaterOrEqual: 20},
							PanelOptions:      PanelOptions().LegendFormat("duration").Unit(Seconds),
							PossibleSolutions: "none",
						},
						{
							Name:              "code_intel_api_errors",
							Description:       "code intel api errors every 5m",
							Query:             `sum(sum by (op)(increase(src_precise_code_intel_api_server_code_intel_api_errors_total[5m])))`,
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
							Query:             `histogram_quantile(0.99, sum by (le)(rate(src_precise_code_intel_api_server_db_duration_seconds_bucket[5m])))`,
							DataMayNotExist:   true,
							Warning:           Alert{GreaterOrEqual: 20},
							PanelOptions:      PanelOptions().LegendFormat("duration").Unit(Seconds),
							PossibleSolutions: "none",
						},
						{
							Name:              "db_errors",
							Description:       "db errors every 5m",
							Query:             `sum(sum by (op)(increase(src_precise_code_intel_api_server_db_errors_total[5m])))`,
							DataMayNotExist:   true,
							Warning:           Alert{GreaterOrEqual: 5},
							Critical:          Alert{GreaterOrEqual: 20},
							PanelOptions:      PanelOptions().LegendFormat("errors"),
							PossibleSolutions: "none",
						},
					},
					{
						{
							Name:              "resetter_errors",
							Description:       "resetter errors every 5m",
							Query:             `sum(sum by (op)(increase(src_precise_code_intel_api_server_resetter_errors[5m])))`,
							DataMayNotExist:   true,
							Warning:           Alert{GreaterOrEqual: 5},
							Critical:          Alert{GreaterOrEqual: 20},
							PanelOptions:      PanelOptions().LegendFormat("errors"),
							PossibleSolutions: "none",
						},
						{
							Name:              "resetter_old_dumps",
							Description:       "resetter stalled jobs every 5m",
							Query:             `sum(sum by (op)(increase(src_precise_code_intel_api_server_resetter_stalled_jobs[5m])))`,
							DataMayNotExist:   true,
							Warning:           Alert{GreaterOrEqual: 5},
							Critical:          Alert{GreaterOrEqual: 20},
							PanelOptions:      PanelOptions().LegendFormat("errors"),
							PossibleSolutions: "none",
						},
					},
					{
						sharedFrontendInternalAPIErrorResponses("precise-code-intel-api-server"),
					},
				},
			},
			{
				Title:  "Container monitoring (not available on k8s or server)",
				Hidden: true,
				Rows: []Row{
					{
						sharedContainerRestarts("precise-code-intel-api-server"),
						sharedContainerMemoryUsage("precise-code-intel-api-server"),
						sharedContainerCPUUsage("precise-code-intel-api-server"),
					},
				},
			},
		},
	}
}
