package api

type Meta struct {
	TotalCount    int64 `json:"total_count,omitempty"`
	FilteredCount int64 `json:"filtered_count,omitempty"`
}
