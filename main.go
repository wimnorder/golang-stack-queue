package main

import "fmt"

type Node struct {
	Value int
}

func (node *Node) String() string {
	return fmt.Sprint(node.Value)
}

func NewStack() *Stack {
	return &Stack{}
}

// Resizable Stack (LIFO) - Last In First Out
type Stack struct {
	nodes []*Node
	count int
}

func (stack *Stack) Push(node *Node) {
	stack.nodes = append(stack.nodes[:stack.count], node)
	stack.count++
}

func (stack *Stack) Pop() *Node {
	if stack.count == 0 {
		return nil
	}
	stack.count--
	return stack.nodes[stack.count]
}

func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]*Node, size),
		size:  size,
	}
}

// Resizable Queue (FIFO) - First In First Out
type Queue struct {
	nodes []*Node
	size  int
	head  int
	tail  int
	count int
}

func (queue *Queue) Push(node *Node) {
	if queue.head == queue.tail && queue.count > 0 {
		nodes := make([]*Node, len(queue.nodes)+queue.size)
		copy(nodes, queue.nodes[queue.head:])
		copy(nodes[len(queue.nodes)-queue.head:], queue.nodes[:queue.head])
		queue.head = 0
		queue.tail = len(queue.nodes)
		queue.nodes = nodes
	}
	queue.nodes[queue.tail] = node
	queue.tail = (queue.tail + 1) % len(queue.nodes)
	queue.count++
}

func (queue *Queue) Pop() *Node {
	if queue.count == 0 {
		return nil
	}
	node := queue.nodes[queue.head]
	queue.head = (queue.head + 1) % len(queue.nodes)
	queue.count--
	return node
}

func main() {
	stack := NewStack()
	stack.Push(&Node{1})
	stack.Push(&Node{2})
	stack.Push(&Node{3})
	stack.Push(&Node{4})
	stack.Push(&Node{5})
	stack.Push(&Node{6})
	stack.Push(&Node{7})
	stack.Push(&Node{8})
	stack.Push(&Node{9})
	stack.Push(&Node{10})
	stack.Push(&Node{11})

	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop(),
		stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())

	queue := NewQueue(1)
	queue.Push(&Node{1})
	queue.Push(&Node{2})
	queue.Push(&Node{3})
	queue.Push(&Node{4})
	queue.Push(&Node{5})
	queue.Push(&Node{6})
	queue.Push(&Node{7})
	queue.Push(&Node{8})
	queue.Push(&Node{9})
	queue.Push(&Node{10})

	fmt.Println(queue.Pop(), queue.Pop(), queue.Pop(), queue.Pop(), queue.Pop(),
		queue.Pop(), queue.Pop(), queue.Pop(), queue.Pop(), queue.Pop())
}
