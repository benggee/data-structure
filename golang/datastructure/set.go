package datastructure

type set struct {
	data *linkList
}

func Set() *set {
	return &set{data: Linklist()}
}

func (s *set) Add(data interface{}) {
	s.data.AddHead(data)
}

func (s *set) Del(data interface{}) {
	s.data.RemoveByValue(data)
}

func (s *set) Contains(data interface{}) bool {
	return s.data.Contains(data)
}

func (s *set) Size() int {
	return s.data.size
}

func (s *set) IsEmpty() bool {
	return s.data.size == 0
}

func (s *set) ToString() {
	s.data.ToString()
}
