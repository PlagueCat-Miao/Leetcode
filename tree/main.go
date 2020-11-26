package main



type Stack struct{
	Vals []rune
}


func (s *Stack)Add(a rune){
	s.Vals  = append( s.Vals,a)

}
func (s *Stack)Eraser(){
	ll:=len(s.Vals)
	if ll !=0{
		s.Vals  = append( s.Vals[:ll-1])
	}
}

func  (s *Stack)Equal(b Stack)bool{
	ls := len(s.Vals)
	lb := len(b.Vals)
	if ls != lb {return false}
	for i:=0;i<ls;i++{
		if s.Vals[i] != b.Vals[i]{
			return false
		}
	}


	return true
}
func backspaceCompare(S string, T string) bool {
	SS :=Stack{make([]rune ,0)}
	TT :=Stack{make([]rune ,0)}

	for _, word:= range S{
		if word == '#'{
			SS.Eraser()
		}else{
			SS.Add(word)
		}
	}

	for _, word:= range T{
		if word == '#'{
			TT.Eraser()
		}else{
			TT.Add(word)
		}
	}

	return TT.Equal(SS)

}