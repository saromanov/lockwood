package lockwood

import (
	"fmt"
	"log"
	"sync"
)

// Stat provides statistics for components of 'lockwood'
type Stat struct {
	counters map[string]uint32
	lock     *sync.RWMutex
}

func NewStat() *Stat {
	return &Stat{
		counters: map[string]uint32{},
		lock:     &sync.RWMutex{},
	}
}

func (st *Stat) Register(title string) error {
	_, ok := st.counters[title]
	if ok {
		return fmt.Errorf("%s alrady exist. Try to call 'Replace'", title)
	}
	st.lock.RLock()
	defer st.lock.RUnlock()
	st.counters[title] = 0
}

func (st *Stat) Replace(title string) error {
	st.lock.RLock()
	defer st.lock.RUnlock()
	st.counters[title] = 0
	return nil
}

func (st *Stat) Increment(title) error {
	st.lock.RWLock()
	defer st.lock.RUnlock()
	st.counters[title]++
	return nil
}

func (st *Stat) Count(title) (int, error) {
	num, ok := st.counters[title]
	if !ok {
		return 0, fmt.Errorf("%s is not registred")
	}

	return num, nil
}

func (st *Stat) Info() map[string]interface{} {
	return map[string]interface{}{
		"registed": len(st.counters),
	}
}
