package main

import "fmt"

func main() {
	res := answer()
	fmt.Println(res)
}

func answer() []int {
	res := []int{}
	for i := 10; i <= 1000; i++ {
		if isPalindrome(i) {
			res = append(res, i)
		}
	}
	return res
}

func isPalindrome(n int) bool {
	temp := n
	res := 0
	for temp > 0 {
		res = res*10 + temp%10
		temp /= 10
	}
	return map[bool]bool{true: true, false: false}[n == res]
}

/*
package main

import "fmt"

func main() {
	for i := 10; i <= 1000; i++ {
		fmt.Println(palindrome(i))
	}
}

func palindrome(i int) int {
	n := i
	z := n
	temp := 0
	res := 0
	result := 0
	for n > 0 {
		temp = temp % 10
		res = res*10 + temp
		i /= 10
	}
	if z == res {
		result = res
	}
	return result
}

/*
package main

import "fmt"

func main() {
	var m, n int
	fmt.Println("enter the first number:")
	fmt.Scanf("%d", &m)
	fmt.Println("enter the second number:")
	fmt.Scanf("%d", &n)
	fmt.Println(palindrome(m, n))
}
func palindrome(m int, n int) []int {
	result := []int{}
	for i := m; i <= n; i++ {
		temp := i
		rev := 0
		for a := i; a > 0; a = a / 10 {
			rev = rev*10 + a%10
		}
		if rev == temp {
			result = append(result, temp)
		}
	}
	return result

}
*/
