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
package mx

import (
	"net"

	"github.com/rangertaha/urlinsane/internal"
	"github.com/rangertaha/urlinsane/internal/pkg"
	"github.com/rangertaha/urlinsane/internal/plugins/collectors"
	log "github.com/sirupsen/logrus"
)

const (
	ORDER       = 2
	CODE        = "mx"
	NAME        = "MX Records"
	DESCRIPTION = "DNS MX Records"
)

type Plugin struct {
	conf internal.Config
	log  *log.Entry
}

func (n *Plugin) Id() string {
	return CODE
}

func (n *Plugin) Order() int {
	return ORDER
}

func (i *Plugin) Init(c internal.Config) {
	i.log = log.WithFields(log.Fields{"plugin": CODE, "method": "Exec"})
	i.conf = c
}

func (n *Plugin) Description() string {
	return DESCRIPTION
}

func (n *Plugin) Headers() []string {
	return []string{"MX"}
}

func (i *Plugin) Exec(acc internal.Accumulator) (err error) {
	l := i.log.WithFields(log.Fields{"domain": acc.Domain().String()})
	if acc.Domain().Cached() {
		// l.Debug("Returning cache domain: ", acc.Domain().String())
		return acc.Next()
	}

	
	// if acc.Domain().Cached() {
	// l.Debug("Returning cache domain: ", acc.Domain().String())

	// if err = acc.Unmarshal("DNS", dns); err == nil {
	// 	// Update simple table data
	// }
	// }
	// dns := make(pkg.DnsRecords, 0)
	// if err = acc.Unmarshal("DNS", dns); err == nil {
	// 	// Update simple table data
	// 	acc.SetMeta("MX", dns.String("MX"))
	// 	return acc.Next()
	// }
	dns := make(pkg.DnsRecords, 0)
	if err := acc.Unmarshal("DNS", &dns); err != nil {
		l.Error("Unmarshal DNS: ", err)
	}
	records, err := net.LookupMX(acc.Domain().String())
	if err != nil {
		log.Error("MX Lookup: ", err)
	}
	for _, record := range records {
		dns.Add("MX", 0, record.Host)
	}
	if len(records) > 0 {
		acc.SetMeta("MX", dns.String("MX"))
		acc.SetJson("DNS", dns.Json())
		acc.Domain().Live(true)
	}

	// }
	return acc.Next()
}

func (i *Plugin) Close() {}

// Register the plugin
func init() {
	collectors.Add(CODE, func() internal.Collector {
		return &Plugin{}
	})
}
