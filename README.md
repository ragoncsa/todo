# todo

## Overview

This is a sample application for REST service development with golang. Libraries used:

- [gorm](https://gorm.io/) as ORM library
- [viper](https://github.com/spf13/viper) for configuration management
- [gin](https://github.com/gin-gonic/gin) as web framework
- [gin-swagger](https://github.com/swaggo/gin-swagger) to generate OpenAPI spec from go comments
- [resty](https://github.com/go-resty/resty) for REST client implementation (for example to talk to OPA)
- [Open Policy Agent (OPA)](https://www.openpolicyagent.org/) for authorization decisions
- [Google APIs Client Library for Go](https://pkg.go.dev/google.golang.org/api) to validate Google JWT tokens

![Overview](assets/screenshot-swagger-ui.png?raw=true)

## Run from container

To build

`docker build -t todo .`

Docker-compose starts the built container with a database

`docker-compose up`

Go to Swagger UI <http://localhost:8080/swagger/index.html>

### Reset the database

`docker-compose down --volumes`

## Run without container

### Start dependencies

```shell
docker-compose up db
opa build authz -o authz/bundle.tar.gz --ignore 'taskservice_authz_test.rego'
docker-compose up bundle_server
docker-compose up opa
```

To access the database:

```shell
$ docker exec -it todo_db_1 /bin/bash
root@187961c81d2e:/# psql -U postgres
psql (14.2 (Debian 14.2-1.pgdg110+1))
Type "help" for help.
```

### Start the server

`go run main.go`

Go to Swagger UI <http://localhost:8080/swagger/index.html>

## Testing

### Test the application

`go test ./...`

### Test the authorization rules

Run unit tests

`opa test authz -v --ignore '*.tar.gz'`

Test rules on the server

```shell
echo "{\"input\": {\"method\":\"POST\",\"owner\":\"johndoe\",\"path\":[\"tasks\"],\"user\":\"johndoe\"}}" \
| http -v POST http://127.0.0.1:8181/v1/data/authz
```

## Generate OpenAPI spec

`swag init`

For more see: <https://github.com/swaggo/gin-swagger>

## Authentication

When authn.notEnforced is true you can simply pass a user ID in the CallerId HTTP header.

If you set Authorizaton HTTP header in your calls, you must use the Bearer schema and include a
valid Google OIDC ID token.

### Testing with Swagger UI

You can use Swagger UI to make a test call. To do this you will need to get an ID token issued by Google. (This demo supports only one identity provider, so it must be Google). The steps to get an ID token is described in the next section.

1. Go to swagger UI <http://localhost:8080/swagger/index.html> and click on the "Authorize" button.

![Authorize button](assets/screenshot-swagger-ui-authorize-button.png?raw=true)

2. Enter "Bearer [YOUR ID TOKEN]" to the text field and click on "Authorize"

![Logging in](assets/screenshot-swagger-ui-login.png?raw=true)

and then click on "Close".

Now Swagger UI will add an Authorization header to all of your operations that have the "lock icon" (all the operations in our case).

3. Try some API calls

Try to create a new task, by going to "POST /tasks/" and clicking on "Try out"

![Try out creating a task](assets/screenshot-swagger-ui-try-out.png?raw=true)

You will receive 403 in response, since with the default sample request you were trying to create a task for user "johndoe" and users are only permitted to create tasks for themselves. (See the authorization rules created in the [earlier article](https://medium.com/enlear-academy/open-policy-agent-opa-to-externalize-authorization-decisions-in-rest-api-implemented-in-go-faee67d29053#9dce).)

![403 response to sample call](assets/screenshot-swagger-ui-403.png?raw=true)

Now change the user to the email address of the user you signed in to retrieve the token.

![Modify request for creating a task](assets/screenshot-swagger-ui-modify-request.png?raw=true)

This time your task will be created and you receive a 201 in your response.

### Generate an ID token for testing the API

1. First you need to create OAuth client ID credentials. You can find documents online how to do this, for example <https://developers.google.com/workspace/guides/create-credentials#oauth-client-id>.

Once you are done you should have something similar to this

![Setting up OAuth Client ID credentials](assets/google-cloud-credentials-setup.png?raw=true)

Make sure you add <https://developers.google.com/oauthplayground> as authorized redirect URI.

2. Generate a token with Google's OAuth Playground

You can find documentation online to see how to do this, for example <https://developers.google.com/google-ads/api/docs/oauth/playground>, but here is brief summary.

Copy and paste your client ID and client secret from the previous step into the playground

![OAuth playground - setting credentials](assets/google-oauth-playground-settings.png?raw=true)

Then enter "openid email" to the scopes and click on "Authorize APIs".

![OAuth playground - setting scopes](assets/google-oauth-playground-scopes.png?raw=true)

Then on the next page click on "Exchange authorization code for tokens"

On the last page you can find the id_token returned that you can use to authenticate with the service (see above).

![OAuth playground - getting id_token](assets/google-oauth-playground-id-token.png?raw=true)

## Deploy to AWS

Create database and set up IAM based authentication.

```bash
aws cloudformation validate-template --template-body file://deploy/cloudformation/database.yaml
aws cloudformation deploy --template-file deploy/cloudformation/database.yaml --stack-name todo-service-database

export PGPASSWORD=$( \
  aws secretsmanager get-secret-value \
  --secret-id "todo/postgres" \
  --query "SecretString" \
  --output text \
  | jq -r .password)

HOST=$( \
  aws secretsmanager get-secret-value \
  --secret-id "todo/postgres" \
  --query "SecretString" \
  --output text \
  | jq -r .host)

psql -h ${HOST} -U postgres -c "GRANT rds_iam TO postgres;"
unset PGPASSWORD
```

Then build and deploy the service.

```bash
cd deploy/aws-sam/todo-service

sam build

VPC=$(aws ec2 describe-vpcs --filters "Name=is-default,Values=true" --query "Vpcs[].VpcId" --output text)
SEC_GROUP=$(aws ec2 describe-security-groups --query "SecurityGroups[?VpcId=='${VPC}']".GroupId --output text)
SUBNETS=$(aws ec2 describe-subnets --query "Subnets[?VpcId=='${VPC}'].SubnetId" --output text | sed 's/\t/,/g')
DB_RESOURCE_ID=$(aws rds describe-db-clusters --db-cluster-identifier todo-service-database --query "DBClusters[].DbClusterResourceId" --output text)

sam deploy --stack-name todo-service \
--resolve-image-repos \
--resolve-s3 \
--capabilities CAPABILITY_IAM \
--parameter-overrides VpcSecurityGroupIds=${SEC_GROUP} VpcSubnetIds=${SUBNETS} TodoDbClusterResourceId=${DB_RESOURCE_ID}
```
