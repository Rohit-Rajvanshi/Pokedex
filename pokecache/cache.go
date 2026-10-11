package pokecache 

import (
	"time"
	"sync"	
)

type cacheEntry struct {
	createdAt time.Time
	val []byte 
}

type Cache struct {
	entry map[string]cacheEntry
	mu *sync.Mutex
}

func NewCache(interval time.Duration) Cache{
	
	c := Cache{
		entry : map[string]cacheEntry{} ,
		mu : &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c 
}

func (c *Cache) Add(key string , val []byte) {
	c.mu.Lock()
	c.entry[key] = cacheEntry{
				createdAt : time.Now(),
				val : val ,
			}

			
	c.mu.Unlock()
}

func (c *Cache) Get(key string)  ([]byte , bool){
	c.mu.Lock()
	defer c.mu.Unlock()
	exists := true 
	elem , ok := c.entry[key] 
	if !ok {
		exists = false 
	}
	return elem.val , exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C{
		c.mu.Lock()
		for key , val := range c.entry {
			if time.Since(val.createdAt) > interval {
				delete(c.entry , key)
			}
		}
		c.mu.Unlock()
		
	}
	ticker.Stop()
}

