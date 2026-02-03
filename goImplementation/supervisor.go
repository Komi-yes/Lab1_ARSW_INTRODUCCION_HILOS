package main

// Supervisor represents a goroutine that checks a section of blacklist servers
// Author: Daniel palacios Moreno & Sofia Nicolle Ariza Goenaga
type Supervisor struct {
	host        string
	section     int
	sectionSize int
	validator   *HostBlackListsValidator
}

// NewSupervisor creates a new Supervisor instance
func NewSupervisor(host string, section int, sectionSize int, validator *HostBlackListsValidator) *Supervisor {
	return &Supervisor{
		host:        host,
		section:     section,
		sectionSize: sectionSize,
		validator:   validator,
	}
}

// Run executes the supervisor's task of checking blacklist servers
func (s *Supervisor) Run() {
	defer s.validator.NotifyThreadCompletion()

	skds := GetHostBlacklistsDataSourceFacadeInstance()

	start := s.section * s.sectionSize
	end := (s.section + 1) * s.sectionSize

	for i := start; i < end; i++ {
		if !s.validator.ShouldContinueSearching() {
			break
		}

		s.validator.IncrementCheckedCount()
		if skds.IsInBlackListServer(i, s.host) {
			s.validator.AddBlackListOccurrence(i)
		}
	}
}
