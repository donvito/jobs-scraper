package jobs

import (
	"fmt"
	"github.com/rockset/rockset-go-client"
	models "github.com/rockset/rockset-go-client/lib/go"
	"os"
)

func SaveToRockset(docs []interface{}) (err error) {

	defaultAPIServer := os.Getenv("ROCKSET_API_SERVER")
	apiKey := os.Getenv("ROCKSET_API_KEY")

	client, err := rockset.NewClient(
		rockset.WithAPIKey(apiKey),
		rockset.WithAPIServer(defaultAPIServer),
	)
	if err != nil {
		fmt.Println("cannot create rockset client")
	}

	documents := models.AddDocumentsRequest{
		Data: docs,
	}

	res, _, err := client.Documents.Add(RocksetWorkspace,
		RocksetCollection,
		documents)

	if err != nil {
		fmt.Printf("error in adding document %v\n", err)
	}
	res.PrintResponse()

	return
}
