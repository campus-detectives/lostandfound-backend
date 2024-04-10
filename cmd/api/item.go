package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"github.com/campus-detectives/lostandfound-backend/internal/data"
)

func (app *application) addItem(w http.ResponseWriter, r *http.Request) {
	// check if user is guard else send
	usr := app.contextGetUser(r)
	log.Println(usr)
	if !usr.IsGuard {
		app.invalidCredentialsResponse(w, r)
		return
	}

	type input struct {
		Location string `json:"location"`
		Category string `json:"category"`
		Image    string `json:"image"`
	}

	// decode input
	var i input
	err := app.readJSON(w, r, &i)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// insert found item
	found := &data.Item{
		FoundBy:  usr.ID,
		Location: i.Location,
		Category: i.Category,
		Image:    i.Image,
	}
	err = app.models.Found.Insert(found)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	cmd := exec.Command("python3", "../models/scripts/upload.py", fmt.Sprint(found.ID))

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out), err)
	}

}

func (app *application) getAllUnclaimed(w http.ResponseWriter, r *http.Request) {
	usr := app.contextGetUser(r)
	if usr.IsAnonymous() {
		app.invalidCredentialsResponse(w, r)
		return
	}
	var err error

	founds, err := app.models.Found.GetAllItems()
	// get all found items
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// encode found items
	err = app.writeJSON(w, http.StatusOK, envelope{"found": founds}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) searchByImage(w http.ResponseWriter, r *http.Request) {
	usr := app.contextGetUser(r)
	if usr.IsAnonymous() {
		app.invalidCredentialsResponse(w, r)
		return
	}
	type input struct {
		Image     string  `json:"image"`
		Threshold float64 `json:"threshold"`
	}
	var i input
	err := app.readJSON(w, r, &i)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	log.Println("test")
	cmd := exec.Command("python3", "../models/scripts/comparator.py", fmt.Sprint(i.Threshold))
	cmd.Stdin = strings.NewReader(i.Image)
	output, err := cmd.CombinedOutput()
	log.Println("[Debug]", string(output))

	if err != nil {
		log.Println(string(output))
		app.serverErrorResponse(w, r, err)
		return
	}
	var ids []int64
	if len(strings.Trim(string(output), "\r\n")) > 0 {

		for _, id := range strings.Split(strings.Trim(string(output), "\r\n"), " ") {
			if id == "" {
				continue
			}
			i, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				app.serverErrorResponse(w, r, err)
				return
			}
			ids = append(ids, i)
		}
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"match": ids}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) claimItem(w http.ResponseWriter, r *http.Request) {
	usr := app.contextGetUser(r)
	if !usr.IsGuard {
		app.invalidCredentialsResponse(w, r)
		return
	}
	var input struct {
		Id        int64  `json:"id"`
		ClaimedBy string `json:"claimed_by"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.models.Found.Claim(input.Id, input.ClaimedBy)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "Claimed successfully"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}

func (app *application) contestClaim(w http.ResponseWriter, r *http.Request) {
	usr := app.contextGetUser(r)
	if usr.IsAnonymous() {
		app.invalidCredentialsResponse(w, r)
		return
	}
	var input struct {
		Id          int64  `json:"id"`
		ContestedBy string `json:"contested_by"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	log.Printf("User %d: %v %v contested claim on item %d", usr.ID, usr.Username, input.ContestedBy, input.Id)
}
