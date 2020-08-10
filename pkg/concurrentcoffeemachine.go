package pkg

type Result struct{
  Bebida Bebida
  Error error
}

type ConcurrentCoffeeMachine struct {
  createCoffee func() Result
  createLagrima func() Result
  createCortado func() Result
  chIn chan string // Recibir ordenes 'cortado' 'lagrima' 'cafe'
  chOut chan Result // Entrega bebidas
}

func CrearConcurrentMachine(ch chan string, createCoffee func() Result, createLagrima func() Result, createCortado func() Result) ConcurrentCoffeeMachine{
  return ConcurrentCoffeeMachine{
    chIn: ch,
    chOut: make(chan Result, 10),
    createCortado: createCortado,
    createCoffee: createCoffee,
    createLagrima: createLagrima,
  }
}

func (c ConcurrentCoffeeMachine) Start() {
    for in := range c.chIn {
        switch in {
          case "Cafe" :
            c.chOut <- c.createCoffee()
          case "Cortado" :
            c.chOut <- c.createCortado()
          case "Lagrima" :
            c.chOut <- c.createLagrima()
          }
      }
}

func(c ConcurrentCoffeeMachine) CreateOrder(b string){
        c.chIn <- b
}

func(c ConcurrentCoffeeMachine) ReceiveOrder() Result{
   return <-c.chOut
}

