## ELK Stack
- Docker Images https://www.docker.elastic.co/
- Documents https://www.elastic.co/guide/index.html

## Elasticsearch
- Download images https://www.docker.elastic.co/r/elasticsearch

- Configure settings https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-discovery-settings.html
- REST APIs https://www.elastic.co/guide/en/elasticsearch/reference/current/rest-apis.html

### Setup Environtment
```
discovery.type=single-node
```

## Kibana
- Download images https://www.docker.elastic.co/r/kibana
- Configure settings https://www.elastic.co/guide/en/kibana/current/settings.html

## Logstash
- Download images https://www.docker.elastic.co/r/logstash
- Configure settings https://www.elastic.co/guide/en/logstash/current/index.html

## Elasticsearch commands

### Create index (index same as table)
- Field type https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-types.html
- Built-in analyzer reference https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-analyzers.html

``` javascript
// users mean index name
PUT users
{
  "mappings": {
    "properties": {
      "title_eng": {
        "type": "text", // Field type
        "analyzer": "standard" // Built-in analyzer (preset)
      },
      "title_thai": {
        "type": "text", // Field type
        "analyzer": "thai" // Built-in analyzer (preset)
      }
    }
  }
}
```

### Create a custom analyzer
- Character filters reference https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-charfilters.html
- Tokenizer reference https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-tokenizers.html
- Token filter reference https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-tokenfilters.html

``` javascript
// Create index with custom analyzer 
PUT users
{
  "settings": {
    "analysis": {
      "analyzer": {
        "my_custom_analyzer": {
          "type": "custom", 
          "tokenizer": "standard",
          "char_filter": [
            "html_strip"
          ],
          "filter": [
            "lowercase",
            "asciifolding"
          ]
        },
        "hobbie_analyzer": {
          "type": "custom", 
          "tokenizer": "standard",
          "filter": [
            "lowercase"
          ]
        },
        "detail_analyzer": {
          "type": "custom", 
          "tokenizer": "thai",
          "char_filter": [
            "html_strip"
          ]
        }
      }
    }
  },
  "mappings": {
    "properties": {
      "name": {
        "type": "text",
        "analyzer": "my_custom_analyzer"
      },
      "hobbies": {
        "type": "text",
        "analyzer": "hobbie_analyzer"
      },
      "detail": {
        "type": "text",
        "analyzer": "detail_analyzer"
      }
    }
  }
}
```

``` javascript
// Create index
PUT programs
{
  "mappings": {
    "properties": {
      "lists": {
        "properties": {
          "description": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "name": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          }
        }
      },
      "message": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "tags": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      }
    }
  }
}
```

### Delete index
``` javascript
DELETE users
```

### List index
``` javascript
GET _cat/indices
```

### Create data
``` javascript
// 1 mean id
POST /customer/_doc/1
{
  "firstname": "Name",
  "lastname": "Test"
}
// or
POST /customer/_create/1
{
  "firstname": "ABC",
  "lastname": "DEF"
}
```

### Update data
``` javascript
// 1 mean id
POST /customer/_doc/1
{
  "firstname": "New",
  "lastname": "Name"
}
```

### Deletr data
``` javascript
// 1 mean id
DELETE /customer/_doc/1
```

### Get document
``` javascript
// 1 mean id
GET /customer/_doc/1
```

### Bulk insert
- NDJSON http://ndjson.org/
``` javascript
// data must be newline-delimited JSON (NDJSON). Each line must end in a newline character (\n), including the last line.
PUT customer/_bulk
{ "create": { "_id": 10 } }
{ "firstname": "AAA","lastname":"BBB"}
{ "create": { "_id": 11 } }
{ "firstname": "CCC","lastname":"DDD"}
{ "create": { "_id": 12 } }
{ "firstname": "EEE","lastname":"FFF"}
{ "create": { "_id": 13 } }
{ "firstname": "GGG","lastname":"HHH"}

```

### Arrays
- https://www.elastic.co/guide/en/elasticsearch/reference/current/array.html
``` javascript
PUT programs/_doc/1
{
  "message": "some arrays in this document",
  "tags":  [ "data_1", "data_2" ], 
  "lists": [ 
    {
      "name": "list_1",
      "description": "list desc"
    },
    {
      "name": "list_2",
      "description": "list desc"
    }
  ]
}

PUT programs/_doc/2 
{
  "message": "no arrays in this document",
  "tags":  "data_1",
  "lists": {
    "name": "list_1",
    "description": "list desc"
  }
}

GET programs/_search
{
  "query": {
    "match": {
      "tags": "data_1" 
    }
  }
}
```

### Search data
``` javascript
// Full-text search
GET customer/_search
{
  "query" : {
    "match" : { "firstname": "Name" }
  }
}
```

## Elasticsearch analyzer
- https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-analyzers.html

### Language analyzers
- https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-lang-analyzer.html
``` javascript
// Thai tokenizer
POST _analyze
{
  "tokenizer": "thai",
  "text": "CORS เป็นกลไกที่ web browser ใช้เวลาที่ client ส่ง request ไปยัง server ที่มี domain ต่างกัน"
}
```

## Logstash

### Input
``` javascript
input {
  http {
    port => 8080
  }
}
```

### Filter
``` javascript

```

### Output
``` javascript

```

VSCode Extensions
- Logstash Editor