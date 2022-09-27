package lru_test

type String string

func (d String) Len() int {
	return len(d)
}
