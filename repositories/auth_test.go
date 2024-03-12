package repositories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIfUsersCanHaveTheSameUsername(t *testing.T) {
	// Gegeven dat een gebruiker een gebruikersnaam wil aanmaken die al bestaan.
	// Wanneer een gebruiker een gebruikersnaam wil aanmaken die al bestaat, DAN moet de applicatie dit weigeren.
	/*
		got := repositories.CreateUser("MaxVersteeg", "Wachtwoord", "Max")
		want := ""
		assert.NotEqual(t, "", repositories.CreateUser("MaxVersteeg", "Wachtwoord", "Max"))
		assert.Error(t)
	*/

	err := CreateUser("user", "Wachtwoord", "user-displayname")
	assert.Nil(t, err)
	err = CreateUser("user", "Wachtwoord", "user-displayname")
	assert.NotNil(t, err)
	/*
		var user *models.User
		result := repositories.DB().Where("Username = ?", "David").Find(&user)
		if result.RowsAffected > 1 {
			t.Errorf("got %v results, want %v results", result, 1)
		}
	*/
}
