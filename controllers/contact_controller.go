package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/luisalfredosv/golang-gorilla/models"
	"github.com/luisalfredosv/golang-gorilla/utils"
	"github.com/gorilla/mux"
)

func GetContact(w http.ResponseWriter, r *http.Request){
	contact := models.Contact{}
	id := mux.Vars(r)["id"]
	db := utils.GetConnection()
	defer db.Close()
	db.Find(&contact, id)

	if contact.ID > 0 {
		j, _ := json.Marshal(contact)
		utils.SendResponse(w, http.StatusOK, j)
	}else{
		utils.SendErr(w, http.StatusNotFound)
	}
}

func GetContacts(w http.ResponseWriter, r *http.Request){
	contacts := []models.Contact{}
	db := utils.GetConnection()
	defer db.Close()

	db.Find(&contacts)

	j, _ := json.Marshal(contacts)
	utils.SendResponse(w, http.StatusOK, j)

}

func StoreContact(w http.ResponseWriter, r *http.Request){
	contact := models.Contact{}
	db := utils.GetConnection()
	defer db.Close()

	err := json.NewDecoder(r.Body).Decode(&contact)

	if err != nil {
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

func UpdateContact(w http.ResponseWriter, r *http.Request){
	contactFind := models.Contact{}
	contactData := models.Contact{}

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
		j, _ := json.Marshal(contactFind)
		utils.SendResponse(w, http.StatusOK, j)

	}else{
		utils.SendErr(w, http.StatusNotFound)
	}

}

func DeleteContact(w http.ResponseWriter, r *http.Request){
	contact := models.Contact{}

	id := mux.Vars(r)["id"]
	db := utils.GetConnection()
	defer db.Close()

	db.Find(&contact, id)

	if contact.ID > 0 {
		db.Delete(contact)
		utils.SendResponse(w, http.StatusOK, []byte(`{}`))
	}else{
		utils.SendErr(w, http.StatusNotFound)
	}
}
