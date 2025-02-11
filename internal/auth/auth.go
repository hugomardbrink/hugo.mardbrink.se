package auth

import (
    "os"
    "crypto/sha512"
    "crypto/subtle"
    "encoding/hex"
    "net/http"
    "log"

    "github.com/labstack/echo/v4"
)

const (
    envUsernameHash = "HUGOMARDBRINK_USERNAME"
    envPasswordHash = "HUGOMARDBRINK_PASSWORD"
)

func BasicAuth(next echo.HandlerFunc) echo.HandlerFunc {
    return echo.HandlerFunc(func(c echo.Context) error {
        username, password, ok := c.Request().BasicAuth()

        if !ok {
            c.Response().Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
            return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
        }
 
        usernameHash := sha512.Sum512([]byte(username)) 
        passwordHash := sha512.Sum512([]byte(password))
    
        expectedUsernameHash, err := hex.DecodeString(os.Getenv(envUsernameHash))
        if err != nil { 
            log.Fatal(err) 
        }
        expectedPasswordHash, err := hex.DecodeString(os.Getenv(envPasswordHash))
        if err != nil {
            log.Fatal(err)
        }

        usernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1
        passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1

        if usernameMatch && passwordMatch {
            return next(c)
        } else {
            c.Response().Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
            return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
        }
    })
}
