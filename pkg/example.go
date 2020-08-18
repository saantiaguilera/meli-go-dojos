package pkg

import (
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
