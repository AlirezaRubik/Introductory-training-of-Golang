package TestTdd

import "math"

func SetSalary(BaseSalary float64, ExtraHourse float64) float64 {
	return math.Ceil(BaseSalary/30/7*1.4*ExtraHourse)
}
func SetInsurance(BaseSalary float64, ExtraHourse float64) float64{
	return math.Ceil((BaseSalary+ExtraHourse)*0.03)
}
func SetTax(BaseSalary float64, ExtraHourse float64) float64{
	return math.Ceil((BaseSalary+ExtraHourse)*0.09)
}
func FinalSalary(BaseSalary float64, ExtraHourse float64,Karaneh float64)float64{
	insurance:=SetInsurance(BaseSalary, ExtraHourse)
	taxamount:=SetTax(BaseSalary, ExtraHourse)
	return ((BaseSalary+ExtraHourse+Karaneh)-taxamount)-insurance
}