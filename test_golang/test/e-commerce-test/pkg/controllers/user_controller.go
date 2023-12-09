package controllers

import (
	"encoding/json"
	"fmt"

	// "fmt"
	"net/http"
	"test/e-commerce-test/pkg/models"
	"test/e-commerce-test/pkg/utils"

	"github.com/google/uuid"
)

// store variable in base_controller file

type user_login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type user_register struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
}

type status struct {
	Check string
}

// display the login interface
func (server *Server) Login(w http.ResponseWriter, r *http.Request) {

}

func (server *Server) DoLogin(w http.ResponseWriter, r *http.Request) {
	user_temp := &user_login{}
	utils.ParseBody(r, user_temp)

	userModel := models.User{}
	user, err := userModel.FindByEmail(server.DB, user_temp.Email)
	if err != nil {
		// SetFlash(w, r, "error", "email or password invalid")
		res, _ := json.Marshal(&status{Check: "email or password invalid"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	}

	if !ComparePassword(user_temp.Password, user.Password) {
		// SetFlash(w, r, "error", "email or password invalid")
		res, _ := json.Marshal(&status{Check: "email or password invalid"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	}

	session, _ := store.Get(r, sessionUser)
	session.Values["id"] = user.ID
	session.Save(r, w)
	fmt.Println(session.Values["id"])

	res, _ := json.Marshal(&status{Check: "ok"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	// http.Redirect(w, r, "/", http.StatusSeeOther)
}

// display register interface
func (server *Server) Register(w http.ResponseWriter, r *http.Request) {

}

func (server *Server) DoRegister(w http.ResponseWriter, r *http.Request) {

	user_temp := &user_register{}
	utils.ParseBody(r, user_temp)

	userModel := models.User{}
	existUser, _ := userModel.FindByEmail(server.DB, user_temp.Email)
	if existUser != nil {
		// SetFlash(w, r, "error", "Sorry, email already registered")
		res, _ := json.Marshal(&status{Check: "Sorry, email already registered"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	}

	hashedPassword, _ := MakePassword(user_temp.Password)
	params := &models.User{
		ID:        uuid.New().String(),
		FirstName: user_temp.FirstName,
		LastName:  user_temp.LastName,
		Email:     user_temp.Email,
		Password:  hashedPassword,
	}

	user, err := userModel.CreateUser(server.DB, params)
	if err != nil {
		// SetFlash(w, r, "error", "Sorry, registration failed")
		res, _ := json.Marshal(&status{Check: "Sorry, registration failed"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return
	}

	session, _ := store.Get(r, sessionUser)
	session.Values["id"] = user.ID
	session.Save(r, w)

	res, _ := json.Marshal(&status{Check: "ok"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	// http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (server *Server) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, sessionUser)

	session.Values["id"] = nil
	session.Save(r, w)

	res, _ := json.Marshal(&status{Check: "ok"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

	// http.Redirect(w, r, "/", http.StatusSeeOther)
}
