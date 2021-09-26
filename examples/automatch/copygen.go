// Code generated by github.com/switchupcb/copygen
// DO NOT EDIT.
package copygen

import (
	"github.com/switchupcb/copygen/examples/automatch/domain"
	"github.com/switchupcb/copygen/examples/automatch/models"
)

// ModelsToDomain copies a T, User, Account to a Account.
func ModelsToDomain(tA domain.Account, fT domain.T, fU models.User, fA models.Account) {
	// Account fields
	tA.ID = fA.ID
	tA.Name = fA.Name
	tA.Email = fA.Email

}
