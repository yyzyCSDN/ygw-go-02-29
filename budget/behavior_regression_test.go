package budget

import (
	"testing"
	"time"
)

func TestRegressionBehavior(t *testing.T) {
	if got := Budget(2*time.Second, 5*time.Second); got != 2*time.Second {
		t.Fatalf("Budget=%v", got)
	}
}
