package main

import(
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  loggly "github.com/jamespearly/loggly"
  "os"
  //"time"
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

  //lifeTimeStats []map[string]string `json:"lifeTimeStats"`

func main() {


  name := os.Getenv("ACCOUNT_NAME")
  platform := os.Getenv("PLATFORM")

  url := "https://api.fortnitetracker.com/v1/profile/" + platform + "/" + name
  apiKey := "TRN-Api-Key"
  apiValue := "d8b929fb-27a1-48dd-a6ac-ef2092db1291"

  playerProfile := new(PlayerProfile)
  //playerLifeTimeStats := new(PlayerLifetimeStats)
  getContent(url, apiKey, apiValue, playerProfile)
  fmt.Printf("%+v\n", playerProfile)

}

func getContent(url string, headerKey string, headerValue string, tmp interface{}){

  log := loggly.New("go-polling-worker")

  client := &http.Client{}
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    err := log.EchoSend("error", err.Error())
    fmt.Println("err:", err)
  }

  if headerKey != ""{
    req.Header.Add(headerKey, headerValue)
  }

  //json.Marshal(strct)
  resp, err := client.Do(req)
  if err != nil {
    err := log.EchoSend("error", err.Error())
    fmt.Println("err:", err)
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    err := log.EchoSend("error", err.Error())
    fmt.Println("err:", err)
  }
  defer resp.Body.Close()

  json.Unmarshal(body, &tmp)
  err = log.EchoSend("info", "Request to " + url + " succeeded")
  if err != nil{
    fmt.Println("err:", err)
  }
}
