package main

import(
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  loggly "github.com/jamespearly/loggly"
  "os"
  "time"
  "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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

  players := [1]string{"Tfue"}

  //name := os.Getenv("ACCOUNT_NAME")
  //platform := os.Getenv("PLATFORM")

  apiKey := "TRN-Api-Key"
  apiValue := os.Getenv("TRN_API_KEY")


  for {
    for _, name := range players{
      url := "https://api.fortnitetracker.com/v1/profile/pc" + "/" + name
      playerProfile := new(PlayerProfile)
      //playerLifeTimeStats := new(PlayerLifetimeStats)
      getContent(url, apiKey, apiValue, playerProfile)
      fmt.Printf("%+v\n", playerProfile)
      sendResponseToDynamoDB(*playerProfile)
      time.Sleep(time.Minute * 15)
    }
    //time.Sleep(time.Minute * 15)
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

  func sendResponseToDynamoDB(playerProfile PlayerProfile) {

	   sess, err := session.NewSession(&aws.Config{
		     Region: aws.String("us-east-1")},
	      )

        if err != nil{
          fmt.Println("err:", err)
        }

	        // Create DynamoDB client
	       svc := dynamodb.New(sess)

		    key, err := dynamodbattribute.MarshalMap(playerProfile)
        id := key["epicUserHandle"].String

        fmt.Println(id)

        if err != nil {
          fmt.Println("err:", err)
        }

        fmt.Println("Key: ", key)

		    input := &dynamodb.PutItemInput{
			        Item:      key,
			        TableName: aws.String("TRN-Stats"),
		}

		_, err = svc.PutItem(input)

    if err != nil{
      fmt.Println("err:", err)
    }

		fmt.Print("Succesfully added item to DB")

	}

  func database_init() *dynamodb.Server {

    auth := aws_auth()

    region := aws.USEast1

    ddbs := NewFrom(auth, region)

    return ddbs
}

func NewFrom(auth aws.Auth, region aws.Region) *dynamodb.Server {
    return &dynamodb.Server{auth, region}
}

func aws_auth() aws.Auth {
    auth, err := aws.EnvAuth()
    return auth
}
