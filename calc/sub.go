package main

import "fmt"

type Calc struct{ atai1, atai2 int }

func (p *Calc) Add() int {
	return p.atai1 + p.atai2
}

func (p *Calc) Sub() int {
	return p.atai1 - p.atai2
}

func (p *Calc) Multi() int {
	return p.atai1 * p.atai2
}

func (p *Calc) Div() int {
	if p.atai2 == 0 {
		return 0
	}
	return p.atai1 / p.atai2
}

func main() {
	p := &Calc{5, 2}       // 5 ? 2
	fmt.Println(p.Add())   // 7
	fmt.Println(p.Sub())   // 3
	fmt.Println(p.Multi()) // 10
	fmt.Println(p.Div())   // 2

	q := &Calc{5, 2}     // 5 ? 2
	fmt.Println(q.Add()) // 7
}
