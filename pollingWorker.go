package main

import(
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  loggly "github.com/jamespearly/loggly"
  //"os"
  "time"
)

type Store struct{
  Items []Items
}
type Items struct{
  Metadata []Metadata `json:"metadata"`
}

type Metadata struct{
  Key string `json:"key"`
  Value string `json:value`
}

type PlayerProfile struct{
  AccountId string `json:"accountId"`
  PlatformName string `json:"platformName"`
  EpicUserHandle string `json:"epicUserHandle"`
  LifeTimeStats []LifeTimeStatsStruct `json:"lifeTimeStats"`
}

type LifeTimeStatsStruct struct{
  Key string `json:"key"`
  Value string `json:"value"`
}

func main() {

  players := [4]string{"ninja", "couragejd", "nickmercs", "vadnay%20on%20mixer"}

  //name := os.Getenv("ACCOUNT_NAME")
  //platform := os.Getenv("PLATFORM")

  apiKey := "TRN-Api-Key"
  apiValue := "d8b929fb-27a1-48dd-a6ac-ef2092db1291"


  for {
    for _, name := range players{
      url := "https://api.fortnitetracker.com/v1/profile/pc" + "/" + name
      playerProfile := new(PlayerProfile)
      //playerLifeTimeStats := new(PlayerLifetimeStats)
      getContent(url, apiKey, apiValue, playerProfile)
      fmt.Printf("%+v\n", playerProfile)
      time.Sleep(time.Second * 10)
    }
    time.Sleep(time.Second * 10)
  }

}

func getContent(url string, headerKey string, headerValue string, tmp interface{}){


  log := loggly.New("go-polling-worker")

  //create new request
  client := &http.Client{}
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    err := log.EchoSend("error", err.Error())
    fmt.Println("err:", err)
  }

  //api key is a header
  if headerKey != ""{
    req.Header.Add(headerKey, headerValue)
  }

  //execute the request
  resp, err := client.Do(req)
  if err != nil {
    err := log.EchoSend("error", err.Error())
    fmt.Println("err:", err)
  }

  //read the bytes od of the response
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    err := log.EchoSend("error", err.Error())
    fmt.Println("err:", err)
  }
  defer resp.Body.Close()

  //unmarshal the body into our structs and log them
  json.Unmarshal(body, &tmp)
  err = log.EchoSend("info", "Request to " + url + " succeeded")
  if err != nil{
    fmt.Println("err:", err)
  }
}
