package fetcher

import (
	"fmt"
	"strings"
)

func NewOperations(selectors []Selector) []Operation {
	selectorMap := make(map[SelectorId]Selector)
	for _, selector := range selectors {
		selectorMap[selector.GetId()] = selector
	}

	rootSelectors := findRootSelectors(selectors)
	if len(rootSelectors) == 0 {
		panic("root selector not found")
	}

	findChildrenMap := newFindChildrenMap(selectors)

	operations := make([]Operation, 0)
	for _, root := range rootSelectors {
		operations = append(operations, newOperationTree(findChildrenMap, root, nil))
	}

	return operations
}

type findChildrenMap map[SelectorId][]Selector

// -- helpers for NewOperations

func findRootSelectors(selectors []Selector) []Selector {
	rootSelectors := make([]Selector, 0)
	for _, selector := range selectors {
		for _, parentSelector := range selector.GetParentSelectors() {
			if parentSelector == "_root" {
				rootSelectors = append(rootSelectors, selector)
				continue
			}
		}
	}
	return rootSelectors
}

func newFindChildrenMap(selectors []Selector) findChildrenMap {
	findChildrenMap := make(findChildrenMap)
	for _, selector := range selectors {
		for _, parentSelector := range selector.GetParentSelectors() {
			findChildrenMap[parentSelector] = append(findChildrenMap[parentSelector], selector)
		}
	}
	return findChildrenMap
}

// --

type Operation struct {
	selector Selector
	children []Operation
	parent   *Operation
	executed bool
}

func newOperationTree(childrenMap findChildrenMap, selector Selector, parent *Operation) Operation {
	childrenOperations := make([]Operation, 0)

	childrenSelectors, ok := childrenMap[selector.GetId()]
	if !ok {
		return Operation{selector, childrenOperations, parent, false}
	}

	var me Operation
	for _, childSelector := range childrenSelectors {
		childrenOperations = append(childrenOperations, newOperationTree(childrenMap, childSelector, &me))
	}
	me = Operation{selector, childrenOperations, parent, false}
	return me
}

func (o *Operation) Selector() Selector {
	return o.selector
}

func (o *Operation) Type() SelectorType {
	return o.selector.GetType()
}

func (o *Operation) Children() []Operation {
	return o.children
}

func (o *Operation) Parent() *Operation {
	return o.parent
}

func (o *Operation) Executed() bool {
	return o.executed
}

func debug(depth int, tabWidth int, operation Operation) {
	var tabs []string
	for i := 0; i < tabWidth; i++ {
		tabs = append(tabs, "\t")
	}
	fmt.Printf("%s id:%q type:%q selector:%q delay:%q",
		strings.Join(tabs, ""),
		operation.selector.GetId(),
		operation.selector.GetType(),
		operation.selector.GetSelector(),
		operation.selector.GetDelay(),
	)
	for _, child := range operation.Children() {
		debug(depth-1, tabWidth+1, child)
	}
}

func (o *Operation) Debug(depth int) {
	debug(depth, 0, *o)
}
