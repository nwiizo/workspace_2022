package main

//dependencies

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"

	dapr "github.com/dapr/go-sdk/client"
)

//code
func main() {
	for i := 0; i < 10; i++ {
		time.Sleep(5000)
		rand.Seed(time.Now().UnixNano())
		orderId := rand.Intn(1000-1) + 1
		client, err := dapr.NewClient()
		STATE_STORE_NAME := "statestore"
		if err != nil {
			panic(err)
		}
		defer client.Close()
		ctx := context.Background()
		log.Println("Result before get: ")
		log.Println(orderId)
		//Using Dapr SDK to save and get state
		for n := 0; n < 10; n++ {
			KEY_NAME := "order_" + strconv.Itoa(n)
			if err := client.SaveState(ctx, STATE_STORE_NAME, KEY_NAME, []byte(strconv.Itoa(orderId))); err != nil {
				panic(err)
			}
			result, err := client.GetState(ctx, STATE_STORE_NAME, KEY_NAME)
			if err != nil {
				panic(err)
			}
			log.Println("Result after get: ")
			log.Println(result)
		}
	}
}
