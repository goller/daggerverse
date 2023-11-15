package main

type Flags struct {
	Flag1 string
	Flag2 string
}

func (m *Flags) WithFlag1(flag string) *Flags {
	m.Flag1 = flag
	return m
}

func (m *Flags) WithFlag2(flag string) *Flags {
	m.Flag2 = flag
	return m
}

// example usage: "dagger call with-flag-1 --flag 1 style-1"
func (m *Flags) Style1() *Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", m.Flag1, m.Flag2})
}

// example usage: "dagger call style-2 --flag-1 1"
func (m *Flags) Style2(flag1 Optional[string], flag2 Optional[string]) *Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", flag1.GetOr(""), flag2.GetOr("")})
}
