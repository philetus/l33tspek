package main

import (
	"fmt"
	"github.com/philetus/eezl"
	"github.com/philetus/l33tspek/klv"
	"github.com/philetus/l33tspek/fig"
)

func butterfly() *fig.Fig {

    // build left profile of butterfly as fig
    lp := fig.Fig{S: "lp"}
    lp.Swlo(fig.Pin{
    	S: "p0", 
    	At: lp.Swlo(fig.Vek{
    		V: lp.Swlo(fig.Skalr{N: 35.0}),
    		U: lp.Swlo(fig.Skalr{N: 25.0}),
    	}),
    })
    lp.Swlo(fig.Pin{
    	S: "p1", 
    	At: lp.Swlo(fig.Vek{
    		V: lp.Swlo(fig.Skalr{N: 0.0}),
    		U: lp.Swlo(fig.Skalr{N: 0.0}),
    	}),
    })
    lp.Swlo(fig.Pin{
    	S: "p2", 
    	At: lp.Swlo(fig.Vek{
    		V: lp.Swlo(fig.Skalr{N: 35.0}),
    		U: lp.Swlo(fig.Skalr{N: 25.0}),
    	}),
    })
    lp.Swlo(fig.Ra{S:"r0", A: klv.Xd{"p0"}, B: klv.Xd{"p1"}})
    lp.Swlo(fig.Ra{S:"r1", A: klv.Xd{"p1"}, B: klv.Xd{"p2"}})
    lp.Swlo(fig.Yok{
    	S: "y0",
    	A: klv.Xd{"r0"},
    	B: klv.Xd{"r1"},
    	Van: lp.Swlo(fig.Vek{
    		S: "vDown",
    		V: lp.Swlo(fig.Skalr{N: 1.0}),
    		U: lp.Swlo(fig.Skalr{N: 0.0}),
    	}),
    })
    
    // joint for left profile
    lp.Swlo(fig.LatWarp{
    	S: "Joint",
    	Dlta: lp.Swlo(fig.Vek{
    		S: "vJnt",
    		V: lp.Swlo(fig.Skalr{N: 0.0}),
    		U: lp.Swlo(fig.Skalr{S: "sWaist", N: 15.0}),
    	}),
    })
    
    // right profile is new fig with left profile as ark and reflect joint
    rp := fig.Fig{S: "rp", A: lp}
    
    // palmp joint to give right profile mirrored position
    rp.Swlo(fig.CmboWarp{
    	S: "Joint",
    	A: lp.Swlo(fig.LatWarp{Dlta: klv.Xd{"vJnt"}}),
    	B: lp.Swlo(fig.FlktWarp{Sym: klv.Xd{"vDown"}})
    })
    
	// build main butterfly fig
    b := fig.Fig{S: "bfly"}
    
    // joint left and right profiles into butterfly fig
    b.Jnt(lp)
    b.Jnt(rp)
        
    // add centerline pins to top fig
    b.Swlo(fig.Pin{
    	S: "p0", 
    	At: lp.Swlo(fig.Vek{
    		V: lp.Swlo(fig.Skalr{N: 20.0}),
    		U: lp.Swlo(fig.Skalr{N: 0.0}),
    	}),
    })
    b.Swlo(fig.Pin{
    	S: "p1", 
    	At: lp.Swlo(fig.Vek{
    		V: lp.Swlo(fig.Skalr{N: -15.0}),
    		U: lp.Swlo(fig.Skalr{N: 0.0}),
    	}),
    })
	
	// add top and bottom ras, use XX to address pins in subfigs
    b.Swlo(fig.Ra{S:"r0", A: klv.XX{"lp", klv.Xd{"p0"}}, B: klv.Xd{"p0"}})
    b.Swlo(fig.Ra{S:"r1", A: klv.Xd{"p0"}, B: klv.XX{"rp", klv.Xd{"p0"}}})
    b.Swlo(fig.Ra{S:"r2", A: klv.XX{"rp", klv.Xd{"p2"}}, B: klv.Xd{"p1"}})
    b.Swlo(fig.Ra{S:"r3", A: klv.Xd{"p1"}, B: klv.XX{"lp", klv.Xd{"p2"}}})
	
	// add yoks
    b.Swlo(fig.Yok{
    	S: "y0",
    	A: klv.XX{"lp", klv.Xd{"r0"}},
    	B: klv.Xd{"r0"},
    	Van: lp.Swlo(fig.Vek{
    		S: "vDown",
    		V: lp.Swlo(fig.Skalr{N: 1.0}),
    		U: lp.Swlo(fig.Skalr{N: 0.0}),
    	}),
    })
    b.Swlo(fig.Yok{
    	S: "y1",
    	A: klv.Xd{"r0"},
    	B: klv.Xd{"r1"},
    	Van: klv.Xd{"vDown"},
    })
    b.Swlo(fig.Yok{
    	S: "y2",
    	A: klv.Xd{"r1"},
    	B: klv.XX{"rp", klv.Xd{"r1"}},
    	Van: klv.Xd{"vDown"},
    })
    b.Swlo(fig.Yok{
    	S: "y3",
    	A: klv.XX{"rp", klv.Xd{"r1"}},
    	B: klv.Xd{"r2"},
    	Van: klv.Xd{"vDown"},
    })
    b.Swlo(fig.Yok{
    	S: "y4",
    	A: klv.Xd{"r2"},
    	B: klv.Xd{"r3"},
    	Van: klv.Xd{"vDown"},
    })
    b.Swlo(fig.Yok{
    	S: "y5",
    	A: klv.Xd{"r3"},
    	B: klv.XX{"lp", klv.Xd{"r1"}},
    	Van: klv.Xd{"vDown"},
    })
    
    // wings flit
    b.Swlo(fig.Flit{
    	S: "fWings",
    	Yoks: []X{
    		klv.Xd{"y0"}, 
    		klv.Xd{"y1"}, 
    		klv.Xd{"y2"}, 
    		klv.XX{"rp", klv.Xd{"y0"}},
    		klv.Xd{"y3"}, 
    		klv.Xd{"y4"}, 
    		klv.Xd{"y5"}, 
    		klv.XX{"lp", klv.Xd{"y0"}},
    	},
    })
    
    return &b
}

