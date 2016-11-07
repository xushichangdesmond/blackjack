package sliceView

type View interface {
	Len() int
	At(index int) (element interface{}, ok bool)
}
