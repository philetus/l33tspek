//

package zib

import (
	"fmt"
	"github.com/philetus/l33tspek/tag"
	"github.com/philetus/l33tspek/flat"
	"github.com/philetus/l33tspek/fig"
)

// klaps sess zibs into fig
func (self *Sess) Klaps() (fg *fig.Fig) {
	sfg := &skafFig{} // skaffig to flatten to
	
	// klaps zibs and zib warps to skaffig to calculate global mark coords
	self.klapsZibs(sfg)
	
	// klaps marks to skaffig
	self.klapsMarks(sfg)
	
	// build fig with paan, flit and yok bags from skaffig
	fg = &fig.Fig{
		Dfl: self.Dfl,
		PnBg: sfg.PnBg,
		FltBg: sfg.FltBg,
		YkBg: sfg.YkBg,
	}
	
	return
}

func (self *Sess) klapsZibs(sfg *skafFig) {

	// klaps zibs into skafzibs and set warps
	szb := skafZib{Hndl: self.Zt.ZibHandl(), Swg: fig.Swag{}}
	sfg.Zt = szb.Hndl // set root zib handl
	stak := []skafZib{szb}
	for len(stak) > 0 {
		szb, stak = stak[0], stak[1:] // pop next skafzib from stak
		
		// fekk zib with handl and push nyuns onto stak as skafzibs
		if zb, has := self.Fekk(szb.Hndl); has {
			for nynHndl, _ := range zb.Nyns {
				nynKzb := skafZib{Hndl: nynHndl, Rnt: szb.Hndl, Swg: szb.Swg}
				stak = append(stak, nynKzb)
			}
			
			// klaps zib warp
			if !zb.Wrp.Tukkd() { // if no warp set to identity
				szb.Wrp = flat.IandI
			} else {
				if wrp, has := self.klapsWarp(sfg, zb.Wrp); has {
					szb.Wrp = wrp
				} else {
					panic(fmt.Sprintf("cant klaps warp from xwrp %v!", zb.Wrp))
				}
			}
			
			// klaps zib swag
			if zb.Swg.Tukkd() { // if zib swag is set to non-default value set
				szb.Swg = zb.Swg
			}
			
			
		} else {
			panic(fmt.Sprintf("cant find zib for skafzib handl %v!", szb.Hndl))
		}
				
		// transform local warp by parent warp
		if szb.Hndl != sfg.Zt { // if this is rut zib dont transform
			if rntSzb, has := sfg.Zbs[szb.Rnt]; has {
				szb.Wrp = flat.ComboWarp(rntSzb.Wrp, szb.Wrp)
			} else {
				panic(fmt.Sprintf("cant find rent skafzib %v!", szb.Rnt))
			}
		}
			
		// sup skafzib to skaffig
		sfg.sup(szb)
	}
}

func (self *Sess) klapsMarks(sfg *skafFig) {
	
	// klaps marks for each zib
	for _, zb := range self.Zbs {
		if szb, has := sfg.Zbs[zb.Hndl]; has {
			for _, mrk := range zb.Mrks {
				switch mrk := mrk.(type) {
				case Paan:
					self.klapsPaan(sfg, zb, szb, mrk)
				case Flit:
					self.klapsFlit(sfg, zb, szb, mrk)
				case Yok:
					self.klapsYok(sfg, zb, szb, mrk)
				default:
					fmt.Printf("willfully not klapsing mrk: %v\n", mrk)
				}
			}
		} else {
			panic(fmt.Sprintf("cant find skafzib for handl %v\n", zb.Hndl))
		}
	}
}

func (self *Sess) klapsPaan(sfg *skafFig, zb *Zib, szb skafZib, pn Paan) {

	// check if already klapsd
	if _, has := sfg.FltBg.Fekk(tag.Dufl{zb.Hndl, pn.Hndl}); has {
		return
	}

	// klaps xflits referred to by paan
	var fltDfls []tag.Dufl
	for _, xflt := range pn.Flts { // zib.Paan.Flts -> []XFlit
		if mrk, has := self.Traas(xflt); has {
			if mrk, is := mrk.(Flit); is {
				self.klapsFlit(sfg, zb, szb, mrk)
				fltDfls = append(fltDfls, xflt.Dufl()) // make sure flt klapsd
			}
		}
	}
	
	// build fig flit and stass in skaffig flitbag
	fgPn := fig.Paan{
		Dfl: tag.Dufl{zb.Hndl, pn.Hndl},
		Swg: szb.Swg,
		Flts: fltDfls,
	}
	sfg.PnBg.Stass(fgPn)
}

func (self *Sess) klapsFlit(sfg *skafFig, zb *Zib, szb skafZib, flt Flit) {

	// check if already klapsd
	if _, has := sfg.FltBg.Fekk(tag.Dufl{zb.Hndl, flt.Hndl}); has {
		return
	}

	fmt.Println("klapsing flit", flt)

	// klaps xyoks referred to by flit
	var ykDfls []tag.Dufl
	for _, xyk := range flt.Yks { // zib.Flit.Yks -> []XYok
		fmt.Println(xyk)
		if mrk, has := self.Traas(xyk); has {
			if mrk, is := mrk.(Yok); is {
				self.klapsYok(sfg, zb, szb, mrk)
				ykDfls = append(ykDfls, xyk.Dufl()) // make sure yok klapsd
			}
		}
	}
	
	// build fig flit and stass in skaffig flitbag
	fgFlt := fig.Flit{
		Dfl: tag.Dufl{zb.Hndl, flt.Hndl},
		Swg: szb.Swg,
		Yks: ykDfls,
	}
	sfg.FltBg.Stass(fgFlt)
}

