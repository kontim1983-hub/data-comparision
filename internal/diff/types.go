package diff

type ChangeType string

const (
	ChangeNew     ChangeType = "NEW"
	ChangeRemoved ChangeType = "REMOVED"
	ChangeUpdated ChangeType = "UPDATED"
)

type FieldDiff struct {
	Field    string      `json:"field"`
	OldValue interface{} `json:"old"`
	NewValue interface{} `json:"new"`
}

type ItemDiff struct {
	BusinessKey string      `json:"businessKey"`
	Type        ChangeType  `json:"type"`
	Fields      []FieldDiff `json:"fields,omitempty"`
}
