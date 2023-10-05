package handlers

import (
	"context"
	"fmt"

	"github.com/jmmesquitacardoso/apitest/spec"
)

type pet struct {
}

func NewStrict() spec.StrictServerInterface {
	return &pet{}
}

func (p *pet) FindPets(ctx context.Context, request spec.FindPetsRequestObject) (spec.FindPetsResponseObject, error) {
	fmt.Println("finding some pets!!")

	return spec.FindPets200JSONResponse{}, nil
}

func (p *pet) AddPet(ctx context.Context, request spec.AddPetRequestObject) (spec.AddPetResponseObject, error) {
	fmt.Println("adding some pet!!")

	return spec.AddPet200JSONResponse{}, nil
}

func (p *pet) DeletePet(ctx context.Context, request spec.DeletePetRequestObject) (spec.DeletePetResponseObject, error) {
	fmt.Println("deleting some pet!!")

	return spec.DeletePet204Response{}, nil
}

func (p *pet) FindPetById(ctx context.Context, request spec.FindPetByIdRequestObject) (spec.FindPetByIdResponseObject, error) {
	fmt.Println("finding some pet by id!!")

	return spec.FindPetById200JSONResponse{}, nil
}
