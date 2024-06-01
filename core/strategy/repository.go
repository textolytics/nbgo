package strategy

type StrategyMetrics struct {
}

type StrategyMetricsRepository interface {
	Store(strategyMetrics *StrategyMetrics) error
}
