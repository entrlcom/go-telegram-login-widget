package telegram_login_widget_test

import (
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	telegram_login_widget "entrlcom.dev/telegram-login-widget"
)

func Test(t *testing.T) {
	t.Parallel()

	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Telegram Login Widget Test Suite")
}

const (
	firstName = "Pavel"
	hash      = "4254bed5d5d1af83166a1e08c1270296f1777af60a69f611a9f1a8d8cf856f40"
	id        = 1
	lastName  = "Durov"
	photoURL  = "https://t.me/i/userpic/320/0000000000000000000000000000000000000000000.jpg"
	username  = "durov"
)

const token = "XXXXXXXX:XXXXXXXXXXXXXXXXXXXXXXXX"

var authDate = time.Date(1984, time.October, 10, 10, 20, 0, 0, time.UTC) //nolint:gochecknoglobals // OK.

var _ = ginkgo.Describe("NewAuthorizationDataFromURL", func() {
	ginkgo.It("", func() {
		v := newURLValues()
		uri := "https://example.com/?" + v.Encode()
		authorizationData, err := telegram_login_widget.NewAuthorizationDataFromURL(uri)
		gomega.Expect(err).To(gomega.Succeed())
		gomega.Expect(authorizationData).To(gomega.Equal(newAuthorizationData()))
	})
})

var _ = ginkgo.Describe("NewAuthorizationDataFromURLValues", func() {
	ginkgo.It("", func() {
		v := newURLValues()
		authorizationData, err := telegram_login_widget.NewAuthorizationDataFromURLValues(v)
		gomega.Expect(err).To(gomega.Succeed())
		gomega.Expect(authorizationData).To(gomega.Equal(newAuthorizationData()))
	})
})

var _ = ginkgo.Describe("Validate", func() {
	ginkgo.It("ok", func() {
		gomega.Expect(newAuthorizationData().Validate(token)).To(gomega.Succeed())
	})

	ginkgo.It("error invalid hash (empty hash)", func() {
		authorizationData := newAuthorizationData()
		authorizationData.Hash = ""
		gomega.Expect(authorizationData.Validate(token)).To(gomega.MatchError(telegram_login_widget.ErrInvalidHash))
	})

	ginkgo.It("error invalid hash", func() {
		authorizationData := newAuthorizationData()
		authorizationData.Hash = "$(invalid hash)"
		gomega.Expect(authorizationData.Validate(token)).To(gomega.MatchError(telegram_login_widget.ErrInvalidHash))
	})
})

var _ = ginkgo.Describe("IsExpired", func() {
	ginkgo.It("not expired", func() {
		authorizationData := telegram_login_widget.AuthorizationData{ //nolint:exhaustruct // OK.
			AuthDate: time.Now().UTC().Unix(),
		}
		gomega.Expect(authorizationData.IsExpired(time.Second * 10)).To(gomega.BeFalse())
	})

	ginkgo.It("expired (future)", func() {
		authorizationData := telegram_login_widget.AuthorizationData{ //nolint:exhaustruct // OK.
			AuthDate: time.Now().Add(time.Second * 10).UTC().Unix(),
		}
		gomega.Expect(authorizationData.IsExpired(time.Second * 10)).To(gomega.BeTrue())
	})

	ginkgo.It("expired", func() {
		authorizationData := telegram_login_widget.AuthorizationData{ //nolint:exhaustruct // OK.
			AuthDate: authDate.Unix(),
		}
		gomega.Expect(authorizationData.IsExpired(time.Second * 10)).To(gomega.BeTrue())
	})
})

func newAuthorizationData() telegram_login_widget.AuthorizationData {
	return telegram_login_widget.AuthorizationData{
		AuthDate:  authDate.Unix(),
		FirstName: firstName,
		Hash:      hash,
		ID:        id,
		LastName:  lastName,
		PhotoURL:  photoURL,
		Username:  username,
	}
}

func newURLValues() url.Values {
	v := url.Values{}
	v.Set(telegram_login_widget.KeyAuthDate, strconv.FormatInt(authDate.Unix(), 10))
	v.Set(telegram_login_widget.KeyFirstName, firstName)
	v.Set(telegram_login_widget.KeyHash, hash)
	v.Set(telegram_login_widget.KeyID, strconv.FormatInt(id, 10))
	v.Set(telegram_login_widget.KeyLastName, lastName)
	v.Set(telegram_login_widget.KeyPhotoURL, photoURL)
	v.Set(telegram_login_widget.KeyUsername, username)

	return v
}
