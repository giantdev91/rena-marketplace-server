package collection

import (
	"rena-server-v2/internal/collection/model"
	"time"
)

type CollectionResponse struct {
	Collection Collection `json:"collection"`
}

type CollectionsResponse struct {
	Collection       []Collection `json:"collections"`
	CollectionsCount int64        `json:"collectionsCount"`
}

type Collection struct {
	ID              uint   `json:"id"`
	ContractAddress string `json:"contractAddress"`
	Slug            string `json:"slug"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	ImageURL        string `json:"imageURL"`
	Royalty         int    `json:"royalty"`
	FeeRecipient    string `json:"feeRecipient"`
	Category        string `json:"category"`
	Website         string `json:"website"`
	Discord         string `json:"discord"`
	TwitterUrl      string `json:"twitterUrl"`
	InstagramUrl    string `json:"instagramUrl"`
	MediumUrl       string `json:"mediumUrl"`
	TelegramUrl     string `json:"telegramUrl"`
	ContactEmail    string `json:"contactEmail"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewCollectionsResponse converts collection models and total count to CollectionsResponse
func NewCollectionsResponse(collections []*model.Collection, total int64) *CollectionsResponse {
	var a []Collection
	for _, collection := range collections {
		a = append(a, NewCollectionResponse(collection).Collection)
	}

	return &CollectionsResponse{
		Collection:       a,
		CollectionsCount: total,
	}
}

// NewCollectionResponse converts collection model to CollectionResponse
func NewCollectionResponse(a *model.Collection) *CollectionResponse {
	return &CollectionResponse{
		Collection: Collection{
			ID:              a.ID,
			ContractAddress: a.ContractAddress,
			Name:            a.Name,
			Slug:            a.Slug,
			Description:     a.Description,
			ImageURL:        a.ImageURL,
			Royalty:         a.Royalty,
			FeeRecipient:    a.FeeRecipient,
			Category:        a.Category,
			Website:         a.Website,
			Discord:         a.Discord,
			TwitterUrl:      a.TwitterUrl,
			InstagramUrl:    a.InstagramUrl,
			MediumUrl:       a.MediumUrl,
			ContactEmail:    a.ContactEmail,
			CreatedAt:       a.CreatedAt,
			UpdatedAt:       a.UpdatedAt,
		},
	}
}
