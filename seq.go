package genapp

type Sequence []interface{}

var H = NewSequence

func NewSequence(handlers ...interface{}) Sequence {
	return handlers
}
