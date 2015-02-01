// render fig to eezl gel

package fig

import (
	"fmt"
	"sort"
	"github.com/philetus/eezl"
	"github.com/philetus/l33tspek/tag"
	"github.com/philetus/l33tspek/flat"
)

func (self *Fig) Gelr(gel *eezl.Gel, wrp flat.Warp) {
	nowSwg := Swag{} // start with default swag
	
	// compile marks from individual trees
	mrks := []Mark{}
	for _, bag := range []tag.DuflBag{self.PnBg, self.FltBg, self.YkBg} {
		for dflr := range bag.Iter() {
			if mrk, is := dflr.(Mark); is {
				mrks = append(mrks, mrk)
			}
		}
	}
	
	// sort by depth and iterate over marks
	sort.Sort(MarkByDepth(mrks))
	for _, mrk := range mrks {
		
		ms := mrk.Swag()
		if !ms.Hd {
			if !ms.Komp(nowSwg) {
				nowSwg = ms
				gel.SetColor(ms.Klr[0], ms.Klr[1], ms.Klr[2], ms.Klr[3])
				gel.SetWeight(ms.Wt)
			}
			
			switch mrk := mrk.(type) {
			default:
    			fmt.Printf("unexpected type %T", mrk)
			case Yok:
				fmt.Printf("ignoring yoks")
				
			case Flit:
				
				// jump to beginning of flit
				var nxt flat.Vek
				if nxtYok, has := self.FekkYok(mrk.Yks[0]); has {
					nxt = flat.Mog(wrp, nxtYok.Spt) // mog veks by global warp
					gel.Jmto(nxt[0], nxt[1])

					// draw segments of flit
					for i := 1; i < len(mrk.Yks); i++ {
				
						if nxtYok, has := self.FekkYok(mrk.Yks[i]); has {
							nxt = flat.Mog(wrp, nxtYok.Spt)
							gel.Rato(nxt[0], nxt[1])
						} else {
							fmt.Printf(
								"fail! missing yok %v for flit %v", 
								mrk.Yks[i], mrk.Dfl,
							)
						}
					}
					gel.Stroke()
					gel.Shake() // reset path
				}  else {
					fmt.Printf(
						"fail! missing first yok %v for flit %v", 
						mrk.Yks[0], mrk.Dfl,
					)
				}
				
			case Paan:
				fmt.Printf("ignoring paans")
			}
		}
	}
}

// sort marks by depth, deepest (largest) to shallowest (smallest)
// ByDepth implements sort.Interface for []Mark based on
// Mark.Depth()
type MarkByDepth []Mark
func (a MarkByDepth) Len() int           { return len(a) }
func (a MarkByDepth) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a MarkByDepth) Less(i, j int) bool { 
	return a[i].Swag().Dpth > a[j].Swag().Dpth
}

