// model for collapsed fig

package klap

import (
	//"fmt"
	"github.com/philetus/l33tspek/klv"
	"github.com/philetus/l33tspek/flat"
	"github.com/philetus/l33tspek/fig"
)

// sigs from root hok
type Sigs []klv.Sig
func (self Sigs) Splice(kid Sigs) Sigs {
	return append(self[:len(self) - 1], kid...)
}
func (self Sigs) Sig() klv.Sig {
	return self[len(self)-1]
}

// sigs from root hok
type Sigs []klv.Sig
func (self Sigs) Splice(kid Sigs) Sigs {
	return append(self[:len(self) - 1], kid...)
}

type Lop struct {
	Rut klv.Sig
	Figs map[klv.Sig]Fig
}

type Fig struct {
	Sgs Sigs
	Kids []klv.Sig
	Nuks map[klv.Sig]Nuk
	Joint flat.Warp
}

type Nuk interface {
	Sig klv.Sig()
}

type Delta struct {
	Sg klv.Sig
	Vek flat.Vek
}

type Warp struct {
	Sg klv.Sig
	Wrp flat.Warp
}

type Gid struct {
	Sg klv.Sig
	Org flat.Vek // origin vek
	Hed flat.Vek // heading vek
}

type Pin struct {
	Sg klv.Sig
	Crds flat.Vek
	Van flat.Vek
}

type Flit struct {
	Sg klv.Sig
	Loop bool
	Pins []Pin
}

// ******

// collapse fig tree within fig lop to flat coords in klapsd lop
func Klpsd(lop fig.Lop) (klpsd Lop) {
	
	// populate klap.lop with klap.figs
	klpsd.populateFigs(lop)
	
	// klap deltas first
	for skrt := range klpsd.iterNuks(lop) {
		if dlta, is := skrt.nuk.(fig.Delta); is { // if nuk is delta klap it
			skrt.fig.Nuks[dlta.Sg] = Delta{
				Sg: dlta.Sg, 
				Vek: flat.Vek{dlta.Cmps...}
			}
		}	
	}
	
	// klap warps next
	for skrt := range klpsd.iterNuks(lop) {
		klpsd.klapWarp(skrt) // if its warp klap it recursively
	}
	
	// descend tree and klap fig joints
	klpsd.klapJoints(lop)
	
	// klap gids
	for skrt := range klpsd.iterNuks(lop) {
		if nuk, is := skrt.nuk.(fig.Gid); is { // if nuk is gid klap it
			 if gd, ok := klpsd.klapGid(skrt.fig, nuk); ok {
			 	skrt.fig.Nuks[gd.Sg] = gd
			 }
		}
	}
	
	// klap pins
	for skrt := range klpsd.iterNuks(lop) {
		if pn, is := skrt.nuk.(fig.Pin); is { // if nuk is pin klap it
			skrt.fig.Nuks[pn.Sg] = klpsd.klapGid(skrt.fig, pn)
		}
	}
	
	// klap flits
	for skrt := range klpsd.iterNuks(lop) {
		if gd, is := skrt.nuk.(fig.Flit); is { // if nuk is gid klap it
			skrt.fig.Nuks[gd.Sg] = klpsd.klapPin(skrt.fig, pn)
		}
	}
	
	// klap arks
}

// ******

func (self Lop) populateFigs(lop fig.Lop) {
	self.Rut = lop.Rut // sig of root fig
	self.Figs = map[klv.Sig]Fig
	nowFig = klap.Fig{Sgs: klap.Sigs{self.Rut}}
	self.Figs[self.Rut] = nowFig
	stak := []Fig{nowFig} // fifo stak to build fig tree
	for len(stak) > 0 {
		nowFig, stak = stak[0], stak[1:len(stak)] // shift next fig off stak
		
		// loop over fig.figs kids and build klap.figs
		// assign each fig a sgs address representing parent tree
		for sig, _ := range lop.Figs[nowFig.Sgs.Sig()].Kids {
			nowFig.Kids = append(nowFig.Kids, sig)
			kidFig = klap.Fig{Sgs: append(nowFig, sig)}
			self.Figs[sig] = kidFig
			stak = append(stak, kidFig) // push kid onto fifo stak
		}
	}
}

// return a channel to iterator over klvd nuks escorted by rent klpsd fig
// * loops over klpsd figs
// * for each klpsd fig loops over klvd nuks from fig.fig with same sig
// * only returns klvd nuks that havent been klpsd to klap.fig
// * channel holds buffer of skorts of klvd nuks and their rent klap figs
//   that nuk should be klpsd to
func (self Lop) iterNuks(lop fig.Lop) <-chan nukskort {
	ch := make(chan nukskort, 8); // buffer for parallel iterator gen
	go func () {
		for fSg, fg := range self.Figs { // klpsd figs
			for nSg, nk := range lop.Figs[fSg].Nuks { // klvd nuks
				if _, has := fg.Nuks[nSg]; !has { // if fig doesnt have nuk
					ch <- nukskort{fig: fg, nuk: nk}
				}
			}
		}
		close(ch)
	}()
	return ch
}
type nukskort struct {
	fig Fig
	nuk fig.Nuk
}

