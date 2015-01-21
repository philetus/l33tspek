// handls & dufls to tag spek

package tag

import (
	"fmt"
)

type Handl string
type Dufl []Handl
func (self Dufl) Komp(othr Dufl) bool {
	if len(self) != len(othr) {
		return false
	}
	for i := 0; i < len(self); i++ {
		if self[i] != othr[i] {
			return false
		}
	}
	return true
}

type Duflr interface {
	Dufl() Dufl
}

type nodMap map[Handl]*duflrNod
type duflrNod struct {
	brnchs nodMap
	lf Duflr
}
type DuflBag struct {
	rut nodMap
}
func (self *DuflBag) Iter() <-chan Duflr {
	ch := make(chan Duflr, 8)
	
	
    go func() {
		
		// push rut nods on stak
		var stak []*duflrNod
		if self.rut != nil {
			for _, dflr := range self.rut {
				stak = append(stak, dflr)
			}
		}
		
        for len(stak) > 0 {
			var nod *duflrNod
        	nod, stak = stak[0], stak[1:]
        	for _, brnch := range nod.brnchs {
        		stak = append(stak, brnch)
        	}
        	if nod.lf != nil {
            	ch <- nod.lf
            }
        }
        close(ch) // Remember to close or the loop never ends!
    }()
	return ch
}
func (self *DuflBag) Stass(dflr Duflr) {
	
	dfl := dflr.Dufl()

	if len(dfl) == 0 {
		panic(fmt.Sprintf("cant stass zero length dufl in bag: %v", dfl))
	}
	
	// get rut nod with first handl from dufl (create if doesnt exist)
	nowHndl := dfl[0]
	var nowNod *duflrNod
	var has bool
	if self.rut == nil {
		self.rut = nodMap{}
	}
	if nowNod, has = self.rut[nowHndl]; !has {
		nowNod = &duflrNod{brnchs: nodMap{}}
		self.rut[nowHndl] = nowNod
	}
	
	// if there are more handls get (or create) branch nod for each
	for i := 1; i < len(dfl); i++ {
		nowHndl = dfl[i]
		var nxtNod *duflrNod
		if nxtNod, has = nowNod.brnchs[nowHndl]; !has {
			nxtNod = &duflrNod{brnchs: nodMap{}}
			nowNod.brnchs[nowHndl] = nxtNod
		}
		nowNod = nxtNod
	}
	
	// add duflr as leef to now nod
	nowNod.lf = dflr
}

func (self *DuflBag) Fekk(dfl Dufl) (Duflr, bool) {
	if len(dfl) == 0 {
		panic("cant fekk duflr with zero length dufl")
	}
	nowHndl := dfl[0]
	if nowNod, has := self.rut[nowHndl]; has {
	
		// iterate over rest of handls in dufl and get last branch nod
		for i := 1; i < len(dfl); i++ {
			nowHndl = dfl[i]
			if nowNod, has = nowNod.brnchs[nowHndl]; !has {
				return nil, false
			}
		}
		
		// try to find leef mark in last branch nod
		if nowNod.lf != nil {
			return nowNod.lf, true
		}
	}
	return nil, false
}
