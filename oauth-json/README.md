# Introduction

In this example we authenticate against the Google AI Platform API using a service account, reading the account's private key from a file. This isn't the best way to authenticate against the AI platform; see the oauth-default example.

## Creating a service account

Create a service account using the `IAM and admin` section of the Google Cloud console. This service account needs to have the following roles:

* Project: `viewer`
* Storage Object: `Admin`
* ML Engine: `Admin`

Once you've created the service account, download the key JSON file.

## Getting an OAuth client in Go

The Google OAuth client is documented [here](https://godoc.org/golang.org/x/oauth2/google). Here's the general idea:

```
import (
	"google.golang.org/api/ml/v1"
	"log"
	"net/http"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"
	"golang.org/x/oauth2/google"
)

...

conf := &jwt.Config{
    Email: "<your service account>@<your project name>.iam.gserviceaccount.com",
    PrivateKey: []byte("-----BEGIN PRIVATE KEY-----\n......\n-----END PRIVATE KEY-----\n"),
    Scopes: []string{"https://www.googleapis.com/auth/cloud-platform"},
    TokenURL: google.JWTTokenURL,
}
client := conf.Client(oauth2.NoContext)
```

* `Email` needs to match the address of your service account
* The private key is the base-64 encoded private key for the service account. Obviously, you should load this from an environment variable or file. *Never hard-code private keys and never add them to version control*
* The `cloud-platform` scope is required - see https://developers.google.com/identity/protocols/googlescopes#mlv1
