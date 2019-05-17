package models

// Register data structure
type Register struct {
	Entity
	Title       string `json:"title"`
	Category    string `json:"category"`
	SubCategory string `json:"subcategory"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

// NewRegister constructor
func NewRegister() *Register {
	e := new(Register)
	e.Type = "register"
	return e
}

func (e *Register) parse(intf map[string]interface{}) {
	e.parseEntity(intf)
	e.Title = intf["title"].(string)
	e.Category = intf["category"].(string)
	e.SubCategory = intf["subcategory"].(string)
	e.Description = intf["description"].(string)
	e.Private = intf["private"].(bool)
}

// GetByID get Entity by ID
func (e *Register) GetByID() error {
	entity, err := DB.Get(e.ID, nil)
	if err != nil {
		return err
	}

	e.parse(entity)

	return nil
}

// Get by attribute and value
func (e *Register) Get(attr string, value string) error {
	resp, err := e.getBy(attr, value)

	if err != nil {
		return err
	}

	e.parse(resp)
	return nil
}
