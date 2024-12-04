# Snippetbox

**Update: finished Sep 28, 2024**

App developed by Alex Edwards in his book "Let's Go".

## How to use it

```sh
go run mage.go
```

or

```sh
mage
```

### Potential problems

1. TLS dir missing.

```bash
mkdir tls && cd tls
go run /nix/store/gdp1f96y3rhrpncllbngi1chmpc9k0gd-go-1.22.6/share/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

Note that you may have to replace the path of `generate_cert.go`.

