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
package text

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/rangertaha/urlinsane"
	"github.com/rangertaha/urlinsane/plugins/outputs"
	"golang.org/x/term"
)

const CODE = "text"

type Text struct {
	table  table.Writer
	config urlinsane.Config
}

func (n *Text) Id() string {
	return CODE
}

func (n *Text) Description() string {
	return "Text outputs one record per line and is the default"
}

func (n *Text) Init(conf urlinsane.Config) {
	n.config = conf
	width, _, err := term.GetSize(int(os.Stdout.Fd()))

	n.table = table.NewWriter()
	n.table.SetAllowedRowLength(width)
	n.table.SetOutputMirror(os.Stdout)
	n.table.AppendHeader(n.getHeader())

	n.activeRow()

	// width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

}

func (n *Text) getHeader() (row table.Row) {
	row = append(row, "ID")
	row = append(row, "TYPE")
	row = append(row, "TYPO")

	for _, info := range n.config.Information() {
		for _, headers := range info.Headers() {
			row = append(row, headers)
		}
	}

	return
}

func (n *Text) getRow(typo urlinsane.Typo) (row table.Row) {
	row = append(row, typo.Id())
	if n.config.Verbose() {
		row = append(row, typo.Algorithm().Name())
	} else {
		row = append(row, strings.ToUpper(typo.Algorithm().Id()))
	}

	row = append(row, typo.Variant().Repr())

	for _, info := range n.config.Information() {
		for _, header := range info.Headers() {
			meta := typo.Variant().Meta()
			row = append(row, meta[header])
		}
	}

	return
}

func (n *Text) activeRow() (row table.Row) {
	nameTransformer := text.Transformer(func(val interface{}) string {
		return text.Bold.Sprint(val)
	})

	// n.table.SetRowPainter()

	n.table.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:              "TYPE",
			Align:             text.AlignLeft,
			AlignFooter:       text.AlignLeft,
			AlignHeader:       text.AlignLeft,
			Colors:            text.Colors{text.BgBlack, text.FgRed},
			ColorsHeader:      text.Colors{text.BgRed, text.FgBlack, text.Bold},
			ColorsFooter:      text.Colors{text.BgRed, text.FgBlack},
			Hidden:            false,
			Transformer:       nameTransformer,
			TransformerFooter: nameTransformer,
			TransformerHeader: nameTransformer,
			VAlign:            text.VAlignMiddle,
			VAlignFooter:      text.VAlignTop,
			VAlignHeader:      text.VAlignBottom,
			WidthMin:          6,
			WidthMax:          64,
		},
	})
	return
}

func (n *Text) Write(in urlinsane.Typo) {
	n.table.AppendRow(n.getRow(in))
}

func (n *Text) Save() {
	n.table.AppendFooter(table.Row{"Total", n.config.Count()})
	n.table.SetStyle(StyleDefault)
	output := n.table.Render()

	if n.config.File() != "" {
		results := []byte(output)
		if err := os.WriteFile(n.config.File(), results, 0644); err != nil {
			fmt.Printf("Error: %s", err)
		}
	}

}

// Register the plugin
func init() {
	outputs.Add(CODE, func() urlinsane.Output {
		return &Text{}
	})
}
