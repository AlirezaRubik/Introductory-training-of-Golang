package main 
import(
	"os"
	"io"
	"fmt"
)
type Errors struct {
	FileName string 
	Message string
}
func (e *Errors) ErrorFunc() {
	 fmt.Printf("The Error Message is: %s The Error FileName is: %s", e.Message,e.FileName)
}
func main() {

err:=CopyFile2("//home//alireza//Desktop//Learn-Golang-p1//Errors//WrappingErrors//Tip3//logs.log","//home//alireza//Desktop//Learn-Golang-p1//Errors//WrappingErrors//Tip3//BackUplogs.log")
if err != nil {
	fmt.Println(err)
}
}

func Constructor(FileName string, Message string){
	ErrorMaker:=Errors{}
	ErrorMaker.FileName=FileName
	ErrorMaker.Message=Message
	fmt.Println()
	ErrorMaker.ErrorFunc()
}
func CopyFile2(Source string, Destination string)error {
	File1, err := os.Open(Source)
	if err != nil {
		Constructor(Source,"When I Want Opening I Saw The Source File IS CLOSED")
         return err
	} 
	defer File1.Close()
	File2,err:=os.Create(Destination)
	if err != nil {
		Constructor(Destination,"When I Want Creating I Saw The Destination File IS CLOSED")
		return err
	}
    defer File2.Close() 
	_,err=io.Copy(File2,File1)
	if err != nil {
		Constructor(Destination,"When I Want Copying I Saw The Destination File IS CLOSED")
		return err
	}
	return nil
}