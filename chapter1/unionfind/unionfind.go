package unionfind

import "fmt"

type UF struct {
	id    []int
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

func (uf *UF) QuickFind(p int) int {
	return uf.id[p]
}

func (uf *UF) Union(p, q int) {
	pID := uf.QuickFind(p)
	qID := uf.QuickFind(q)
	if pID == qID {
		return
	}
	for i := 0; i < len(uf.id); i++ {
		if uf.id[i] == pID {
			uf.id[i] = qID
		}
		uf.count--
	}
}

func (uf *UF) Find(p int) int {
	for p != uf.id[p] {
		p = uf.id[p]
	}
	return p
}

func (uf *UF) FindForPathCompression(p int) int {
	if p == uf.id[p] {
		return p
	}
	uf.id[p] = uf.FindForPathCompression(uf.id[p])
	return uf.id[p]
}

func (uf *UF) QuickUnion(p, q int) {
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)
	if pRoot == qRoot {
		return
	}
	uf.id[pRoot] = qRoot
	uf.count--
}

func main() {
	fmt.Println("vim-go")
}
