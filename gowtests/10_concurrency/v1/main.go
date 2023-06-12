package main

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func() {
			// this will cause issue as the function has different url in the loop and the value is not copied
			results[url] = wc(url)
		}()
	}
	return results
}
