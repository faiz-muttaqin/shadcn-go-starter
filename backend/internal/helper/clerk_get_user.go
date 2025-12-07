package helper

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/database"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/model"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/kvstore"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/types"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ClerkGetUser(c *gin.Context) (*model.User, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("authorization header missing")
	}

	// Expect header like: "Bearer <token>"
	token := ""
	if t, ok := strings.CutPrefix(authHeader, "Bearer "); ok {
		token = t
	}
	if token == "" {
		return nil, fmt.Errorf("invalid token")
	}

	ctx := c.Request.Context()
	userID, err := kvstore.GetKey(token)
	if err != nil {
		claims, ok := clerk.SessionClaimsFromContext(ctx)
		if !ok {
			logrus.Println("OK?", ok)
			logrus.Println("CTX:", c.Request.Context())
			logrus.Println("CTX:", ctx)
			logrus.Println("token")
			logrus.Println(token)
			return nil, fmt.Errorf("unauthorized")
		}
		userID = claims.Subject
		kvstore.SetKey(token, userID, 5*time.Minute)
	}
	usr, err := CachedClerkUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Safely extract email
	email := ""
	if len(usr.EmailAddresses) > 0 {
		email = usr.EmailAddresses[0].EmailAddress
	}

	// Safely check email verification
	verificationStatus := "unverified"
	if len(usr.ExternalAccounts) > 0 && usr.ExternalAccounts[0].Verification != nil {
		verificationStatus = usr.ExternalAccounts[0].Verification.Status
	}

	// Safely extract optional fields with nil checks
	name := ""
	if usr.Username != nil {
		name = *usr.Username
	}

	firstName := ""
	if usr.FirstName != nil {
		firstName = *usr.FirstName
	}

	lastName := ""
	if usr.LastName != nil {
		lastName = *usr.LastName
	}

	picture := ""
	if usr.ImageURL != nil {
		picture = *usr.ImageURL
	}
	if len(superUserEmails) == 0 {
		superUserEmails = strings.Split(util.Getenv("superUserEmails", "muttaqinfaiz@gmail.com"), ",")
	}
	isSuperUser := util.Contains(superUserEmails, email)

	var existingUser model.User
	err = database.DB.Preload("Role.AbilityRules").Where("email = ?", email).First(&existingUser).Error
	if err == nil {
		if isSuperUser && existingUser.RoleID != 1 {
			existingUser.RoleID = 1
			existingUser.Status = "active"
			database.DB.Save(&existingUser)
		}
		return &existingUser, nil
	}
	usrNew := model.User{
		ExternalID:         usr.ID,
		Avatar:             types.Avatar(picture),
		Email:              types.Email(email),
		FirstName:          firstName,
		LastName:           lastName,
		Username:           name,
		VerificationStatus: verificationStatus,
		Status:             "inactive",
		RoleID:             2,
	}
	if isSuperUser {
		usrNew.RoleID = 1
		usrNew.Status = "active"
	}
	if err = database.DB.Preload("Role.AbilityRules").Where("email = ?", email).FirstOrCreate(&usrNew).Error; err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}
	return &usrNew, nil
}

var clerkUserCache sync.Map

func CachedClerkUser(ctx context.Context, id string) (*clerk.User, error) {
	if user, ok := clerkUserCache.Load(id); ok {
		return user.(*clerk.User), nil
	}

	user, err := user.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	clerkUserCache.Store(id, user)
	return user, nil
}
