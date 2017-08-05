package filter

import "sync"

var matchRoute map[string]bool = make(map[string]bool)
var lock sync.RWMutex

func AddMatchRoute(url string, flag bool) {
	lock.Lock()
	matchRoute[url] = flag
	lock.Unlock()
}

func FilterMatchRoute(url string) bool {
	lock.RLock()
	defer lock.RUnlock()
	return matchRoute[url]
}
