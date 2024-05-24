package _struct

// ConcreteList is a list of entities.
type ConcreteListGood struct {
	list AbstractList
}

// Add adds an entity to the list.
func (l *ConcreteListGood) Add(e int) {
	l.list.Add(e)
}

// Remove removes an entity from the list.
func (l *ConcreteListGood) Remove(e int) {
	l.list.Remove(e)
}
