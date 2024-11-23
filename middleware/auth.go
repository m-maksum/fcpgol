package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Auth middleware untuk memvalidasi token JWT
func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		// Abaikan middleware untuk rute tertentu
		if ctx.Request.URL.Path == "/user/register" || ctx.Request.URL.Path == "/user/login" {
			ctx.Next()
			return
		}

		// Ambil cookie session_token
		cookie, err := ctx.Cookie("session_token")
		if err != nil {
			handleCookieError(ctx, err)
			return
		}

		// Parse token JWT dengan klaim
		tokenClaims := model.Claims{}
		token, err := jwt.ParseWithClaims(cookie, &tokenClaims, func(token *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})
		if err != nil || !token.Valid {
			handleTokenError(ctx, err)
			return
		}

		// Set Email ke context jika valid
		if tokenClaims.Email == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: invalid user email",
			})
			return
		}

		// Set user email in context
		ctx.Set("email", tokenClaims.Email)

		// Panggil middleware berikutnya
		ctx.Next()
	})
}

// handleCookieError menangani error saat membaca cookie
// func handleCookieError(ctx *gin.Context, err error) {
// 	if err == http.ErrNoCookie {
// 		ctx.Redirect(http.StatusSeeOther, "/login")
// 		return
// 	}

// 	// if err == http.ErrNoCookie {
// 	// 	// Kembalikan status 401 jika cookie tidak ada
// 	// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: no session token"})
// 	// 	return
// 	// }
// 	ctx.JSON(http.StatusBadRequest, gin.H{
// 		"error": "Bad Request",
// 	})
// 	ctx.Abort()
// }

func handleCookieError(ctx *gin.Context, err error) {
	if err == http.ErrNoCookie {
		// Redirect ke /login jika cookie tidak ditemukan
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}
	// Jika bukan error karena cookie, kembalikan respons Bad Request
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": "Bad Request",
	})
	ctx.Abort()
}



// // handleTokenError menangani error saat memvalidasi token
// func handleTokenError(ctx *gin.Context, err error) {
// 	if err == jwt.ErrSignatureInvalid {
// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 			"error": "Unauthorized: invalid signature",
// 		})
// 		return
// 	}
// 	ctx.JSON(http.StatusBadRequest, gin.H{
// 		"error": "Bad Request",
// 	})
// 	ctx.Abort()
// }

func handleTokenError(ctx *gin.Context, err error) {
	if err == jwt.ErrSignatureInvalid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized: invalid signature",
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": "Bad Request",
	})
	ctx.Abort()
}


// package middleware

// import (
// 	"a21hc3NpZ25tZW50/model"
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt"
// )

// // Auth middleware untuk memvalidasi token JWT
// func Auth() gin.HandlerFunc {
// 	return gin.HandlerFunc(func(ctx *gin.Context) {
// 		// Abaikan middleware untuk rute tertentu
// 		if ctx.Request.URL.Path == "/user/register" || ctx.Request.URL.Path == "/user/login" {
// 			ctx.Next()
// 			return
// 		}

// 		// Ambil cookie session_token
// 		cookie, err := ctx.Cookie("session_token")
// 		if err != nil {
// 			handleCookieError(ctx, err)
// 			return
// 		}

// 		// Parse token JWT dengan klaim
// 		tokenClaims := model.Claims{}
// 		token, err := jwt.ParseWithClaims(cookie, &tokenClaims, func(token *jwt.Token) (interface{}, error) {
// 			return model.JwtKey, nil
// 		})
// 		if err != nil || !token.Valid {
// 			handleTokenError(ctx, err)
// 			return
// 		}

// 		// Set Email ke context jika valid
// 		if tokenClaims.Email == "" {
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error": "Unauthorized: invalid user email",
// 			})
// 			return
// 		}

// 		// Set user email in context
// 		ctx.Set("email", tokenClaims.Email)

// 		// Panggil middleware berikutnya
// 		ctx.Next()
// 	})
// }

// // handleCookieError menangani error saat membaca cookie
// func handleCookieError(ctx *gin.Context, err error) {
// 	if err == http.ErrNoCookie {
// 		// Mengembalikan status 303 jika cookie tidak ada
// 		ctx.Redirect(http.StatusSeeOther, "/login")
// 		return
// 	}
// 	ctx.JSON(http.StatusBadRequest, gin.H{
// 		"error": "Bad Request",
// 	})
// 	ctx.Abort()
// }

// // handleTokenError menangani error saat memvalidasi token
// func handleTokenError(ctx *gin.Context, err error) {
// 	if err == jwt.ErrSignatureInvalid {
// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 			"error": "Unauthorized: invalid signature",
// 		})
// 		return
// 	}
// 	ctx.JSON(http.StatusBadRequest, gin.H{
// 		"error": "Bad Request",
// 	})
// 	ctx.Abort()
// }