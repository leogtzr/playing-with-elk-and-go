package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

type Account struct {
	ID            string `json:"id"`
	AccountNumber int    `json:"account_number"`
	Address       string `json:"address"`
	Age           int    `json:"age"`
	Balance       int    `json:"balance"`
	City          string `json:"city"`
	Email         string `json:"email"`
	Employer      string `json:"employer"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	Gender        string `json:"gender"`
	State         string `json:"state"`
}

func main() {
	ctx := context.Background()

	// Obtain a client and connect to the default Elasticsearch installation
	// on 127.0.0.1:9200. Of course you can configure your client to connect
	// to other hosts and configure it in various other ways.
	client, err := elastic.NewClient()
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println(client)
	fmt.Println(ctx)

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists("accounts").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println(exists)

	// Index an account (using JSON serialization)
	account1 := Account{
		ID:            "leoIDX1",
		AccountNumber: 23,
		Balance:       123456,
		FirstName:     "El_Leo",
		LastName:      "El Gtz",
		Gender:        "M",
		Address:       "Nueva XXX 4414",
		Employer:      "OXXO",
		Email:         "leogtzr@abc.ch",
		City:          "Chihuas",
		State:         "CH",
	}

	put1, err := client.Index().
		Index("accounts").
		Type("_doc").
		Id("995533").
		BodyJson(account1).
		Do(ctx)

	/*
		GET accounts/_search
		{
		  "query": {
		    "match": {
		      "email": "leogtzr@...."
		    }
		  }
		}

		GET accounts/_search
		{
		  "query": {
		    "match": {
		      "firstname": "El_Leo"
		    }
		  }
		}
	*/
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed account %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
