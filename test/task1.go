package main

import "fmt"

type PepeSchnele struct { speed int; charisma int; wisdom int }

func NewPepeSchnele(s, c, w int) *PepeSchnele {
	return &PepeSchnele{s, c, w}
}

func (p *PepeSchnele) GetRating() int {
	return p.speed*2 + p.charisma*3 + p.wisdom
}

func main() {
	p1 := NewPepeSchnele(5, 7, 4)
	p2 := NewPepeSchnele(6, 4, 8)

	fmt.Println("пепе [скорость:", p1.speed, "харизма:", p1.charisma, "мудрость:", p1.wisdom, "] рейтинг:", p1.GetRating())
	fmt.Println("пепе [скорость:", p2.speed, "харизма:", p2.charisma, "мудрость:", p2.wisdom, "] рейтинг:", p2.GetRating())

	if p1.GetRating() > p2.GetRating() {
		fmt.Println("у первого рейтинг выше")
	} else if p1.GetRating() < p2.GetRating() {
		fmt.Println("у второго рейтинг выше")
	} else {
		fmt.Println("рейтинги равны")
	}
}
