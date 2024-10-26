package none

import (
	"github.com/rangertaha/urlinsane"
	"github.com/rangertaha/urlinsane/plugins/algorithms"
)

type BitFlipping struct {
types []string
}

func (n *BitFlipping) Code() string {
	return "bf"
}
func (n *BitFlipping) IsType(str string) bool {
	return algorithms.IsType(n.types, str)
}

func (n *BitFlipping) Name() string {
	return "Bit Flipping"
}

func (n *BitFlipping) Description() string {
	return "Relies on random bit-errors to redirect connections"
}

func (n *BitFlipping) Fields() []string {
	return []string{}
}

func (n *BitFlipping) Headers() []string {
	return []string{}
}

func (n *BitFlipping) Exec(urlinsane.Typo) (results []urlinsane.Typo) {
	return
}

// Register the plugin
func init() {
	algorithms.Add("bf", func() urlinsane.Algorithm {
		return &BitFlipping{
			
			types: []string{algorithms.ENTITY, algorithms.DOMAINS},
		}
	})
}
