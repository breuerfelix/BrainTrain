package models

// User data structure
type User struct {
	Entity
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Date     string `json:"date"`
}

func (u *User) parse(intf map[string]interface{}) {
	u.parseEntity(intf)
	u.Email = intf["email"].(string)
	u.Name = intf["name"].(string)
	u.Password = intf["password"].(string)
}

// GetByID a user by ID
func (u *User) GetByID() error {
	user, err := db.Get(u.ID, nil)
	if err != nil {
		return err
	}

	u.parse(user)

	return nil
}

// Get by attribute and value
func (u *User) Get(attr string, value string) error {
	resp, err := u.GetBy(attr, value)

	if err != nil {
		return err
	}

	u.parse(resp)
	return nil
}
