package CustomTypes

type Set struct {
	m map[interface{}]bool
}

func NewSet(init ...interface{}) *Set {
	set := &Set{
		m: make(map[interface{}]bool),
	}
	if init != nil && len(init) > 0 {
		for _, v := range init {
			set.Add(v)
		}
	}
	return set
}

func (s *Set) Add(item interface{}) {
	if item == nil {
		return
	} else {
		s.m[item] = true
	}
}

func (s *Set) Remove(item interface{}) {
	delete(s.m, item)
}

func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

func (s *Set) Vals() []string {
	vals := make([]string, 0)
	for k, _ := range s.m {
		vals = append(vals, k.(string))
	}
	return vals
}
