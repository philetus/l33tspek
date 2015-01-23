package main

import (
	"fmt"
	"github.com/philetus/l33tspek/flat"
)

func main() {
	a := flat.Vek{2.0, 3.0}
	fmt.Printf("a: %v\n", a)
	
	i := flat.Mog(flat.IandI, a)
	fmt.Printf("i: %v\n", i)

	lw := flat.LatWarp(a)
	fmt.Printf("lw: %v\n", lw)
	
	b := flat.Mog(lw, a)
	fmt.Printf("b: %v\n", b)
	
	fw := flat.FlektWarp(flat.Vek{1.0, 1.0})
	fmt.Printf("fw: %v\n", fw)

	c := flat.Mog(fw, a)
	fmt.Printf("c: %v\n", c)
	
	hdng := flat.Vek{0.0, 1.0}
	fmt.Printf("hdng: %v\n", hdng)

	angl := flat.Hed(hdng)
	fmt.Printf("angl: %v\n", angl)
		
	rw := flat.RotWarp(hdng)
	fmt.Printf("rw: %v\n", rw)

	d := flat.Mog(rw, a)
	fmt.Printf("d: %v\n", d)
	
	lrw := flat.ComboWarp(lw, rw)
	fmt.Printf("lrw: %v\n", lrw)
	
	rlw := flat.ComboWarp(rw, lw)
	fmt.Printf("rlw: %v\n", rlw)
	
	e := flat.Mog(lrw, a)
	fmt.Printf("e: %v\n", e)

	f := flat.Mog(rlw, a)
	fmt.Printf("f: %v\n", f)
	
}
