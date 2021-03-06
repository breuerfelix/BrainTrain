package models

import (
	"encoding/json"
	"fmt"
)

// User data structure
type User struct {
	Entity
	Email    string             `json:"email"`
	Name     string             `json:"name"`
	Password string             `json:"password"`
	Date     string             `json:"date"`
	Progress []RegisterProgress `json:"progress"`
	Picture  string             `json:"picture"`
}

// NewUser constructor
func NewUser() *User {
	e := new(User)
	e.Type = "user"
	return e
}

func (e *User) parse(intf map[string]interface{}) {
	data, _ := json.Marshal(intf)
	json.Unmarshal(data, e)
}

// GetAllUser gets all registers
func (e *User) GetAllUser() []User {
	eListMap, _ := e.GetAll()

	eList := make([]User, len(eListMap))

	for index, element := range eListMap {
		parsed := NewUser()
		parsed.parse(element)
		eList[index] = *parsed
	}

	return eList
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

// Insert new user
func (e *User) Insert() error {
	eMap := toMap(e)

	// Delete _id and _rev from map, otherwise DB access will be denied (unauthorized)
	delete(eMap, "_id")
	delete(eMap, "_rev")

	_, _, err := DB.Save(eMap, nil)

	if err != nil {
		fmt.Printf("[Add] error: %s", err)
	}

	return err
}

// Update user
func (e *User) Update() error {
	eMap := toMap(e)

	_, _, err := DB.Save(eMap, nil)

	if err != nil {
		fmt.Printf("[Add] error: %s", err)
	}

	return err
}

func (e *User) Delete() error {
	register := NewRegister()

	registers := register.GetAllRegister()

	userRegisters := make([]Register, 0)

	for _, element := range registers {
		if element.Owner == e.ID {
			userRegisters = append(userRegisters, element)
		}
	}

	allUsers := e.GetAllUser()

	for _, otherUser := range allUsers {
		copyOfProgress := append(otherUser.Progress[:0:0], otherUser.Progress...)
		for i := len(otherUser.Progress)-1; i >= 0; i-- {
			registerID := otherUser.Progress[i].Register
			for _, delReg := range userRegisters {
				if delReg.ID == registerID {
					s := copyOfProgress
					s[i] = s[len(s)-1]
					copyOfProgress = s[:len(s)-1]
				}
			}
		}
		otherUser.Progress = copyOfProgress
		otherUser.Update()
	}

	for _, reg := range userRegisters {
		reg.Delete()
	}

	err := DB.Delete(e.ID)

	return err
}