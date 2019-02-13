package UIModels

type Project struct {
	ID          uint
	Title       string
	Description string
	Skills      []string
	BudgetType  uint
	BudgetLevel uint
}
