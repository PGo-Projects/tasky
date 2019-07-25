package hooks

import (
	"github.com/PGo-Projects/tasky/internal/taskdb"
	"github.com/PGo-Projects/tasky/internal/userdb"
	"github.com/PGo-Projects/webauth"
)

func ActivateWebauth() {
	webauth.AddRegisterSuccessHook(taskdb.AddIndexGuardHook)
	webauth.AddRegisterSuccessHook(userdb.GenerateUserEntryHook)
}
