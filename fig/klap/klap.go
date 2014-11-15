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

type Fig {
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
func Klapsd(lop fig.Lop) (klpsd Lop) {
	
	// populate klap.lop with klap.figs
	klpsd.Rut = lop.Rut // sig of root fig
	klpsd.Figs = map[klv.Sig]Fig
	nowFig = klap.Fig{Sgs: klap.Sigs{klpsd.Rut}}
	klpsd.Figs[klpsd.Rut] = nowFig
	stak := []klap.Fig{nowFig} // fifo stak to build fig tree
	for len(stak) > 0 {
		nowFig, stak = stak[0], stak[1:len(stak)] // shift next fig off stak
		
		// loop over fig.figs kids and build klap.figs
		for sig, _ := range lop.Figs[nowFig.Sgs.Sig()].Kids {
			nowFig.Kids = append(nowFig.Kids, sig)
			kidFig = klap.Fig{Sgs: append(nowFig, sig)}
			klpsd.Figs[sig] = kidFig
			stak = append(stak, kidFig) // push kid onto fifo stak
		}
	}
	
	// klap deltas first
	klpsd.klapDeltas(lop.Figs)
	
	// klap warps next
	klpsd.klapWarps(lop.Figs)
	
	// klap fig joints
	
	// klap gids
	
	// klap pins
	
	// klap flits
	
}

// ******

func (self Lop) klapDeltas(figs []fig.Figs) {
	for fSg, fig := range self.Figs {
		for nSg, nuk := range figs[fSg].Nuks {
			if dlta, is := nuk.(fig.Delta); is {
				fig.Nuks[nSg] = Delta{
					Sg: nSg, 
					Vek: flat.Vek{dlta.Cmps...}
				}
			}
		}
	}
}

func (self Lop) klapWarps(figs []fig.Figs) {
	for fSg, fig := range klpsd.Figs {
		for nSg, nuk := range figs[fSg].Nuks {
			self.klapWrp(fig, nuk)
		}
	}
}

func (self Lop) klapWrp(fig Fig, wrp fig.Nuk) (fltWrp flat.Warp) {
	switch wrp := wrp.(type) {
	case LatWarp:
		if rentSg, dltaSg, ok := klv.Dox(fig, wrp.Lat) {
			if rent, ok := self.Figs[rentSg]; ok {
				if dlta, ok := rent.Nuks[dltaSg] {
					fltWrp = flat.LatWarp(dlta.Vek)
					fig.Nuks[nSg] = klap.Warp{Sg: nSg, Wrp: fltWrp}
				}
			}
		}
	case FlktWarp:
		if rentSg, dltaSg, ok := klv.Dox(fig, wrp.Sym) {
			if rent, ok := self.Figs[rentSg]; ok {
				if dlta, ok := rent.Nuks[dltaSg] {
					fltWrp = flat.FlktWarp(dlta.Vek)
					fig.Nuks[nSg] = klap.Warp{Sg: nSg, Wrp: fltWrp}
				}
			}
		}
	case RotWarp:
		if rentSg, dltaSg, ok := klv.Dox(fig, wrp.Hed) {
			if rent, ok := self.Figs[rentSg]; ok {
				if dlta, ok := rent.Nuks[dltaSg] {
					fltWrp = flat.RotWarp(dlta.Vek)
					fig.Nuks[nSg] = klap.Warp{Sg: nSg, Wrp: fltWrp}
				}
			}
		}
	case SkalWarp:
		if rentSg, dltaSg, ok := klv.Dox(fig, wrp.Skal) {
			if rent, ok := self.Figs[rentSg]; ok {
				if dlta, ok := rent.Nuks[dltaSg] {
					fltWrp = flat.SkalWarp(dlta.Vek)
					fig.Nuks[nSg] = klap.Warp{Sg: nSg, Wrp: fltWrp}
				}
			}
		}
	case CmboWarp:
		warps := []flat.Warp{}
		for _, x := range wrp.Wrps {
			if rentSg, kidSg, ok := klv.Dox(fig, x) {
				if rent, ok := self.Figs[rentSg]; ok {
					if kid, ok := rent.Nuks[kidSg] {
						self.klapWrp(rent, kid)
					}
				}
			}
		}
	}
}

