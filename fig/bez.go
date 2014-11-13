// klap hok nuks into bezs with global flat.veks

package fig

import (
	//"fmt"
	//"reflect"
	"github.com/philetus/l33tspek/klv"
	"github.com/philetus/l33tspek/flat
)

type Bez struct {
	Coords [2]flat.Vek
	Vans [2]flat.Vek
}

type bezSkaf struct {
	yok Yok
	yokSs klv.Xsig

	ras []Ra
	raSss []klv.Xsig

	pins []Pin
	pinSss []klv.Xsig

	vekSss []klv.Xsig

	vanSS klv.Xsig

	bz Bez
}

type warpNod struct {
	sig klv.Sig
	kids map[klv.Sig]warpNod
	warp flat.Warp
}

func (self Fig) Bezize(flit klv.Sig) bzs []Bez {

	// get flit by sig or fail
	var f Flit	
	if nuk, has := self.deks[flit]; has {
		if tst, ok := nuk.(Flit); ok {
			f = tst
		} else {
			return nil
		}
	} else {
		return nil
	}
	
	// build slice of yoks and equal sized slice of blank bezs to return
	// (if flit isnt a loop there is extra bez to be removed at end)
	skafs := self.klapSkafs(f.Yoks)
	
	// loop over skafs and order ras
	lstRa := skafs[0].ras[0]
	lstRaSss := skafs[0].raSss[0]
	nxtRa := skafs[0].ras[1]
	nxtRaSss := skafs[0].raSss[1]
	
	switch {
	case nxtRaSss == skafs[1].raSss[0]: // order is ok already
	case nxtRaSss == skafs[1].raSss[1]:
	case lstRaSss == skafs[1].raSss[0]: fallthrough;// flop nxt & lst
	case lstRaSss == skafs[1].raSss[1]: 
		lstRas, lstRaSss = nxtRas, nxtRaSss
		nxtRas, nxtRaSss = skafs[1].ras[0], skafs[1].raSss[0]
	default:
		panic(fmt.Sprintf("first and second yoks not connected!"))		
	}
	
	for i, skf := range skafs {
		switch {
		case nxtRaSs == skf.raSss[0]: // order ok, set nxt ra
			nxtRa, nxtRaSs = skf.ras[1], skf.raSss[1]
		case nxtRaSs == skf.raSss[1]: // flip skf ras then set nxt ra
			skf.ras[1], skf.raSss[1] = skf.ras[0], skf.raSss[0]
			skf.ras[0], skf.raSss[0] = nxtRa, nxtRaSs
			nxtRa, nxtRaSs = skf.ras[1], skf.raSss[1]
		default:
			panic(fmt.Sprintf("yok %d not connected to yok %d!", i - 1, i))
		}
		
		// dox pins then pick the right two
		if pins, pinSss, ok := self.klapPins(skf.ras); ok {
			skf.pins, skf.pinSss = pins, pinSss
		} else {
			panic(fmt.Sprintf("klap pins fail at yok %d!", i))
		}
		
		// klap flat veks from pins
		for j := 0; j < 2; j++ {
			if skf.bz.Coords[j], skf.vekSss[j], ok := 
				self.KlapVek(skf.pins[j].At); !ok {
				panic(fmt.Sprintf("klap pins fail at yok %d, pin %d!", i, j))
			}
		}
		
		// klap flat veks from vans
		if skf.vanVeks, ok := self.klapVan(skf.yok.Van); !ok {
			 panic(fmt.Sprintf("klap vans fail at yok %d!", i))
		}
	}
	
	// warp veks thru fig tree
	warpRt := warpNod{
		kids: make(map[klv.Sig]warpNod),
		warp: self.klapJoint()}
	
	// loop over coord veks and mog with warp tree
	for i, skf := range skafs {
	
		// descend warp tree for each coord vek and mog
		for j := 0; j < 2; j++ {
			wrp := self.klapWarpTree(warpRt, skf.VekSss[j])
			skf.bz.Coords[j] = flat.Mog(wrp, skf.bz.Coords[j])
		}
		
		// add coords[0] to vanVeks[1]
		skf.bz.Vans[0] = flat.Add(skf.vanVeks[1], skf.bz.Coords[0])
		if i != 0 {
			preSkf = skafs[i - 1]
			preSkf.bz.Vans[1] = flat.Add(skf.vanVeks[0], preSkf.bz.Coords[0])
		}
	}
	lstSkf = skafs[len(skaf) - 1]
	lstSkf.bz.Vans[1] = flat.Add(skafs[0].vanVeks[0], lstSkf.bz.Vans[1])
	
	// build slice of bezs from skafs and return
	for _, skf := range skafs {
		bzs = append(bzs, skf.bz)	
	}
	return
}

