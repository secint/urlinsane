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
package spanish

import (
	"github.com/rangertaha/urlinsane/internal"
	"github.com/rangertaha/urlinsane/internal/plugins/languages"
)

type Keyboard struct {
	lang        string
	code        string
	name        string
	description string
	layout      []string
}

func (k *Keyboard) Id() string {
	return k.code
}
func (k *Keyboard) Language() string {
	return k.lang
}
func (k *Keyboard) Name() string {
	return k.name
}
func (k *Keyboard) Description() string {
	return k.description
}
func (k *Keyboard) Layouts() []string {
	return k.layout
}
func (k *Keyboard) Adjacent(char string) []string {
	return languages.Adjacent(k.layout, char)
}
func (k *Keyboard) Languages() []internal.Language {
	return languages.Languages(k.lang)
}

var Keyboards = []Keyboard{
	{
		lang:        LANGUAGE,
		code:        "es1",
		name:        "QWERTY",
		description: "Spanish keyboard layout",
		layout: []string{
			"1234567890-",
			"qwertyuiop ",
			"asdfghjklñ ",
			"zxcvbnm  ç ",
		},
	},
	{
		lang:        LANGUAGE,
		code:        "es2",
		name:        "QWERTY",
		description: "Spanish ISO keyboard layout",
		layout: []string{
			"1234567890 ¡",
			"qwertyuiop  ",
			"asdfghjklñ  ",
			"zxcvbnm  -  ",
		},
	},
}

func init() {
	for _, kb := range Keyboards {
		languages.AddKeyboard(kb.code, func() internal.Keyboard {
			return &kb
		})
	}

}
