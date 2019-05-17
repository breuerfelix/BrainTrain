package models

// Card data structure
type Card struct {
	Entity
	Title    string `json:"title"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Progress int    `json:"progress"`
	Register string `json:"register"`
}

// NewCard constructor
func NewCard() *Card {
	e := new(Card)
	e.Type = "card"
	return e
}

func (e *Card) parse(intf map[string]interface{}) {
	e.parseEntity(intf)
	e.Title = intf["title"].(string)
	e.Question = intf["question"].(string)
	e.Answer = intf["answer"].(string)
	e.Progress = intf["progress"].(int)
	e.Register = intf["register"].(string)
}

// GetByID get Entity by ID
func (e *Card) GetByID() error {
	entity, err := DB.Get(e.ID, nil)
	if err != nil {
		return err
	}

	e.parse(entity)

	return nil
}

// Get by attribute and value
func (e *Card) Get(attr string, value string) error {
	resp, err := e.getBy(attr, value)

	if err != nil {
		return err
	}

	e.parse(resp)
	return nil
}
