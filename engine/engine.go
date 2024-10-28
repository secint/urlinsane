// Copyright (C) 2024 Rangertaha
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
package engine

import (
	"sync"
	"time"

	"github.com/rangertaha/urlinsane"
	"github.com/rangertaha/urlinsane/config"
	"github.com/schollz/progressbar/v3"
)

type (

	// DomainTypos ...
	DomainTypos struct {
		Config config.Config
		Typos  map[string]urlinsane.Typo
		Count  int64

		algoWG sync.WaitGroup
		infoWG sync.WaitGroup

		progress *progressbar.ProgressBar
		// stats <-chan Statser
		// errs <-chan interface{}
	}
)

// NewDomainTypos ...
func NewDomainTypos(conf config.Config) DomainTypos {
	return DomainTypos{
		Config: conf,
	}
}

// GenOptions typo config options
func (t *DomainTypos) GenOptions() <-chan urlinsane.Typo {
	out := make(chan urlinsane.Typo)
	go func() {
		for _, lang := range t.Config.Languages() {
			for _, board := range t.Config.Keyboards() {
				for _, algo := range t.Config.Algorithms() {
					domain := NewDomain(t.Config.Target())
					out <- &Typo{
						language:  lang,
						keyboard:  board,
						algorithm: algo,
						original:  domain,
					}
				}
			}
		}
		close(out)
	}()
	return out
}

// Algorithms generate typos using the algorithm plugins
func (ts *DomainTypos) Algorithms(in <-chan urlinsane.Typo) <-chan urlinsane.Typo {
	out := make(chan urlinsane.Typo)

	for w := 1; w <= ts.Config.Concurrency(); w++ {
		ts.algoWG.Add(1)
		go func(id int, in <-chan urlinsane.Typo, out chan<- urlinsane.Typo) {
			defer ts.algoWG.Done()
			for typo := range in {
				// Execute typo algorithm returning typos
				for _, typ := range typo.Algorithm().Exec(typo) {
					if typ.Variant() != nil {
						out <- typ
					}
				}
			}
		}(w, in, out)
	}
	go func() {
		ts.algoWG.Wait()
		close(out)
	}()

	return ts.Dedup(out)
}

func (ts *DomainTypos) Dedup(in <-chan urlinsane.Typo) <-chan urlinsane.Typo {
	out := make(chan urlinsane.Typo)
	var typos = make(map[string]urlinsane.Typo)

	go func() {

		// Create map of unique domains
		for typo := range in {
			typos[typo.Variant().Repr()] = typo
		}
		var count int64 = 0
		for _, typ := range typos {
			count++
			typ.Id(count) // Set typo record number
		}
		// Save the total count in the config for output plugins to use
		ts.Config.Count(count)

		// Return all typos via channels
		for _, typ := range typos {
			out <- typ
		}
		close(out)
	}()

	return out
}

func (t *DomainTypos) Information(in <-chan urlinsane.Typo) <-chan urlinsane.Typo {
	out := make(chan urlinsane.Typo)
	for w := 1; w <= t.Config.Concurrency(); w++ {
		t.infoWG.Add(1)
		go func(in <-chan urlinsane.Typo, out chan<- urlinsane.Typo) {
			defer t.infoWG.Done()
			output := t.InfoChain(t.Config.Information(), in)
			for c := range output {
				out <- c
			}
		}(in, out)
	}
	go func() {
		t.infoWG.Wait()
		close(out)
	}()
	return out
}

// InfoChain creates a chain of information-gathering functions
func (t *DomainTypos) InfoChain(funcs []urlinsane.Information, in <-chan urlinsane.Typo) <-chan urlinsane.Typo {
	var xfunc urlinsane.Information
	out := make(chan urlinsane.Typo)
	xfunc, funcs = funcs[len(funcs)-1], funcs[:len(funcs)-1]
	go func() {
		for i := range in {
			time.Sleep(t.Config.Random() * t.Config.Delay())
			out <- xfunc.Exec(i)
		}
		close(out)
	}()

	if len(funcs) > 0 {
		return t.InfoChain(funcs, out)
	}
	return out
}

func (t *DomainTypos) Storage(in <-chan urlinsane.Typo) <-chan urlinsane.Typo {


	return in
}

func (t *DomainTypos) Progress(in <-chan urlinsane.Typo) <-chan urlinsane.Typo {
	if t.Config.Progress() {
		out := make(chan urlinsane.Typo)
		go func(in <-chan urlinsane.Typo, out chan<- urlinsane.Typo) {
			for c := range in {
				if t.Config.Count() != 0 && t.progress == nil {
					t.progress = progressbar.Default(t.Config.Count())
				}
				out <- c
				t.progress.Add(1)
			}
			close(out)
		}(in, out)
		return out
	}
	return in
}

func (t *DomainTypos) Filter(in <-chan urlinsane.Typo) <-chan urlinsane.Typo {


	return in
}

func (t *DomainTypos) Output(in <-chan urlinsane.Typo) {
	t.Config.Output().Init(&t.Config)
	for c := range in {
		// Write output
		if c != nil {
			t.Config.Output().Write(c)
		}
	}
	t.Config.Output().Save()
}

func (t *DomainTypos) Execute() {
	typos := t.GenOptions()
	typos = t.Algorithms(typos)
	typos = t.Information(typos)
	typos = t.Storage(typos)
	typos = t.Progress(typos)
	typos = t.Filter(typos)
	t.Output(typos)
}
