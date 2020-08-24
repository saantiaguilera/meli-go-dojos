package pkg

import (
	"errors"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_shipping-dx-dojo/pkg/external"
)

/*
-> OrderID
<- A procesar

-- Obtener orden
-- crear ruta + Obtener drivers candidatos para la orden
-- Hacer rostering de drivers
- trackear
*/

func Example() {
	orderID := uint(1)
	tracker := external.NewMetricTracker()

	_ = tracker.Track("order", orderID)

	orderProvider := external.NewOrderProvider()
	order := orderProvider.Get(orderID)

	routeCreator := external.NewRouteCreator()
	route := routeCreator.Create(order)

	driverSelector := external.NewDriverSelector()
	drivers, _ := driverSelector.Select(order)

	driverService := external.NewDriverService()

	var ds []external.Driver
	for _, driverID := range drivers {
		d, err := driverService.GetDriverInfo(driverID)

		if err == nil {
			ds = append(ds, d)
		}
	}

	rosteringService := external.NewRosteringService()
	selectedDriver, _ := rosteringService.PerformRostering(route, ds)

	_ = tracker.Track("candidate", selectedDriver)
}

func NewOrder(orderId uint) {
	var wg sync.WaitGroup
	wg.Add(1)
	var order external.Order
	go trackOrder(orderId)

	go func(wg *sync.WaitGroup) {
		order = getOrder(orderId)
		wg.Done()
	}(&wg)
	wg.Wait()

	wg.Add(2)
	//	var route external.Route
	go func(wg *sync.WaitGroup) {
		_ = createRoute(order)
		wg.Done()
	}(&wg)

	//var drivers []external.Driver
	var err error
	go func(wg *sync.WaitGroup) {
		driversId, e := getDrivers(order)
		_ = getDriversById(driversId)
		err = e
		wg.Done()
	}(&wg)
	wg.Wait()
	if err != nil {
		tracker := external.NewMetricTracker()
		tracker.Track("Error while getting drivers", err)
		return
	}
}

func Println(msg string) {
	fmt.Println(msg)
}

func getDriversById(ids []uint) []external.Driver {
	var wg sync.WaitGroup
	wg.Add(len(ids))
	driverService := external.NewDriverService()
	var ds []external.Driver
	var lock sync.Mutex
	for _, driverID := range ids {
		go func(id uint, wg *sync.WaitGroup, l *sync.Mutex) {
			defer wg.Done()
			d, err := getInfo(driverService, id)
			if err != nil {
				l.Lock()
				ds = append(ds, d)
				l.Unlock()
			}
		}(driverID, &wg, &lock)
	}

	wg.Wait()
	return ds
}

type DriverService interface {
	GetDriverInfo(driverID uint) (external.Driver, error)
}

func getInfo(driverService DriverService, driverID uint) (external.Driver, error) {
	d, err := driverService.GetDriverInfo(driverID)
	if err != nil {
		tracker := external.NewMetricTracker()
		tracker.Track("Error while getting drivers", err)
		return external.Driver{}, err
	}
	return d, nil
}

func createRoute(o external.Order) external.Route {
	r := external.NewRouteCreator()
	route := r.Create(o)

	return route
}

func getDrivers(o external.Order) ([]uint, error) {
	s := external.NewDriverSelector()
	return s.Select(o)
}

func getOrder(orderId uint) external.Order {
	orderProvider := external.NewOrderProvider()
	return orderProvider.Get(orderId)
}

func trackOrder(orderId uint) {
	tracker := external.NewMetricTracker()

	_ = tracker.Track("order", orderId)
}

type User struct {
	ID int
}
type UserService interface {
	GetUsers() []User
}

type GetUserUseCase struct {
	userService UserService
}

var ErrorNotFound = errors.New("no user found")

func (g GetUserUseCase) GetUser(userId int) (User, error) {
	users := g.userService.GetUsers()

	for _, u := range users {
		if u.ID == userId {
			return u, nil
		}
	}
	return User{}, ErrorNotFound
}

func NewGetUserUseCase(u UserService) GetUserUseCase {
	return GetUserUseCase{
		userService: u,
	}
}

func PostUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"user_id": 200,
		"name":    "the-it-crowd",
	})
}
