package pms

// OccupancyFilters is used to filter Occupancy queries. At the moment only some parameters are exposed, while the others are hardcoded
type OccupancyFilters struct {
}

type Occupancy struct {
}

func (u *Unit) Occupancy(filters OccupancyFilters) ([]Occupancy, error) {
	return nil, nil
}
