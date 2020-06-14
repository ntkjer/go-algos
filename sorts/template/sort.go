package template

//TODO: Generics and Comparable interfaces for 1.15? too verbose.

type IntSort interface {
	Less(int) bool
	IsSorted([]int) bool
	Sort([]int)
	Show([]int)
}

type StrSort interface {
	Less(string) bool
	IsSorted([]string) bool
	Sort([]string)
	Show([]string)
}

type InterfaceSort interface {
	Less(interface{}) bool
	IsSorted([]interface{}) bool
	Sort([]interface{})
	Show([]interface{})
}
