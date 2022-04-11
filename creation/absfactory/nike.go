package main

type nike struct{}

func (n *nike) makeShot() iShot {
	return &nikeShot{
		shot: shot{
			logo: "nike",
			size: 12,
		},
	}
}

func (n *nike) makeShirt() iShirt {
	return &nikeShirt{
		shirt: shirt{
			logo: "nike",
			size: 12,
		},
	}
}
