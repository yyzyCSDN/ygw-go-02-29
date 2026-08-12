package query

import (
	"fmt"
	"sort"
	"strings"

	"example.com/requestgateway/model"
)

type Filter struct {
	Status      string
	Tag         string
	Active      *bool
	MinPriority *int
	MaxPriority *int
	Name        string
}

type SortField string

const (
	SortByID       SortField = "id"
	SortByName     SortField = "name"
	SortByPriority SortField = "priority"
	SortByAmount   SortField = "amount"
)

func Select(values []model.Request, filter Filter) []model.Request {
	result := make([]model.Request, 0)
	for _, value := range values {
		if filter.Status != "" && value.Status != filter.Status {
			continue
		}
		if filter.Tag != "" && !value.HasTag(filter.Tag) {
			continue
		}
		if filter.Active != nil && value.Active != *filter.Active {
			continue
		}
		if filter.MinPriority != nil && value.Priority < *filter.MinPriority {
			continue
		}
		if filter.MaxPriority != nil && value.Priority > *filter.MaxPriority {
			continue
		}
		if filter.Name != "" && !strings.Contains(strings.ToLower(value.Name), strings.ToLower(filter.Name)) {
			continue
		}
		result = append(result, value.Clone())
	}
	return result
}

func Sort(values []model.Request, field SortField, descending bool) ([]model.Request, error) {
	result := make([]model.Request, len(values))
	for index, value := range values {
		result[index] = value.Clone()
	}
	less := func(i, j int) bool { return false }
	switch field {
	case SortByID:
		less = func(i, j int) bool { return result[i].ID < result[j].ID }
	case SortByName:
		less = func(i, j int) bool { return result[i].Name < result[j].Name }
	case SortByPriority:
		less = func(i, j int) bool { return result[i].Priority < result[j].Priority }
	case SortByAmount:
		less = func(i, j int) bool { return result[i].Amount < result[j].Amount }
	default:
		return nil, fmt.Errorf("unsupported sort field %q", field)
	}
	sort.SliceStable(result, func(i, j int) bool {
		if descending {
			return less(j, i)
		}
		return less(i, j)
	})
	return result, nil
}

func Page(values []model.Request, offset, limit int) ([]model.Request, error) {
	if offset < 0 || limit < 0 {
		return nil, fmt.Errorf("offset and limit cannot be negative")
	}
	if offset >= len(values) || limit == 0 {
		return []model.Request{}, nil
	}
	end := offset + limit
	if end > len(values) {
		end = len(values)
	}
	result := make([]model.Request, end-offset)
	for index := range result {
		result[index] = values[offset+index].Clone()
	}
	return result, nil
}

func GroupByStatus(values []model.Request) map[string][]model.Request {
	result := map[string][]model.Request{}
	for _, value := range values {
		result[value.Status] = append(result[value.Status], value.Clone())
	}
	return result
}

func IDs(values []model.Request) []string {
	result := make([]string, len(values))
	for index, value := range values {
		result[index] = value.ID
	}
	return result
}
