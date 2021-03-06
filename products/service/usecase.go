package service_products

import (
	"errors"
	domain_products "ppob/products/domain"
	"sort"

	"ppob/helper/slug"
)

type ProductService struct {
	Repository domain_products.Repository
}

func NewProductsService(repo domain_products.Repository) domain_products.Service {
	return ProductService{
		Repository: repo,
	}
}

// GetProduct implements domain_products.Service
func (ps ProductService) GetProduct(id int) (domain_products.Products, error) {
	data, err := ps.Repository.GetByID(id)
	if err != nil {
		return domain_products.Products{}, errors.New("bad request")
	}
	return data, nil
}

// GetProductByCategory implements domain_products.Service
func (ps ProductService) GetProductByCategory(id int) []domain_products.Products {
	data := ps.Repository.GetByCategory(id)
	return data
}

// GetProducts implements domain_products.Service
func (ps ProductService) GetProducts() ([]domain_products.Products, error) {
	datas, err := ps.Repository.GetAll()
	if err != nil {
		return []domain_products.Products{}, errors.New("internal server error")
	}
	return datas, nil
}

// GetProductTransaction implements domain_products.Service
func (ps ProductService) GetProductTransaction(code string) (domain_products.Products, error) {
	data, err := ps.Repository.GetProductTransaction(code)
	if err != nil {
		return domain_products.Products{}, errors.New("bad request")
	}
	return data, nil
}

// InsertData implements domain_products.Service
func (ps ProductService) InsertData(category_id int, cat domain_products.Category_Product, domain domain_products.Products) error {
	productName := cat.Name + " " + domain.Name

	domain.Product_Slug = slug.GenerateSlug(productName)
	domain.Category_Id = category_id
	err := ps.Repository.Store(domain)
	if err != nil {
		return errors.New("internal server error")
	}

	return nil
}

// Delete implements domain_products.Service
func (ps ProductService) Destroy(id int) error {
	data, err := ps.Repository.GetByID(id)
	if err != nil {
		return errors.New("bad request")
	}

	err = ps.Repository.Delete(data.ID)
	if err != nil {
		return errors.New("delete failed")
	}

	err = ps.Repository.DeleteDetails(data.Product_Slug)
	if err != nil {
		return errors.New("delete detail product")
	}

	return nil
}

// Edit implements domain_products.Service
func (ps ProductService) Edit(id int, domain domain_products.Products) error {
	data, err := ps.Repository.GetByID(id)
	if err != nil {
		return errors.New("bad request")
	}

	cat, err := ps.Repository.GetCategoryById(domain.Category_Id)
	if err != nil {
		return errors.New("bad request")
	}

	if domain.Image == "" {
		domain.Image = data.Image
	}

	domain.Product_Slug = slug.GenerateSlug(cat.Name + " " + domain.Name)
	err = ps.Repository.Update(id, domain)
	if err != nil {
		return errors.New("update failed, product")
	}
	err = ps.Repository.UpdateDetails(data.Product_Slug, domain.Product_Slug)
	if err != nil {
		return errors.New("update failed, detail")
	}
	return nil
}

// GetDetails implements domain_products.Service
func (ps ProductService) GetDetails(code string) []domain_products.Detail_Product {
	data, err := ps.Repository.GetDetailsByCode(code)
	sort.Slice(data, func(i, j int) bool {
		return data[i].Price < data[j].Price
	})
	if err != nil {
		return []domain_products.Detail_Product{}
	}
	return data
}

// GetDetail implements domain_products.Service
func (ps ProductService) GetDetail(code string) (domain_products.Detail_Product, error) {
	data, err := ps.Repository.GetDetail(code)
	if err != nil {
		return domain_products.Detail_Product{}, errors.New("bad request")
	}
	return data, nil
}

// InsertDetail implements domain_products.Service
func (ps ProductService) InsertDetail(code string, domain domain_products.Detail_Product) error {
	domain.Detail_Slug = slug.GenerateSlug(domain.Name)
	err := ps.Repository.StoreDetail(code, domain)
	if err != nil {
		return errors.New("internal server error")
	}
	return nil
}

// EditDetail implements domain_products.Service
func (ps ProductService) EditDetail(id int, domain domain_products.Detail_Product) error {
	domain.Detail_Slug = slug.GenerateSlug(domain.Name)
	err := ps.Repository.UpdateDetail(id, domain)
	if err != nil {
		return err
	}
	return nil
}

// DestroyDetail implements domain_products.Service
func (ps ProductService) DestroyDetail(id int) error {
	err := ps.Repository.DeleteDetail(id)
	if err != nil {
		return errors.New("delete failed")
	}
	return nil
}

// GetCategory implements domain_products.Service
func (ps ProductService) GetCategory(id int) (domain_products.Category_Product, error) {
	data, err := ps.Repository.GetCategoryById(id)
	if err != nil {
		return domain_products.Category_Product{}, errors.New("bad request")
	}
	return data, nil
}

// GetCategories implements domain_products.Service
func (ps ProductService) GetCategories() ([]domain_products.Category_Product, error) {
	data, err := ps.Repository.GetCategories()
	if err != nil {
		return []domain_products.Category_Product{}, errors.New("bad request")
	}
	return data, nil
}

// InsertCategory implements domain_products.Service
func (ps ProductService) InsertCategory(domain domain_products.Category_Product) error {
	domain.Category_Slug = slug.GenerateSlug(domain.Name)
	err := ps.Repository.StoreCategory(domain)
	if err != nil {
		return errors.New("internal server error")
	}
	return nil
}

// EditCategory implements domain_products.Service
func (ps ProductService) EditCategory(id int, domain domain_products.Category_Product) error {
	err := ps.Repository.UpdateCategory(id, domain)
	if err != nil {
		return err
	}
	return nil
}

// DestroyCategory implements domain_products.Service
func (ps ProductService) DestroyCategory(id int) error {
	err := ps.Repository.DeleteCategory(id)
	if err != nil {
		return errors.New("delete failed")
	}
	return nil
}

// CountProducts implements domain_products.Service
func (ps ProductService) CountProducts() int {
	count := ps.Repository.Count()
	return count
}
