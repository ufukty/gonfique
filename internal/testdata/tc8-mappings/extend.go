package config

func (scs *SpecContainers) Append(sc SpecContainer) { // if this doesn't throw error in compilation, test will complete
	*scs = append(*scs, sc)
}