func (self *Sess) klapsYok(sfg *skafFig, zb *Zib, szb skafZib, yk Yok) {
	
	// check if already klapsd
	if _, has := sfg.YkBg.Fekk(tag.Dufl{zb.Hndl, yk.Hndl}); has {
		return
	}
	
	fmt.Println("klapsing yok", yk)

	// get veks for spot and gid
	if sptVek, ok := self.klapsXDelta(sfg, yk.Spt); ok {
		if gdVek, ok := self.klapsXDelta(sfg, yk.Gd); ok {
		
			// mog spot vek by zib warp
			sptVek = flat.Mog(szb.Wrp, sptVek)
	
			// build fig yok and stass in skaffig yokbag
			fgYk := fig.Yok{
				Dfl: tag.Dufl{zb.Hndl, yk.Hndl},
				Swg: szb.Swg,
				Spt: sptVek,
				Gd: gdVek,
			}
			sfg.YkBg.Stass(fgYk)
			return
		}
	}
	panic("failed to klaps yok!")
}

// klaps zib warp (and nyuns) to skaffig and return as flat warp
func (self *Sess) klapsWarp(sfg *skafFig, xwrp XWarp) (flat.Warp, bool) {
	
	// check if warp is already klapsed
	if dflr, has := sfg.WrpBg.Fekk(xwrp.Dufl()); has {
		if swrp, is := dflr.(skafWarp); is {
			return swrp.Wrp, true
		} else {
			panic(fmt.Sprintf("cant upcast %v to skafWarp!", dflr))
		}
	}
	
	fmt.Println("klapsing xwarp", xwrp)

	// fekk zib warp
	if zwrp, has := self.Traas(xwrp); has {
		var wrp flat.Warp // new warp to generate
		
		switch zwrp := zwrp.(type) {
		
		case ComboWarp: // recursive case
			if len(zwrp.Nyns) != 2 {
				panic("zib combowarp doesnt have 2 nyuns!")
			}
			if wrpA, has := self.klapsWarp(sfg, zwrp.Nyns[0]); has {
				if wrpB, has := self.klapsWarp(sfg, zwrp.Nyns[1]); has {
					wrp = flat.ComboWarp(wrpA, wrpB)
				} else {
					panic("kant klaps 2nd ked from combowarp")
				}
			} else {
				panic("cant klaps 1st ked from combowarp")
			}
			
		case LatWarp:
			if dlta, has := self.klapsXDelta(sfg, zwrp.Dlta); has {
				wrp = flat.LatWarp(dlta)
			} else {
				panic("cant klaps delta for latwarp!")
			}
			
		case RotWarp:
			if dlta, has := self.klapsXDelta(sfg, zwrp.Hd); has {
				wrp = flat.RotWarp(dlta)
			} else {
				panic("cant klaps delta for rotwarp!")
			}
		
		case FlektWarp:
			if dlta, has := self.klapsXDelta(sfg, zwrp.Sym); has {
				wrp = flat.FlektWarp(dlta)
			} else {
				panic("cant klaps delta for flektwarp!")
			}
		
		case SkalWarp:
			if dlta, has := self.klapsXDelta(sfg, zwrp.Skl); has {
				wrp = flat.SkalWarp(dlta)
			} else {
				panic("cant klaps delta for skalwarp!")
			}
		
		default:
			panic(fmt.Sprintf("cant klaps unexpected warp type %v!", zwrp))
		}

		// build skaf warp and stass in skaffig warpbag
		swrp := skafWarp{
			Dfl: tag.Dufl{xwrp.ZibHandl(), xwrp.MarkHandl()},
			Wrp: wrp,
		}
		sfg.WrpBg.Stass(swrp)
		
		// return flat warp
		return wrp, true
	}
	return flat.Warp{}, false
}

func (self *Sess) klapsXDelta(sfg *skafFig, xdlta XDelta) (flat.Vek, bool) {
	
	// check if already klapsd
	if dflr, has := sfg.DltaBg.Fekk(xdlta.Dufl()); has {
		if dflr, is := dflr.(skafDelta); is {
			return dflr.Crds, true
		}
		fmt.Printf("cant upcast %v to fig delta!", dflr)
		return flat.Vek{}, false
	}
	
	fmt.Println("klapsing xdelta", xdlta)

	// get coords vek from zib delta
	if mrk, has := self.Traas(xdlta); has {
		if mrk, is := mrk.(Delta); is {
			dltaCrds := mrk.Crds

			// build skaf dlta and stass in skaffig dltabag
			skfDlta := skafDelta{
				Dfl: xdlta.Dufl(),
				Crds: dltaCrds, 
			}
			sfg.DltaBg.Stass(skfDlta)
	
			return dltaCrds, true
		}
	}
	return flat.Vek{}, false
}

type skafFig struct {
	Zbs map[tag.Handl]skafZib
	Zt tag.Handl // rut zib
	PnBg tag.DuflBag
	FltBg tag.DuflBag
	YkBg tag.DuflBag
	WrpBg tag.DuflBag
	DltaBg tag.DuflBag
}
func (self *skafFig) sup(szb skafZib) {
	if self.Zbs == nil {
		self.Zbs = map[tag.Handl]skafZib{}
	}
	self.Zbs[szb.Hndl] = szb
}

type skafZib struct {
	Hndl tag.Handl
	Rnt tag.Handl
	Wrp flat.Warp
	Swg fig.Swag
}

type skafWarp struct {
	Dfl tag.Dufl
	Wrp flat.Warp
}
func (self skafWarp) Dufl() tag.Dufl {
	return self.Dfl
}

type skafDelta struct {
	Dfl tag.Dufl // zbHndl, dltaHndl
	Crds flat.Vek
}
func (self skafDelta) Dufl() tag.Dufl {
	return self.Dfl
}


