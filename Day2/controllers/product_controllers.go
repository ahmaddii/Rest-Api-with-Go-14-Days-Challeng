package controllers

import (
	"Day2/database"
	"Day2/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

// now get Product by specific id

func GetProduct(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id := params["id"]

	var product models.Product

	query := `SELECT id,namee,descriptionn,price,quantity FROM products WHERE id = ?`

	err := database.DB.QueryRow(query, id).Scan(

		&product.ID,
		&product.Namee,
		&product.Descriptionn,
		&product.Price,
		&product.Quantity,
	)

	if err != nil {

		http.Error(w, "Produc Not found ", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)

}

// now if I want to update sepecifc task wtr to its id

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id := params["id"]

	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {

		http.Error(w, "Invalid data from body", http.StatusBadRequest)
		return

	}

	query := `UPDATE products SET namee = ? , descriptionn = ? , price = ? , quantity = ? WHERE id = ?`

	_, err = database.DB.Exec(query, product.Namee, product.Descriptionn, product.Price, product.Quantity, id)

	if err != nil {

		http.Error(w, "Error while Updating Data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "task upadted Succesfully"})

}

// Last crud function is to delete specific product with respect to its id

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id := params["id"]

	query := `DELETE FROM products WHERE id = ?`

	result, err := database.DB.Exec(query, id)

	if err != nil {

		http.Error(w, "Error while Deleting a Task from Menu", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {

		http.Error(w, "Error while Checking the Row ", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {

		http.Error(w, "Product not found ", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "product deleted succesfully"})

}
