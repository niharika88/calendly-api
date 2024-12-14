package api

type UpdateUserRequest struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
	Timezone  *string `json:"timezone"`
} // @name UpdateUserRequest
