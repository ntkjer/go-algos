package unionfind

import "fmt"

type UF struct {
	id    []int
	size  []int
	count int
}

func NewUF(n int) *UF {
	uf := UF{count: n}
	for i := 0; i < n; i++ {
		uf.id[i] = i
	}
	return &uf
}

func (uf *UF) Count() int {
	return uf.count
}

func (uf *UF) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *UF) Union(p, q int) {
	i := uf.Find(p)
	j := uf.Find(q)

	if i == j {
		return
	}

	if uf.size[i] > uf.size[j] {
		uf.id[j] = i
		uf.size[i] += uf.size[j]
	} else {
		uf.id[i] = j
		uf.size[j] += uf.size[i]
	}
	uf.count--

}

func (uf *UF) Find(p int) int {
	for p != uf.id[p] {
		p = uf.id[p]
	}
	return p
}

func main() {
	fmt.Println("vim-go")
}