func main() {
	xscreen := eezl.Xconnect() // get connection to x server screen
	ez := xscreen.NewEezl(600, 600) // open window
	fmt.Printf("created new eezl!\n")
	
	var pressed_flag bool = false
	
    var ln_thk float64 = 10.0
    ln_clr := [4]float64{1.0, 0.0, 0.0, 0.6} // slightly translucent red
    //var bnd_thk float64 = 6.0
    //bnd_clr := [4]float64{0.0, 0.0, 0.0, 0.4} // translucent gray
    bg_clr := [4]float64{1.0, 1.0, 1.0, 1.0} // opaque white
    
	for {
		select {
			
			case inpt := <-ez.InputPipe:
				//fmt.Printf(".%d", inpt.Flavr)
				switch inpt.Flavr {
				
				case eezl.PointerPress:
					pressed_flag = true
					fmt.Printf("pressed pointer!\n")


				case eezl.PointerRelease:
					pressed_flag = false
					fmt.Printf("pressed pointer!\n")
					
					//ez.Stain() // trigger eezl redraw
					
				case eezl.PointerMotion:
					if pressed_flag {
						fmt.Printf("dragged pointer!\n")
						
						//ez.Stain() // trigger eezl redraw
					}
				
				case eezl.KeyPress:
					fmt.Println(inpt.Stroke.Name)
				}
				
			case gel := <- ez.GelPipe:
			
				// fill background
				gel.SetColor(bg_clr[0], bg_clr[1], bg_clr[2], bg_clr[3])
				gel.Coat()
				
				// set line weight and color
				gel.SetColor(ln_clr[0], ln_clr[1], ln_clr[2], ln_clr[3])
				gel.SetWeight(ln_thk)
				
				// build butterfly
				bf := butterfly()
				
				// klap wing flit of fig into global veks in list of bezs
				bzs := bf.Bezize("fWings")
				
				// draw wing bezs to eezl
				gel.Jmto(bzs[0].Coords[0][0], bzs[0].Coords[0][1])
				for _, bz := range bzs {
					gel.Beto(bz.Coords[1][0], bz.Coords[1][1],
							 bz.Gids[0][0], bz.Gids[0][1],
							 bz.Gids[1][0], bz.Gids[1][1])
				}
				gel.Seal()
				gel.Stroke()
				gel.Shake()
				
				// draw band
				//if pressed_flag {
				//}
				
				// send trigger sig
				gel.Trigger()
		}
	}
}
