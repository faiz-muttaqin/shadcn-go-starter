package helper

import (
	"context"
	"fmt"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/database"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/model"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/types"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/util"
	"github.com/gin-gonic/gin"
)

var FirebaseAuth *auth.Client

type FirebaseAuthData struct {
	UID           string  `json:"uid"`
	AuthTime      float64 `json:"auth_time"`
	Email         string  `json:"email"`
	EmailVerified bool    `json:"email_verified"`
	Firebase      struct {
		Identities     map[string][]string `json:"identities"`
		SignInProvider string              `json:"sign_in_provider"`
	} `json:"firebase"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
	UserID  string `json:"user_id"`
}

var superUserEmails []string

func GetFirebaseUser(c *gin.Context) (*model.User, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("authorization header missing")
	}
	// Expect header like: "Bearer <token>"
	idToken := ""
	if t, ok := strings.CutPrefix(authHeader, "Bearer "); ok && t != "" {
		idToken = t
	} else {
		return nil, fmt.Errorf("invalid token")
	}

	token, err := FirebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, err
	}

	claims := token.Claims
	firebaseAuthData := &FirebaseAuthData{}

	// Map claims to struct
	if v, ok := claims["auth_time"].(float64); ok {
		firebaseAuthData.AuthTime = v
	}
	if v, ok := claims["email"].(string); ok {
		firebaseAuthData.Email = v
	}
	if v, ok := claims["email_verified"].(bool); ok {
		firebaseAuthData.EmailVerified = v
	}
	if v, ok := claims["firebase"].(map[string]interface{}); ok {
		if identities, ok := v["identities"].(map[string]interface{}); ok {
			firebaseAuthData.Firebase.Identities = make(map[string][]string)
			for k, arr := range identities {
				if arrSlice, ok := arr.([]interface{}); ok {
					strSlice := make([]string, len(arrSlice))
					for i, s := range arrSlice {
						strSlice[i], _ = s.(string)
					}
					firebaseAuthData.Firebase.Identities[k] = strSlice
				}
			}
		}
		if provider, ok := v["sign_in_provider"].(string); ok {
			firebaseAuthData.Firebase.SignInProvider = provider
		}
	}
	if v, ok := claims["name"].(string); ok {
		firebaseAuthData.Name = v
	}
	if v, ok := claims["picture"].(string); ok {
		firebaseAuthData.Picture = v
	}
	if v, ok := claims["user_id"].(string); ok {
		firebaseAuthData.UserID = v
	}
	firebaseAuthData.UID = token.UID

	if len(superUserEmails) == 0 {
		superUserEmails = strings.Split(util.Getenv("SUPER_ADMIN_EMAILS", "muttaqinfaiz@gmail.com"), ",")
	}
	isSuperUser := util.Contains(superUserEmails, firebaseAuthData.Email)

	var existingUser model.User
	err = database.DB.Preload("UserRole.AbilityRules").Where("email = ?", firebaseAuthData.Email).First(&existingUser).Error
	if err == nil {
		if isSuperUser && existingUser.RoleID != 1 {
			existingUser.RoleID = 1
			existingUser.Status = "active"
			database.DB.Save(&existingUser)
		}
		return &existingUser, nil
	}
	verificationStatus := "unverified"
	if firebaseAuthData.EmailVerified {
		verificationStatus = "verified"
	}
	usrNew := model.User{
		ExternalID:         firebaseAuthData.UID,
		Avatar:             types.Avatar(firebaseAuthData.Picture),
		Email:              types.Email(firebaseAuthData.Email),
		FirstName:          firebaseAuthData.Name,
		LastName:           firebaseAuthData.Name,
		Username:           firebaseAuthData.Name,
		VerificationStatus: verificationStatus,
		Status:             "inactive",
		RoleID:             2,
	}
	if isSuperUser {
		usrNew.RoleID = 1
		usrNew.Status = "active"
	}
	if err = database.DB.Preload("UserRole.AbilityRules").Where("email = ?", firebaseAuthData.Email).FirstOrCreate(&usrNew).Error; err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}
	return &usrNew, nil
}
