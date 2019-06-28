package models

import (
	"encoding/json"
)

// Register data structure
type Register struct {
	Entity
	Title       string `json:"title"`
	Category    string `json:"category"`
	SubCategory string `json:"subcategory"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
	Owner       string `json:"owner"`
	Cards       []Card `json:"cards"`
	Misc        map[string]interface{}
}

// NewRegister constructor
func NewRegister() *Register {
	e := new(Register)
	e.Type = "register"
	e.Misc = make(map[string]interface{})
	return e
}

func (e *Register) parse(intf map[string]interface{}) {
	data, _ := json.Marshal(intf)
	json.Unmarshal(data, e)
}

// GetAllRegister gets all registers
func (e *Register) GetAllRegister() []Register {
	eListMap, _ := e.GetAll()
	eList := make([]Register, len(eListMap))

	for index, element := range eListMap {
		parsed := NewRegister()
		parsed.parse(element)
		eList[index] = *parsed
	}

	return eList
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
