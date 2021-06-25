package statemachine

type State struct {
	Initial     string
	End         string
	Intermedite []string
}

type States struct {
	iState State
	Set    map[string]bool
}

func (st *States) Add(s State) {
	ss := make(map[string]bool)
	ss[s.Initial] = true
	ss[s.End] = true
	for _, v := range s.Intermedite {
		ss[v] = true
	}
	st.iState = s
}
func (st *States) IsValid(s string) bool {
	_, prs := st.Set[s]
	if prs {
		return true
	}
	return false
}

func (st *States) IsEndState(s string) bool {
	return st.iState.End == s
}
