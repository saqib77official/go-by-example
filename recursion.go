package main

import "fmt"

// Basic factorial recursion
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Fibonacci recursion (inefficient but demonstrates recursion)
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Recursive sum of array
func sumArray(arr []int, index int) int {
	if index >= len(arr) {
		return 0
	}
	return arr[index] + sumArray(arr, index+1)
}

// Recursive reverse string
func reverseString(s string) string {
	if len(s) <= 1 {
		return s
	}
	return reverseString(s[1:]) + string(s[0])
}

// Recursive power function
func power(base, exponent int) int {
	if exponent == 0 {
		return 1
	}
	if exponent < 0 {
		return 0 // Not handling negative exponents in this simple example
	}
	return base * power(base, exponent-1)
}

// Recursive binary search
func binarySearch(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}
	
	mid := left + (right-left)/2
	
	if arr[mid] == target {
		return mid
	} else if arr[mid] > target {
		return binarySearch(arr, target, left, mid-1)
	} else {
		return binarySearch(arr, target, mid+1, right)
	}
}

// Recursive GCD (Greatest Common Divisor)
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Recursive palindrome check
func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	if s[0] != s[len(s)-1] {
		return false
	}
	return isPalindrome(s[1 : len(s)-1])
}

// Recursive tree traversal simulation
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) inOrderTraversal() []int {
	if n == nil {
		return []int{}
	}
	
	result := []int{}
	result = append(result, n.Left.inOrderTraversal()...)
	result = append(result, n.Value)
	result = append(result, n.Right.inOrderTraversal()...)
	
	return result
}

// Recursive directory structure simulation
type File struct {
	Name     string
	IsDir    bool
	Children []File
}

func printFileStructure(files []File, indent string) {
	for _, file := range files {
		fmt.Printf("%s%s\n", indent, file.Name)
		if file.IsDir {
			printFileStructure(file.Children, indent+"  ")
		}
	}
}

// Recursive permutations
func permutations(arr []int) [][]int {
	if len(arr) <= 1 {
		return [][]int{arr}
	}
	
	var result [][]int
	
	for i, num := range arr {
		rest := make([]int, 0, len(arr)-1)
		rest = append(rest, arr[:i]...)
		rest = append(rest, arr[i+1:]...)
		
		for _, perm := range permutations(rest) {
			newPerm := append([]int{num}, perm...)
			result = append(result, newPerm)
		}
	}
	
	return result
}

// Recursive combination
func combinations(arr []int, r int) [][]int {
	if r == 0 {
		return [][]int{{}}
	}
	if len(arr) < r {
		return [][]int{}
	}
	if len(arr) == r {
		return [][]int{arr}
	}
	
	// Include first element
	withFirst := combinations(arr[1:], r-1)
	for i := range withFirst {
		withFirst[i] = append([]int{arr[0]}, withFirst[i]...)
	}
	
	// Exclude first element
	withoutFirst := combinations(arr[1:], r)
	
	// Combine both results
	return append(withFirst, withoutFirst...)
}

// Recursive Tower of Hanoi
func towerOfHanoi(n int, source, destination, auxiliary string, moves *[]string) {
	if n == 1 {
		*moves = append(*moves, fmt.Sprintf("Move disk 1 from %s to %s", source, destination))
		return
	}
	
	towerOfHanoi(n-1, source, auxiliary, destination, moves)
	*moves = append(*moves, fmt.Sprintf("Move disk %d from %s to %s", n, source, destination))
	towerOfHanoi(n-1, auxiliary, destination, source, moves)
}

// Recursive depth-first search
type Graph struct {
	Vertices map[int][]int
	Visited  map[int]bool
}

func (g *Graph) DFS(vertex int) []int {
	g.Visited[vertex] = true
	result := []int{vertex}
	
	for _, neighbor := range g.Vertices[vertex] {
		if !g.Visited[neighbor] {
			result = append(result, g.DFS(neighbor)...)
		}
	}
	
	return result
}

// Recursive memoization example
func memoizedFibonacci(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	
	if val, exists := memo[n]; exists {
		return val
	}
	
	memo[n] = memoizedFibonacci(n-1, memo) + memoizedFibonacci(n-2, memo)
	return memo[n]
}

