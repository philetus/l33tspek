package main

import (
	"fmt"
	"github.com/philetus/l33tspek/x"
)

type Fig struct {
	x.Spam // hok mixin
}
func NewFig(sg x.Sig, ark *x.Hok) *Fig {
	return &Fig{
		x.Spam{
			S: sg,
			A: ark,
			Kids: make(map[x.Sig]*x.Hok),
			Deks: make(map[x.Sig]*x.Bon),
		},
	}
}
type Skalr struct {
	S x.Sig
	Val float64
}
func (self Skalr) Sig() x.Sig {
	return self.S
}

type Vek struct {
	S x.Sig
	V x.X
	U x.X
}
func (self Vek) Sig() x.Sig {
	return self.S
}

type Pin struct {
	S x.Sig
	Vek x.X
}
func (self Pin) Sig() x.Sig {
	return self.S
}

func main() {
	f := NewFig("root", nil)
	fmt.Println("created f!: ", f)
		
	f.Et(Vek{S: "vTest", 
			 V: f.Et(Skalr{S:"sV", Val:42.0}),
			 U: f.Et(Skalr{S:"sU", Val:17.0}),
			})
	
	fmt.Println("f: ", f)
}
