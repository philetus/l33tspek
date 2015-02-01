// sup i heard you like to spek 2d forms

package zib

import (
	"fmt"
	"github.com/philetus/l33tspek/tag"
	"github.com/philetus/l33tspek/flat"
	"github.com/philetus/l33tspek/fig"
)

type Sess struct {
	Dfl tag.Dufl // Dufl -> []Handl
	//Rnt XSess TODO
	//Dds []Deed TODO
	Zbs map[tag.Handl]*Zib
	Zt XZib // zut -> rut zib
}
func (self *Sess) Sup(zb *Zib) {
	if self.Zbs == nil {
		self.Zbs = map[tag.Handl]*Zib{}
	}
	self.Zbs[zb.Handl()] = zb
	return
}
func (self *Sess) Zut(zb *Zib) {
	self.Zt = XZib{zb.Hndl}
	return
}
func (self *Sess) Fekk(hndl tag.Handl) (zb *Zib, has bool) {
	if self.Zbs != nil {
		zb, has = self.Zbs[hndl]
	}
	return
}
func (self *Sess) Traas(xmrk XMark) (mrk Mark, has bool) {

	if zb, hs := self.Fekk(xmrk.ZibHandl()); hs {
	
		// if markhandl is empty return just return zib
		if xmrk.MarkHandl() == "" {
			mrk, has = zb, true
			return
		}
	
		for !has { // just loop forever?
			
			// try to get mrk from current zib
			if zb.Mrks != nil {
				if mrk, has = zb.Mrks[xmrk.MarkHandl()]; has {
					return // found mark, stop
				}
			}
			
			// otherwise check for ark
			if zb, hs = self.Fekk(zb.Ark.ZbHndl); !hs {
				return // no ark zib, fail
			}
		}
	}
	return
}

type Zib struct {
	Hndl tag.Handl
	Ark XZib
	Wrp XWarp
	Swg fig.Swag
	Nyns map[tag.Handl]bool
	Mrks map[tag.Handl]Mark
}
func (self *Zib) Handl() tag.Handl {
	return self.Hndl
}
func (self *Zib) String() string {
	return fmt.Sprintf("{Zib %v}", self.Hndl)
}
func (self *Zib) Skrib(mrk Mark) { // write mark to zib
	if self.Mrks == nil {
		self.Mrks = map[tag.Handl]Mark{}
	}
	self.Mrks[mrk.Handl()] = mrk
}
func (self *Zib) Nyun(zb *Zib) {
	if self.Nyns == nil {
		self.Nyns = map[tag.Handl]bool{}
	}
	self.Nyns[zb.Hndl] = true
}
type Mark interface {
	Handl() tag.Handl
	String() string
}
type Paan struct {
	Hndl tag.Handl
	Flts []XFlit
}
func (self Paan) Handl() tag.Handl {
	return self.Hndl
}
func (self Paan) String() string {
	return fmt.Sprintf("{Paan %v}", self.Hndl)
}
type Flit struct {
	Hndl tag.Handl
	Yks []XYok
}
func (self Flit) Handl() tag.Handl {
	return self.Hndl
}
func (self Flit) String() string {
	return fmt.Sprintf("{Flit %v}", self.Hndl)
}
type Yok struct {
	Hndl tag.Handl
	Spt XDelta // spot
	Gd XDelta // gid
}
func (self Yok) Handl() tag.Handl {
	return self.Hndl
}
func (self Yok) String() string {
	return fmt.Sprintf("{Yok %v}", self.Hndl)
}
type LatWarp struct {
	Hndl tag.Handl
	Dlta XDelta
}
func (self LatWarp) Handl() tag.Handl {
	return self.Hndl
}
func (self LatWarp) String() string {
	return fmt.Sprintf("{ComboWarp %v}", self.Hndl)
}
type RotWarp struct {
	Hndl tag.Handl
	Hd XDelta // hed for rotation
}
func (self RotWarp) Handl() tag.Handl {
	return self.Hndl
}
func (self RotWarp) String() string {
	return fmt.Sprintf("{RotWarp %v}", self.Hndl)
}
type FlektWarp struct {
	Hndl tag.Handl
	Sym XDelta // line of symmetry for reflection
}
func (self FlektWarp) Handl() tag.Handl {
	return self.Hndl
}
func (self FlektWarp) String() string {
	return fmt.Sprintf("{FlektWarp %v}", self.Hndl)
}
type SkalWarp struct {
	Hndl tag.Handl
	Skl XDelta // skal y, x
}
func (self SkalWarp) Handl() tag.Handl {
	return self.Hndl
}
func (self SkalWarp) String() string {
	return fmt.Sprintf("{SkalWarp %v}", self.Hndl)
}
type ComboWarp struct {
	Hndl tag.Handl
	Nyns []XWarp
}
func (self ComboWarp) Handl() tag.Handl {
	return self.Hndl
}
func (self ComboWarp) String() string {
	return fmt.Sprintf("{ComboWarp %v}", self.Hndl)
}
type Delta struct {
	Hndl tag.Handl
	Crds flat.Vek
}
func (self Delta) Handl() tag.Handl {
	return self.Hndl
}
func (self Delta) String() string {
	return fmt.Sprintf(
		"{Delta %v [%.2f %.2f]}", 
		self.Hndl, 
		self.Crds[0], self.Crds[1],
	)
}

