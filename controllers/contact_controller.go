package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/luisalfredosv/golang-gorilla/models"
	"github.com/luisalfredosv/golang-gorilla/utils"
	"github.com/gorilla/mux"
)

func GetContact(w http.ResponseWrite, r *http.Request){
	contact := models.Contact{}
	id := mux.Vars(r)["id"]
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&contact, id)

	if contact.ID > 0 {
		j, _ := json.Marshal(contact)
		utils.SendResponse(w, http.StatusOK, j)
	}else{
		utils.SendErr(w, StatusNotFound)
	}
}

func GetContacts(w http.ResponseWrite, r *http.Request){
	contacts := []models.Contact{}
	db := utils.GetConnection()
	defer db.Close()

	db.Find(&contacts)

	j, _ := json.Marshal(contacts)
	utils.SendResponse(w, http.StatusOK, j)

}

func StoreContact(w http.ResponseWrite, r *http.Request){
	contact := []models.Contact{}
	db := utils.GetConnection()
	defer db.Close()

	err := json.NewDecoder(r.Body).Decode(&contact)

	if err != {
		fmt.Println(err)
		utils.SendErr(w, http.StatusBadRequest)
		return
	}

	err = db.Create(&contact).Error

	if err != nil {
		fmt.Println(err)
		utils.SendErr(w, http.StatusInternalServerError)
		return
	}

	j, _ := json.Marshal(contact)
	utils.SendResponse(w, http.StatusCreated, j)
}

func UpdateContact(w http.ResponseWrite, r *http.Request){
	contactFind := []models.Contact{}
	contactData := []models.Contact{}

	id := mux.Vars(r)["id"]

	db := utils.GetConnection()
	defer db.Close()

	db.Find(&contactFind, id)

	if contactFind.ID > 0 {

		err := json.NewDecoder(r.Body).Decode(&contactData)
			
			if err != nil {
				utils.SendErr(w, http.StatusBadRequest)
				return
			}

		db.Model(&contactFind).Updates(contactData)
		j, _ := json.Marshal(contact)
		utils.SendResponse(w, http.StatusOK, j)

	}else{
		utils.SendErr(w, StatusNotFound)
	}

}

func DeleteContact(w http.ResponseWrite, r *http.Request){
	contact := models.Contact{}

	id := mux.Vars(r)["id"]
	db := utils.GetConnection()
	defer db.Close()

	db.Find(&contact, id)

	if contact.ID > 0 {
		db.Delete(contact)
		utils.SendResponse(w, http.StatusOK, []byte(`{}`))
	}else{
		utils.SendErr(w, StatusNotFound)
	}
}
