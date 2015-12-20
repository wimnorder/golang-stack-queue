package main

import (
  "testing"
  "fmt"
)

func TestStack(t *testing.T) {

  t.Log("Stack.Push(1,2,3) -> Stack.Pop(3,2,1)")
  
  stack := NewStack()
  stack.Push(&Node{1})
  stack.Push(&Node{2})
  stack.Push(&Node{3})

  if stack_test := fmt.Sprintf("%s %s %s", stack.Pop(), stack.Pop(), stack.Pop()); stack_test != "3 2 1"  {
      t.Errorf("Stack is not equal: 3 2 1")
  }
}

func TestQueue(t *testing.T) {

  t.Log("Queue.Push(4,5,6) -> Queue.Pop(4,5,6)")

  queue := NewQueue(1)
  queue.Push(&Node{1})
  queue.Push(&Node{2})
  queue.Push(&Node{3})

  if queue_test := fmt.Sprintf("%s %s %s", queue.Pop(), queue.Pop(), queue.Pop()); queue_test != "1 2 3" {
      t.Errorf("Queue is not equal: 1 2 3")
  }
}
