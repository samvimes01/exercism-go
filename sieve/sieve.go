package sieve

import "math"

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
			prime, ok := <-ch
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

func gorout(limit int) []int {
	primes := sieve(limit)
	var result []int
	for prime := range primes {
		result = append(result, prime)
	}
	return result
}

func boolArr(limit int) []int {
	composite := make([]bool, limit+1)
	primes := make([]int, limit/2)
	primeIndex := 0

	for number := 2; number < limit; number++ {
		if !composite[number] {
			primes[primeIndex] = number
			primeIndex++
			for idx := number + number; idx <= limit; idx += number {
				composite[idx] = true
			}
		}
	}
	return primes[:primeIndex]
}
func Sieve(limit int) []int {
	// return gorout(limit)
	// return boolArr(limit)
	return wiki(limit)
}
func wiki(n int) []int {
	// Step 1: Create a boolean array "prime[0..n]" and initialize
	// all entries as true. A value in prime[i] will finally be false if i is Not a prime, otherwise true.
	prime := make([]bool, n+1)
	for i := range prime {
		prime[i] = true
	}

	// Step 2: Change prime[0] and prime[1] to false as 0 and 1 are not prime numbers
	if n >= 0 {
		prime[0] = false
	}
	if n >= 1 {
		prime[1] = false
	}

	// Step 3: Iterate over each number starting from 2 to sqrt(n)
	for i := 2; float64(i) <= math.Sqrt(float64(n)); i++ {
		// If prime[i] is not changed, then it is a prime
		if prime[i] {
			// Checking for prime numbers by marking their multiples as false
			for j := i * i; j <= n; j += i {
				prime[j] = false
			}
		}
	}

	// Step 4: Collect all prime numbers
	var primes []int
	for i := 2; i <= n; i++ {
		if prime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}