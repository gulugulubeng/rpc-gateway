package main

func main() {

}


type User struct {
	Id       int      `json:"id" xml:"id"`
	Name,Birthday     string   `json:"name" json:"birthday" xml:"name" xml:"birthday"`
	Location Location `json:"location" xml:"location"`
}

type Location struct {
	Year int `json:"year" xml:"year"`
}
