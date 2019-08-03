package maps

type Dictionary map[string]string

func (dic *Dictionary) Search(word string) string {
	return (*dic)[word]
}
