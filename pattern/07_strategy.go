package pattern

import "fmt"

type StratAlgo interface {
	DoMagicAlgo()
}
type DarkMagic struct {
}

func (dm *DarkMagic) DoMagicAlgo() {
	fmt.Println("Using some dark magic algorithms")
}

type ArcaneMagic struct {
}

func (am *ArcaneMagic) DoMagicAlgo() {
	fmt.Println("using some arcane magic algorithms")
}

type Context struct {
	strat StratAlgo
}

func (c *Context) SetAlgo(strat StratAlgo) {
	c.strat = strat
}
func (c *Context) UseMagic() {
	c.strat.DoMagicAlgo()
}
