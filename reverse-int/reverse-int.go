package main

func reverseInt(n int) int {
	var reverse int
	for n != 0 {
		remain := n % 10
		reverse = reverse*10 + remain
		n = n / 10
	}
	return reverse
}

func main() {
	println(reverseInt(123))
}
