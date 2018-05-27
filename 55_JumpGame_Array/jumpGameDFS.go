package jumpGame

func canJumpDFS(nums []int) bool {

	if len(nums) == 1 {
		return true
}

	if len(nums) == 2 {
		if nums[0] < 1 {
			return false
		} else {
			return true
		}
	}

	graph := make(map[int][]int, 0)

	for index, value := range nums[:len(nums)-1] {

		
		size := min(len(nums)-index-1, value)
		for i :=1; i <= size; i++ {
			graph[index] = append(graph[index], index+i)
		}
	}

	var visited []bool
	for i := 0; i < len(nums); i++ {
		visited = append(visited, false)
	}
	key := len(nums)-1
	result := findDFS(graph, 0, visited, key )

	//graph[1] = []int{1}

	return result
}

func findDFS(graph map[int][]int, v int, visited []bool, key int) bool {

	found := false
	visited[v] = true

	edges := graph[v]

	for _, vertex := range edges {

		if vertex == key {
			return true
		}
		
		if visited[vertex] == false {
			found = findDFS(graph, vertex, visited, key)
		}
	}

	return found
}


func min(a int, b int) int {
	if a < b  {
		return a
	}
	return b
}