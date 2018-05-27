package jump

import ("math")

func jump(nums []int) int {

	if len(nums) < 1 {
		return 0
	}

	minDistance := math.MaxInt32
	graph := make(map[int][]int, 0)

	for index, value:= range nums[:len(nums)-1] {
		maxJump := min(len(nums)-index-1, value)
		for i :=1; i<= maxJump; i++ {
			graph[index] = append(graph[index], i+index )
		}
	}

	distance := make([]int, len(nums))

	for i := range distance {
		distance[i] = -1
	}

	distance[0] = 0

	vertex, queue := 0, []int{0}

	for len(queue) > 0 {

		vertex, queue = queue[0], queue[1:]

		edges := graph[vertex]

		for _ , v := range edges {
			if distance[v] == -1 {
				queue = append(queue, v)
				distance[v] = distance[vertex] + 1
			}

			if v == len(nums) - 1 {
					minDistance = min(minDistance, distance[v])
			}
		}
	}

    return minDistance
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}