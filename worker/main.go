// main.go

package main

func main() {
	w := Worker{}
	w.Initialize()

	w.Run(":8020")
}
