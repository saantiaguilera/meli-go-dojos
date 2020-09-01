package pkg_test

import (
	"errors"
	"testing"

	"github.com/mercadolibre/fury_shipping-dx-dojo/go/pkg"
	"github.com/stretchr/testify/assert"
)

func TestHacerCortado(t *testing.T) {
	m := pkg.CrearMachine(
		func() (pkg.Bebida, error) {
			return pkg.Bebida{}, errors.New("not supported")
		},
		func() (pkg.Bebida, error) {
			return pkg.Bebida{}, errors.New("not supported")
		},
		func() (pkg.Bebida, error) {
			return pkg.CreateCortado(), nil
		},
	)
	b, err := m.GetCortado()

	assert.Nil(t, err)
	assert.Equal(t, pkg.Bebida{
		Cafe:  250,
		Leche: 150,
	}, b)
}

func TestHacerDosALaVezFalla(t *testing.T) {

}
