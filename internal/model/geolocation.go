package model

type GeoLocation struct {
	CountryCode string
	Country string
	City string
	Timezone string
	Location *Coordinates
}

type Coordinates struct {
	Latitude float64
	Longitude float64
}
