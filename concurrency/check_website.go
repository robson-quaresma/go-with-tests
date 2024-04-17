package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		// New: working with go routines
		go func() {
			results[url] = wc(url)
		}()
	}

	return results
}
