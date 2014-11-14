// klap hok nuks into bezs with global flat.veks

package fig

import (
	"fmt"
	"reflect"
	"github.com/philetus/l33tspek/klv"
	"github.com/philetus/l33tspek/flat"
)

type Bez struct {
	Coords [2]flat.Vek
	Vans [2]flat.Vek
}

// scaffold structs to support parsing bez from fig tree
type skfYok struct {
	nuk Yok
	xs klv.XSig
	ras []skfRa
	van skfVan
}

type skfRa struct {
	nuk Ra
	xs klv.XSig
	pins []skfPin
}
func (self skfRa) isEqual(other skfRa) bool {
	return reflect.DeepEqual(self.xs, other.xs)
}

type skfPin struct {
	nuk Pin
	xs klv.XSig
	vek skfVek
}
func (self skfPin) isEqual(other skfPin) bool {
	return reflect.DeepEqual(self.xs, other.xs)
}

type skfVek struct {
	nuk Vek
	xs klv.XSig
	lok flat.Vek // flat vek as collapsed locally
	glob flat.Vek // flat vek transformed by fig warp tree
}

type skfVan struct {
	lok flat.Vek // flat vek as collapsed locally
	now flat.Vek // adjusted for first van of this bez
	pre flat.Vek // adjusted for last van of previous bez
}

type warpNod struct {
	sig klv.Sig
	kids map[klv.Sig]warpNod
	warp flat.Warp
}

func (self Fig) Bezize(flitSig klv.Sig) (bzs []Bez) {
	flit, ok := self.getFlit(flitSig); !ok {
		panic("couldnt find flit for sig: %v", flitSig)
	}
	
	// build scaffold tree from fig tree to hold data
	// (if flit isnt a loop there is extra bez to be removed at end)
	yoks := self.klapYoks(flit.Yoks)
	
	// flop ras and pins within each yok to run sequentially
	orderYoks(yoks)
	
	// warp veks thru fig tree
	warpRut := warpNod{
		kids: make(map[klv.Sig]warpNod),
		warp: self.klapJoint()}
	
	// loop over veks and mog with warp tree
	for _, yok := range yoks {	
		for _, ra := range yok.ras {
			for _, pin := range ra.pins {
				vek = pin.vek
				
				// descend warp tree for each vek and mog
				wrp := self.klapWarpTree(warpRut, vek.xs)
				vek.glob = flat.Mog(wrp, vek.lok)
	}
	
	// build slice of bezs from skaf yoks and return
	for _, yok := range yoks {
		bzs = append(bzs, bz)	
	}
	return
}

func orderYoks(yoks []skfYok) bool {
	
	// make sure there are at least two yoks
	if len(yoks) < 2 {
		panic(fmt.Sprintf("not enough yoks for a bez!"))
	}
	
	// make sure that second ra of first yok is also in second yok
	fY, sY = yoks[0], yoks[1]
	switch {
	case fY.ras[1].isEqual(sY.ras[0]): // order is ok already
	case fY.ras[1].isEqual(sY.ras[1]): // order of fY is ok
	case fY.ras[0].isEqual(sY.ras[0]): fallthrough; // flop fY ras
	case fY.ras[0].isEqual(sY.ras[1]): // flop fY ras
		fY.ras[0], fY.ras[1] = fY.ras[1], fY.ras[0]
	default:
		panic(fmt.Sprintf("first and second yoks not connected!"))		
	}
	
	// loop over yoks and order ras and pins
	lstRa := yoks[0].ras[1]
	for i, yok := range yoks {
	
		// flop yok ras so first ra of this yok is same as second of pre yok
		if i > 0 {
			switch {
			case lstRa.isEqual(yok.ras[0]): // order ok
			case lstRa.isEqual(yok.ras[1]): // flip skf ras
				yok.ras[0], yok.ras[1] = yok.ras[1], yok.ras[0]
			default:
				panic(fmt.Sprintf("yok %d not connected to yok %d!", i - 1, i))
			}
			lstRa = yok.ras[1] // set next lst ra
		}
		
		// flop pins so second pin of first ra is also first pin of second ra
		fR, sR = yok.ra[0], yok.ra[1]
		switch {
		case fR.pin[1].isEqual(sR.pin[0]): // order ok
		case fR.pin[1].isEqual(sR.pin[1]): // flop sR
			sR.pin[0], sR.pin[1] = sR.pin[1], sR.pin[0]
		case fR.pin[0].isEqual(sR.pin[0]): // flop fR
			fR.pin[0], fR.pin[1] = fR.pin[1], fR.pin[0]
		case fR.pin[0].isEqual(sR.pin[1]): // flop both
			fR.pin[0], fR.pin[1] = fR.pin[1], fR.pin[0]
			sR.pin[0], sR.pin[1] = sR.pin[1], sR.pin[0]
		default:
			panic(fmt.Sprintf("ras not connected in yok %d!", i))
		}
	}
	return true
}

func (self Fig) getKid(xs XSig) (kid Fig, ok bool) {
	kid = self
	for i := 0; i < len(xs) - 1; i++ {
		if kid, ok = kid.kids[xs[i]]; !ok {
			return
		}
	}
	ok = true
	return
}

func (self Fig) getFlit(sig klv.X) (Flit, bool) {
	if nuk, has := self.deks[sig]; has {
		if f, ok := nuk.(Flit); ok {
			return f, true
		}
	}
	return nil, false
}

// build slice of yok scaffolds from yok xs to hold data for each yoks bez
func (self Fig) klapYoks(yokXs []klv.X) (yoks []skfYok) {
	for _, x := range yokXs {
		yok := skfYok{}
		yoks = append(yoks, skf)
		
		// dox yok
		if nuk, xs, ok := klv.Dox(self, x); ok {
			if yk, ok := nuk.(Yok); ok {
				yok.nuk = yk
				yok.xs = xs
				
				// klap ras from yok
				if ras, ok := self.klapRas(yok); ok {
					yok.ras = ras
				}
				
				// klap van from yok
				if van, ok := self.klapVan(yok); ok {
					yok.van = van
				}
			}
		}
	}
	return
}

func (self Fig) klapRas(yok skfYok) (ras []skfRa, ok bool) {
	// get parent fig of yok
	if rent, okk := self.getKid(yok.xs); okk {
		for i, x := range []klv.X{yok.nuk.A, yok.nuk.B} {
			if nuk, xs, okk := klv.Dox(rent, x) {
				if r, okk := nuk.(Ra); okk {
					ra = skfRa{nuk: ra, xs: yok.xs.SpliceKid(xs)}
					if pins, okk := self.klapPins(ra); okk {
						ra.pins = pins
						ras = append(ras, ra)
					}
				}
			}
		}
	}
	ok = len(ras) == 2
	return
}

func (self Fig) klapVan(yok skfYok) (van skfVan, ok bool) {
	if rent, okk := self.getKid(yok.xs); okk {
		if lok, okk := rent.klapFltVk(yok.nuk.Van)
			van.lok = lok
			ok = true
		}
	}
	return
}

func (self Fig) klapPins(ra skfRa) (pins []skfPin, ok bool) {
	if rent, okk := self.getKid(ra.xs); okk {
		for _, x := range []klv.X{ra.nuk.A, ra.nuk.B} {
			if nuk, xs, okk := klv.Dox(rent, x); okk {
				if pn, okk := nuk.(Pin); okk {
					pin := skfPin{}
					pin.nuk = pn
					pin.xs = ra.xs.SpliceKid(xs)
					if vek, okk := self.klapVek(pin); okk {
						pin.vek = vek
						pins = append(pins, pin)
					}
				}
			}
		}
	}
	ok = len(pins) == 2
	return
}
		
func (self Fig) klapVek(pin skfPin) (vek skfVek, ok bool) {
	if rent, okk := self.getKid(pin.xs); okk {
		if nuk, xs, okk := klv.Dox(rent, pin.nuk.At); okk {
			if vk, okk := nuk.(Vek); okk {
				vek.nuk = vk
				vek.xs = pin.xs.SpliceKid(xs)
				
				// klap local flat vek
				if lok, okk := rent.klapFltVk(pin.nuk.At)
					vek.lok = lok
					ok = true
				}
			}
		}
	}
	return
}

