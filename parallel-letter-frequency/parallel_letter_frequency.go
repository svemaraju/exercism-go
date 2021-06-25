package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// newFrequency is run as a goroutine for concurrently calculating the frequency.
func newFrequency(s string, c chan map[rune]int) {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	c <- m
}

// ConcurrentFrequency concurrently counts the frequency of each rune in a give array of strings
func ConcurrentFrequency(inputs []string) FreqMap {
	var c chan map[rune]int = make(chan map[rune]int)
	totalCounts := FreqMap{}
	// Trigger the goroutines.
	for _, s := range inputs {
		go newFrequency(s, c)
	}
	// read the channel for outputs, and merge them into totalCounts
	for i := 0; i < len(inputs); i++ {
		counts := <-c
		for k, v := range counts {
			totalCounts[k] += v
		}
	}
	return totalCounts
}
