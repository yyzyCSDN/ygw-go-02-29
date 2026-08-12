package batch

import (
	"testing"

	"example.com/requestgateway/model"
)

func TestApplyOperations(t *testing.T) {
	current := []model.Request{{ID: "one", Name: "One"}, {ID: "two", Name: "Two"}}
	result, err := Apply(current, []Operation{
		{Kind: Upsert, Value: model.Request{ID: "one", Name: "Updated"}},
		{Kind: Delete, ID: "two"},
		{Kind: Upsert, Value: model.Request{ID: "three", Name: "Three"}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Created != 1 || result.Updated != 1 || result.Deleted != 1 || len(result.Values) != 2 {
		t.Fatalf("unexpected result: %#v", result)
	}
	if current[0].Name != "One" {
		t.Fatal("Apply mutated input")
	}
}

func TestApplyFailureReturnsNoPartialResult(t *testing.T) {
	result, err := Apply([]model.Request{{ID: "one", Name: "One"}}, []Operation{
		{Kind: Upsert, Value: model.Request{ID: "two", Name: "Two"}},
		{Kind: Delete, ID: "missing"},
	})
	if err == nil {
		t.Fatal("expected operation error")
	}
	if len(result.Values) != 0 || result.Created != 0 {
		t.Fatalf("partial result leaked: %#v", result)
	}
	if err := ValidateUnique([]model.Request{{ID: "x"}, {ID: "x"}}); err == nil {
		t.Fatal("expected duplicate error")
	}
}
