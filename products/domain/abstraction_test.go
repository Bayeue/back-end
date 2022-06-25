package domain_products_test

import (
	"errors"
	"os"
	domain_products "ppob/products/domain"
	"ppob/products/domain/mocks"
	service_products "ppob/products/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	productsService domain_products.Service
	productDomain   domain_products.Products
	detailDomain    domain_products.Detail_Product
	categoryDomain  domain_products.Category_Product
	repoProduct     mocks.Repository
)

func TestMain(m *testing.M) {
	productsService = service_products.NewProductsService(&repoProduct)
	productDomain = domain_products.Products{
		ID:          1,
		Name:        "Paket Data XL",
		Image:       "example.jpg",
		Category_Id: 1,
		Status:      true,
	}
	detailDomain = domain_products.Detail_Product{
		ID:    1,
		Price: 20000,
	}
	categoryDomain = domain_products.Category_Product{
		ID:    1,
		Name:  "Payment",
		Image: "example_cat.jpg",
	}
	os.Exit(m.Run())
}

func TestInsertData(t *testing.T) {
	t.Run("success insert data", func(t *testing.T) {
		repoProduct.On("Store", mock.Anything).Return(nil).Once()
		err := productsService.InsertData(categoryDomain.ID, productDomain)

		assert.NoError(t, err)
		assert.Equal(t, nil, err)
	})

	t.Run("failed insert data", func(t *testing.T) {
		repoProduct.On("Store", mock.Anything).Return(errors.New("status internal error")).Once()
		err := productsService.InsertData(categoryDomain.ID, productDomain)

		assert.Error(t, err)
		assert.Equal(t, err, err)
	})
}
func TestGetProducts(t *testing.T) {
	t.Run("success get products", func(t *testing.T) {
		repoProduct.On("GetAll").Return([]domain_products.Products{productDomain}, nil).Once()
		res, err := productsService.GetProducts()

		assert.NoError(t, err)
		assert.Equal(t, productDomain.ID, res[0].ID)
	})
	t.Run("failed get products", func(t *testing.T) {
		repoProduct.On("GetAll").Return([]domain_products.Products{}, errors.New("internal status error")).Once()
		res, err := productsService.GetProducts()

		assert.Error(t, err)
		assert.Equal(t, []domain_products.Products{}, res)
	})
}
