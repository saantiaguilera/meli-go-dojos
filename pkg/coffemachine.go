package pkg

import (
    "errors"
)

/*
 * Armar una maquina de cafe, que procesa de a una sola unidad por vez
 * Posibles bebidas a prepar son:
 * Cafe:
 *  200ml de cafe
 * Cortado:
 *  250ml de cafe
 *  150ml de leche
 * Lagrima:
 *  250ml de leche
 *  150ml de cafe
 */

type CoffeeMachine interface {
    GetCoffee() (Bebida, error)
    GetLagrima() (Bebida, error)
    GetCortado() (Bebida, error)
}

type Bebida struct {
    Cafe int
    Leche int
}

type machine struct {
    ocupado bool
    createCoffee func() (Bebida, error)
    createLagrima func() (Bebida, error)
    createCortado func() (Bebida, error)
}

func CreateCoffee() Bebida{
    return Bebida{
    	Cafe: 200,
    }
}

func CreateLagrima() Bebida{
    return Bebida{
    	Cafe: 150,
    	Leche: 250,
    }
}

func CreateCortado() Bebida{
    return Bebida{
        Cafe: 250,
        Leche: 150,
    }
}

type BebidaFactory func() (Bebida, error)

func CrearMachine(
    createCoffee BebidaFactory,
    createLagrima BebidaFactory,
    createCortado BebidaFactory,
) CoffeeMachine {

    return &machine{
        ocupado : false,
        createCoffee: createCoffee,
        createLagrima: createLagrima,
        createCortado: createCortado,
    }
}

func (m *machine) desocupar() {
	m.ocupado = false
}

func (m *machine) procesar(bebidaFactory BebidaFactory) (Bebida, error) {
    if !m.ocupado {
        m.ocupado = true
        defer m.desocupar()
        return bebidaFactory()
    } else {
        return Bebida{}, errors.New("maquina en uso")
    }
}

func (m *machine) GetCoffee() (Bebida, error) {
    return m.procesar(func() (Bebida, error) {
        return m.createCoffee()
    })
}

func (m *machine) GetLagrima() (Bebida, error) {
    return m.procesar(func() (Bebida, error) {
        return m.createLagrima()
    })
}

func (m *machine) GetCortado() (Bebida, error) {
    return m.procesar(func() (Bebida, error) {
        return m.createCortado()
    })
}

