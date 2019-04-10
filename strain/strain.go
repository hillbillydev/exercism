package strain

type Ints []int
type Strings []string
type Lists [][]int

type intPred func(int) bool
type stringPred func(string) bool
type listPred func([]int) bool

func (ints Ints) Keep(p intPred) Ints {
	var res Ints
	for _, i := range ints {
		if ok := p(i); ok {
			res = append(res, i)
		}
	}
	return res
}

func (ints Ints) Discard(p intPred) Ints {
	var res Ints
	for _, i := range ints {
		if ok := p(i); !ok {
			res = append(res, i)
		}
	}
	return res
}

func (strings Strings) Keep(p stringPred) Strings {
	var res Strings
	for _, i := range strings {
		if ok := p(i); ok {
			res = append(res, i)
		}
	}
	return res
}

func (lists Lists) Keep(p listPred) Lists {
	var res Lists
	for _, i := range lists {
		if ok := p(i); ok {
			res = append(res, i)
		}
	}
	return res
}
