package hooks

import (
	"github.com/PGo-Projects/tasky/internal/categorydb"
	"github.com/PGo-Projects/tasky/internal/userdb"
	"github.com/PGo-Projects/webauth"
)

func ActivateWebauth() {
	webauth.AddRegisterSuccessHook(categorydb.AddIndexGuardHook)
	webauth.AddRegisterSuccessHook(userdb.GenerateUserEntryHook)
}
