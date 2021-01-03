package side_packages

import(
	"fmt"
	"sort"
)

type Human struct {
	Name string
	Age int
}

// необходимо создать новый тип, который будем сортировать
type People []Human

// ВСЕ эти три метода обязательно прописывать.
// без них сортировка невозможна
func (ppl People) Len() int{
	return len(ppl)
}
func (ppl People) Less(i, j int) bool {
	return ppl[i].Name < ppl[j].Name
}
func (ppl People) Swap(i, j int) {
	ppl[i], ppl[j] = ppl[j], ppl[i]
}

func Sorting(){
	kids := []Human{
		{"Ken", 18},
		{"Morgana", 21},
	}
	// Применяем новый тип к списку
	sort.Sort(People(kids))
	fmt.Println(kids)
}

