package main
import(
"fmt"
)
type Ticket struct{
	Id int 
	Origin string
	Destination string
	Date string
	PassangerName string
}
type BustTicket struct{
	Ticket
	Terminal string
}
type AirportTicket struct{
	Ticket
	Gate int
}
type ITicket interface{
	PrintTicket()
}
func main(){

bustTicker:=new(BustTicket)
bustTicker.Id=1
bustTicker.Origin="Tehran"
bustTicker.Destination="Sweeden"
bustTicker.Date="2022-01-01"
bustTicker.PassangerName="Alireza"
bustTicker.Terminal="Terminal 1"
bustTicker.PrintTicket()


airportTicker:=new(AirportTicket)
airportTicker.Id=2
airportTicker.Origin="Istambul"
airportTicker.Destination="Tabriz"
airportTicker.Date="2022-01-05"
airportTicker.PassangerName="Nima"
airportTicker.Gate=10
airportTicker.PrintTicket()


}
func (t *BustTicket)PrintTicket(){
	fmt.Printf("\n\n")
    fmt.Printf("Ticket ID: %d\nOrigin: %s\nDestination: %s\nDate: %s\nPassanger Name: %s\nTerminal:%s\n",t.Id,t.Origin,t.Destination,t.Date,t.PassangerName,t.Terminal)
}
func (t *AirportTicket)PrintTicket(){
	fmt.Printf("\n\n")
    fmt.Printf("Ticket ID: %d\nOrigin: %s\nDestination: %s\nDate: %s\nPassanger Name: %s\nTerminal:%d\n",t.Id,t.Origin,t.Destination,t.Date,t.PassangerName,t.Gate)
}
