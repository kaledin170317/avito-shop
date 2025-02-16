package imp

import (
	"awesomeProject1/internal/entities/inventory"
	"awesomeProject1/internal/storage"
	"context"
	"fmt"
)

type InventoryServiceImp struct {
	rep storage.Repository
}

func NewInventoryServiceImp(rep storage.Repository) *InventoryServiceImp {
	return &InventoryServiceImp{rep: rep}
}

func (s InventoryServiceImp) Get(ctx context.Context, username string) ([]inventory.InventoryResponse, error) {

	inventoryRecords, err := s.rep.InventoryRepository.Get(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get inventory: %w", err)
	}

	var response []inventory.InventoryResponse
	for _, record := range inventoryRecords {
		response = append(response, inventory.InventoryResponse{
			Name:     record.Product.Name,
			Quantity: record.Quantity,
		})
	}

	return response, nil
}
