# jwt-cli

A simple command-line tool I built in Go to decode JWT tokens. It takes a JWT string, splits it into its three parts, decodes the header and payload, and checks whether the token has expired.

This was part of my Go fundamentals lab (Lab II.1) — the goal was to get hands-on with Go basics like slices, error handling, JSON unmarshalling, and writing unit tests, using a JWT decoder as the practical exercise.

## What it does

- Takes a JWT token as a command-line argument
- Splits it into header, payload, and signature
- Base64url-decodes the header and payload
- Pretty-prints both as JSON
- Reads the `exp` claim and tells you if the token is expired or still valid

Note: this tool only decodes the token, it doesn't verify the signature. That's a separate concern (would need the signing secret/key) and was out of scope for this lab.

## Installation

Clone the repo and build it:

```bash
git clone https://github.com/kssreelakshmi/jwt-cli.git
cd jwt-cli
go build -o jwt-cli
```

Or just run it directly without building a binary:

```bash
go run main.go <token>
```

## Usage

```bash
./jwt-cli <your-jwt-token>
```

If you don't pass a token, it'll show a usage message:

```
Usage: jwt-cli <token>
```

## Example output

Running it against a token with a past expiry:

```
$ go run main.go eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJleHAiOjE3MDAwMDAwMDB9.YOUR_SIGNATURE_HERE

=== HEADER ===
{
  "alg": "HS256",
  "typ": "JWT"
}
=== PAYLOAD ===
{
  "exp": 1700000000,
  "iat": 1516239022,
  "name": "John Doe",
  "sub": "1234567890"
}
=== STATUS ===
TOKEN EXPIRED (expired at: 2023-11-15 03:43:20 +0530 IST)
```

And with a token that's still valid, the status section instead shows:

```
=== STATUS ===
TOKEN VALID (expires at: 2033-05-18 09:03:20 +0530 IST)
```

## Running tests

```bash
go test -v ./...
```

This runs 4 unit tests covering:
- A valid (non-expired) token
- An expired token
- Successfully decoding a valid base64url segment
- Handling an invalid/malformed segment

![go test output](./test-output.png)
## What I learned

This was my first real Go project. Coming from Python, the things that stood out most were Go's explicit error handling (no exceptions, just checking `err != nil` everywhere), how JSON numbers always unmarshal as `float64` (so I had to cast `exp` to `int64` before using `time.Unix()`), and that JWT uses unpadded base64url encoding, not standard base64 — I had to use `base64.RawURLEncoding` instead of `base64.StdEncoding`.