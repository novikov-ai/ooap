# Задание 2.

## 2.1. Опишите АТД LinkedList с предложенным набором операций.

~~~
type AbstractLinkedList interface {
// Commands:

	// Precondition: list is not empty
	// Postcondition: cursor set on the first node in the list
	Head()

        // Precondition: list is not empty
	// Postcondition: cursor set on the last node in the list
	Tail()

	// Precondition: list is not empty and next node exists
	// Postcondition: cursor shifted 1 node right
	Right()

	// Precondition: current node exists
	// Postcondition: new node with value inserted after current node
	PutRight(value any)

	// Precondition: current node exists
	// Postcondition: new node with value inserted before current node
	PutLeft(value any)

	// Precondition: current node exists
	// Postcondition: current node removed and cursor shifted right if node exists or left
	Remove()

	// Postcondition: linked list is empty
	Clear()

	// Postcondition: new node inserted in tail
	AddTail(value any)

	// Precondition: current node exists
	// Postcondition: current node value replaced with given
	Replace(value any)

	// Postcondition: cursor set to the next node with given value if exists
	Find(value any)

	// Postcondition: all nodes which equal given value removed
	RemoveAll(value any)

// Queries:

	// Precondition: list is not empty
	Get() any

	Size() int

	// Cursor set at the start of the list
	IsHead() bool

	// Cursor set at the end of the list
	IsTail() bool

	// Cursor set on the node of the list
	IsValue() bool

// Additional queries:
	
	GetRightStatus() bool
}
~~~

## 2.2. Почему операция tail не сводима к другим операциям (если исходить из эффективной реализации)?

- Для достижения такого же результата необходимо будет использовать операцию `Right()`, которая сдвигает курсор всего на 1 узел, что дает нам довольно неэффективный алгоритм с линейной сложностью O(n)

## 2.3. Операция поиска всех узлов с заданным значением, выдающая список таких узлов, уже не нужна. Почему? 

- Комбинируя операции `Head()`, `Find(...)` и `Get()` можно получить операцию поиска всех узлов с заданным значением