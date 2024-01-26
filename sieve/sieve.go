package sieve


// Send the sequence 2, 3, 4, ... to returned channel
func generate(n int) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; i <= n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

// Filter out input values divisible by prime, send rest to returned channel
func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			ispr := i%prime != 0
			if ispr {
				out <- i
			}
		}
		close(out)
	}()
	return out
}

func sieve(n int) chan int {
	out := make(chan int)
	go func() {
		ch := generate(n)
		for {
			prime, ok := <- ch
			if !ok {
				break
			}
			ch = filter(ch, prime)
			out <- prime
		}
		close(out)
	}()
	return out
}


func Sieve(limit int) []int {
	primes := sieve(limit)
	var result []int
	for prime := range primes {
		result = append(result, prime)
	}
	return result
}
