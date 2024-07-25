// services/product_service.go

package services

import (
	"company/models"
	"database/sql"
)

type ProductService struct {
	DB *sql.DB
}

func (s *ProductService) GetProductList() ([]models.Product, error) {
	var products []models.Product
	rows, err := s.DB.Query("SELECT pno, cate, pname, pcontent, img1, img2, img3, resdate, hits FROM product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.Pno, &product.Cate, &product.Pname, &product.Pcontent, &product.Img1, &product.Img2, &product.Img3, &product.ResDate, &product.Hits); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (s *ProductService) GetProductDetail(pno int) (*models.Product, error) {
	var product models.Product
	row := s.DB.QueryRow("SELECT pno, cate, pname, pcontent, img1, img2, img3, resdate, hits FROM product WHERE pno = ?", pno)
	if err := row.Scan(&product.Pno, &product.Cate, &product.Pname, &product.Pcontent, &product.Img1, &product.Img2, &product.Img3, &product.ResDate, &product.Hits); err != nil {
		return nil, err
	}
	return &product, nil
}

func (s *ProductService) InsertProduct(product *models.Product) error {
	_, err := s.DB.Exec("INSERT INTO product (cate, pname, pcontent, img1, img2, img3, resdate, hits) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		product.Cate, product.Pname, product.Pcontent, product.Img1, product.Img2, product.Img3, product.ResDate, product.Hits)
	return err
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
	_, err := s.DB.Exec("UPDATE product SET cate = ?, pname = ?, pcontent = ?, img1 = ?, img2 = ?, img3 = ?, resdate = ?, hits = ? WHERE pno = ?",
		product.Cate, product.Pname, product.Pcontent, product.Img1, product.Img2, product.Img3, product.ResDate, product.Hits, product.Pno)
	return err
}

func (s *ProductService) DeleteProduct(pno int) error {
	_, err := s.DB.Exec("DELETE FROM product WHERE pno = ?", pno)
	return err
}
