package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	collectionDB "rena-server-v2/internal/collection/database"
	collectionModel "rena-server-v2/internal/collection/model"
	"rena-server-v2/internal/database"
	itemDB "rena-server-v2/internal/item/database"
	itemModel "rena-server-v2/internal/item/model"
	"strings"
	"time"

	"github.com/gosimple/slug"
	"github.com/spf13/cobra"
)

var factoryCmd = &cobra.Command{
	Use: "factory",
	Run: func(cmd *cobra.Command, args []string) {
		runFactory()
	},
}

type TokenResponse struct {
	Status string `json:"status"`
	Data   struct {
		Tokens []itemModel.Item `json:"tokens"`
		Total  int64            `json:"total"`
	}
}

type MetaData struct {
	Compiler     string `json:"compiler"`
	Description  string `json:"description"`
	DNA          string `json:"dna"`
	Image        string `json:"image"`
	Date         string `json:"date"`
	Royalty      int    `json:"royalty"`
	FeeRecipient string `json:"fee_recipient"`
	Category     string `json:"category"`
	Website      string `json:"website"`
	Discord      string `json:"discord"`
	TwitterUrl   string `json:"twitter_url"`
	InstagramUrl string `json:"instagram_url"`
	MediumUrl    string `json:"medium_url"`
	TelegramUrl  string `json:"telegram_url"`
	ContactEmail string `json:"contact_email"`
}

func request(target string, from int, count int, sortBy string, cc chan string) {
	query := fmt.Sprintf(`{"from":"%d", "count":"%d", "type":"%s", "sortby":"%s"}`, from, count, target, sortBy)
	var jsonStr = []byte(query)
	req, err := http.NewRequest("POST", "https://fetch-tokens.vercel.app/api", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	cc <- string(body)
}

func runFactory() {
	fmt.Println("Add initial nft items to database...")

	cfg, _ := loadConfig()

	db, _ := database.NewDatabase(cfg)
	collectionDB := collectionDB.NewCollectionDB(db)
	itemDB := itemDB.NewItemDB(db)

	cc := make(chan string, 1)
	go request("all", 0, 1, "createdAt", cc)
	msg := <-cc
	var token TokenResponse
	json.Unmarshal([]byte(msg), &token)

	fmt.Println(token.Data.Total)
	var meta MetaData
	re := regexp.MustCompile(`^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`)

	for index := 1; index < int(token.Data.Total/6); index++ {
		go func() {
			go request("all", 6*index, 6, "createdAt", cc)
			msg = <-cc
			json.Unmarshal([]byte(msg), &token)

			for _, item := range token.Data.Tokens {
				token, err := itemDB.FindItemByAddress(context.Background(), item.ContractAddress, item.TokenID)
				if err != nil {
					fmt.Println(err)
					continue
				}

				fmt.Println(item.Name)
				if token == nil {
					var collection collectionModel.Collection
					// Save collection
					tokenURI := item.TokenURI
					if !re.MatchString(item.TokenURI) {
						tokenURI = strings.Replace(tokenURI, "ipfs://", "https://ipfs.io/ipfs/", -1)
					}
					if tokenURI == "ipfs://" {
						continue
					}

					resp, err := http.Get(tokenURI)
					if err != nil {
						fmt.Println(err)
						continue
					}
					defer resp.Body.Close()
					body, err := io.ReadAll(resp.Body)

					json.Unmarshal([]byte(body), &meta)
					if err != nil {
						continue
					}

					if meta.Compiler != "" {
						col, _ := collectionDB.FindCollectionBySlug(context.Background(), slug.Make(meta.Compiler))

						if col == nil {
							fmt.Println(meta.Compiler)
							// save collection
							collection = collectionModel.Collection{
								ContractAddress: meta.DNA,
								Slug:            slug.Make(meta.Compiler),
								Name:            meta.Compiler,
								Description:     meta.Description,
								ImageURL:        meta.Image,
								Royalty:         meta.Royalty,
								FeeRecipient:    meta.FeeRecipient,
								Category:        meta.Category,
								Website:         meta.Website,
								Discord:         meta.Discord,
								TwitterUrl:      meta.TwitterUrl,
								InstagramUrl:    meta.InstagramUrl,
								MediumUrl:       meta.MediumUrl,
								TelegramUrl:     meta.TelegramUrl,
								ContactEmail:    meta.ContactEmail,
							}

							err = collectionDB.SaveCollection(context.Background(), &collection)
							if err != nil {
								if database.IsKeyConflictErr(err) {
									fmt.Println("Duplicated collection name")
								}
								continue
							}
						} else {
							collection = *col
						}
					}

					// save item
					nft := itemModel.Item{
						ContentType:               item.ContentType,
						ContractAddress:           item.ContractAddress,
						ImageURL:                  item.ImageURL,
						IsAppropriate:             item.IsAppropriate,
						LastSalePrice:             item.LastSalePrice,
						LastSalePriceInUSD:        item.LastSalePriceInUSD,
						LastSalePricePaymentToken: item.LastSalePricePaymentToken,
						Liked:                     item.Liked,
						Name:                      item.Name,
						PaymentToken:              item.PaymentToken,
						Price:                     item.Price,
						PriceInUSD:                item.PriceInUSD,
						Supply:                    item.Supply,
						ThumbnailPath:             item.ThumbnailPath,
						TokenID:                   item.TokenID,
						TokenType:                 item.TokenType,
						TokenURI:                  item.TokenURI,
						CollectionID:              collection.ID,
					}

					err = itemDB.SaveItem(context.Background(), &nft)
					if err != nil {
						if database.IsKeyConflictErr(err) {
							fmt.Println("Duplicated collection name")
						}
						continue
					}
				}
			}
		}()
		time.Sleep(3 * time.Second)
	}
}
