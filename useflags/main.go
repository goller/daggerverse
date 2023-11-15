package main

type UseFlags struct{}

// example usage: "dagger call style-1"
func (m *UseFlags) Style1() *Container {
	return dag.
		MyFeature().
		WithFlag1("howdy").
		Style1()
}

// example usage: "dagger call style-2"
func (m *UseFlags) Style2() *Container {
	return dag.MyFeature().Style2(MyFeatureStyle2Opts{Flag1: "howdy"})
}
