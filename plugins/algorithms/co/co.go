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
package co

// Character Omission
// Created by leaving out a character in the name.
//
//For example:
//
// Input: google.com
//
// Output:
// ID     TYPE    TYPO      
// --------------------------
// 1      CO      gogle.com 
// 2      CO      googlecom 
// 5      CO      google.cm 
// 6      CO      google.co 
// 7      CO      oogle.com 
// 8      CO      goole.com 
// 9      CO      googe.com 
// 3      CO      googl.com 
// 4      CO      google.om 
// --------------------------
// TOTAL  9   
//
//
// Input: abcd
//
// Output:
// ID     TYPE    TYPO 
// ---------------------
//  3      CO      abd  
//  4      CO      abc  
//  1      CO      bcd  
//  2      CO      acd  
// ---------------------
//  TOTAL  4            




import (
	"fmt"

	"github.com/rangertaha/urlinsane"
	"github.com/rangertaha/urlinsane/plugins/algorithms"
)

const (
	CODE        = "co"
	NAME        = "Character Omission"
	DESCRIPTION = "Omitting a character from the name"
)

type Algo struct {
	types []string
}

func (n *Algo) Id() string {
	return CODE
}
func (n *Algo) IsType(str string) bool {
	return algorithms.IsType(n.types, str)
}

func (n *Algo) Name() string {
	return NAME
}

func (n *Algo) Description() string {
	return DESCRIPTION
}

func (n *Algo) Exec(typo urlinsane.Typo) (typos []urlinsane.Typo) {
	for _, variant := range n.Func(typo.Original().Domain()) {
		typos = append(typos, typo.New(variant))
	}
	return
}


// Func swaps numbers and carninal numbers
func (n *Algo) Func(name string) (results []string) {
	for i := range name {
		if i <= len(name)-1 {
			variant := fmt.Sprint(
				name[:i],
				name[i+1:],
			)
			if name != variant {
				results = append(results, variant)
			}
		}

	}

	return results
}

// Register the plugin
func init() {
	algorithms.Add(CODE, func() urlinsane.Algorithm {
		return &Algo{
			types: []string{algorithms.ENTITY, algorithms.DOMAIN},
		}
	})
}
