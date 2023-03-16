package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gin-gonic/gin"
)

type Workout struct {
	Name      string
	Rest      int
	Exercices []interface{}
}

func main() {
	var mongoUser string
	var mongoPassword string
	var mongoHost string
	var mongoPort string

	flag.StringVar(&mongoUser, "mongoUser", "admin", "username")
	flag.StringVar(&mongoPassword, "mongoPassword", "password", "password")
	flag.StringVar(&mongoHost, "mongoHost", "localhost", "Mongo DB host")
	flag.StringVar(&mongoPort, "mongoPort", "27017", "username")

	flag.Parse()

	fmt.Println(mongoUser)
	uri := "mongodb://" + mongoUser + ":" + mongoPassword + "@" + mongoHost + ":" + mongoPort + "/workout"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	r := gin.Default()

	r.GET("/workouts", func(c *gin.Context) {
		fmt.Println(client.Database("workout").Collection("workouts").Find(context.TODO(), bson.D{}))
		cursor, _ := client.Database("workout").Collection("workouts").Find(context.TODO(), bson.D{})
		fmt.Println(cursor)
		var results []bson.M
		if err = cursor.All(context.TODO(), &results); err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"data": results,
		})
	})

	r.POST("/workouts", func(c *gin.Context) {
		var newWorkout Workout
		fmt.Println("CREATE")
		fmt.Println(c.Request.Body)
		if err := c.BindJSON(&newWorkout); err != nil {
			fmt.Println(err)
			return
		}
		client.Database("workout").Collection("workouts").InsertOne(context.TODO(), newWorkout)
		c.IndentedJSON(http.StatusCreated, newWorkout)

	})

	if err != nil {
		panic(err)
	}
	r.Run()
}
