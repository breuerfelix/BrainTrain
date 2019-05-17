package models

// User data structure
type User struct {
	Entity
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Date     string `json:"date"`
}

func (e *User) parse(intf map[string]interface{}) {
	e.parseEntity(intf)
	e.Email = intf["email"].(string)
	e.Name = intf["name"].(string)
	e.Password = intf["password"].(string)
}

// GetByID a Entity by ID
func (e *User) GetByID() error {
	entity, err := DB.Get(e.ID, nil)
	if err != nil {
		return err
	}

	e.parse(entity)

	return nil
}

// Get by attribute and value
func (e *User) Get(attr string, value string) error {
	resp, err := e.getBy(attr, value)

	if err != nil {
		return err
	}

	e.parse(resp)
	return nil
}
