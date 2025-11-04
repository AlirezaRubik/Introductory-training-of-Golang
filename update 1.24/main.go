package main

type MyAlias[T any] = []T


func main() {
    var a MyAlias[int] = []int{1,2,3}
    _ = a
}
