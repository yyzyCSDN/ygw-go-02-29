package service

import (
	"testing"

	"example.com/requestgateway/batch"
	"example.com/requestgateway/model"
	"example.com/requestgateway/query"
)

func TestServiceWorkflow(t *testing.T) {
	service := New()
	if _, err := service.Save("alice", model.Request{ID: "one", Name: "First", Status: "ready", Priority: 2}); err != nil {
		t.Fatal(err)
	}
	snapshot := service.Capture("before batch")
	err := service.Apply("bob", []batch.Operation{
		{Kind: batch.Upsert, Value: model.Request{ID: "one", Name: "Updated", Status: "ready", Priority: 3}},
		{Kind: batch.Upsert, Value: model.Request{ID: "two", Name: "Second", Status: "new", Priority: 1}},
	})
	if err != nil {
		t.Fatal(err)
	}
	values, err := service.List(query.Filter{}, query.SortByPriority, true)
	if err != nil || len(values) != 2 || values[0].ID != "one" {
		t.Fatalf("List = %#v, %v", values, err)
	}
	if err := service.Restore("alice", snapshot.ID); err != nil {
		t.Fatal(err)
	}
	if _, exists := service.Get("two"); exists {
		t.Fatal("restore retained later value")
	}
	if len(service.Audit()) != 4 {
		t.Fatalf("unexpected audit: %#v", service.Audit())
	}
}

func TestServiceRejectsInvalidOperations(t *testing.T) {
	service := New()
	if _, err := service.Save("alice", model.Request{}); err == nil {
		t.Fatal("expected validation error")
	}
	if err := service.Remove("alice", "missing"); err == nil {
		t.Fatal("expected missing error")
	}
	if err := service.Restore("alice", 99); err == nil {
		t.Fatal("expected snapshot error")
	}
}
