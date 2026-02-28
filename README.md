# hashpass

Command-line tool to generate and validate SHA-512 salted password hashes.

## Usage

Generate a hash:
```
hashpass -g <password>
```

Validate a password against a hash:
```
hashpass -v <password> <hash>
```

## Examples
```
$ hashpass -g mypassword
String: mypassword
Hash: abc123...==

$ hashpass -v mypassword abc123...==
Password valido

$ hashpass -v wrongpassword abc123...==
Password invalido
```

## How it works

The hash is generated as follows:

1. A random salt of 4 to 7 non-zero bytes is generated.
2. The salt is appended to the UTF-8 encoded password.
3. SHA-512 is applied to the combined bytes.
4. The resulting hash is concatenated with the salt.
5. The final value is Base64 encoded.

## Build

Windows:
```
go build -o hashpass.exe .
```

Linux / macOS:
```
go build -o hashpass .
```

Cross-compile for Windows from Linux/macOS:
```
GOOS=windows GOARCH=amd64 go build -o hashpass.exe .
```

## Requirements

Go 1.18 or higher.