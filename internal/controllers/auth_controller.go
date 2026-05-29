package controllers
func Register(w http.ResponseWriter, r *http.Request) {

	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	user.Password = hashedPassword

	newUser, err := repositories.UserCreate(r.Context(), user)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(newUser)
}