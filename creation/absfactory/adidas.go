package main

type adidas struct{}

func (a *adidas) makeShot() iShot {
	return &adidasShot{
		shot: shot{
			logo: "adidas",
			size: 111,
		},
	}
}

func (a *adidas) makeShirt() iShirt {
	return &adidasShirt{
		shirt: shirt{
			logo: "adidas",
			size: 111,
		},
	}
}
