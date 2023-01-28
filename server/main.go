package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
    //"encoding/json"
    "github.com/jmatth11/yfusion"
    "fmt"
    "os"
    "github.com/joho/godotenv"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Gorilla!\n"))
}

// Get Destination Handler function which takes in a location given from the user and prints out relevant info
// Working on getting result information put into json so I can return it for the front-end
func GetDestination(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/json")

    // Using godotenv to hide API keys
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Could not load .env file")
    }

    // Setting key var to YELP
    yelpAPIKey := os.Getenv("YELP_REST_API_KEY")
    // secretKey := os.Getenv("YELP_REST_SECRET_KEY")

    // db := connect()
    // defer db.Close()

    // Taking the location value typed by the user, to use for computations
    params := mux.Vars(r)
    destinationLocation := params["location"]

    // Starting new Yelp client, and setting all relevant filters for specific search
    yelp := yfusion.NewYelpFusion(yelpAPIKey)
    businessSearch := &yfusion.BusinessSearchParams{}
    businessSearch.SetLocation(destinationLocation)
    businessSearch.SetTerm("food")
    businessSearch.SetLimit(10)
    businessSearch.SetRadius(15000)
    businessSearch.SetSortBy("rating")
    foodResult, err := yelp.SearchBusiness(businessSearch)
    if err != nil {
        fmt.Println("X")
    }

    // Printing out results for example, this snippet will be taken out and replaced with lines of code that
    // Take all result info and put into model structs for returns
    for _, b := range foodResult.Businesses {
        if len(b.Price) != 0 {
            fmt.Println("Name:", b.Name, "\nPrice:", b.Price, "\nRating:", b.Categories[0].Title)
        } else {
            fmt.Println("Name:", b.Name, "\nPrice:", "Not listed.", "\nRating:", b.Categories[0].Title)
        }
    }

    // Starting new Yelp client, and setting all relevant filters for specific search
    businessSearch.SetLocation(destinationLocation)
    businessSearch.SetTerm("shopping")
    businessSearch.SetLimit(10)
    businessSearch.SetRadius(15000)
    businessSearch.SetSortBy("rating")
    shoppingResult, err := yelp.SearchBusiness(businessSearch)
    if err != nil {
        fmt.Println("X")
    }

    fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

    // Printing out results for example, this snippet will be taken out and replaced with lines of code that
    // Take all result info and put into model structs for returns
    for _, b := range shoppingResult.Businesses {
        if len(b.Price) != 0 {
            fmt.Println("Name:", b.Name, "\nPrice:", b.Price, "\nRating:", b.Categories[0].Title)
        } else {
            fmt.Println("Name:", b.Name, "\nPrice:", "Not listed.", "\nRating:", b.Categories[0].Title)
        }
    }

    // for index := 0; index < len(ratings); index++ {
    //     fmt.Println(ratings)
    // }

    //destination := &Destination{Location: destinationLocation}

}

func main() {

    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", YourHandler)

    // New function handler for "GET" requests
    r.HandleFunc("/newDestination/{location}", GetDestination)

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8080", r))

}