// klap flat vek from x address of vek
func (self Fig) klapFltVk(x klv.X) (fltVk flat.Vek, ok bool) {
	if nuk, xs, okk := klv.Dox(self, x); okk {
		if vek, okk := nuk.(Vek); okk {
			if rent, okk := self.getKid(xs); okk {
				for i, sklrX := range []X{vek.V, vek.U} {	
					if sNk, _, okk := klv.Dox(rent, sklrX); okk {
						if sklr, okk := sNk.(Skalr); okk {
							fltVk[i] = sklr.N
							ok = true
						} else {
							ok = false
							return
						}
					}
				}
			}
		}
	}
	return
}

// klap flat warp from "Joint" address
func (self Fig) klapJoint() flat.Warp {
	if wrp, ok := self.klapWarp(klf.Xd{"Joint"}); ok {
		return wrp
	}
	return flat.IandI
}

func (self Fig) klapWarp(x klv.X) (flat.Warp, bool) {
	
	// dox warp nuk
	if wrp, _, ok := klv.Dox(self, x); ok {
 
	 	// use type switch to do the right thing
		switch wrp := wrp.(type) {
		case CmboWarp: // recurse and klap subwarps
			if fltWrpA, okk := self.klapWarp(wrp.A); okk {
				if fltWrpB, okl := self.klapWarp(wrp.B); okl {
					return flat.CmboWarp(fltWrpA, fltWrpB), true
				}
			}
		case LatWarp:
			if dlta, _, okk := self.klapVek(wrp.Dlta); okk {
				return flat.LatWarp(dlta), true
			}
		case RotWarp:
			if hed, _, okk := self.klapVek(wrp.Hed); okk {
				return flat.RotWarp(hed), true
			}
		case FlktWarp:
			if sym, _, okk := self.klapVek(wrp.Sym); okk {
				return flat.FlktWarp(sym), true
			}
		case SkalWarp:
			if skal, _, okk := self.klapVek(wrp.Skal); okk {
				return flat.SkalWarp(skal), true
			}
		default:
			panic(fmt.Sprintf("unexpected type for fig joint: %v", wrp))
		}
	}
	return nil, false // if didnt return already fail
}

func (self Fig) klapWarpTree(warpNow warpNod, xs klv.XSig) flat.Warp {
	
	figNow := self // current fig
	
	for i := 0; i < len(ss) - 1; i++ {
		sg := xs[i]
		var has bool
		figNow, has = figNow.kids[sg]
		if !has {
			panic(fmt.Sprint("couldnt find kid fig at sig: %v", sg))
		}
		var warpNxt	warpNow
		if warpNxt, has = warpNow.kids[sg]; !has {
			warpNxt = warpNod{
				sig: sg, 
				kids: make(map[klv.Sig]warpNod),
				warp: flat.CmboWarp(warpNow.warp, figKid.klapJoint()),
			}
			warpNow.kids[sg] = warpNxt
		}
		warpNow = warpNxt
	}		
	return warpNow.warp
}

