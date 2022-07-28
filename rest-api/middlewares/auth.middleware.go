package middlewares

import (
	"cesargdd/rest-api/controllers"
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type authString string

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			next.ServeHTTP(w, r)
			return
		}
		bearer := "Bearer"
		auth = auth[len(bearer):]

		validate, err := controllers.JwtValidate(context.Background(), auth)
		if err != nil || !validate.Valid {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		customClaim, _ := validate.Claims.(*controllers.JwtCustomClaim)

		ctx := context.WithValue(r.Context(), authString("auth"), customClaim)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func IsAuth(ctx *fiber.Ctx) error {

	auth := ctx.Get("Authorization")
	bearer := "Bearer"
	auth = auth[len(bearer):]

	validate, err := controllers.JwtValidate(context.Background(), auth)
	if err != nil || !validate.Valid {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"message": "not authenticated",
		})
	}

	return ctx.Next()
}

func CtxValue(ctx context.Context) *controllers.JwtCustomClaim {
	raw, _ := ctx.Value(authString("auth")).(*controllers.JwtCustomClaim)
	return raw
}
