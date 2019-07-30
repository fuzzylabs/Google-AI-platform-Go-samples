# Introduction

In this example we authenticate against the Google AI Platform API using the [application-default](https://cloud.google.com/docs/authentication/production) credentials. This is the most common and simple way to authenticate against Google APIs including the AI platform.

## Running locally

If you want to use your own Google Cloud account, authenticate using the `gcloud auth application-default login`. If you want to assume the role of a service account, use `gcloud auth activate-service-account`.

## Running in the cloud

When deploying to Compute Engine, Kubernetes Engine, App Engine, or Cloud Function you'll need to specify a service account for the application to run as, which is surfaced as the application-default context.

## Creating a service account

Create a service account using the `IAM and admin` section of the Google Cloud console. This service account needs to have the following roles:

* Project: `viewer`
* Storage Object: `Admin`
* ML Engine: `Admin`
