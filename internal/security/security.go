package security

import (
	"context"
	"encoding/base64"

	"github.com/PGo-Projects/output"
	"github.com/PGo-Projects/tasky/internal/config"
	"github.com/PGo-Projects/tasky/internal/securitydb"
	"github.com/PGo-Projects/webauth"
	"github.com/go-chi/chi"
	"github.com/spf13/viper"
)

func MustDecodeBase64(s string) []byte {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		output.ErrorAndPanic(err)
	}
	return decoded
}

func MustSetup(mux *chi.Mux) {
	// Setup security middleware
	mux.Use(webauth.ExpirationMiddleware)

	authenticationKey := MustDecodeBase64(viper.GetString(config.AuthenticationKey))
	encryptionKey := MustDecodeBase64(viper.GetString(config.EncryptionKey))
	csrfAuthenticationKey := MustDecodeBase64(viper.GetString(config.CSRFAuthenticationKey))

	mongoClient := securitydb.MustMongoClient(context.TODO(), "mongodb://localhost:27017")
	webauth.RegisterDatabase(mongoClient)

	webauth.SetupSessions(authenticationKey, encryptionKey)
	webauth.SessionOptions.Secure = config.ProdRun

	allowedOrigins := viper.GetStringSlice(config.AllowedOrigins)
	webauth.SetupCORS(mux, false, allowedOrigins)
	webauth.SetupCSRF(mux, csrfAuthenticationKey, config.ProdRun)
	webauth.RegisterEndPoints(mux)
}
