# Telegram Login Widget

## Table of Content

- [Examples](#examples)
  - [JSON](#json)
  - [NewAuthorizationDataFromURL](#newauthorizationdatafromurl)
  - [NewAuthorizationDataFromURLValues](#newauthorizationdatafromurlvalues)
- [FAQ](#faq)
- [License](#license)
- [Links](#links)

## Examples

#### JSON

```go
package main

import (
	"net/http"
	"time"

	http_request_body_json_decoder "entrlcom.dev/http-request-body-json-decoder"
	telegram_login_widget "entrlcom.dev/telegram-login-widget"
	"entrlcom.dev/unit"
)

const limit = unit.B * 512 // 512 B.

// TODO: Update values.
const (
	timeout = time.Second * 10
	token = "XXXXXXXX:XXXXXXXXXXXXXXXXXXXXXXXX"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var authorizationData telegram_login_widget.AuthorizationData

	// Decode HTTP request body to authorization data struct.
	if err := http_request_body_json_decoder.Decode(w, r, &authorizationData, limit); err != nil {
		// TODO: Handle error.
		return
	}

	// Validate authorization data.
	if err := authorizationData.Validate(token); err != nil {
		// TODO: Handle error.
		return
	}
	
	// Check whether authorization data is expired.
	if authorizationData.IsExpired(timeout) {
		// TODO: Handle.
		return
	}

	// ...
}

```

#### NewAuthorizationDataFromURL

```go
package main

import (
	"net/http"
	"time"

	telegram_login_widget "entrlcom.dev/telegram-login-widget"
)

// TODO: Update values.
const (
	timeout = time.Second * 10
	token = "XXXXXXXX:XXXXXXXXXXXXXXXXXXXXXXXX"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authorizationData, err := telegram_login_widget.NewAuthorizationDataFromURL(r.URL.String())
	if err != nil {
		// TODO: Handle error.
		return
	}

	// Validate authorization data.
	if err = authorizationData.Validate(token); err != nil {
		// TODO: Handle error.
		return
	}
	
	// Check whether authorization data is expired.
	if authorizationData.IsExpired(timeout) {
		// TODO: Handle.
		return
	}

	// ...
}

```

#### NewAuthorizationDataFromURLValues

```go
package main

import (
	"net/http"
	"time"

	telegram_login_widget "entrlcom.dev/telegram-login-widget"
)

// TODO: Update values.
const (
	timeout = time.Second * 10
	token = "XXXXXXXX:XXXXXXXXXXXXXXXXXXXXXXXX"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authorizationData, err := telegram_login_widget.NewAuthorizationDataFromURLValues(r.URL.Query())
	if err != nil {
		// TODO: Handle error.
		return
	}

	// Validate authorization data.
	if err = authorizationData.Validate(token); err != nil {
		// TODO: Handle error.
		return
	}
	
	// Check whether authorization data is expired.
	if authorizationData.IsExpired(timeout) {
		// TODO: Handle.
		return
	}

	// ...
}

```

## FAQ

#### What is Telegram Login Widget?

The Telegram login widget is a simple way to authorize users on your website.
Check out [this page](https://core.telegram.org/widgets/login) for a general overview of the widget.

#### How to validate hash?

Call `Validate` method on `AuthorizationData` struct to validate hash.

#### How to prevent the use of outdated authentication data?

Call `IsExpired` method on `AuthorizationData` struct to check whether authentication data is expired.

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Links

* [Telegram Login Widget](https://core.telegram.org/widgets/login)
