package diff

import (
	"data-comparision/internal/parser/models"
	"reflect"
)

func DiffItems(prev, curr []models.LeasingItem) []ItemDiff {
	result := make([]ItemDiff, 0)

	prevMap := make(map[string]models.LeasingItem)
	currMap := make(map[string]models.LeasingItem)

	for _, p := range prev {
		prevMap[p.BusinessKey] = p
	}
	for _, c := range curr {
		currMap[c.BusinessKey] = c
	}

	// NEW и UPDATED
	for key, currItem := range currMap {
		prevItem, exists := prevMap[key]

		if !exists {
			result = append(result, ItemDiff{
				BusinessKey: key,
				Type:        ChangeNew,
			})
			continue
		}

		fields := compareFields(prevItem, currItem)
		if len(fields) > 0 {
			result = append(result, ItemDiff{
				BusinessKey: key,
				Type:        ChangeUpdated,
				Fields:      fields,
			})
		}
	}

	// REMOVED
	for key := range prevMap {
		if _, exists := currMap[key]; !exists {
			result = append(result, ItemDiff{
				BusinessKey: key,
				Type:        ChangeRemoved,
			})
		}
	}

	return result
}

func compareFields(a, b models.LeasingItem) []FieldDiff {
	var diffs []FieldDiff

	if !floatEqual(a.ApprovedPrice, b.ApprovedPrice) {
		diffs = append(diffs, FieldDiff{
			Field:    "approved_price",
			OldValue: a.ApprovedPrice,
			NewValue: b.ApprovedPrice,
		})
	}

	if a.Status != b.Status {
		diffs = append(diffs, FieldDiff{
			Field:    "status",
			OldValue: a.Status,
			NewValue: b.Status,
		})
	}

	if a.Location != b.Location {
		diffs = append(diffs, FieldDiff{
			Field:    "location",
			OldValue: a.Location,
			NewValue: b.Location,
		})
	}

	// JSON-поля
	for k, v := range a.Data {
		if !reflect.DeepEqual(v, b.Data[k]) {
			diffs = append(diffs, FieldDiff{
				Field:    k,
				OldValue: v,
				NewValue: b.Data[k],
			})
		}
	}

	for k, v := range b.Data {
		if _, exists := a.Data[k]; !exists {
			diffs = append(diffs, FieldDiff{
				Field:    k,
				OldValue: nil,
				NewValue: v,
			})
		}
	}

	return diffs
}

func floatEqual(a, b *float64) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}
