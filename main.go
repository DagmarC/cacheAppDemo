package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}

var randomNumber = rand.New(rand.NewSource(time.Now().UnixNano()))

const numberOfBooks int = 7

func main() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	cacheChannel := make(chan Book)
	dbChannel := make(chan Book)

	for i := 0; i < numberOfBooks; i++ {
		id := randomNumber.Intn(numberOfBooks) + 1 // Ids are from 1 to 7
		wg.Add(2)
		go func(id int, channel chan<- Book, wg *sync.WaitGroup, m *sync.RWMutex) {
			if book, ok := queryCache(id, m); ok {
				channel <- book
			}
			wg.Done()
		}(id, cacheChannel, wg, m)

		go func(id int, channel chan<- Book, wg *sync.WaitGroup, m *sync.RWMutex) {
			if book, ok := queryDatabase(id); ok {
				m.Lock()
				cache[id] = book
				m.Unlock()
				channel <- book
			}
			wg.Done()
		}(id, dbChannel, wg, m)

		select {
		case book := <-dbChannel:
			fmt.Print("DB:", book.ID, "\t", book.String())
		case book := <-cacheChannel:
			fmt.Println("CACHE:", book.ID, "\t", book.String())
			<-dbChannel
		}
		go func(cacheChannel, dbChannel <-chan Book) {
		}(cacheChannel, dbChannel)
	}
	wg.Wait()
	close(dbChannel)
	close(cacheChannel)
	fmt.Println("Main finished.")
}

func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	m.RLock()
	book, ok := cache[id] // Read.
	m.RUnlock()
	return book, ok
}

func queryDatabase(id int) (Book, bool) {
	for _, book := range books {
		if book.ID == id {
			return book, true
		}
	}
	return Book{}, false
}
