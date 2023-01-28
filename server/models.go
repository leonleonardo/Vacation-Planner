package main

type Business struct {
    Name        string `json: "name"`
    Price       string `json: "price"`
    Rating      int    `json: "rating"`
    Location    string `json: "location"`
    Type        string `json: "type"`
}

type Destination struct {
    Location        string      `json: "location"`
    Restaurants     *[5]Business `json: "restaurants"`
    Entertainment   *[5]Business `json: "entertainment"`
}