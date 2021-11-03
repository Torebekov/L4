package tetris

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)
func Run() {
	fmt.Println(get("test/data/struct.go", 3))
}
func getPermutations(x int) [][]int {
	var res [][]int
	nums := make([]int, 0, x)
	for i := 0; i < x; i++ {
		nums = append(nums, i)
	}
	n := len(nums)
	var permute func(int)
	permute = func(first int) {
		if first == n {
			temp := make([]int, n)
			copy(temp, nums)
			res = append(res, temp)
		}
		for i := first; i < n; i++ {
			nums[first], nums[i] = nums[i], nums[first]
			permute(first + 1)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
	permute(0)
	return res
}
func checkSize(s string) int {
	sizes := map[string]int{
		"bool":       1,
		"int8":       1,
		"int16":      2,
		"int32":      4,
		"int64":      8,
		"uint8":      1,
		"uint16":     2,
		"uint32":     4,
		"uint64":     8,
		"float32":    4,
		"float64":    8,
		"complex64":  8,
		"complex128": 16,
		"byte":       1,
		"rune":       4,
		"int":        8,
		"uint":       8,
		"uintptr":    8,
	}
	if strings.Contains(s, "*") || strings.Contains(s, "map") || strings.Contains(s, "func") || strings.Contains(s, "chan") {
		return 8
	} else if strings.Contains(s, "interface") || strings.Contains(s, "string") {
		return 16
	} else if strings.Contains(s, "[]") {
		return 24
	} else {
		v, ok := sizes[s]
		if ok {
			return v
		} else {
			fmt.Println("Unknown type detected")
		}
	}
	return 0
}
func getTopN(lines []string, perms [][]int, n int) []string {
	topN := make([]string, 0, n)
	for i := 0; i < n; i++ {
		tmp := fmt.Sprintf("#%v struct\n", i+1)
		for j := range perms[i] {
			tmp += lines[perms[i][j]] + "\n"
		}
		topN = append(topN, tmp)
	}
	return topN
}
func get(s string, num int) []string {
	if num == 0 {
		return nil
	}
	input, _ := ioutil.ReadFile(s)
	var start, end int
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if strings.Contains(line, "struct {") {
			start = i + 1
		} else if strings.Contains(line, "}") {
			end = i
			break
		}
	}

	fieldSizes := make([]int, 0, end-start)
	for i := start; i < end; i++ {
		fieldSizes = append(fieldSizes, checkSize(strings.Split(lines[i], " ")[1]))
	}

	perms := getPermutations(end - start)

	structSizes := make([]int, 0, len(perms))
	for i := range perms {
		weight := 0
		for j := range perms[i] {
			weight += (8-weight%8)%fieldSizes[perms[i][j]]%8 + fieldSizes[perms[i][j]]
		}
		weight += (8 - weight%8) % 8
		structSizes = append(structSizes, weight)
	}

	sort.Slice(perms, func(i, j int) bool {
		return structSizes[i] < structSizes[j]
	})

	topN := getTopN(lines[start:end], perms[:num], num)
	topFirstFields := strings.Split(topN[0], "\n")[1:]
	for i := start; i < end; i++ {
		lines[i] = topFirstFields[i-start]
	}

	output := strings.Join(lines, "\n")
	ioutil.WriteFile(s, []byte(output), 0644)

	return topN
}