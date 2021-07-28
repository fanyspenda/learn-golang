package module

//Students ...
type Students struct {
	Name  []string
	Score []int
}

//Average ...
func (s Students) Average() (avg int) {
	for _, v := range s.Score {
		avg += v
	}
	avg /= len(s.Score)
	return avg
}

//Min ...
func (s Students) Min() (min int, name string) {
	studentsMap := map[string]int{}
	for i, v := range s.Name {
		studentsMap[v] = s.Score[i]
	}

	min = s.Score[0]
	name = s.Name[0]
	for key, value := range studentsMap {
		if min > value {
			min = value
			name = key
		}
	}
	return min, name
}

//Max ...
func (s Students) Max() (max int, name string) {
	studentsMap := map[string]int{}
	for i, v := range s.Name {
		studentsMap[v] = s.Score[i]
	}

	max = s.Score[0]
	name = s.Name[0]
	for key, value := range studentsMap {
		if max < value {
			max = value
			name = key
		}
	}
	return max, name
}
