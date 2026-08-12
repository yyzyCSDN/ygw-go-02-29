package budget

import (
	"testing"
	"time"
)

func TestBudgetShortRequest(t *testing.T) {
	if Budget(5*time.Second, time.Second) != time.Second {
		t.Fatal("short request changed")
	}
}
