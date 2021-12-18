package main

type trip struct {
	id            int32
	originId      int32
	destinationId int32
	dates         string
	price         float64
}

type city struct {
	id   int32
	name string
}

var trips = []trip{
	{id: 1, originId: 1, destinationId: 2, dates: "Mon Tue Wed Fri", price: 40.55},
	{id: 2, originId: 2, destinationId: 1, dates: "Sat Sun", price: 40.55},
	{id: 3, originId: 3, destinationId: 6, dates: "Mon Tue Wed Thu Fri", price: 32.10},
}
