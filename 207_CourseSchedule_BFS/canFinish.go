package canFinish

func canFinish(numCourses int, prerequisites [][]int) bool {

	dict := make(map[int][]int)

	if prerequisites == nil {
		return false
	}

	for key, pair := range prerequisites {
		if val, ok := dict[key]; !ok {
			dict[pair]
		}
	}
	return false
}