func (self Fig) klapWarpTree(warpNow warpNod, ss klv.Xsig) flat.Warp {
	
	figNow := self // current fig
	
	for i := 0; i < len(ss) - 1; i++ {
		sg := ss[i]
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

func (self Fig) klapSkafs(yokXs []X) (skafs []bezSkaf) {
	for _, x := range yokXs {
		skf := bezSkaf{}
		skafs = (append(skafs, skf))
		if nuk, ss, ok := klv.Dox(self, x); ok {
			if yok, ok := nuk.(Yok); ok {
				skf.yok, skf.yokSs = yok, ss
				
				// dox ras
				for i, raX := range []X{skf.yok.A, skf.Yok.B} {
					if rNk, rSs, ok := klv.Dox(self, raX); ok {
						if ra, ok := rNk.(Ra); ok {
							skf.ras[i], skf.raSss[i] = ra, rSs
						}
					}	
				}
			}
		}
	}
	return
}

func (self Fig) klapVan(vekX X) (vanVeks []flat.Vek) {
	vanVek, _, ok := klapVek(vekX); 
	if !ok {
		panic(fmt.Sprint("couldnt find van at x: %v", vekX))
	}
	
	// rotate vans pos/neg 1/4 turn
	vanVeks = append(
		vanVeks, 
		flat.Mog(flat.RotWarp(flat.Vek{0.0, 1.0}), vanVek))
	vanVeks = append(
		vanVeks, 
		flat.Mog(flat.RotWarp(flat.Vek{0.0, -1.0}), vanVek))
	return
}

func (self Fig) klapPins(ras [2]Ra) (pins []Pin, pinSss []klv.Xsig, ok bool) {
	pns = [2][2]Pin
	pnSss = [2][2]klv.Xsig
	for i := 0; i < 2; i++ {
		for j, pnX := range []X{ras[i].A, ras[i].B} {
			if nuk, ss, okk := klv.Dox(self, pnX); okk {
				if pn, okk := nuk.(Pin); okk {
					pns[i][j], pnSss[i][j] = pn, ss
				}
			}
		}
	}
	switch {
	case pnSss[0][0] == pnSss[1][0]: fallthrough; // order ok
	case pnSss[0][1] == pnSss[1][0]:
		for j := 0; j < 2; j++ {
			pins = append(pins, pns[1][j])
			pinSss = append(pinSss, pinSss[1][j])
		}
	case pnSss[0][0] == pnSss[1][1]: fallthrough; // flop (pins[1])
	case pnSss[0][1] == pnSss[1][0]:
		for j := 0; j < 2; j++ {
			pins = append(pins, pns[1][1 - j])
			pinSss = append(pinSss, pnSss[1][1 - j])
		}
	default:
		ok = false
		return
	}
	ok = true
	return
}

func (self Fig) klapVek(vekX X) (fltVk flat.Vek, vkSs klv.Xsig, ok bool) {
	if nuk, ss, okk := klv.Dox(self, vekX); okk {
		if vek, okk := nuk.(Vek); okk {
			fltVk = flat.Vek{}
			vkSs = ss
			for k, sklrX := range []X{vek.V, vek.U} {	
				if nuk, ss, ok := klv.Dox(self, sklrX); ok {
					if sklr, ok := nuk.(Skalr); ok {
						fltVk[k] = sklr.N
					}
				}
			}
			ok = true
			return
		}
	}
	return nil, nil, false
}

func (self Fig) klapJoint() Warp {

}

