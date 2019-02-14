package UIModels

type Project struct {
	ID          uint     `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Skills      []string `json:"skills"`
	BudgetType  uint     `json:"budget_type"`
	BudgetLevel uint     `json:"budget_level"`
	User        User     `json:"user"`
}
