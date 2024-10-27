package sd

import (
	"github.com/rangertaha/urlinsane"
	"github.com/rangertaha/urlinsane/plugins/algorithms"
)

const CODE = "sd"

type StripDash struct {
	types []string
}

func (n *StripDash) Code() string {
	return CODE
}
func (n *StripDash) IsType(str string) bool {
	return algorithms.IsType(n.types, str)
}

func (n *StripDash) Name() string {
	return "Strip Dash"
}

func (n *StripDash) Description() string {
	return "created by omitting a single dash from the domain"
}

func (n *StripDash) Fields() []string {
	return []string{}
}

func (n *StripDash) Headers() []string {
	return []string{}
}

func (n *StripDash) Exec(urlinsane.Typo) (results []urlinsane.Typo) {
	return
}

// Register the plugin
func init() {
	algorithms.Add(CODE, func() urlinsane.Algorithm {
		return &StripDash{
			types: []string{algorithms.ENTITY, algorithms.DOMAINS},
		}
	})
}
