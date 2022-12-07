package main

import "errors"

type DirectoryStack struct {
	entries []*directory
}

func (this *DirectoryStack) Push(value *directory) {
	this.entries = append(this.entries, value)
}

func (this *DirectoryStack) Peek() (*directory, error) {
	size := len(this.entries)
	if size > 0 {
		return this.entries[size-1], nil
	}
	return nil, errors.New("empty stack")
}

func (this *DirectoryStack) Pop() (*directory, error) {
	size := len(this.entries)
	if size == 0 {
		return nil, errors.New("empty stack")
	}
	index := size - 1
	value := this.entries[index]
	this.entries = this.entries[:index]
	return value, nil
}

func (this *DirectoryStack) IsEmpty() bool {
	return len(this.entries) == 0
}
