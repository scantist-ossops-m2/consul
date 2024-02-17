// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package adaptive

import "sync"

type Node256[T any] struct {
	partialLen  uint32
	artNodeType uint8
	numChildren uint8
	partial     []byte
	children    [256]*Node[T]
	mutex       *sync.RWMutex
}

func (n *Node256[T]) getPartialLen() uint32 {
	return n.partialLen
}

func (n *Node256[T]) setPartialLen(partialLen uint32) {
	n.partialLen = partialLen
}

func (n *Node256[T]) getArtNodeType() uint8 {
	return n.artNodeType
}

func (n *Node256[T]) setArtNodeType(artNodeType uint8) {
	n.artNodeType = artNodeType
}

func (n *Node256[T]) getNumChildren() uint8 {
	return n.numChildren
}

func (n *Node256[T]) setNumChildren(numChildren uint8) {
	n.numChildren = numChildren
}

func (n *Node256[T]) getPartial() []byte {
	return n.partial
}

func (n *Node256[T]) setPartial(partial []byte) {
	n.partial = partial
}

func (n *Node256[T]) isLeaf() bool {
	return false
}

// Iterator is used to return an iterator at
// the given node to walk the tree
func (n *Node256[T]) Iterator() *Iterator[T] {
	stack := make([]Node[T], 0)
	stack = append(stack, n)
	nodeT := Node[T](n)
	return &Iterator[T]{stack: stack, root: &nodeT, mutex: n.mutex}
}

func (n *Node256[T]) PathIterator(path []byte) *PathIterator[T] {
	nodeT := Node[T](n)
	return &PathIterator[T]{parent: &nodeT, path: path, mutex: n.mutex}
}

func (n *Node256[T]) matchPrefix(_ []byte) bool {
	// No partial keys in NODE256, always match
	return true
}

func (n *Node256[T]) getChild(index int) *Node[T] {
	if index < 0 || index >= 256 {
		return nil
	}
	return n.children[index]
}

func (n *Node256[T]) getTreeMutex() *sync.RWMutex {
	return n.mutex
}
