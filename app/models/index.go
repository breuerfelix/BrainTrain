package models

import (
	"encoding/json"
	"fmt"

	couchdb "github.com/leesper/couchdb-golang"
)

// global package db handle
var db *couchdb.Database

func init() {
	var err error
	db, err = couchdb.NewDatabase("http://localhost:5984/braintrain")
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to DB")
}

// Entity is the base database item struct
type Entity struct {
	ID   string `json:"_id"`
	Rev  string `json:"_rev"`
	Type string `json:"type"`
	couchdb.Document
}

func (e *Entity) toMap() map[string]interface{} {
	var doc map[string]interface{}
	tJSON, _ := json.Marshal(e)
	json.Unmarshal(tJSON, &doc)
	return doc
}

func (e *Entity) parseEntity(intf map[string]interface{}) {
	e.ID = intf["_id"].(string)
	e.Rev = intf["_rev"].(string)
	e.Type = intf["type"].(string)
}

// Add a new entity to the db
func (e *Entity) Add() error {
	eMap := e.toMap()

	// Delete _id and _rev from map, otherwise DB access will be denied (unauthorized)
	delete(eMap, "_id")
	delete(eMap, "_rev")

	// Add todo to DB
	_, _, err := db.Save(eMap, nil)

	if err != nil {
		fmt.Printf("[Add] error: %s", err)
	}

	return err
}

// GetAll entity from db
func (e *Entity) GetAll() ([]map[string]interface{}, error) {
	allEntities, err := db.QueryJSON(fmt.Sprintf(`{
		"selector": {
			"type": {
				"$eq": "%s"
			}
		}
	}`, e.Type))

	if err != nil {
		return nil, err
	}

	return allEntities, nil
}

// GetBy  entity from db
func (e *Entity) GetBy(attr string, value string) (map[string]interface{}, error) {
	entity, err := db.QueryJSON(fmt.Sprintf(`{
		"selector": {
			"%s": {
				"$eq": "%s"
			}
		}
	}`, attr, value))

	if err != nil {
		return nil, err
	}

	return entity[0], nil
}
