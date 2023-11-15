package main

type Useflags struct{}

// example usage: "dagger call style-1"
func (m *Useflags) Style1() *Container {
	return dag.
		Flags().
		WithFlag1("howdy").
		Style1()
}

// example usage: "dagger call style-2"
func (m *Useflags) Style2() *Container {
	return dag.Flags().Style2(FlagsStyle2Opts{Flag1: "howdy"})
}
