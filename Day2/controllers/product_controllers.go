package controllers

import (
	"Day2/database"
	"Day2/models"
	"encoding/json"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid Data from Body", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO products(namee, descriptionn, price, quantity) VALUES (?, ?, ?, ?)"

	_, err = database.DB.Exec(query, product.Namee, product.Descriptionn, product.Price, product.Quantity)
	if err != nil {
		http.Error(w, "Error while inserting data: ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ Product Created Successfully!"})
}

// now if you want to show all the products catalog in json
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := database.DB.Query("SELECT id, namee, descriptionn, price, quantity FROM products")
	if err != nil {
		http.Error(w, "Error while getting data: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var pro models.Product
		err := rows.Scan(&pro.ID, &pro.Namee, &pro.Descriptionn, &pro.Price, &pro.Quantity)
		if err != nil {
			http.Error(w, "Error while scanning data: "+err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, pro)
	}

	json.NewEncoder(w).Encode(products)
}
