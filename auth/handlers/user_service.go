package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"auth/user_service"
)

func UserService(w http.ResponseWriter, r *http.Request){
	if r.URL.Query().Get("type") == "get_user_list" {
		users, err := user_service.GetUserList()
		if err != nil {
		    http.Error(w, "no users in mongo :<", http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}


	if r.URL.Query().Get("type") == "get_user_info" {
		var req user_service.UserInfRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request :<", http.StatusBadRequest)
			return
		}

		user, err := user_service.GetUserInfo(req.UserId)
		if err != nil {
		    http.Error(w, "no such user :<", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}


	if r.URL.Query().Get("type") == "update_full_name" {
		var req user_service.UpdateUserNameRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request :<", http.StatusBadRequest)
			return
		}

		if err := user_service.UpdateUserName(req.UserId, req.NewName); err != nil {
			http.Error(w, "Couldn't update user :<", http.StatusNotFound)
			log.Println("Couldn't update user :<", err)
			return
		}
		w.Write([]byte("Username successfully updated!"))
	}


	if r.URL.Query().Get("type") == "get_user_roles" {
		var req user_service.GetRolesRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request :<", http.StatusBadRequest)
			return
		}

		roles, err := user_service.GetUserRoles(req.UserId)
		if err != nil {
		    http.Error(w, "no roles here :<", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(roles)
	}
}