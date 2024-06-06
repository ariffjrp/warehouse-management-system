package inbound

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"a21hc3NpZ25tZW50/models"

	"github.com/gorilla/mux"
)

func ReadInbound(w http.ResponseWriter, r *http.Request) {
	var products []models.Inbound

	if err := models.DB.Find(&products).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func ShowInboundId(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(400)
		err := models.ErrorResponse{Error: "invalid show inbound request"}
		json, _ := json.Marshal(err)
		w.Write(json)
		return
	}

	var product models.Inbound
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			w.WriteHeader(404)
			err := models.ErrorResponse{Error: "Inbound not found"}
			json, _ := json.Marshal(err)
			w.Write(json)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Internal Server Error"})
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func CreateInbound(w http.ResponseWriter, r *http.Request) {

	var inbound models.Inbound

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&inbound); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "invalid inbound request"})
		return
	}

	if inbound.Products == "" || inbound.Date == "" || inbound.Qty == "" || inbound.Location == "" || inbound.Building == "" || inbound.Room == "" || inbound.Floor == "" || inbound.Area == "" || inbound.Rack == "" || inbound.Racklevel == "" {
		w.WriteHeader(400)
		err := models.ErrorResponse{Error: "invalid add inbound request"}
		json, _ := json.Marshal(err)
		w.Write(json)
		return
	}

	defer r.Body.Close()

	if err := models.DB.Create(&inbound).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	response := map[string]string{"message": "Product succes add"}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}

func UpdateInboundId(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(400)
		err := models.ErrorResponse{Error: "invalid update inbound request"}
		json, _ := json.Marshal(err)
		w.Write(json)
		return
	}

	var product models.Inbound

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	defer r.Body.Close()

	if models.DB.Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		w.WriteHeader(400)
		err := models.ErrorResponse{Error: "invalid update inbound request"}
		json, _ := json.Marshal(err)
		w.Write(json)
		return
	}

	product.Id = uint64(id)

	response := map[string]string{"message": "Product success diupdate"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func DeleteInboundId(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(400)
		err := models.ErrorResponse{Error: "invalid show inbound request"}
		json, _ := json.Marshal(err)
		w.Write(json)
		return
	}

	var product models.Inbound
	if err := models.DB.Delete(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			w.WriteHeader(404)
			err := models.ErrorResponse{Error: "Inbound not found"}
			json, _ := json.Marshal(err)
			w.Write(json)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Internal Server Error"})
			return
		}
	}

	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(product)

	response := map[string]string{"message": "Product success dihapus"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
