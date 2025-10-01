package data

import "scalvid/entities"

type Product = entities.Product

var Products []Product = []Product{
	{ID: 1, Title: "Product 1", Price: 100, Description: "Product 1 description", Category: "Category 1", ImageUrl: "https://via.placeholder.com/150"},
	{ID: 2, Title: "Product 2", Price: 200, Description: "Product 2 description", Category: "Category 2", ImageUrl: "https://via.placeholder.com/150"},
}
