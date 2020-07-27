package pkg

/**
 * Definition of a calculator
 */
type Calculator interface {
	Sum(a, b int) int
}

/**
 * A basic calculator, knows how to sum
 */
type basicCalculator struct{}

func (c basicCalculator) Sum(a, b int) int {
	return a + b
}

func NewBasicCalculator() Calculator {
	return basicCalculator{}
}

/**
 * A result-aware calculator, knows how to sum but keeps records
 */
type resultAwareCalculator struct {
	result int
}

func (r *resultAwareCalculator) Sum(a, b int) int {
	returnValue := r.result + a + b
	r.result += a + b
	return returnValue
	// More complex but works exactly the same:
	//defer func() {
	//	r.result += a + b
	//}()
	//return r.result + a + b
}

func NewResultAwareCalculator() Calculator {
	return &resultAwareCalculator{
		result: 0,
	}
}