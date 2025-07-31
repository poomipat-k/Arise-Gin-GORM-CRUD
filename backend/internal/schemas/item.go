package schemas

// represent the input data for the creation of item
type CreateItemSchemaInput struct {
	Name string `json:"name" binding:"required,min=1,max=100"`
}
