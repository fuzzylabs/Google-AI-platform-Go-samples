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

To make API calls you'll need to be authenticated using OAuth. There are plenty of ways to do this and some are quite complex. We cover two simple examples:

* [Using the application-default context](oauth-default). The [application-default context](https://cloud.google.com/docs/authentication/production) is the most common and simple way to authenticate against Google APIs.
* [Directly using a service account by loading a key](oauth-json). This is provided as an example but it's not recommended, as application-default should be preferred for accessing the AI platform.

