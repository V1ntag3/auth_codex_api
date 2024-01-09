package utilities

import (
	"auth_codex_api/database"
	"auth_codex_api/models"
)

// Remove token
func UnauthorizedToken(token string) {

	var revokeToken models.ValidToken

	revokeToken.Token = token

	database.DB.Where("token = ?", token).First(&revokeToken).Delete(&revokeToken)
}

// Register a valid token
func AuthorizedToken(token string) {

	var revokeToken models.ValidToken

	revokeToken.Token = token

	database.DB.Create(&revokeToken)

}

// Check if a token is not revoked
func IsAuthorizedToken(token string) bool {

	var ValidToken models.ValidToken

	database.DB.Where("token = ?", token).First(&ValidToken)
	if ValidToken.Token == token {
		// O token está na tabela
		return true
	}
	print(ValidToken.Token)
	// O token não esta na tabela
	return false
}
