package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Dog struct {
	Name  string
	Price float64
	Breed string
}

type Cat struct {
	Name  string
	Price float64
	Breed string
}

var dogs []Dog
var cats []Cat

func HomeDogs(c *gin.Context) {
	dogs = []Dog{
		{Name: "Wilson", Price: 19.99, Breed: "Golden Retriever"},
		{Name: "Roberto", Price: 19.99, Breed: "Shi tzu"},
		{Name: "Matilda", Price: 19.99, Breed: "Mixed"},
	}
}

func HomeCats(c *gin.Context) {
	cats = []Cat{
		{Name: "Rodisney", Price: 15.99, Breed: "Ragdoll"},
		{Name: "Clovis", Price: 9.99, Breed: "British Shorthair"},
		{Name: "Beerus", Price: 17.99, Breed: "Sphynx"},
	}
}

func main() {
	r := gin.New()
	r.GET("/dogs", HomeDogs)
	r.POST("/dogs", PostDogs)
	r.GET("/cats", HomeCats)
	r.POST("/cats", PostCats)
	r.Run()
}

func saveDog(dog Dog) error {
	db, err := sql.Open("mysql", "dogs.db")
	if err != nil {
		fmt.Println(err.Error())
	} 
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO dogs (name, price, breed) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec(dog.Name, dog.Price, dog.Breed)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func PostDogs(c *gin.Context) {
	dog := new(Dog)
	if err := c.Bind(dog); err != nil {
		fmt.Println(err.Error())
	}
	dogs = append(dogs, *dog)
	saveDog(*dog)
	return c.JSON(200, gin.H{dogs})
}