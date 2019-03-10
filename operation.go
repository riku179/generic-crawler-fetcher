package fetcher

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
		operations = append(operations, newOperation(findChildrenMap, root))
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
}

func newOperation(childrenMap findChildrenMap, selector Selector) Operation {
	childOperations := make([]Operation, 0)

	children, ok := childrenMap[selector.GetId()]
	if !ok {
		return Operation{selector, childOperations}
	}
	for _, child := range children {
		childOperations = append(childOperations, newOperation(childrenMap, child))
	}
	return Operation{selector, childOperations}
}

func (o *Operation) Selector() Selector {
	return o.selector
}

func (o *Operation) Children() []Operation {
	return o.children
}

func (o *Operation) Type() SelectorType {
	return o.selector.GetType()
}

func (o *Operation) Next(f func(Operation)) Operation {
	return Operation{}
}
