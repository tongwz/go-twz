package pkgC

type CS struct {
	Name    string
	Age     uint32
	Company string
}

func (c *CS) SetName(name string) {
	c.Name = name
}
