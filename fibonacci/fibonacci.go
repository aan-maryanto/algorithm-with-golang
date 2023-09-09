package main

func fibonacci(n int) int64 {
	if n <= 1 {
		return int64(n)
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	n := 1
	for n < 100 {
		var hasil = fibonacci(n)
		println(hasil)
		n++
	}
}
