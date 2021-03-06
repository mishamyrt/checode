package file

import (
	"sync"

	"github.com/mishamyrt/checode/v1/pkg/bit"
	"github.com/mishamyrt/checode/v1/pkg/types"
)

// Parsing results.
type Parsing struct {
	Matches []Matches
	Config  types.Config
	Flags   bit.Map
	mutex   sync.Mutex
}

// Append match to parsing.
func (p *Parsing) Append(file Matches) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.Matches = append(p.Matches, file)
	p.Flags |= file.Flags
}

// Run parsing.
func (p *Parsing) Run(files []string) {
	var wg sync.WaitGroup
	wg.Add(len(files))

	for _, path := range files {
		go func(filePath string) {
			defer wg.Done()
			file := ParseFile(filePath, &p.Config)
			if len(file.Matches) > 0 {
				p.Append(file)
			}
		}(path)
	}
	wg.Wait()
}
