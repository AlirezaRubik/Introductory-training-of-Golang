package Fuzz

func CalculationRoomPrice(RoomType string, Night int, Persons int) (FullPayment int) {
	switch RoomType {
	case "sweat":
		FullPayment = (Night * Persons) * 4500
	case "double":
		FullPayment = (Night * Persons) * 4000
	case "standard":
		FullPayment = (Night * Persons) * 3500
	}
	return FullPayment
}