//Package letter, Count the frequency of letters in texts using parallel computation.
package letter

// goroutine with channel
func ConcurrentFrequency(slice []string) FreqMap {
	freqChan := make(chan FreqMap, len(slice))

	for _, s := range slice {
		a := s
		go func() {
			f := Frequency(a)
			freqChan <- f
		}()
	}
	freq := FreqMap{}

	for range slice {
		for k, v := range <-freqChan {
			freq[k] += v
		}
	}

	return freq
}
