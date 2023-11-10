package contacts

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Archiver struct {
	mu              sync.Mutex
	archiveStatus   string
	archiveProgress float64
}

var (
	instance *Archiver
	once     sync.Once
)

func GetArchiver() *Archiver {
	once.Do(func() {
		instance = &Archiver{archiveStatus: "Waiting"}
	})
	return instance
}

func (a *Archiver) Status() string {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.archiveStatus
}

func (a *Archiver) Progress() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.archiveProgress
}

func (a *Archiver) ProgressPercent() int {
	return int(a.Progress() * 100)
}

func (a *Archiver) Run() {
	a.mu.Lock()
	if a.archiveStatus == "Waiting" {
		a.archiveStatus = "Running"
		a.archiveProgress = 0
		go a.runImpl()
	}
	a.mu.Unlock()
}

func (a *Archiver) runImpl() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Float64() * float64(time.Second)))
		a.mu.Lock()
		if a.archiveStatus != "Running" {
			a.mu.Unlock()
			return
		}
		a.archiveProgress = float64(i+1) / 10
		fmt.Printf("Here... %.2f\n", a.archiveProgress)
		a.mu.Unlock()
	}
	time.Sleep(time.Second)
	a.mu.Lock()
	if a.archiveStatus == "Running" {
		a.archiveStatus = "Complete"
	}
	a.mu.Unlock()
}

func (a *Archiver) ArchiveFile() string {
	return "contacts.json"
}

func (a *Archiver) Reset() {
	a.mu.Lock()
	a.archiveStatus = "Waiting"
	a.mu.Unlock()
}
