package cost

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/uber-go/tally"
)

const (
	defaultCostExceededErrorFmt = "%s exceeds limit of %s"
	customCostExceededErrorFmt  = "%s exceeds limit of %s: %s"
)

var (
	noopManager = NewStaticLimitManager(
		NewLimitManagerOptions().
			SetDefaultLimit(Limit{
				Threshold: MaxCost,
				Enabled:   false,
			},
			),
	)
	noopEnforcer = NewEnforcer(noopManager, NewNoopTracker(), nil)
)

// Report is a report on the cost limits of an Enforcer.
type Report struct {
	Cost
	Error error
}

// Enforcer enforces cost limits for operations.
type Enforcer struct {
	LimitManager
	tracker Tracker

	costMsg string
	metrics enforcerMetrics
}

// NewEnforcer returns a new enforcer for cost limits.
func NewEnforcer(m LimitManager, t Tracker, opts EnforcerOptions) Enforcer {
	if opts == nil {
		opts = NewEnforcerOptions()
	}

	return Enforcer{
		LimitManager: m,
		tracker:      t,
		costMsg:      opts.CostExceededMessage(),
		metrics:      newEnforcerMetrics(opts.InstrumentOptions().MetricsScope(), opts.ValueBuckets()),
	}
}

// Add adds the cost of an operation to the enforcer's current total. If the operation exceeds
// the enforcer's limit the enforcer will return a CostLimit error in addition to the new total.
func (e Enforcer) Add(op Costable) (Report, error) {
	cost, err := op.Cost()
	if err != nil {
		return Report{}, fmt.Errorf("unable to calculate cost of operation: %v", err)
	}
	e.metrics.cost.RecordValue(float64(cost))
	current := e.tracker.Add(cost)
	return Report{
		Cost:  current,
		Error: e.checkLimit(current, e.Limit()),
	}, nil
}

// State returns the current state of the enforcer.
func (e Enforcer) State() (Report, Limit) {
	cost := e.tracker.Current()
	l := e.Limit()
	err := e.checkLimit(cost, l)
	r := Report{
		Cost:  cost,
		Error: err,
	}
	return r, l
}

// Clone clones the current Enforcer. The new Enforcer uses the same Estimator and LimitManager
// as e buts its Tracker is independent.
func (e Enforcer) Clone() Enforcer {
	return Enforcer{
		LimitManager: e.LimitManager,
		tracker:      NewTracker(),
		costMsg:      e.costMsg,
		metrics:      e.metrics,
	}
}

func (e Enforcer) checkLimit(cost Cost, limit Limit) error {
	if cost < limit.Threshold {
		return nil
	}

	// Emit metrics on number of operations that are over the limit even when not enabled.
	e.metrics.overLimit.Inc(1)
	if !limit.Enabled {
		return nil
	}

	e.metrics.overLimitAndEnabled.Inc(1)

	if e.costMsg == "" {
		return defaultCostExceededError(cost, limit)
	}
	return costExceededError(e.costMsg, cost, limit)
}

func defaultCostExceededError(cost Cost, limit Limit) error {
	p := message.NewPrinter(language.Make("en"))

	return fmt.Errorf(
		defaultCostExceededErrorFmt,
		p.Sprintf("%v", float64(cost)),
		p.Sprintf("%v", float64(limit.Threshold)),
	)
}

func costExceededError(customMessage string, cost Cost, limit Limit) error {
	p := message.NewPrinter(language.Make("en"))

	return fmt.Errorf(
		customCostExceededErrorFmt,
		p.Sprintf("%v", float64(cost)),
		p.Sprintf("%v", float64(limit.Threshold)),
		customMessage,
	)
}

// NoopEnforcer returns a new Enforcer that always returns a current cost of 0 and
//  is always disabled.
func NoopEnforcer() Enforcer {
	return noopEnforcer
}

type enforcerMetrics struct {
	cost                tally.Histogram
	overLimit           tally.Counter
	overLimitAndEnabled tally.Counter
}

func newEnforcerMetrics(s tally.Scope, b tally.ValueBuckets) enforcerMetrics {
	return enforcerMetrics{
		cost:                s.Histogram("estimate", b),
		overLimit:           s.Counter("over-limit"),
		overLimitAndEnabled: s.Counter("over-limit-and-enabled"),
	}
}
