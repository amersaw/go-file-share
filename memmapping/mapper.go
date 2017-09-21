package memmapping

import (
	"fmt"

	"github.com/amersaw/go-file-share/datamapping"
)

//Mapper struct is in-memory mapper
type Mapper struct {
	items map[int64]datamapping.Entity
}

//NewMapper function will initialize a new MemoryDataMapper object
func NewMapper() Mapper {

	return Mapper{
		items: map[int64]datamapping.Entity{},
	}
}

//Get function retreive the item with id equals to id
func (m Mapper) Get(id int64) datamapping.Entity {
	i, ok := m.items[id]
	if !ok {
		return nil
	}
	return i
}

//Post funcation allows adding an item to the memory data mapper
func (m Mapper) Post(entity datamapping.Entity) bool {
	m.items[entity.GetID()] = entity
	fmt.Printf("Added %d\n", entity.GetID())
	return true
}
