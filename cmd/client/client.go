package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/szmulinho/prescription/internal/model"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type client struct {
	client http.Client
}

type JwtToken struct {
	Token string `json:"token"`
}

func (c *client) postPrescription() {
	newPresc := model.CreatePrescInput{
		PreId:      "2",
		Drugs:      []string{"1", "3"},
		Expiration: time.Time{},
	}
	myJsonValue, _ := json.Marshal(newPresc)

	resp, err := c.client.Post("http://localhost:8080/presc", "application/json", bytes.NewBuffer(myJsonValue))
	if err != nil {
		fmt.Errorf("Error %s", err)
		return
		log.Printf("new presc added")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Error %s", err)
		return
		log.Println("new presc added")
	}
	fmt.Printf("Body : %s", body)
	resp.Body.Close()
	log.Println("new presc added")
}

func main() {
	cli := client{
		client: http.Client{},
	}
	cli.postPrescription()
}
