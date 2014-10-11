// geometry for l33tspek

package flat

import (
	//"fmt"
)

// unique name for tracking instances
type Sig uint

func makeKrystr(nxt uint) func() Sig {
	return func() Sig {
		sig := Sig(nxt)
		nxt++
		return sig
	}
}

var (
	krystr = makeKryster(1) // closure to gen unique ids for tings
)

// implements flake interface
type Ting struct {
	sig Sig
}
func (self *Ting) Kryst() {
	self.sig = krystr() // use closure to gen unique id
}

// every ting is a special snowflake with a unique sig ;)
type Flake interface {
	Sig() Sig
}
func (self *Ting) Sig() Sig {
	return self.sig
}

// 2d vector
type Vek [2]float64

// position in 2d plane; implements spot
type Pin struct {
	Ting
	coords Vek
}
func NewPin(crds Vek) (pn *Pin) {
	pn.Kryst() // set unique id
	pn.coords = crds
}
func (self *Pin) Stick(crds Vek) {
	self.coords = crds
}

// interface for endpoint of a flit
type Spot interface {
	Coords() Vek
}
func (self *Pin) Coords() Vek {
	return self.coords
}

// straight line between 2 spots; implements flit
type Ra struct {
	Ting
	ends [2]*Spot
}
func NewRa(spts ...*Spot) (r *Ra) {
	r.Kryst()
	for i := 0; i < 2; i++ {
		r.ends[i] = spts[i]
	}
}
// replace one of the ends of ra with new spot, returns old spot
func (self *Ra) Jerk(i int, spt *Spot) (ospt *Spot) {
	ospt = self.ends[i]
	self.ends[i] = spt
}

// bezier curve between 2 spots; curve defined by gid veks
type Be struct {
	Ra
	gids [2]Vek
}
func NewBe(gds []Vek, spts ...*Spot) (b *Be) {
	b.Kryst()
	for i := 0; i < 2; i++ {
		b.gids[i] = gds[i]
		b.ends[i] = spts[i]
	}
}
// overrides ra jerk, replaces both gid & spot, returns old gid & spot
func (self *Be) Jerk(i int, gd Vek, spt *Spot) (ogd Vek, ospt *Spot) {
	ogd, ospt = self.gids[i], self.ends[i]
	self.gids[i], self.ends[i] = gd, spt
}
// access be gids
func (self *Be) Gids() []Vek {
	return self.gids[:2]
}

// interface for a connection between two spots
type Flit interface {
	Ends() []*Spot
}
func (self *Ra) Ends() []*Spot {
	return self.ends[:2]
}

// a contiguous series of flits; subsequent flits should share a spot
type Trak struct {
	Ting
	flits []*Flit
}
func NewTrak(flts ...*Flit) (trk *Trak) {
	trk.Kryst()
	trk.flits = make([]*Flit, len(flts))
	copy(trk.flits, flts)
}
func (self *Trak) Flits() []*Flit {
	return self.flits
}
func (self *Trak) Loops() bool {
	return self.flits[0].Ends()[0] == self.flits[len(self.flits)-1].Ends()[1]
}
