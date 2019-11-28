package structs

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

// JList loads a list (map string->bool) from a json file,
// and can test a string's membership.
type JList struct {
	List map[string]bool
	sync.RWMutex
}

func (b *JList) Load(path string) error {
	var newList map[string]bool

	bFd, err := os.Open(path)
	if err != nil {
		return err
	}
	defer bFd.Close()

	bBytes, err := ioutil.ReadAll(bFd)
	if err != nil {
		return err
	}

	b.Lock()
	defer b.Unlock()

	err = json.Unmarshal([]byte(bBytes), &newList)
	if err != nil {
		return err
	}

	b.List = newList
	return nil
}

func (b *JList) Has(member string) bool {
	b.RLock()
	defer b.RUnlock()
	return b.List[member]
}

func (b *JList) Len() int {
	b.RLock()
	defer b.RUnlock()
	return len(b.List)
}
