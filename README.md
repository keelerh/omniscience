# omniscience
Search Everything

## Elasticsearch

### Install

```
$ brew update
$ brew install elasticsearch
```

### Using local ES server

To avoid getting a cors related error when connecting the UI to your local
elasticsearch instance, add the following to your `config/elasticsearch.yml' file.

```
http.cors.enabled : true  
http.cors.allow-origin : "*"
http.cors.allow-methods : OPTIONS, HEAD, GET, POST, PUT, DELETE
http.cors.allow-headers : X-Requested-With,X-Auth-Token,Content-Type, Content-Length
```

### Start local ES server

```$xslt
$ bin/elasticsearch
```
