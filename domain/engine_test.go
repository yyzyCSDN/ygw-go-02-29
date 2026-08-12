package domain

import (
	"encoding/json"
	"testing"
)

func TestGatewayQueueLifecycle(t *testing.T) {
	engine := NewGatewayQueue()
	if err := engine.OpenRoute(); err != nil {
		t.Fatal(err)
	}
	if err := engine.RegisterRequest(GatewayQueueRecord{ID: "primary", Quantity: 4, Labels: map[string]string{"zone": "a"}}); err != nil {
		t.Fatal(err)
	}
	if err := engine.DispatchRequest("primary", 3); err != nil {
		t.Fatal(err)
	}
	if err := engine.CancelRequest("primary", 2); err != nil {
		t.Fatal(err)
	}
	if got := engine.CountPending(); got != 5 {
		t.Fatalf("count = %d; want 5", got)
	}
	if err := engine.CloseRoute(); err != nil {
		t.Fatal(err)
	}
}

func TestGatewayQueuePrioritiesAndExport(t *testing.T) {
	engine := NewGatewayQueue()
	_ = engine.RegisterRequest(GatewayQueueRecord{ID: "low", Quantity: 1})
	_ = engine.RegisterRequest(GatewayQueueRecord{ID: "high", Quantity: 2})
	if err := engine.PrioritizeRequest("high", 9); err != nil {
		t.Fatal(err)
	}
	values := engine.List()
	if len(values) != 2 || values[0].ID != "high" {
		t.Fatalf("unexpected order: %#v", values)
	}
	values[0].Labels = map[string]string{"changed": "yes"}
	data, err := engine.ExportRequests()
	if err != nil || !json.Valid(data) {
		t.Fatalf("invalid export: %s, %v", data, err)
	}
}

func TestGatewayQueueRejectsInvalidOperations(t *testing.T) {
	engine := NewGatewayQueue()
	if err := engine.RegisterRequest(GatewayQueueRecord{}); err == nil {
		t.Fatal("expected blank id error")
	}
	if err := engine.RegisterRequest(GatewayQueueRecord{ID: "one", Quantity: 1}); err != nil {
		t.Fatal(err)
	}
	if err := engine.RegisterRequest(GatewayQueueRecord{ID: "one"}); err == nil {
		t.Fatal("expected duplicate error")
	}
	if err := engine.CancelRequest("one", 2); err == nil {
		t.Fatal("expected insufficient quantity error")
	}
	if err := engine.PrioritizeRequest("missing", 1); err == nil {
		t.Fatal("expected missing record error")
	}
}
