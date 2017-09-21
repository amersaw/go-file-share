package memmapping

//FRecord type include a simple information about a file record
type FRecord struct {
	ID   int64
	Name string
}

//NewFRecord returns a new FRecord object
func NewFRecord(id int64, name string) FRecord {
	return FRecord{ID: id, Name: name}
}

//GetID method returns the record's ID
func (fr FRecord) GetID() int64 {
	return fr.ID
}
