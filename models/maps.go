package models


type Maps struct {
	Geocoded_Waypoints []Geo
	Routes []Router
}

type Geo struct {
	Geocoder_Status string
	Place_Id string
	Types []string
}

type Router struct {
	Bounds Bounder
	Copyrights string
	Legs []Legger
	Overview_Polyline Points
	Summary string
	Warnings []string
	Waypoint_Order []string
}

type Legger struct {
	Distance Distancer
	Duration Distancer
	End_Address string
	End_Location LatLong
	Start_Address string
	Start_Location LatLong
	Steps []Stepper
	Via_Waypoints []string
}

type Stepper struct {
	Distance Distancer
	Duration Distancer
	End_Location LatLong
	Html_Instruction string
	Polyline Points
	Start_Location LatLong
	Travel_Mode string
}

type Points struct {
	Points string
}

type Distancer struct {
	Text string
	Value int
}

type Bounder struct {
	Northeast LatLong
	Southwest LatLong
}

type LatLong struct {
	Lat float64
	Lng float64
}