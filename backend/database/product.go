package database

var products []Product = []Product{
	{ID: 1, Title: "Product 1", Price: 100, Description: "Product 1 description", Category: "Category 1", ImageUrl: "https://via.placeholder.com/150"},
	{ID: 2, Title: "Product 2", Price: 200, Description: "Product 2 description", Category: "Category 2", ImageUrl: "https://via.placeholder.com/150"},
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	ImageUrl    string  `json:"imageUrl"`
}

func GetProducts() []Product {
	return products
}

func CreateProduct(product Product) Product {
	// Auto-generate ID
	if len(products) == 0 {
		product.ID = 1
	} else {
		// Find the highest ID and increment
		maxID := 0
		for _, p := range products {
			if p.ID > maxID {
				maxID = p.ID
			}
		}
		product.ID = maxID + 1
	}

	products = append(products, product)
	return product
}

func GetProduct(id int) Product {
	for _, product := range products {
		if product.ID == id {
			return product
		}
	}
	return Product{}
}

func UpdateProduct(id int, product Product) {
	for i, p := range products {
		if p.ID == id {
			products[i] = product
			return
		}
	}
}

func DeleteProduct(id int) {
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			return
		}
	}
}
