package main

type subject interface {
	register(observer)
	deregister(observer)
	notifyAll()
}

type goods struct {
	observerList []observer
	name         string
	inStock      bool
}

func newgoods(name string) *goods {
	return &goods{
		name: name,
	}
}

func (i *goods) updateAvailability() {
	println("goods availability changed to: ", i.inStock)
	i.notifyAll()
}

func (i *goods) register(o observer) {
	i.observerList = append(i.observerList, o)
}

func (i *goods) deregister(o observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *goods) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

func removeFromSlice(observerList []observer, observerToRemove observer) []observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

type observer interface {
	update(string)
	getID() string
}

type customer struct {
	id string
}

func (c *customer) update(goodsName string) {
	println("Sending email to customer ", c.id, " for goods ", goodsName)
}

func (c *customer) getID() string {
	return c.id
}

func main() {

	shirtgoods := newgoods("Nike Shirt")

	observerFirst := &customer{id: "abc@gmail.com"}
	observerSecond := &customer{id: "xyz@gmail.com"}

	shirtgoods.register(observerFirst)
	shirtgoods.register(observerSecond)

	shirtgoods.updateAvailability()
}
