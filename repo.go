package main

import (
	"fmt"
	"google.golang.org/genproto/googleapis/type/money"
)

var currentId int

var bikes Bikes

// some seed data
func init() {
	RepoCreateBike(Bike{Desc: "Haro Midway", Frame: "3 tubes Crmo, Internal HT & mid BB - 20.5 & 21 TT", Gearing: "25/9", CustomerPrice: money.Money{CurrencyCode: "CHF", Units: 459, Nanos: 50}, SoldOut: false})
	RepoCreateBike(Bike{Desc: "Haro CK AM 2019", Frame: "3 tubes Crmo, Internal HT & mid BB - 20.5 & 21 TT", Gearing: "25/9", CustomerPrice: money.Money{CurrencyCode: "CHF", Units: 759, Nanos: 0}, SoldOut: true})
	RepoCreateBike(Bike{Desc: "Haro Caro 2019", Frame: "All tubes Crmo, Internal HT & mid BB - 20.5 & 21 TT", Gearing: "25/9", CustomerPrice: money.Money{CurrencyCode: "CHF", Units: 652, Nanos: 0}, SoldOut: false})
	RepoCreateBike(Bike{Desc: "Haro Racer TX 2019", Frame: "All tubes Crmo, Internal HT & mid BB - 20.5 & 21 TT", Gearing: "25/9", CustomerPrice: money.Money{CurrencyCode: "CHF", Units: 1250, Nanos: 0}, SoldOut: false})
	RepoCreateBike(Bike{Desc: "Haro Racer Xtreme 2019", Frame: "All tubes Crmo, Internal HT & mid BB - 20.5 & 21 TT", Gearing: "25/9", CustomerPrice: money.Money{CurrencyCode: "CHF", Units: 1768, Nanos: 20}, SoldOut: false})
	RepoCreateBike(Bike{Desc: "Haro Freestyle X 2019", Frame: "All tubes Crmo, Internal HT & mid BB - 20.5 & 21 TT", Gearing: "25/9", CustomerPrice: money.Money{CurrencyCode: "CHF", Units: 980, Nanos: 0}, SoldOut: false})
	RepoCreateBike(Bike{Desc: "Haro Pro 2019", Frame: "All tubes Crmo, Internal HT & mid BB - 20.5 & 21 TT", Gearing: "25/9", CustomerPrice: money.Money{CurrencyCode: "CHF", Units: 520, Nanos: 0}, SoldOut: false})
	RepoCreateBike(Bike{Desc: "Haro Midway Mid 2019", Frame: "All tubes Crmo, Internal HT & mid BB - 20.5 & 21 TT", Gearing: "25/9", CustomerPrice: money.Money{CurrencyCode: "CHF", Units: 1100, Nanos: 0}, SoldOut: false})
	RepoCreateBike(Bike{Desc: "Haro Retro Pro 2019", Frame: "All tubes Crmo, Internal HT & mid BB - 20.5 & 21 TT", Gearing: "25/9", CustomerPrice: money.Money{CurrencyCode: "CHF", Units: 1640, Nanos: 0}, SoldOut: false})
	RepoCreateBike(Bike{Desc: "Haro Beginner 2019", Frame: "All tubes Crmo, Internal HT & mid BB - 20.5 & 21 TT", Gearing: "25/9", CustomerPrice: money.Money{CurrencyCode: "CHF", Units: 1250, Nanos: 0}, SoldOut: false})
	RepoCreateBike(Bike{Desc: "Haro Extreme 2019", Frame: "All tubes Crmo, Internal HT & mid BB - 20.5 & 21 TT", Gearing: "25/9", CustomerPrice: money.Money{CurrencyCode: "CHF", Units: 250, Nanos: 0}, SoldOut: false})
}

func RepoFindBike(id int) (Bike, error) {
	for _, t := range bikes {
		if t.Id == id {
			return t, nil
		}
	}
	// return empty Bike if not found
	return Bike{}, fmt.Errorf("404 not found Bike with id of %d", id)
}

func RepoCreateBike(t Bike) Bike {
	currentId += 1
	t.Id = currentId
	bikes = append(bikes, t)
	return t
}

func RepoDestroyBike(id int) error {
	for i, t := range bikes {
		if t.Id == id {
			bikes = append(bikes[:i], bikes[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Bike with id of %d to delete", id)
}
