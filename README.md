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
4. Once you have created your project, click on the options menu icon in the top left corner of the page, then 'APIs & Services', then 'Credentials'
5. Click on 'Create credentials' and then 'Service account key'
6. Next, select 'New service account', name it anything and select 'Project' and then 'Viewer' as the role from the dropdown list, finally select JSON as the key type and click 'Create'. Upon clicking 'Create', a JSON file will be downloaded; this is important for later so remember where you downloaded it
7. Click on 'Manage service accounts' (on the right-hand side), select your new Service Account, click on the three dots on the right-hand side, and then select Edit
8. Tick the box "Enable GSuite Domain-wide Delegation" and click Configure consent screen
9. Click on the options menu icon in the top left corner again, choose 'APIs & Services' and then 'Dashboard'
10. Click on the Drive API and then click the blue Enable button

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
