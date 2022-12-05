package commons

import "errors"

type Stack struct {
	entries []string
}

func (this *Stack) Push(value string) {
	this.entries = append(this.entries, value)
}

func (this *Stack) Peek() (string, error) {
	size := len(this.entries)
	if size > 0 {
		return this.entries[size-1], nil
	}
	return "", errors.New("empty stack")
}

func (this *Stack) Pop() (string, error) {
	size := len(this.entries)
	if size == 0 {
		return "", errors.New("empty stack")
	}
	index := size - 1
	value := this.entries[index]
	this.entries = this.entries[:index]
	return value, nil
}

func (this *Stack) IsEmpty() bool {
	return len(this.entries) == 0
}
