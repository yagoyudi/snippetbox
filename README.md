# Snippetbox

App developed by Alex Edwards in his book "Let's Go".

## How to use it

At the moment:

```go
go run ./cmd/web
```

### Potential problems

1. TLS dir missing.

```bash
mkdir tls && cd tls
go run /nix/store/gdp1f96y3rhrpncllbngi1chmpc9k0gd-go-1.22.6/share/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

Note that you may have to replace the path of `generate_cert.go`.

## Routes setup

| Method | Pattern                 | Handler          | Action                                         |
|--------|-------------------------|------------------|------------------------------------------------|
| GET    | /                       | home             | Display the home page                          |
| GET    | /snippet/view/:id       | snippetView      | Display a specific snippet                     |
| GET    | /snippet/create         | snippetCreate    | Display a HTML form for creating a new snippet |
| POST   | /snippet/create         | snippetCreatePost| Create a new snippet                           |
| GET    | /user/signup            | userSignup       | Display a HTML form for signing up a new user  |
| POST   | /user/signup            | userSignupPost   | Create a new user                              |
| GET    | /user/login             | userLogin        | Display a HTML form for logging in a user      |
| POST   | /user/login             | userLoginPost    | Authenticate and login the user                |
| POST   | /user/logout            | userLogoutPost   | Logout the user                                |
| GET    | /static/*filepath       | http.FileServer  | Serve a specific static file                   |

