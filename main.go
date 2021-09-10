package main

import "fmt"

func main () {
  var publisher, writer, artist, title string
  var year, pageNumber uint
  var grade float32

  publisher = "Marvel Comics"
  writer = "Stan Lee"
  artist = "Steve Ditko"
  title = "The Amazing Spiderman"
  year = 1963
  pageNumber = 20
  grade = 9.8

  fmt.Println(title, "written by", writer, "drawn by", artist, "year in ", year, "published by", publisher,"number of pages", pageNumber, "condition", grade)

  title = "Mr. GoToSleep"
  writer = "Tracy Hatchet"
  artist = "Jewel Tampson"
  publisher = "DizzyBooks Publishing Inc."
  year = 1997 
  pageNumber = 14
  grade = 6.5

    fmt.Println(title, "written by", writer, "drawn by", artist, "year in ", year, "published by", publisher,"number of pages", pageNumber, "condition", grade)

  title = "Epic Vol.1"
  writer = "Ryan N. Shawn"
  artist = "Phoebe Paperclips"
  publisher = "DizzyBooks Publishing Inc."
  year = 2013
  pageNumber = 160
  grade = 9.0

    fmt.Println(title, "written by", writer, "drawn by", artist, "year in ", year, "published by", publisher,"number of pages", pageNumber, "condition", grade)
}