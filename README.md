# omniscience
Search Everything

## Elasticsearch

### Install

```
$ brew update
$ brew install elasticsearch
```

### Using local ES server

To avoid getting a cors related error when connecting the UI to your local elasticsearch instance, add the following to your `config/elasticsearch.yml` file (if you installed Elasticsearch using brew, this file can be found at `/usr/local/Cellar/elasticsearch/6.2.2/libexec/config/elasticsearch.yml`).

```
http.cors.enabled : true  
http.cors.allow-origin : "*"
http.cors.allow-methods : OPTIONS, HEAD, GET, POST, PUT, DELETE
http.cors.allow-headers : X-Requested-With,X-Auth-Token,Content-Type, Content-Length
```

### Start local ES server

```
$ elasticsearch
```

## Google Drive

### Create a Google service account

1. Go to [cloud.google.com/console](http://cloud.google.com/console)
2. Click Create Project
3. Enter a project name and click Create
4. Once you have created your project, click on the options menu in the top left of the page, then 'API Manager', then 'Credentials'. Click on 'New credentials' and then 'Service account key'
5. Click on the 'hamburger' menu icon (next to "Google Cloud Platform" in the top left of the page), then 'API Manager', then 'Credentials'
6. Click on 'New credentials', then 'Service account key'
7. Next, select 'New service account', name it anything and select 'Project' and then 'Viewer' as the role from the dropdown list, finally select JSON as the key type and click 'Create'. Upon clicking 'Create', a JSON file will be downloaded; this is important for later so remember where you downloaded it
8. Take note of the `client_email`
8. Click on 'Manage service accounts' (on the right-hand side), then select your new Service Account, click on the three dots beside, and select Edit
9. Tick the box "Enable GSuite Domain-wide Delegation" and click Configure consent screen
10. Add a Project name (it can be anything) and click Save
11. Click on the three-lines icon again, choose 'API Manager' and in the 'Overview'
12. Click on the Drive API and then click the blue Enable button

### Share folders with Google service account

For the Google service account to have access to index your files you must share the relevant folders with the service account's `client_email`.

## Golang backend

### Start Go server

From the project root run:

```
$ go run cmd/omniscience_server/main.go --fGoogleServiceAccountPath=<path-to-service-account-json>
```

### Manually invoke indexing

```
$ go run cmd/omniscience_client/main.go
```

## React app

### Install dependencies

From the project root run:

```
$ cd client && npm i
```

### Start local node server

From the project root run:

```
$ cd client && npm start
```

## Search everything!

Navigate to [http://localhost:3000](http://localhost:3000)
