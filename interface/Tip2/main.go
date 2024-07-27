package main

import "fmt"
type Player struct {
	Id int
	Name string
}
type Walk interface{
	Walk()
}
type Shooting interface{
	Shooting()
}
type Runing interface{
	Runing()
}
type PlayingPlayer interface{
	Walk
	Shooting
	Runing
}
func main(){
var Iplayer PlayingPlayer
Iplayer=&Player{Id: 1,Name: "Alireza"}
Iplayer.Walk()
Iplayer.Shooting()
Iplayer.Runing()

}
func (p *Player)Walk(){
	fmt.Println("The Player With Name:",p.Name,"and Id:",p.Id,"Walking")
}
func (p *Player)Shooting(){
	fmt.Println("The Player With Name:",p.Name,"and Id:",p.Id,"Shooting")
}
func (p *Player)Runing(){
	fmt.Println("The Player With Name:",p.Name,"and Id:",p.Id,"Running")
}