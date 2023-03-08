package registry

import (
	registryModel "rena-server-v2/internal/registry/model"
)

type RegistryResponse struct {
	Registry Registry `json:"registry"`
}

type RegistrysResponse struct {
	Registrys      []Registry `json:"registrys"`
	RegistrysCount int64      `json:"registrysCount"`
}

type Registry struct {
	ID              uint   `json:"id"`
	ContractType    string `json:"contractType"`
	ContractAddress string `json:"contractAddress"`
}

// NewItemsResponse converts registry models and total count to ItemsResponse
func NewItemsResponse(registrys []*registryModel.Registry, total int64) *RegistrysResponse {
	var a []Registry
	for _, registry := range registrys {
		a = append(a, NewRegistryResponse(registry).Registry)
	}

	return &RegistrysResponse{
		Registrys:      a,
		RegistrysCount: total,
	}
}

// NewRegistryResponse converts registry model to RegistryResponse
func NewRegistryResponse(a *registryModel.Registry) *RegistryResponse {
	return &RegistryResponse{
		Registry: Registry{
			ID:              a.ID,
			ContractType:    a.ContractType,
			ContractAddress: a.ContractAddress,
		},
	}
}
