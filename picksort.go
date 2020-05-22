package picksort

import (
	"sort"
)

// PickInt 从数据类型中取整数数据
type PickInt func(interface{}) int

// PickString 从数据类型中取字符串数据
type PickString func(interface{}) string

// Sort 取数据后排序
func (pi PickInt) Sort(ItemsToSort []interface{}, less func(i, j int) bool) {
	sort.Sort(&sortHolder{
		Items:           ItemsToSort,
		Pi:              pi,
		LessInternalInt: less,
	})
}

// Sort 取数据后排序
func (ps PickString) Sort(ItemsToSort []interface{}, less func(i, j string) bool) {
	sort.Sort(&sortHolder{
		Items:              ItemsToSort,
		Ps:                 ps,
		LessInternalString: less,
	})
}

type sortHolder struct {
	Items              []interface{}
	Pi                 PickInt
	Ps                 PickString
	LessInternalInt    func(i, j int) bool
	LessInternalString func(i, j string) bool
}

func (s *sortHolder) Len() int {
	return len(s.Items)
}

func (s *sortHolder) Swap(i, j int) {
	s.Items[i], s.Items[j] = s.Items[j], s.Items[i]
}

func (s *sortHolder) Less(i, j int) bool {
	if s.Pi != nil {
		return s.LessInternalInt(s.Pi(s.Items[i]), s.Pi(s.Items[j]))
	}

	if s.Ps != nil {
		return s.LessInternalString(s.Ps(s.Items[i]), s.Ps(s.Items[j]))
	}

	panic("Ps Pi both nil")
}
