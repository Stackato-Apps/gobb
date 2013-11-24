package utils

import (
    "net/http"
    "sirjtaa/models"
)

// TODO: Cache this!
func GetCurrentUser(r *http.Request) *models.User {
	session, _ := Store.Get(r, "sirsid")

    if session.Values["username"] == nil || session.Values["password"] == nil {
        return nil
    }
    err, current_user := models.AuthenticateUser(session.Values["username"].(string), session.Values["password"].(string))

	if err != nil {
		return nil
	}
	return current_user
}