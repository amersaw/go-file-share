package datamapping

//DataMapper is an interface that shows how basic
//data mappers should be implemented
type DataMapper interface {
	Get(int64) Entity
	Post(Entity) bool
	CreateNew() DataMapper
}
