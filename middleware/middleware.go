package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/triwira-joel/technical-test-bosshire/helper"
)

func JwtAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			t := strings.Split(authHeader, " ")
			if len(t) == 2 {
				authToken := t[1]
				authorized, err := helper.IsAuthorized(authToken)
				if err != nil {
					c.JSON(http.StatusInternalServerError, err.Error())
					return err
				}
				if authorized {
					id, _, err := helper.ExtractDataFromToken(authToken)
					if err != nil {
						c.JSON(http.StatusUnauthorized, err.Error())
						return err
					}
					c.Set("x-user-id", id)
					next(c)
					return nil
				}
			}
			c.JSON(http.StatusUnauthorized, "Not Authorized")
			return nil
		}
	}
}

func JwtTalentAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			t := strings.Split(authHeader, " ")
			if len(t) == 2 {
				authToken := t[1]
				authorized, err := helper.IsAuthorized(authToken)
				if err != nil {
					c.JSON(http.StatusInternalServerError, err.Error())
					return err
				}
				if authorized {
					id, role, err := helper.ExtractDataFromToken(authToken)
					if err != nil {
						c.JSON(http.StatusUnauthorized, err.Error())
						return err
					}
					if role != "TALENT" {
						c.JSON(http.StatusUnauthorized, "Not Authorized")
						return nil
					}
					c.Set("x-user-id", id)
					next(c)
					return nil
				}
			}
			c.JSON(http.StatusUnauthorized, "Not Authorized")
			return nil
		}
	}
}

func JwtEmployerAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			t := strings.Split(authHeader, " ")
			if len(t) == 2 {
				authToken := t[1]
				authorized, err := helper.IsAuthorized(authToken)
				if err != nil {
					c.JSON(http.StatusInternalServerError, err.Error())
					return err
				}
				if authorized {
					id, role, err := helper.ExtractDataFromToken(authToken)
					if err != nil {
						c.JSON(http.StatusUnauthorized, err.Error())
						return err
					}
					if role != "EMPLOYER" {
						c.JSON(http.StatusUnauthorized, "Not Authorized")
						return nil
					}
					c.Set("x-user-id", id)
					next(c)
					return nil
				}
			}
			c.JSON(http.StatusUnauthorized, "Not Authorized")
			return nil
		}
	}
}
