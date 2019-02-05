package main

import(
  "fmt"
  "net/http"
  "io/ioutil"
  //"time"
)

func main() {

  
  url := "https://api.fortnitetracker.com/v1/store"
  client := &http.Client{}
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    fmt.Println(err)
  }

  //form request
  req.Header.Add("TRN-Api-Key", "d8b929fb-27a1-48dd-a6ac-ef2092db1291")
  resp, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
  }

  //read response of body
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println(err)
  }

  //get read body in string form
  bodyString := string(body)
  fmt.Println(bodyString)

  resp.Body.Close()
}
