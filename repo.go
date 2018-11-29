package main

import (
	"fmt"
	"google.golang.org/genproto/googleapis/type/money"
)

var currentId int

var bikes Bikes

// some seed data
func init() {
	RepoCreateBike(Bike{Desc: "Haro Midway", Frame: "3 tubes Crmo, Internal HT & mid BB - 20.5 & 21 TT", Gearing: "25/9", CustomerPrice: money.Money{CurrencyCode: "CHF", Units: 459, Nanos: 0}})
	RepoCreateBike(Bike{Desc: "Haro CK AM 2019", Frame: "3 tubes Crmo, Internal HT & mid BB - 20.5 & 21 TT", Gearing: "25/9", CustomerPrice: money.Money{CurrencyCode: "CHF", Units: 759, Nanos: 0}})
}

func RepoFindBike(id int) Bike {
	for _, t := range bikes {
		if t.Id == id {
			return t
		}
	}
	// return empty Bike if not found
	return Bike{}
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
