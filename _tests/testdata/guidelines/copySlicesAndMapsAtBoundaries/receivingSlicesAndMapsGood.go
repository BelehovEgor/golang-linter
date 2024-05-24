package copySlicesAndMapsAtBoundaries

type Driver struct {
	trips []Trip
}

type Trip struct {
	name string
}

func (d *Driver) SetTripsG(trips []Trip) {
	d.trips = make([]Trip, len(trips))
	copy(d.trips, trips)
}

func exampleRecSlice() {
	trips := []Trip{{name: "a"}, {name: "b"}}
	d1 := Driver{trips: trips}
	d1.SetTripsG(trips)

	// We can now modify trips[0] without affecting d1.trips.
	trips[0] = Trip{name: "c"}
}
