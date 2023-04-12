package authmodel

import (
	"github.com/dinhlockt02/cs_video_call_app_server/components/tokenprovider"
)

type AuthToken struct {
	AccessToken    *tokenprovider.Token `json:"access_token"`
	RefreshToken   *tokenprovider.Token `json:"refresh_token"`
	EmailVerified  bool                 `json:"email_verified"`
	ProfileUpdated bool                 `json:"profile_updated"`
}
