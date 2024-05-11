package avoidEmeddingTypesInPublicStructs

// ConcreteList is a list of entities.
type ConcreteListBad struct {
	*AbstractList
}