func main() {
	fmt.Println("=== Recursion Examples ===")

	// 1. Factorial
	fmt.Println("\n1. Factorial:")
	fmt.Printf("Factorial of 5: %d\n", factorial(5))
	fmt.Printf("Factorial of 6: %d\n", factorial(6))
	fmt.Printf("Factorial of 0: %d\n", factorial(0))

	// 2. Fibonacci
	fmt.Println("\n2. Fibonacci:")
	fmt.Printf("Fibonacci of 10: %d\n", fibonacci(10))
	fmt.Printf("Fibonacci of 7: %d\n", fibonacci(7))

	// 3. Sum of array
	fmt.Println("\n3. Sum of array:")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Sum of %v: %d\n", numbers, sumArray(numbers, 0))

	// 4. Reverse string
	fmt.Println("\n4. Reverse string:")
	original := "Hello, World!"
	reversed := reverseString(original)
	fmt.Printf("Original: %s\n", original)
	fmt.Printf("Reversed: %s\n", reversed)

	// 5. Power function
	fmt.Println("\n5. Power function:")
	fmt.Printf("2^8 = %d\n", power(2, 8))
	fmt.Printf("3^4 = %d\n", power(3, 4))
	fmt.Printf("5^0 = %d\n", power(5, 0))

	// 6. Binary search
	fmt.Println("\n6. Binary search:")
	sorted := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 7
	index := binarySearch(sorted, target, 0, len(sorted)-1)
	fmt.Printf("Found %d at index %d in %v\n", target, index, sorted)

	target = 8
	index = binarySearch(sorted, target, 0, len(sorted)-1)
	fmt.Printf("Found %d at index %d in %v\n", target, index, sorted)

	// 7. GCD
	fmt.Println("\n7. Greatest Common Divisor:")
	fmt.Printf("GCD of 48 and 18: %d\n", gcd(48, 18))
	fmt.Printf("GCD of 17 and 23: %d\n", gcd(17, 23))

	// 8. Palindrome check
	fmt.Println("\n8. Palindrome check:")
	palindromes := []string{"racecar", "madam", "hello", "level"}
	for _, word := range palindromes {
		fmt.Printf("Is '%s' a palindrome? %t\n", word, isPalindrome(word))
	}

	// 9. Tree traversal
	fmt.Println("\n9. Tree traversal:")
	root := &TreeNode{
		Value: 5,
		Left: &TreeNode{
			Value: 3,
			Left:  &TreeNode{Value: 2},
			Right: &TreeNode{Value: 4},
		},
		Right: &TreeNode{
			Value: 7,
			Left:  &TreeNode{Value: 6},
			Right: &TreeNode{Value: 8},
		},
	}
	fmt.Printf("In-order traversal: %v\n", root.inOrderTraversal())

	// 10. File structure
	fmt.Println("\n10. File structure:")
	fileSystem := File{
		Name:  "root",
		IsDir: true,
		Children: []File{
			{Name: "file1.txt", IsDir: false},
			{Name: "folder1", IsDir: true, Children: []File{
				{Name: "file2.txt", IsDir: false},
				{Name: "file3.txt", IsDir: false},
			}},
			{Name: "folder2", IsDir: true, Children: []File{
				{Name: "subfolder", IsDir: true, Children: []File{
					{Name: "file4.txt", IsDir: false},
				}},
			}},
		},
	}
	fmt.Println("Directory structure:")
	printFileStructure([]File{fileSystem}, "")

	// 11. Permutations
	fmt.Println("\n11. Permutations:")
	arr := []int{1, 2, 3}
	perms := permutations(arr)
	fmt.Printf("Permutations of %v:\n", arr)
	for i, perm := range perms {
		fmt.Printf("  %d: %v\n", i+1, perm)
	}

	// 12. Combinations
	fmt.Println("\n12. Combinations:")
	combs := combinations([]int{1, 2, 3, 4}, 2)
	fmt.Printf("Combinations of [1,2,3,4] taken 2 at a time:\n")
	for i, comb := range combs {
		fmt.Printf("  %d: %v\n", i+1, comb)
	}

	// 13. Tower of Hanoi
	fmt.Println("\n13. Tower of Hanoi:")
	var moves []string
	towerOfHanoi(3, "A", "C", "B", &moves)
	fmt.Printf("Moves for 3 disks:\n")
	for i, move := range moves {
		fmt.Printf("  %d: %s\n", i+1, move)
	}

	// 14. Graph DFS
	fmt.Println("\n14. Graph DFS:")
	graph := Graph{
		Vertices: map[int][]int{
			0: {1, 2},
			1: {2},
			2: {0, 3},
			3: {3},
		},
		Visited: make(map[int]bool),
	}
	fmt.Printf("DFS starting from vertex 2: %v\n", graph.DFS(2))

	// 15. Memoized Fibonacci
	fmt.Println("\n15. Memoized Fibonacci:")
	memo := make(map[int]int)
	fmt.Printf("Memoized Fibonacci of 50: %d\n", memoizedFibonacci(50, memo))
	fmt.Printf("Memoized Fibonacci of 45: %d\n", memoizedFibonacci(45, memo))
}
