package types

type Project struct {
	Name     string
	Language string
	Modules  []string
	Git      bool
}
