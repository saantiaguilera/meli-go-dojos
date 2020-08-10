package pkg_test

import (
	"testing"

	"github.com/mercadolibre/fury_shipping-dx-dojo/pkg"
	"github.com/stretchr/testify/assert"
)

func TestMakeConcurrentCoffee(t *testing.T) {
	chIn := make(chan string, 10)
	crearBebida := func() pkg.Result {
		return pkg.Result{
			Bebida: pkg.Bebida{},
			Error:  nil,
		}
	}

	m := pkg.CrearConcurrentMachine(chIn, crearBebida, crearBebida, crearBebida)
	go m.Start()
	m.CreateOrder("Cafe")
	m.CreateOrder("Lagrima")

	assert.Equal(t, crearBebida(), m.ReceiveOrder())
	assert.Equal(t, crearBebida(), m.ReceiveOrder())
	assert.Equal(t, 0, len(chIn))
	close(chIn)
}
