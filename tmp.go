package main

//
//import (
//	"fmt"
//	"math/rand"
//	"sync"
//	"time"
//)
//
//// Build a simulation for a db query sytem
//// with in-memory cahce in front of it.
//
//// Storing entries by id.
//var cache = map[int]Book{}
//
//var randomNumber = rand.New(rand.NewSource(time.Now().UnixNano()))
//
//const numberOfBooks int = 7
//
//func main() {
//	// ChanelExample()
//
//	// If cache has result - return it.
//	// If DB has result - return it and update cache.
//
//	// WG. Can be address or normal.
//	// Adress - no need to copy wg when used on multiple places
//	waitGroup := &sync.WaitGroup{}
//
//	// Adresses a Shared memory problem r/w. Only 1 task acess at a time.
//	// Reading - normal mutex is too restricted cuz just multiple reading
//	// is not a bad thing - use RWMutex instead.
//	// mutex := &sync.Mutex{}
//	mutex := &sync.RWMutex{}
//
//	// Got two sending channels - db and cache.
//	// And one receiving decidor.
//
//	// One bi channel let us know that I received value from cache.
//	cacheChannel := make(chan Book)
//	// Second channel let us know I received value from db.
//	dbChannel := make(chan Book)
//
//	for i := 0; i < numberOfBooks; i++ {
//		id := randomNumber.Intn(numberOfBooks) + 1 // Ids are from 1 to 7
//		// Every loop call add method - number of tasks I am waiting on.
//		// Add 1 before each go routine or add 2 at once.
//		fmt.Println("\nID: **********", id)
//		waitGroup.Add(2)
//
//		// Anonymous function
//		go func(id int, waitGroup *sync.WaitGroup, mutex *sync.RWMutex, channel chan<- Book) {
//			// LOOK FOR SYNTAX:
//			if book, ok := queryCache(id, mutex); ok {
//				channel <- book
//			}
//			// 1 concurrent task is being completed.
//			waitGroup.Done()
//		}(id, waitGroup, mutex, cacheChannel)
//
//		go func(id int, waitGroup *sync.WaitGroup, mutex *sync.RWMutex, channel chan<- Book) {
//			if book, ok := queryDatabase(id); ok {
//				fmt.Println("Adding to cache...")
//				// mutex.Lock()
//				// cache[id] = book
//				// mutex.Unlock()
//				channel <- book
//			}
//			waitGroup.Done()
//		}(id, waitGroup, mutex, dbChannel)
//
//		// Create Receiving channel - inside the loop `cuz number of receivers and senders must match.
//		// Select - sth like switch but it has cases channels that waits for input.
//		go func(cacheChannel, dbChannel <-chan Book) {
//			select {
//			case book := <-cacheChannel:
//				fmt.Println("Cache:")
//				fmt.Println(book.String())
//				// Wait until I get msg from db channel to nt to block db.
//				<-dbChannel
//			case book := <-dbChannel:
//				fmt.Println("DB:")
//				fmt.Println(book.String())
//			}
//		}(cacheChannel, dbChannel)
//	}
//	// Wait till our waitGroup is no longer waiting on any concurrent activities (2 activiies)
//	waitGroup.Wait()
//}
//
//// Multiple goroutines are trying to access a shared memory:
//// 1 is writing to chache[id] and the other one is reading from cache[id] ->
//// Problem with shared memory - solition is to use Mutaxes. Lock and RLock().
//
//func queryCache(id int, mutex *sync.RWMutex) (Book, bool) {
//	// RLock -- allows multiple readers to acquire a read lock.
//	// When Writing - all readers will finish their operations and write can be executed.
//	// mutex.RLock()         // Whoever calls it, controls the mutex - owns a read lock.
//	book, ok := cache[id] // Read.
//	// mutex.RUnlock()
//	return book, ok
//}
//
//func queryDatabase(id int) (Book, bool) {
//	time.Sleep(100 * time.Millisecond)
//	for _, book := range books {
//		if book.ID == id {
//			cache[id] = book
//			return book, true
//		}
//	}
//	return Book{}, false
//}
