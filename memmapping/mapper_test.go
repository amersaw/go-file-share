package memmapping

import (
	"testing"
)

//TestEntity a simple struct for testing the data mappers
type testEntity struct {
	ID    int64
	Value string
}

//GetID method returns the record's ID
func (fr testEntity) GetID() int64 {
	return fr.ID
}

func TestMemoryDataMapper(t *testing.T) {
	mapper := NewMapper()

	r := testEntity{ID: 1, Value: "The file Record"}
	mapper.Post(r)

	r2 := testEntity{ID: 2, Value: "The second file Record"}
	mapper.Post(r2)

	element := mapper.Get(1)
	original, ok := element.(testEntity)

	if !ok {
		t.Fatal("Not existed element")
		if original.GetID() != 1 || original.Value != r.Value {
			t.Error("values are not correct")
		}
	}

}