// recursively collapse warps based on type switch -- ignore other types
func (self Lop) klapWarp(skrt nukskort) (wrp flat.Warp) {
	switch nuk := skrt.nuk.(type) {
	case LatWarp:
		if rentSg, dltaSg, ok := klv.Dox(skrt.fig, nuk.Lat); ok {
			wrp = self.flatnWrp(rentSg, dltaSg, flat.LatWarp)
			skrt.fig.Nuks[nuk.Sg] = klap.Warp{Sg: nuk.Sg, Wrp: wrp}
		}
	case FlktWarp:
		if rentSg, dltaSg, ok := klv.Dox(skrt.fig, nuk.Sym) {
			wrp = self.flatnWrp(rentSg, dltaSg, flat.FlktWarp)
			skrt.fig.Nuks[nuk.Sg] = klap.Warp{Sg: nuk.Sg, Wrp: wrp}
		}
	case RotWarp:
		if rentSg, dltaSg, ok := klv.Dox(skrt.fig, nuk.Hed) {
			wrp = self.flatnWrp(rentSg, dltaSg, flat.RotWarp)
			skrt.fig.Nuks[nuk.Sg] = klap.Warp{Sg: nuk.Sg, Wrp: wrp}
		}
	case SkalWarp:
		if rentSg, dltaSg, ok := klv.Dox(skrt.fig, nuk.Skal) {
			wrp = self.flatnWrp(rentSg, dltaSg, flat.LatWarp)
			skrt.fig.Nuks[nuk.Sg] = klap.Warp{Sg: nuk.Sg, Wrp: wrp}
		}
	case CmboWarp:
		wrps := []flat.Warp{}
		for _, x := range nuk.Wrps {
			if rentSg, kidSg, ok := klv.Dox(skrt.fig, x) {
				if rent, ok := self.Figs[rentSg]; ok {
					if kid, ok := rent.Nuks[kidSg] {
						wrps = append(wrps, self.klapWrp(rent, kid)) // recurse
					}
				}
			}
		}
		if len(wrps) == 2 {
			wrp = flat.CmboWarp(wrps...)
			skrt.fig.Nuks[nuk.Sg] = klap.Warp{Sg: nuk.Sg, Wrp: wrp}
		}
	}
	return
}

func (self Lop) flatnWrp(rSg, dSg klv.X, wrpr flat.Warpr) (wrp flat.Warp) {
	if rent, has := self.Figs[rSg]; has {
		if dlta, has := rent.Nuks[dSg]; has {
			wrp = wrpr(dlta.Vek)
		}
	}
	return
}

func (self Lop) klapJoints(lop fig.Lop) {
	nowFig := self.Figs[self.Rut]
	
	// rut figs joint is just flattened from klvd fig joint address
	if klvdRut, has := lop.Figs[self.Rut]; has {
		nowFig.Joint = self.flatnJoint(klvdRut)	
	} else {
		panic(fmt.Sprintf("cant find root fig to klap joint!"))
	}
	
	stak := []Fig{nowFig} // fifo stak to walk fig tree
	for len(stak) > 0 {
		nowFig, stak = stak[0], stak[1:len(stak)] // shift next fig off stak
		
		// klap each childs joint warp and combine it with nowfigs
		for _, kidSig := range nowFig.Kids {
			if klvdFig, has := lop.Figs[kidSig]; has { // get klvd kid by sig
				if kidFig, has := self.Figs[kidSig]; has {
					wrp := self.flatnJoint(klvdFig)
					kidFig.Joint = flat.CmboWarp(nowFig.Joint, wrp)
					stak = append(stak, kidFig) // push kid onto fifo stak
				}
			}
		}
	}
}

// return flat warp for fig joint by finding klpsd warp at joint address
// * warps must already have been klpsd
func (self Lop) flatnJoint(fg fig.Fig) (wrp flat.Warp) {
	if rentSg, wrpSg, okk := klv.Dox(fg, klv.Xd{fig.JointNuk}); okk {
		if rent, has := self.Figs[rentSg]; has {
			if klpsdWrp, has := rent.Nuks[wrpSg]; has {
				if klpsdWrp, is := klpsdWrp.(Warp); is {
					wrp = klpsdWrp.Wrp
				}
			}
		}
	}
	if wrp == nil {
		wrp = flat.IandI // if cant find joint warp return identity matrix
	}
	return
}

func (self Lop) klapGid(fg Fig, gd fig.Gid) {
	if rentSg, wrpSg, okk := klv.Dox(fg, gd.Dlta); okk {
		if rent, has := self.Figs[rentSg]; has {
			if klpsdWrp, has := rent.Nuks[wrpSg]; has {
				if klpsdWrp, is := klpsdWrp.(Warp); is {
	
	}
}

