package main


import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)



type SmartContract struct {
	contractapi.Contract
}


type Bid struct {
	BidderID   string `json:"id"`
	OfferID string `json:"offerID"`
	Value int `json:"value"`
}

type Offer struct {
	ID   string `json:"id"`
	Winner string `json:"winner"`//init to ""
	Bids []Bid `json:"bids"`
}

type QueryResult struct {
	Key    string `json:"Key"`
	Record *Offer
}


func (s *SmartContract) Init(ctx contractapi.TransactionContextInterface) error {

	bid1 := Bid{
		BidderID: "bid_1",
		OfferID: "offer_1",
		Value: 5,
	}
	bidAsBytes1, _ := json.Marshal(bid1)

	bid2 := Bid{
		BidderID: "bid_2",
		OfferID: "offer_1",
		Value: 3,
	}
	bidAsBytes2, _ := json.Marshal(bid2)

	
	bid3 := Bid{
		BidderID: "bid_3",
		OfferID: "offer_1",
		Value: 10,
	}
	bidAsBytes3, _ := json.Marshal(bid3)



	offer := Offer{
		ID: "offer_1",
		Winner:  "",
		Bids: []Bid{bid1,bid2,bid3},
	}
	offerAsBytes, _ := json.Marshal(offer)




	err := ctx.GetStub().PutState("offer_1", offerAsBytes)

	if err != nil {
		return fmt.Errorf("Failed to put to world state. %s", err.Error())
	}





	
	//offerWbids, err := s.QueryOffer(ctx, "offer_1")

	//if err != nil {
	//	return err
	//}

	//offer.Bids = append(offer.Bids, bid1)
	//offer.Bids = append(offer.Bids, bid2)
	//offer.Bids = append(offer.Bids, bid3)



	//offerAsBytesWbids, _ := json.Marshal(offerWbids)


	err = ctx.GetStub().PutState("bid_1", bidAsBytes1)

	if err != nil {
		return fmt.Errorf("Failed to put to world state. %s", err.Error())
	}

	err = ctx.GetStub().PutState("bid_2", bidAsBytes2)

	if err != nil {
		return fmt.Errorf("Failed to put to world state. %s", err.Error())
	}

	err = ctx.GetStub().PutState("bid_3", bidAsBytes3)

	if err != nil {
		return fmt.Errorf("Failed to put to world state. %s", err.Error())
	}

	//err = ctx.GetStub().PutState("offer_1", offerAsBytesWbids)

	//if err != nil {
	//	return fmt.Errorf("Failed to put to world state. %s", err.Error())
	//}


	//return ctx.GetStub().PutState(bidderId, bidAsBytes)




	return nil

}



func (s *SmartContract) CreateOffer(ctx contractapi.TransactionContextInterface, offerId string, winner string) error {

/*
	bid1 := Bid{
		BidderID: "bid_1",
		OfferID: offerId,
		Value: 5,
	}


	bid2 := Bid{
		BidderID: "bid_2",
		OfferID: offerId,
		Value: 3,
	}


	
	bid3 := Bid{
		BidderID: "bid_3",
		OfferID: offerId,
		Value: 10,
	}
*/

	newoffer := Offer{
		ID: offerId,
		Winner:  winner,
		Bids: []Bid{},
	}

	newofferAsBytes, _ := json.Marshal(newoffer)
	//tempID := offerId
	//err := ctx.GetStub().PutState(offerId, offerAsBytes)

	//if err != nil {
	//	return fmt.Errorf("Failed to put to world state. %s", err.Error())
	//}


	//return nil


	err := ctx.GetStub().SetEvent("newOfferCreated", newofferAsBytes)

	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(offerId, newofferAsBytes)
}



func (s *SmartContract) QueryOffer(ctx contractapi.TransactionContextInterface, offerId string) (*Offer, error) {
	offerAsBytes, err := ctx.GetStub().GetState(offerId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if offerAsBytes == nil {
		return nil, fmt.Errorf("Offer %s does not exist", offerId)
	}

	offer := new(Offer)
	_ = json.Unmarshal(offerAsBytes, offer)

	return offer, nil
}




func (s *SmartContract) QueryBid(ctx contractapi.TransactionContextInterface, bidderId string) (*Bid, error) {
	bidAsBytes, err := ctx.GetStub().GetState(bidderId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if bidAsBytes == nil {
		return nil, fmt.Errorf("Bid %s does not exist", bidderId)
	}

	bid := new(Bid)
	_ = json.Unmarshal(bidAsBytes, bid)

	return bid, nil
}



//submit bid fucntion has 2 arguments, machineID (from the mainenance event) and a value (to be stored in either bid1 or bid2 or bid 3)
func (s *SmartContract) CreateBid(ctx contractapi.TransactionContextInterface, bidderId string, offerId string, value string) error {

	temp, err := strconv.Atoi(value)
	_ = err



	bid := Bid{
		BidderID:  bidderId,
		OfferID: offerId,
		Value: temp,
	}
	bidAsBytes, _ := json.Marshal(bid)


	//bidAsBytes, _ := json.Marshal(bid)
	//APIstub.PutState(args[0], bidAsBytes)

	offer, err := s.QueryOffer(ctx, offerId)

	if err != nil {
		return err
	}

	offer.Bids = append(offer.Bids, bid)



	offerAsBytes, _ := json.Marshal(offer)


	err = ctx.GetStub().PutState(offerId, offerAsBytes)

	if err != nil {
		return fmt.Errorf("Failed to put to world state. %s", err.Error())
	}

	return ctx.GetStub().PutState(bidderId, bidAsBytes)
}

	//err := APIstub.SetEvent("newBidCreated", bidAsBytes)
	//if err != nil {
	//	return shim.Error(err.Error())
	//}




func (s *SmartContract) MakeDecision(ctx contractapi.TransactionContextInterface, offerId string) error {




	offer, err := s.QueryOffer(ctx, offerId)

	if err != nil {
		return err
	}




	//sort.Slice(offer.Bids[:], func(i, j int) bool {
	//	return offer.Bids[i].Value < offer.Bids[j].Value
	//  })


	//winner := "the bidderID with the min values in the Bids array of bid"
	
	winner := offer.Bids[0].BidderID
	fmt.Printf("initial winner is : %s", offer.Bids[0].BidderID)
	tempmin := offer.Bids[0].Value
	for _, v := range offer.Bids {
			if (v.Value < tempmin) {
				tempmin = v.Value
				winner = v.BidderID
			}
	}
	
	  

	fmt.Printf("winner is : %s", winner)
	offer.Winner = winner

	offerAsBytes, _ := json.Marshal(offer)

	return ctx.GetStub().PutState(offerId, offerAsBytes)


}




func main() {

	// Create a new Smart Contract

	chaincode, err := contractapi.NewChaincode(new(SmartContract))


	if err != nil {
		fmt.Printf("Error create new chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting new chaincode: %s", err.Error())
	}


}