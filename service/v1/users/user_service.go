package users

import (
	"crypto/sha256"
	"encoding/hex"
	"go-deploy/dto/v1/body"
	"go-deploy/models/model"
	"go-deploy/pkg/db/resources/user_repo"
	"go-deploy/service/utils"
	"go-deploy/service/v1/users/opts"
	"net/http"
	"strings"
	"time"
)

// Get gets a user
//
// It uses service.AuthInfo to only return the model the requesting user has access to
func (c *Client) Get(id string, opts ...opts.GetOpts) (*model.User, error) {
	_ = utils.GetFirstOrDefault(opts)

	if c.V1.Auth() != nil && id != c.V1.Auth().UserID && !c.V1.Auth().IsAdmin {
		return nil, nil
	}

	return c.User(id, user_repo.New())
}

// List lists users
//
// It uses service.AuthInfo to only return the resources the requesting user has access to
// It uses the search param to enable searching in multiple fields
func (c *Client) List(opts ...opts.ListOpts) ([]model.User, error) {
	o := utils.GetFirstOrDefault(opts)

	umc := user_repo.New()

	if o.Pagination != nil {
		umc.WithPagination(o.Pagination.Page, o.Pagination.PageSize)
	}

	if o.Search != nil {
		umc.WithSearch(*o.Search)
	}

	if c.V1.Auth() != nil && !c.V1.Auth().IsAdmin || !o.All {
		user, err := umc.GetByID(c.V1.Auth().UserID)
		if err != nil {
			return nil, err
		}

		return []model.User{*user}, nil
	}

	return c.Users(umc)
}

// Exists checks if a user exists
//
// This does not use AuthInfo
func (c *Client) Exists(id string) (bool, error) {
	return user_repo.New().ExistsByID(id)
}

// Synchronize creates a user or updates an existing user.
// It does nothing if no auth info is provided
func (c *Client) Synchronize() (*model.User, error) {
	if !c.V1.HasAuth() {
		return nil, nil
	}

	roleNames := make([]string, len(c.V1.Auth().Roles))
	for i, role := range c.V1.Auth().Roles {
		roleNames[i] = role.Name
	}

	effectiveRole := c.V1.Auth().GetEffectiveRole()

	params := &model.UserCreateParams{
		Username:  c.V1.Auth().GetUsername(),
		FirstName: c.V1.Auth().GetFirstName(),
		LastName:  c.V1.Auth().GetLastName(),
		Email:     c.V1.Auth().GetEmail(),
		IsAdmin:   c.V1.Auth().IsAdmin,
		EffectiveRole: &model.EffectiveRole{
			Name:        effectiveRole.Name,
			Description: effectiveRole.Description,
		},
	}

	umc := user_repo.New()

	user, err := umc.Synchronize(c.V1.Auth().UserID, params)
	if err != nil {
		return nil, err
	}

	if user.Gravatar.FetchedAt.IsZero() || user.Gravatar.FetchedAt.Add(model.FetchGravatarInterval).Before(time.Now()) {
		gravatarURL, err := c.FetchGravatar(user.ID)
		if err != nil {
			return nil, err
		}

		if gravatarURL == nil {
			err = umc.UnsetGravatar(user.ID)
			if err != nil {
				return nil, err
			}
		} else {
			err = umc.SetGravatar(user.ID, *gravatarURL)
			if err != nil {
				return nil, err
			}
		}
	}

	return user, nil
}

// Discover returns a list of users that the requesting user has access to
//
// It uses search param to enable searching in multiple fields
func (c *Client) Discover(opts ...opts.DiscoverOpts) ([]body.UserReadDiscovery, error) {
	o := utils.GetFirstOrDefault(opts)
	umc := user_repo.New()

	if o.Search != nil {
		umc.WithSearch(*o.Search)
	}

	if o.Pagination != nil {
		umc.WithPagination(o.Pagination.Page, o.Pagination.PageSize)
	}

	users, err := c.Users(umc)
	if err != nil {
		return nil, err
	}

	var usersRead []body.UserReadDiscovery
	for _, user := range users {
		if c.V1.Auth() != nil && user.ID == c.V1.Auth().UserID {
			continue
		}

		usersRead = append(usersRead, body.UserReadDiscovery{
			ID:        user.ID,
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		})
	}

	return usersRead, nil
}

// Update updates a user
//
// It uses service.AuthInfo to only update the model the requesting user has access to
func (c *Client) Update(userID string, dtoUserUpdate *body.UserUpdate) (*model.User, error) {
	umc := user_repo.New()

	if c.V1.Auth() != nil && userID != c.V1.Auth().UserID && !c.V1.Auth().IsAdmin {
		return nil, nil
	}

	userUpdate := model.UserUpdateParams{}.FromDTO(dtoUserUpdate)

	err := umc.UpdateWithParams(userID, &userUpdate)
	if err != nil {
		return nil, err
	}

	return c.RefreshUser(userID, umc)
}

// FetchGravatar checks if the user has a gravatar image and fetches it if it exists
// If the user does not have a gravatar image, it returns nil
func (c *Client) FetchGravatar(userID string) (*string, error) {
	umc := user_repo.New()

	if c.V1.Auth() != nil && userID != c.V1.Auth().UserID && !c.V1.Auth().IsAdmin {
		return nil, nil
	}

	user, err := c.User(userID, umc)
	if err != nil {
		return nil, err
	}

	hasher := sha256.Sum256([]byte(strings.TrimSpace(user.Email)))
	hash := hex.EncodeToString(hasher[:])

	gravatarURL := "https://www.gravatar.com/avatar/" + hash + "?d=404"

	// Check if image exists
	resp, err := http.Head(gravatarURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	// Trim the query string
	gravatarURL = gravatarURL[:strings.Index(gravatarURL, "?")]

	return &gravatarURL, nil
}
