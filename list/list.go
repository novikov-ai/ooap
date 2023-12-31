package list

type List struct {
	list []any
}

func New() *List {
	return &List{
		list: []any{},
	}
}

func (l *List) RemoveAt(index int) {
	if len(l.list) <= index {
		return
	}

	if index == -1 {
		l.list = l.list[0 : len(l.list)-1]
	} else {
		l.list = append(l.list[:index], l.list[index+1:len(l.list)]...)
	}
}

func (l *List) Append(value any) {
	l.list = append(l.list, value)
}

func (l *List) Length() int {
	return len(l.list)
}

func (l *List) IndexAt(index int) any {
	if index == -1 {
		index = l.Length() - 1
	}

	return l.list[index]
}

func (l *List) InsertAfter(index int, value any) {
	if index >= l.Length() {
		return
	}

	left := l.list[:index+1]
	right := l.list[index+1:]

	updatedList := make([]any, 0, l.Length()+1)

	for _, item := range left {
		updatedList = append(updatedList, item)
	}

	updatedList = append(updatedList, value)

	for _, item := range right {
		updatedList = append(updatedList, item)
	}

	l.list = updatedList
}
