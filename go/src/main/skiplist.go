package main

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	DEFAULT_MAX_LEVEL   = 12
	DEFAULT_PROBABILITY = 0.5
)

type elementNode struct {
	next []*Element
}

type Element struct {
	elementNode
	key   float64
	value interface{}
}

func (e *Element) Key() float64 {
	return e.key
}

func (e *Element) Value() interface{} {
	return e.value
}

func (element *Element) Next() *Element {
	return element.next[0]
}

type SkipList struct {
	elementNode
	maxLevel       int
	Length         int
	randSource     rand.Source
	probability    float64
	probTable      []float64
	mutex          sync.RWMutex
	prevNodesCache []*elementNode
}

func (list *SkipList) Front() *Element {
	return list.next[0]
}

func (list *SkipList) Set(key float64, value interface{}) *Element {
	list.mutex.Lock()
	defer list.mutex.Unlock()

	var element *Element
	prevs := list.getPrevElementNodes(key)

	if element = prevs[0].next[0]; element != nil && element.key <= key {
		element.value = value
		return element
	}

	element = &Element{
		elementNode: elementNode{
			next: make([]*Element, list.randLevel()),
		},
		key:   key,
		value: value,
	}

	for i := range element.next {
		element.next[i] = prevs[i].next[i]
		prevs[i].next[i] = element
	}

	list.Length++
	return element
}

func (list *SkipList) Get(key float64) *Element {

	prev := &list.elementNode
	var next *Element

	for i := list.maxLevel - 1; i >= 0; i-- {
		next = prev.next[i]

		for next != nil && key > next.key {
			prev = &next.elementNode
			next = next.next[i]
		}
	}

	if next != nil && next.key <= key {
		return next
	}

	return nil
}

func (list *SkipList) Remove(key float64) *Element {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	prevs := list.getPrevElementNodes(key)

	if element := prevs[0].next[0]; element != nil && element.key <= key {
		for k, v := range element.next {
			prevs[k].next[k] = v
		}

		list.Length--
		return element
	}

	return nil
}

func (list *SkipList) getPrevElementNodes(key float64) []*elementNode {
	prev := &list.elementNode
	var next *Element

	prevs := list.prevNodesCache

	for i := list.maxLevel - 1; i >= 0; i-- {
		next = prev.next[i]

		for next != nil && key > next.key {
			prev = &next.elementNode
			next = next.next[i]
		}

		prevs[i] = prev
	}

	return prevs
}

func (list *SkipList) SetProbability(newProbability float64) {
	list.probability = newProbability
	list.probTable = probabilityTable(list.probability, list.maxLevel)
}

func (list *SkipList) randLevel() (level int) {
	r := float64(list.randSource.Int63()) / (1 << 63)

	level = 1
	for level < list.maxLevel && r < list.probTable[level] {
		level++
	}
	return
}

func probabilityTable(probability float64, MaxLevel int) (table []float64) {
	for i := 1; i <= MaxLevel; i++ {
		prob := math.Pow(probability, float64(i-1))
		table = append(table, prob)
	}
	return table
}

func NewWithMaxLevel(maxLevel int) *SkipList {
	if maxLevel < 1 || maxLevel > 64 {
		panic("maxLevel for a SkipList must be a positive integer <= 64")
	}

	return &SkipList{
		elementNode:    elementNode{next: make([]*Element, maxLevel)},
		prevNodesCache: make([]*elementNode, maxLevel),
		maxLevel:       maxLevel,
		randSource:     rand.New(rand.NewSource(time.Now().UnixNano())),
		probability:    DEFAULT_PROBABILITY,
		probTable:      probabilityTable(DEFAULT_PROBABILITY, maxLevel),
	}
}

func New() *SkipList {
	return NewWithMaxLevel(DEFAULT_MAX_LEVEL)
}
