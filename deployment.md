# Deployment
Core is designed to be semi-stateless and have minimal external dependencies. Core can be deployed on the cloud or on-premise. This document describes deployment in Google Cloud, however there is very little GCP specific features in the deployment, so it is simple to convert these instructions for other cloud providers like Amazon Web Services or on-premise deployment on ACM@UIUC owned servers.

## Continuous Integration & Continuous Deployment

### Unit Tests
Units tests are run on every commit which is pushed to GitHub using GitHub actions. These unit tests ensure that the service level interactions with datastores are correct, however additional testing still needs to be performed in order to verify that the end to end workflows the user experiences is correct.

### Integration Tests
At this time there are no integration tests for automatically verify end to end correctness. In the future, theses tests will be provided and run on every commit pushed to GitHub using GitHub actions.

### Continuous Deployment
Continuous deployment is handled using GCP Cloud Build with a GitHub webhook as the build trigger. The webhook is fired on every commit pushed to the repository's master branch. When triggered the Cloud Build workflow build a new version of the production container and pushes it to the GCP Container Registry for this service. Lastly Cloud Build runs a deployment which creates a new version of the Cloud Run service, redirecting traffic to the latest version of the service and draining the old version of the service.

## Production Resources
Core utilizes two production data stores and a single compute platform. Combined these handle running the entire production deployment

### GitStore
The GitStore needs to provide an HTTPS interface for downloading files in their raw format. The production deployment reads data directly from the `data/` folder in this repository. GitHub provides the HTTPS interface for rectreiving this information.

### Database
All user data is stored in a MySQL data. The production MySQL database is fully managed GCP Cloud SQL database. Migrations should be run against this database using the GCP Console which provides direct access a SQL client that is connected to the database.

### Core Service
The Core service is deployed using GCP Cloud Run which is allows for the serverless deployment of containers. Cloud Run will handle auto scaling within the limits specified in the service definition. Secrets are provided to Core using the environment variable configuration for the service.
