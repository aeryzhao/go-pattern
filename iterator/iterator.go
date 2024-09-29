package main

type collection interface {
	createIterator() iterator
}

type iterator interface {
	hasNext() bool
	getNext() *item
}

type item struct {
	name string
}

type collectionItem struct {
	items []*item
}

func (c *collectionItem) createIterator() iterator {
	return &itemIterator{
		index: 0,
		items: c.items,
	}
}

func (c *collectionItem) addItem(name string) {
	c.items = append(c.items, &item{name: name})
}

type itemIterator struct {
	index int
	items []*item
}

func (i *itemIterator) hasNext() bool {
	if i.index < len(i.items) {
		return true
	}
	return false
}

func (i *itemIterator) getNext() *item {
	if i.hasNext() {
		item := i.items[i.index]
		i.index++
		return item
	}
	return nil
}

func main() {
	collection := &collectionItem{}
	collection.addItem("item 1")
	collection.addItem("item 2")
	collection.addItem("item 3")
	collection.addItem("item 4")
	collection.addItem("item 5")

	iterator := collection.createIterator()
	for iterator.hasNext() {
		item := iterator.getNext()
		println(item.name)
	}
}
