package datamapping

//Entity interface shows how storable entities should be
type Entity interface {
	GetID() int64
}
