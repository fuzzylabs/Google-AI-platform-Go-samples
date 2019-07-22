# Introduction

Google's [AI platform](https://cloud.google.com/ai-platform) brings a number of tools together to enable you to create end-to-end machine learning pipelines. This repository provides an example of using the platform to train a recommendation engine using the Go programming language.

## Motivation

At the time of writing, the AI platform is relatively new and as a result the documentation is incomplete, particularly with respect to Go. The AI platform is not part of the [Google Cloud SDK](https://cloud.google.com/sdk). Instead, we need to use the [Google APIs client](https://github.com/googleapis/google-api-go-client) which is a bit 'lower level', in that it just exposes Google's various REST APIs in raw form. Documentation for [using the AI platform with Python](https://cloud.google.com/ml-engine/docs/tensorflow/python-client-library) is useful as a basis for doing the same thing in Go.

*n.b.* although not part of the SDK, the AI platform's functionality is still available through the `gcloud` command line tool, it just isn't available as part of the SDK for Python, Go, etc.

## Accessing the API

The AI platform's API is described here: https://cloud.google.com/ml-engine/reference/rest. To set up the client do something like this:

```go
import "google.golang.org/api/ml/v1"
client := /* construct oauth'd HTTP client */
service := ml.New(client)
```

Authentication is covered further down.

### Regions

Some features are only available in certain regions; for instance [TensorFlow](https://cloud.google.com/ml-engine/docs/tensorflow/regions).

### REST endpoints

Using the APIs client the endpoints are built up in stages that follow the layout of the REST endpoints. For example:

```go
service.Projects.Jobs.List("projects/<your gcloud project name>").Do()
```

One thing to watch out for is the project name: it *must* be prefixed with `projects/`, or else you'll get a `404 not found` error like `The requested URL <code>/v1/<your gcloud project name>/jobs?alt=json&amp;prettyPrint=false</code> was not found on this server`.

## Authenticaton / OAuth

To make API calls you'll need to be authenticated using OAuth. The Google APIs client Go repository has some [examples](https://github.com/googleapis/google-api-go-client/tree/master/examples) but not for the AI platform, and in any case the OAuth setup in these examples is too complicated for getting started, so our example is simplified.

### Creating a service account

Create a service account using the `IAM and admin` section of the Google Cloud console. This service account needs to have the following roles:

* Project: `viewer`
* Storage Object: `Admin`
* ML Engine: `Admin`

Once you've created the service account, download the key JSON file.

### Getting an OAuth client in Go

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
