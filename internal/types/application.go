package types

// LocationUpdate contains the location update data types as retrieved from the OsmAnd app by default
type LocationUpdate struct {
	Latitude  float64 `json:"latitude"`  // Latitude, 64 bit floating point number
	Longitude float64 `json:"longitude"` // Longitude, 64 bit floating point number
	Timestamp uint64  `json:"timestamp"` // Timestamp, unsigned 64 bit integer representing the Unix timestamp in milliseconds (not seconds)
	Hdop      float64 `json:"hdop"`      // Horizontal Dilution of Precision, represented as a number (1 means excellent, >20 means very poor)
	Altitude  float64 `json:"altitude"`  // Altitude, represented in meters as a 64 bit floating point number
	Speed     float64 `json:"speed"`     // Speed in m/s, as a 64 bit floating point number
}
