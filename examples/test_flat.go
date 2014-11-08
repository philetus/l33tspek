package main

import (
	"fmt"
	"github.com/philetus/l33tspek/flat"
)

func main() {
	a := flat.Vek{2.0, 3.0}
	fmt.Printf("a: %v\n", a)
	
	i := flat.Mog(a, flat.IandI)
	fmt.Printf("i: %v\n", i)

	lw := flat.LatWarp(a)
	fmt.Printf("lw: %v\n", lw)
	
	b := flat.Mog(a, lw)
	fmt.Printf("b: %v\n", b)
	
	fw := flat.FlktWarp(flat.Vek{1.0, 1.0})
	fmt.Printf("fw: %v\n", fw)

	c := flat.Mog(a, fw)
	fmt.Printf("c: %v\n", c)
	
	hdng := flat.Vek{0.0, 1.0}
	fmt.Printf("hdng: %v\n", hdng)

	angl := flat.Hed(hdng)
	fmt.Printf("angl: %v\n", angl)
		
	rw := flat.RotWarp(hdng)
	fmt.Printf("rw: %v\n", rw)

	d := flat.Mog(a, rw)
	fmt.Printf("d: %v\n", d)
	
	lrw := flat.CmboWarp(lw, rw)
	fmt.Printf("lrw: %v\n", lrw)
	
	rlw := flat.CmboWarp(rw, lw)
	fmt.Printf("rlw: %v\n", rlw)
	
	e := flat.Mog(a, lrw)
	fmt.Printf("e: %v\n", e)

	f := flat.Mog(a, rlw)
	fmt.Printf("f: %v\n", f)
	
}
