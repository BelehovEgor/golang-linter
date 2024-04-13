package copySlicesAndMapsAtBoundaries

func (d *Driver) SetTripsB(trips []Trip) {
	d.trips = trips
}

func exampleRecSliceBad() {
	trips := []Trip{{name: "a"}, {name: "b"}}
	d1 := Driver{trips: trips}
	d1.SetTripsB(trips)

	// We can now modify trips[0] without affecting d1.trips.
	trips[0] = Trip{name: "c"}
}
