package analytics

type ClickMetric struct {
	Group      string `json:"group"`
	GroupID    string `json:"group_id"`
	GroupLabel string `json:"group_label"`
	Clicks     int64  `json:"clicks"`
}
