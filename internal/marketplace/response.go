package marketplace

import (
	marketplaceModel "rena-server-v2/internal/marketplace/model"
)

type MarketplaceResponse struct {
	Marketplace Marketplace `json:"marketplace"`
}

type MarketplacesResponse struct {
	Marketplaces      []Marketplace `json:"marketplaces"`
	MarketplacesCount int64         `json:"marketplacesCount"`
}

type Marketplace struct {
	ID              uint   `json:"id"`
	AddressRegistry string `json:"addressRegistry"`
	PlatformFee     string `json:"platformFee"`
	FeeRecipient    string `json:"feeRecipient"`
}

// NewMarketplacesResponse converts marketplace models and total count to MarketplacesResponse
func NewMarketplacesResponse(marketplaces []*marketplaceModel.Marketplace, total int64) *MarketplacesResponse {
	var a []Marketplace
	for _, marketplace := range marketplaces {
		a = append(a, NewMarketplaceResponse(marketplace).Marketplace)
	}

	return &MarketplacesResponse{
		Marketplaces:      a,
		MarketplacesCount: total,
	}
}

// NewMarketplaceResponse converts marketplace model to MarketplaceResponse
func NewMarketplaceResponse(a *marketplaceModel.Marketplace) *MarketplaceResponse {
	return &MarketplaceResponse{
		Marketplace: Marketplace{
			ID:              a.ID,
			AddressRegistry: a.AddressRegistry,
			PlatformFee:     a.PlatformFee,
			FeeRecipient:    a.FeeRecipient,
		},
	}
}
