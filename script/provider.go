package script

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

// Provider loads SQL scripts and provides all SQL statements as commands.
type Provider struct {
	mu         sync.RWMutex
	repository map[string]string
}

// LoadDir loads all script commands from a given directory. Note that all
// scripts should have .sql extension.
func (p *Provider) ReadDir(fs FileSystem) error {
	return fs.Walk("/", func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return fmt.Errorf("Directory does not exist")
		}

		if info.IsDir() {
			return nil
		}

		matched, err := filepath.Match("*.sql", info.Name())
		if err != nil || !matched {
			return err
		}

		file, err := fs.OpenFile(path, os.O_RDONLY, 0)
		if err != nil {
			return err
		}

		defer func() {
			if ioErr := file.Close(); err == nil {
				err = ioErr
			}
		}()

		if err = p.ReadFrom(file); err != nil {
			return err
		}

		return err
	})
}

// ReadFrom reads the script from a reader
func (p *Provider) ReadFrom(r io.Reader) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.repository == nil {
		p.repository = make(map[string]string)
	}

	scanner := &Scanner{}
	stmts := scanner.Scan(r)

	for name, stmt := range stmts {
		if _, ok := p.repository[name]; ok {
			return fmt.Errorf("Command '%s' already exists", name)
		}

		p.repository[name] = stmt
	}

	return nil
}

// Command returns a command for given name and parameters. The operation can
// err if the command cannot be found.
func (p *Provider) Command(name string, params ...Param) (*Cmd, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if query, ok := p.repository[name]; ok {
		return &Cmd{
			query:  query,
			params: params,
		}, nil
	}

	return nil, fmt.Errorf("Command '%s' not found", name)
}
