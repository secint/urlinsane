// Copyright 2024 Rangertaha. All Rights Reserved.
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
package html

// import (
// 	"fmt"
// 	"os"
// 	"strings"

// 	"github.com/jedib0t/go-pretty/v6/table"
// 	"github.com/rangertaha/urlinsane/internal"
// 	"github.com/rangertaha/urlinsane/internal/plugins/outputs"
// 	log "github.com/sirupsen/logrus"
// )

// const (
// 	CODE        = "html"
// 	DESCRIPTION = "HTML formatted output"
// )

// type Plugin struct {
// 	table  table.Writer
// 	config internal.Config
// 	output string
// }

// func (p *Plugin) Id() string {
// 	return CODE
// }

// func (p *Plugin) Description() string {
// 	return DESCRIPTION
// }

// func (p *Plugin) Init(conf internal.Config) {
// 	n.config = conf
// 	n.table = table.NewWriter()
// 	n.table.SetOutputMirror(os.Stdout)
// 	n.table.AppendHeader(n.Header())
// }

// func (p *Plugin) Read(in internal.Domain) {
// 	n.table.AppendRow(n.Row(in))
// }

// func (p *Plugin) Header() (row table.Row) {
// 	row = append(row, "LD")
// 	row = append(row, "TYPE")
// 	row = append(row, "TYPO")

// 	for _, info := range n.config.Collectors() {
// 		for _, header := range info.Headers() {
// 			if n.Filter(header) {
// 				row = append(row, header)
// 			}
// 		}
// 	}
// 	return
// }

// func (p *Plugin) Row(typo internal.Domain) (row table.Row) {
// 	row = append(row, typo.Ld())
// 	if n.config.Verbose() {
// 		row = append(row, typo.Algorithm().Name())
// 	} else {
// 		row = append(row, strings.ToUpper(typo.Algorithm().Id()))
// 	}
// 	row = append(row, typo.String())

// 	for _, info := range n.config.Collectors() {
// 		for _, header := range info.Headers() {
// 			if n.Filter(header) {
// 				meta := typo.Meta()
// 				if col, ok := meta[header]; ok {
// 					row = append(row, col)
// 				} else {
// 					row = append(row, "")
// 				}
// 			}
// 		}
// 	}
// 	return
// }

// func (p *Plugin) Filter(header string) bool {
// 	header = strings.TrimSpace(header)
// 	header = strings.ToLower(header)
// 	for _, filter := range n.config.Filters() {
// 		filter = strings.TrimSpace(filter)
// 		filter = strings.ToLower(filter)
// 		if filter == header {
// 			return true
// 		}
// 	}
// 	return false
// }

// // func (p *Plugin) Progress(typo <-chan internal.Domain) <-chan internal.Domain {
// // 	return typo
// // }

// func (p *Plugin) Write() {
// 	n.output = n.table.RenderHTML()
// }

// func (p *Plugin) Summary(report map[string]string) {
// 	fmt.Println("")
// 	for k, v := range report {
// 		fmt.Printf("%s %s   ", k, v)
// 	}
// 	fmt.Println("")
// }

// func (p *Plugin) Save(fname string) {
// 	results := []byte(n.output)
// 	if err := os.WriteFile(fname, results, 0644); err != nil {
// 		log.Errorf("Error: %s", err)
// 	}
// }

// // Register the plugin
// func init() {
// 	outputs.Add(CODE, func() internal.Output {
// 		return &Plugin{}
// 	})
// }
