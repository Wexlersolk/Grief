package main

import (
	"fmt"
	"net/http"

	"github.com/Nerzal/gocloak/v10"
	"github.com/Wexlersolk/Grief/internal/store"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

type UserWithToken struct {
	*store.User
	Token string `json:"token"`
}

// registerUserHandler godoc
//
//	@Summary		Registers a user
//	@Description	Registers a user
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		RegisterUserPayload	true	"User credentials"
//	@Success		201		{object}	UserWithToken		"User registered"
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Router			/authentication/user [post]
func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Create user in Keycloak
	client := gocloak.NewClient(app.config.keycloak.url)
	token, err := client.LoginAdmin(r.Context(), app.config.keycloak.adminUsername, app.config.keycloak.adminPassword, app.config.keycloak.realm)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	user := gocloak.User{
		Username: gocloak.String(payload.Username),
		Email:    gocloak.String(payload.Email),
		Enabled:  gocloak.Bool(true),
		Credentials: &[]gocloak.CredentialRepresentation{
			{
				Type:      gocloak.String("password"),
				Value:     gocloak.String(payload.Password),
				Temporary: gocloak.Bool(false),
			},
		},
	}

	userID, err := client.CreateUser(r.Context(), token.AccessToken, app.config.keycloak.realm, user)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// Assign the "user" role to the new user
	role, err := client.GetRealmRole(r.Context(), token.AccessToken, app.config.keycloak.realm, "user")
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	err = client.AddRealmRoleToUser(r.Context(), token.AccessToken, app.config.keycloak.realm, userID, []gocloak.Role{*role})
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// Respond with the user ID
	response := map[string]string{
		"userID": userID,
	}

	if err := app.jsonResponse(w, http.StatusCreated, response); err != nil {
		app.internalServerError(w, r, err)
	}
}

type CreateUserTokenPayload struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

// createTokenHandler godoc
//
//	@Summary		Creates a token
//	@Description	Creates a token for a user
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		CreateUserTokenPayload	true	"User credentials"
//	@Success		200		{string}	string					"Token"
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Router			/authentication/token [post]
func (app *application) createTokenHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateUserTokenPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Authenticate user with Keycloak
	client := gocloak.NewClient(app.config.keycloak.url)
	token, err := client.Login(r.Context(), app.config.keycloak.clientID, app.config.keycloak.clientSecret, app.config.keycloak.realm, payload.Email, payload.Password)
	if err != nil {
		app.unauthorizedErrorResponse(w, r, err)
		return
	}

	// Respond with the token
	response := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}

	if err := app.jsonResponse(w, http.StatusOK, response); err != nil {
		app.internalServerError(w, r, err)
	}
}

// validateTokenMiddleware validates the token using Keycloak
func (app *application) validateTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			app.unauthorizedErrorResponse(w, r, fmt.Errorf("missing token"))
			return
		}

		// Validate token with Keycloak
		client := gocloak.NewClient(app.config.keycloak.url)
		rptResult, err := client.RetrospectToken(r.Context(), token, app.config.keycloak.clientID, app.config.keycloak.clientSecret, app.config.keycloak.realm)
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		if !*rptResult.Active {
			app.unauthorizedErrorResponse(w, r, fmt.Errorf("invalid token"))
			return
		}

		// Token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	}
}
