package event

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	collectionDB "rena-server-v2/internal/collection/database"
	"rena-server-v2/internal/config"
	itemContract "rena-server-v2/internal/contract/item"
	itemDB "rena-server-v2/internal/item/database"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	collectionDB collectionDB.CollectionDB
	itemDB       itemDB.ItemDB
}

type MetaData struct {
	ContractAddress string `json:"contract_address"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Image           string `json:"image"`
	Date            string `json:"date"`
	Royalty         int    `json:"royalty"`
}

func events(cfg *config.Config, address string) {
	client, err := ethclient.Dial(cfg.CryptoConfig.WebSocket)
	if err != nil {
		return
	}

	contractAddress := common.HexToAddress(address)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(itemContract.RenaABI)))
	if err != nil {
		return
	}

	for _, vLog := range logs {
		result, err := contractAbi.Unpack("Minted", vLog.Data)
		if err != nil {
			continue
		}

		fmt.Println(result)

		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}
	}
}

func SubscribeCollection(cfg *config.Config, h *Handler, id uint, address string) {
	client, err := ethclient.Dial(cfg.CryptoConfig.WebSocket)
	if err != nil {
		return
	}

	contractAddress := common.HexToAddress(address)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(itemContract.RenaABI)))
	if err != nil {
		return
	}

	for {
		select {
		case err := <-sub.Err():
			fmt.Println(err)
		case vLog := <-logs:
			result, err := contractAbi.Unpack("Minted", vLog.Data)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(result[2])
			tokenURI := strings.Replace(result[2].(string), "ipfs://", "https://rena.mypinata.cloud/ipfs/", -1)

			resp, err := http.Get(tokenURI)
			if err != nil {
				fmt.Println(err)
				continue
			}
			body, err := io.ReadAll(resp.Body)

			var meta MetaData
			json.Unmarshal([]byte(body), &meta)
			if err != nil {
				continue
			}

			tokenID, _ := strconv.Atoi(result[2].(string))
			fmt.Println(tokenID)
		}
	}
}

func RouteV1(cfg *config.Config, h *Handler, r *gin.Engine) {
	criteria := collectionDB.IterateCollectionCriteria{
		Offset: 0,
		Limit:  100,
	}

	collections, _, err := h.collectionDB.FindCollections(context.Background(), criteria)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range collections {
		go SubscribeCollection(cfg, h, v.ID, v.ContractAddress)
	}
}

func NewHandler(collectionDB collectionDB.CollectionDB, itemDB itemDB.ItemDB) *Handler {
	return &Handler{
		collectionDB: collectionDB,
		itemDB:       itemDB,
	}
}
