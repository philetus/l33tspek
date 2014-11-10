// X is interface to address hok bons for deferred lookup

package x

import (
	//"fmt"
	//"reflect"
	//"strconv"
)

type Sig string

type Bon interface {
	Sig() Sig
	//Kids() map[string]X
}

// cleave reality at the joints --> ham hocks
// hoks hold bones referenced by sig
type Hok interface {
	Sig() Sig
	Ark() *Hok
	Kid(Sig) (*Hok, bool)
	Jnt(*Hok) *Hok
	Bon(Sig) (*Bon, bool)
	Et(Bon) X
	Dox(X) (*Bon, bool)
}

// mixin struct for hok definition
type Spam struct {
	S Sig
	A *Hok
	Kids map[Sig]*Hok
	Deks map[Sig]*Bon
}
func (self Spam) Sig() Sig {
	return self.S
}
func (self Spam) Ark() *Hok {
	return self.A
}
func (self Spam) Kid(sg Sig) (kd *Hok, ok bool) {
	kd, ok = self.Kids[sg]
	return
}
func (self Spam) Jnt(kid Hok) *Hok {
	self.Kids[kid.Sig()] = &kid
	return &kid
}
func (self Spam) Bon(sg Sig) (bn *Bon, ok bool) {
	bn, ok = self.Deks[sg]
	return
}
func (self Spam) Et(bn Bon) X {
	sg := bn.Sig()
	self.Deks[sg] = &bn
	return Xd{S: sg}
}

//
func (self Spam) Dox(x X) (bn *Bon, ok bool) {

	// if this is not leaf x try to recurse on kid hok
	if !x.IsLef() {
		if kid, has := self.Kid(x.Sig()); has {
			return (*kid).Dox(x.Kid())
		}
	
	// otherwise if it is leaf try to return bon by x.sig
	} else {
		if bn, ok = self.Deks[x.Sig()]; ok {
			return
		}
	}
	
	// otherwise try to dox with ark
	if self.A != nil {
		return (*self.A).Dox(x)
	}
	
	// otherwise fail
	return nil, false
}

// X is lookup address of bon
type X interface { 
	Kid() X
	Sig() Sig
	IsLef() bool
}

//
//  *** Xd has sig indexed in current fig
//
type Xd struct {
	S Sig
}
func (self Xd) Kid() X {
	panic("x.Xd.Kid() called!")
}
func (self Xd) Sig() Sig {
	return self.S
}
func (self Xd) IsLef() bool {
	return true
}

//
// *** XX has sig of kid node to apply 
//
type XX struct {
	S Sig
	K X
}
func (self XX) Kid() X {
	return self.K
}
func (self XX) Sig() Sig {
	return self.S
}
func (self XX) IsLef() bool {
	return false
}

