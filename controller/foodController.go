package controller

import (
	"PROJECT_1/database"
	"PROJECT_1/models"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/text/number"
)
 var foodCollection *mongo.Collection = database.OpenCollection(database.Client,"food")
 var validate = validator.New()
 func GetFoods() gin.HandlerFunc{
	return func(c *gin.Context){

	}
 }
 
 func GetFood() gin.HandlerFunc{
	return func(c *gin.Context) {
		 var ctx,cancel = context.WithTimeout(contex.Background(),100*time.Second)
		 foodId = c.param("food_id")
		 var food models.Food
		 err := foodcollection.FindOne(ctx,bson.M{"food_id":foodId}).Decode(&food)
		 defer cancel()
         if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error while fetching food"})
		 }
		 c.JSON(http.StatusOK,food)
	} 
}

 func round(num float32)  int {

 }

 func toFixed(num float64,precision,int) float64{
	
 }

 func UpdateFood() gin.HandlerFunc{
	return func(c *gin.Context) {

	}
 }

 func CreateFood() gin.HandlerFunc{
	return func(c *gin.Context) {
      var ctx,cancel  contxt.WithTimeout(context.Background(),100*time.Second)
       var menu models.Menu 
	   var food models.Food

	   c.BindJSON(&food); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	   }
	    validationErr := validate.Struct(food)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":validationErr.Error()})
		}
		 err := menuCollection.FindOne(ctx,bson.M{"menu_id":food.Menu_id}).Decode(&menu)
		defer cancel()
		if err != nil {
			mas := fmt.Sprintf("menu was not found")
			c.JSON(http.StatusInternalServerError,gin.H{"error":msg})
			return 
		}
		food.Createda_at, _ = time.Parse(time.RFC3339,time.Now()).Format(time.RFC3339)
		food.Updated_at, _ = time.Parse(time.RFC3339,time.Now()).Format(time.RFC3330)
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		var num = toFixed(*food.Price,2)
		food.Price = &num
		
		result,insertErr := foodCollection.InsertOne(ctx,food)
		if insertErr != nil {
			msg := fmt.Sprintf("food item was not created")
			c.JSON(http.StatusInternalServerError,gin.H("error":msg))
			return 
		}
		defer cancel()
		c.JSON(http.StatusOK,result)


	}

 }

