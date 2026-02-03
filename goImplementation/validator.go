package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const (
	// BlackListAlarmCount is the threshold for considering a host as not trustworthy
	BlackListAlarmCount = 5
)

// HostBlackListsValidator validates hosts against blacklist servers
// Author: Daniel palacios Moreno & Sofia Nicolle Ariza Goenaga
type HostBlackListsValidator struct {
	blackListOccurrences []int
	occurrencesCount     atomic.Int32
	checkedListsCount    atomic.Int32
	stopFlag             atomic.Bool
	completionWg         sync.WaitGroup
	mutex                sync.Mutex
}

// NewHostBlackListsValidator creates a new HostBlackListsValidator instance
func NewHostBlackListsValidator() *HostBlackListsValidator {
	return &HostBlackListsValidator{
		blackListOccurrences: make([]int, 0),
	}
}

// CheckHost checks the given host's IP address in all available black lists
// and reports it as NOT Trustworthy when such IP was reported in at least
// BLACK_LIST_ALARM_COUNT lists, or as Trustworthy in any other case.
// The search is not exhaustive: When the number of occurrences is equal to
// BLACK_LIST_ALARM_COUNT, the search is finished, the host reported as
// NOT Trustworthy, and the list of the five blacklists returned.
// Parameters:
//   ipaddress - suspicious host's IP address
//   n - number of threads/goroutines to use
// Returns: Blacklists numbers where the given host's IP address was found
func (h *HostBlackListsValidator) CheckHost(ipaddress string, n int) []int {
	// Reset state for new check
	h.blackListOccurrences = make([]int, 0)
	h.occurrencesCount.Store(0)
	h.checkedListsCount.Store(0)
	h.stopFlag.Store(false)

	skds := GetHostBlacklistsDataSourceFacadeInstance()
	threadSectionSize := skds.GetRegisteredServersCount() / n

	// Start n goroutines
	h.completionWg.Add(n)
	for i := 0; i < n; i++ {
		supervisor := NewSupervisor(ipaddress, i, threadSectionSize, h)
		go func() {
			supervisor.Run()
		}()
	}

	// Wait for all goroutines to complete
	h.completionWg.Wait()

	// Report results
	if h.occurrencesCount.Load() >= BlackListAlarmCount {
		skds.ReportAsNotTrustworthy(ipaddress)
	} else {
		skds.ReportAsTrustworthy(ipaddress)
	}

	fmt.Printf("Checked Black Lists: %d of %d\n", h.checkedListsCount.Load(), skds.GetRegisteredServersCount())

	return h.blackListOccurrences
}

// AddBlackListOccurrence adds a blacklist occurrence to the list
func (h *HostBlackListsValidator) AddBlackListOccurrence(index int) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	if h.occurrencesCount.Load() < BlackListAlarmCount {
		h.blackListOccurrences = append(h.blackListOccurrences, index)
		count := h.occurrencesCount.Add(1)

		if count >= BlackListAlarmCount {
			h.stopFlag.Store(true)
		}
	}
}

// NotifyThreadCompletion notifies that a thread/goroutine has completed
func (h *HostBlackListsValidator) NotifyThreadCompletion() {
	h.completionWg.Done()
}

// IncrementCheckedCount increments the count of checked blacklists
func (h *HostBlackListsValidator) IncrementCheckedCount() {
	h.checkedListsCount.Add(1)
}

// ShouldContinueSearching returns whether the search should continue
func (h *HostBlackListsValidator) ShouldContinueSearching() bool {
	return !h.stopFlag.Load()
}
