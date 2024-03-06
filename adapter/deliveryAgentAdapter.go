package adapter

import (
	"encoding/json"
	pb "github.com/rima971/allocator/allocator"
	"io"
	"log"
	"net/http"
)

type Response struct {
	status     string
	statusCode int32
	message    string
	data       []pb.DeliveryAgent
}

func FetchNearestDeliveryAgents(pincode int32) ([]pb.DeliveryAgent, error) {
	url := "/delivery-agents?pincode=" + string(pincode)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	//req.Header.Set("basic")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("error occurred while closing response body: %s", err)
		}
	}(res.Body)
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return nil, err
	}
	var daList Response
	err = json.Unmarshal(body, &daList)
	if err != nil {
		return nil, err
	}
	return daList.data, nil
}
