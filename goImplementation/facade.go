package main

import (
	"fmt"
	"sync"
	"time"
)

// Tuple represents a pair of values
type Tuple struct {
	First  int
	Second string
}

// HostBlacklistsDataSourceFacade is a thread-safe class for blacklist data
// Author: hcadavid
type HostBlacklistsDataSourceFacade struct {
	blistOccurrences sync.Map
	threadHits       sync.Map
	lastConfig       string
	lastIndex        int
	mutex            sync.Mutex
}

var (
	instance *HostBlacklistsDataSourceFacade
	once     sync.Once
)

// GetHostBlacklistsDataSourceFacadeInstance returns the singleton instance
func GetHostBlacklistsDataSourceFacadeInstance() *HostBlacklistsDataSourceFacade {
	once.Do(func() {
		instance = &HostBlacklistsDataSourceFacade{
			lastConfig: "",
			lastIndex:  0,
		}
		instance.initializeBlacklists()
	})
	return instance
}

// initializeBlacklists initializes the blacklist occurrences
func (h *HostBlacklistsDataSourceFacade) initializeBlacklists() {
	// to be found by a single thread
	h.blistOccurrences.Store(Tuple{23, "200.24.34.55"}, struct{}{})
	h.blistOccurrences.Store(Tuple{50, "200.24.34.55"}, struct{}{})
	h.blistOccurrences.Store(Tuple{200, "200.24.34.55"}, struct{}{})
	h.blistOccurrences.Store(Tuple{1000, "200.24.34.55"}, struct{}{})
	h.blistOccurrences.Store(Tuple{500, "200.24.34.55"}, struct{}{})

	// to be found through all threads
	h.blistOccurrences.Store(Tuple{29, "202.24.34.55"}, struct{}{})
	h.blistOccurrences.Store(Tuple{10034, "202.24.34.55"}, struct{}{})
	h.blistOccurrences.Store(Tuple{20200, "202.24.34.55"}, struct{}{})
	h.blistOccurrences.Store(Tuple{31000, "202.24.34.55"}, struct{}{})
	h.blistOccurrences.Store(Tuple{70500, "202.24.34.55"}, struct{}{})

	// to be found through all threads
	h.blistOccurrences.Store(Tuple{39, "202.24.34.54"}, struct{}{})
	h.blistOccurrences.Store(Tuple{10134, "202.24.34.54"}, struct{}{})
	h.blistOccurrences.Store(Tuple{20300, "202.24.34.54"}, struct{}{})
	h.blistOccurrences.Store(Tuple{70210, "202.24.34.54"}, struct{}{})
}

// GetRegisteredServersCount returns the total number of registered servers
func (h *HostBlacklistsDataSourceFacade) GetRegisteredServersCount() int {
	return 80000
}

// IsInBlackListServer checks if an IP is in a specific blacklist server
func (h *HostBlacklistsDataSourceFacade) IsInBlackListServer(serverNumber int, ip string) bool {
	// Track thread hits (for debugging/monitoring purposes)
	threadName := getCurrentGoroutineID()
	
	if val, ok := h.threadHits.Load(threadName); ok {
		h.threadHits.Store(threadName, val.(int)+1)
	} else {
		h.threadHits.Store(threadName, 1)
	}

	// Simulate small delay
	time.Sleep(time.Nanosecond)

	// Check if the tuple exists in the blacklist
	_, exists := h.blistOccurrences.Load(Tuple{serverNumber, ip})
	return exists
}

// ReportAsNotTrustworthy reports a host as not trustworthy
func (h *HostBlacklistsDataSourceFacade) ReportAsNotTrustworthy(host string) {
	fmt.Printf("HOST %s Reported as NOT trustworthy\n", host)
}

// ReportAsTrustworthy reports a host as trustworthy
func (h *HostBlacklistsDataSourceFacade) ReportAsTrustworthy(host string) {
	fmt.Printf("HOST %s Reported as trustworthy\n", host)
}

// getCurrentGoroutineID returns a string identifier for the current goroutine
// This is a simplified version - in production you might want to use runtime.GoID
func getCurrentGoroutineID() string {
	// For simplicity, we'll use a timestamp-based ID
	// In a real implementation, you could use runtime information
	return fmt.Sprintf("goroutine-%d", time.Now().UnixNano())
}
