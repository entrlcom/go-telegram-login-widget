package telegram_login_widget

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var (
	ErrInvalidHash              = errors.New("invalid hash")
	ErrInvalidAuthorizationData = errors.New("invalid telegram authorization data")
)

const (
	KeyAuthDate  = "auth_date"
	KeyFirstName = "first_name"
	KeyHash      = "hash"
	KeyID        = "id"
	KeyLastName  = "last_name"
	KeyPhotoURL  = "photo_url"
	KeyUsername  = "username"
)

type AuthorizationData struct { //nolint:govet // OK.
	AuthDate  int64  `json:"auth_date"`
	FirstName string `json:"first_name"`
	Hash      string `json:"hash"`
	ID        int64  `json:"id"`
	LastName  string `json:"last_name"`
	PhotoURL  string `json:"photo_url"`
	Username  string `json:"username"`
}

func (x AuthorizationData) GetAuthenticationTimestamp() time.Time {
	return time.Unix(x.AuthDate, 0)
}

func (x AuthorizationData) IsExpired(d time.Duration) bool {
	now := time.Now().UTC()
	t := x.GetAuthenticationTimestamp()

	return t.After(now) || t.Add(d).Before(now)
}

func (x AuthorizationData) Validate(token string) error {
	hash, err := hex.DecodeString(x.Hash)
	if err != nil {
		return errors.Join(err, ErrInvalidHash)
	}

	key := sha256.Sum256([]byte(token))

	h := hmac.New(sha256.New, key[:])
	if _, err = h.Write([]byte(x.dataCheckString())); err != nil {
		return errors.Join(err, ErrInvalidHash)
	}

	if subtle.ConstantTimeCompare(h.Sum(nil), hash) != 1 {
		return ErrInvalidHash
	}

	return nil
}

func (x AuthorizationData) dataCheckString() string {
	var b strings.Builder

	b.WriteString(KeyAuthDate + "=" + strconv.FormatInt(x.AuthDate, 10)) //nolint:errcheck // OK.

	if len(x.FirstName) != 0 {
		b.WriteString("\n" + KeyFirstName + "=" + x.FirstName) //nolint:errcheck // OK.
	}

	b.WriteString("\n" + KeyID + "=" + strconv.FormatInt(x.ID, 10)) //nolint:errcheck // OK.

	if len(x.LastName) != 0 {
		b.WriteString("\n" + KeyLastName + "=" + x.LastName) //nolint:errcheck // OK.
	}

	if len(x.PhotoURL) != 0 {
		b.WriteString("\n" + KeyPhotoURL + "=" + x.PhotoURL) //nolint:errcheck // OK.
	}

	if len(x.Username) != 0 {
		b.WriteString("\n" + KeyUsername + "=" + x.Username) //nolint:errcheck // OK.
	}

	return b.String()
}

func NewAuthorizationDataFromURL(uri string) (AuthorizationData, error) {
	u, err := url.ParseRequestURI(uri)
	if err != nil {
		return AuthorizationData{}, err
	}

	return NewAuthorizationDataFromURLValues(u.Query())
}

func NewAuthorizationDataFromURLValues(v url.Values) (AuthorizationData, error) {
	authDate, err := strconv.ParseInt(v.Get(KeyAuthDate), 10, 64)
	if err != nil {
		return AuthorizationData{}, errors.Join(err, ErrInvalidAuthorizationData)
	}

	id, err := strconv.ParseInt(v.Get(KeyID), 10, 64)
	if err != nil {
		return AuthorizationData{}, errors.Join(err, ErrInvalidAuthorizationData)
	}

	x := AuthorizationData{
		AuthDate:  authDate,
		FirstName: v.Get(KeyFirstName),
		Hash:      v.Get(KeyHash),
		ID:        id,
		LastName:  v.Get(KeyLastName),
		PhotoURL:  v.Get(KeyPhotoURL),
		Username:  v.Get(KeyUsername),
	}

	return x, nil
}
