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

## Kibana Dev Tools
http://localhost:5601/app/dev_tools#/console

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
PUT logs
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
        "type_analyzer": {
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
      "type": {
        "type": "text",
        "analyzer": "type_analyzer"
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
// Create index (Nested property)
PUT activity
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

### View index fields
``` javascript
// recipes mean index name
GET recipes/_mapping
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

### Delete data
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
// 1 mean id
PUT activity/_doc/1
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

// 2 mean id
PUT activity/_doc/2 
{
  "message": "no arrays in this document",
  "tags":  "data_1",
  "lists": {
    "name": "list_1",
    "description": "list desc"
  }
}

// search
GET activity/_search
{
  "query": {
    "match": {
      "tags": "data_1" 
    }
  }
}
```

### Search options
- https://www.elastic.co/guide/en/elasticsearch/reference/current/search.html
``` javascript
// Size
GET recipes/_search
{
  "size": 5
}

// Source (same as select command in sql)
GET recipes/_search
{
  "_source": ["title", "description", "ingredients.name"]
}

// Source excludes
GET recipes/_search
{
  "_source": {
    "excludes": ["description","servings","ratings"]
  }
}

// Pagination
GET recipes/_search
{
  "_source": false,
  "size": 5,
  "from": 6
}

// Sort
GET recipes/_search
{
  "_source": ["title","ingredients.name","steps"],
  "sort": [
    {
      "title.keyword": {
        "order": "asc"
      }
    },
    {
      "preparation_time_minutes": {
        "order": "desc"
      }
    }
  ]
}
```

### Searching score explanation
``` javascript
GET products/_search
{
  "query": {
    "term": {
      "name": {
        "value": "tea"
      }
    }
  },
  "_source": ["name"], 
  "explain": true // display score explanation
}
```

### Match search
``` javascript
GET recipes/_search
{
  "query": {
    "match": {
      "title": "Cheese"
    }
  },
  "_source": ["title","description"]
}

GET recipes/_search
{
  "query": {
    "match": {
      "ingredients.name": "cheese"
    }
  },
  "_source": ["title","description"]
}
```

### Keyword search
``` javascript
GET recipes/_search
{
  "query": {
    "match": {
      "title.keyword": "Stovetop Macaroni and Cheese"
    }
  },
  "_source": ["title","description"]
}
```

### Term search
``` javascript
GET recipes/_search
{
  "query": {
    "term": {
      "title": "stovetop"
    }
  },
  "_source": ["title","description"]
}

GET products/_search
{
  "query": {
    "term": {
      "name": {
        "value": "tea"
      }
    }
  }
}
```

### Search with id
``` javascript
GET products/_search
{
  "query": {
    "ids": {
      "values": ["1","2","3","20","30"]
    }
  }
}
```

### Search with multiple terms
``` javascript
GET products/_search
{
  "_source": ["tags"],
  "query": {
    "terms": { // terms
      "tags.keyword": [
        "Coffee",
        "Energy drink"
      ]
    }
  }
}
```

### Search range values
- [Range query](https://www.elastic.co/guide/en/elasticsearch/reference/8.7/query-dsl-range-query.html)
``` javascript
GET products/_search
{
  "query": {
    "range": {
      "in_stock": {
        "gte": 10,
        "lte": 50
      }
    }
  },
  "_source": ["name","in_stock"]
}
```

### Search with date range
- [Date Format](https://www.elastic.co/guide/en/elasticsearch/reference/8.7/mapping-date-format.html)
``` javascript
GET products/_search
{
  "query": {
    "range": {
      "created": {
        "gte": "01/01/2004",
        "lte": "01/12/2005",
        "format": "dd/MM/yyyy" // set date format
      }
    }
  },
  "_source": ["name","in_stock","created"]
}
```

### Date math
- [Date Math](https://www.elastic.co/guide/en/elasticsearch/reference/8.8/common-options.html#date-math)
> +1h: Add one hour<br>
> -1d: Subtract one day
``` javascript
// Subtract one month
GET products/_search
{
  "query": {
    "range": {
      "created": {
        "gte": "01/01/2004||-1M", // Subtract one month
        "format": "dd/MM/yyyy"
      }
    }
  },
  "_source": ["name","in_stock","created"]
}

// Now and subtract 20 year
GET products/_search
{
  "query": {
    "range": {
      "created": {
        "gte": "now-20y", // now and Subtract 20 year
        "format": "dd/MM/yyyy"
      }
    }
  },
  "_source": ["name","in_stock","created"]
}
```

### Round down to the nearest day
> /d mean start at 00:00:00 time of day<br>
> /M mean start at 1st day of month at 00:00:00 time<br>
> /y mean start at 1st day and 1st month of year at 00:00:00 time<br>
``` javascript
// now = 26/05/2023 07:59:18
// now-1y/d => 26/05/2022 00:00:00
// now-1y/M => 01/05/2022 00:00:00
// now-1y/y => 01/01/2022 00:00:00

GET products/_search
{
  "query": {
    "range": {
      "created": {
        "gte": "now/M-20y", // or "now-20y/M"
        "format": "dd/MM/yyyy"
      }
    }
  },
  "_source": ["name","in_stock","created"]
}
```

### Search with non null value
``` javascript
GET products/_search
{
  "query": {
    "exists": {
      "field": "tags"
    }
  }
}
```

### Search with null value
``` javascript
GET products/_search
{
  "query": {
    "bool": {
      "must_not": [
        {
          "exists": {
            "field": "tags"
          }
        }
      ]
    }
  }
}
```

### Search with prefix
``` javascript
GET products/_search
{
  "query": {
    "prefix": {
      "tags.keyword": {
        "value": "Energy"
      }
    }
  }
}
```

### Search with wildcard
> \* Represents zero or more characters<br>
> ? Represents a single character
``` javascript
// *drink
GET products/_search
{
  "query": {
    "wildcard": {
      "tags.keyword": {
        "value": "*drink"
      }
    }
  }
}

// Energy*
GET products/_search
{
  "query": {
    "wildcard": {
      "tags.keyword": {
        "value": "Energy*"
      }
    }
  }
}

// En??gy*
GET products/_search
{
  "query": {
    "wildcard": {
      "tags.keyword": {
        "value": "En??gy*"
      }
    }
  }
}
```

### Search with Regular Expression
- https://regex101.com/
``` javascript
GET products/_search
{
  "query": {
    "regexp": {
      "tags.keyword": "[Ee][\\w]+\\s(drink)"
    }
  }
}
```

### Full-text search
``` javascript
// OR operator 
GET recipes/_search
{
  "query": {
    "match": {
      "title": "Mushrooms Spinach"
    }
  }
}

// AND operator
GET recipes/_search
{
  "query": {
    "match": {
      "title": {
        "query": "Mushrooms Spinach",
        "operator": "and"
      }
    }
  },
  "_source": ["title"]
}
```

### Matching phrases
``` javascript
GET recipes/_search
{
  "query": {
    "match_phrase": {
      "title": "Spinach, and Mushrooms"
    }
  }
}
```

### Multi-match search
``` javascript
GET recipes/_search
{
  "query": {
    "multi_match": {
      "query": "olive oil",
      "operator": "and", 
      "fields": ["title","description"]
    }
  },
  "_source": ["title","description"]
}
```

### Must & Must not & Should & Filter search
``` javascript
GET recipes/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "multi_match": {
            "query": "Spaghetti",
            "fields": ["title", "description"]
          }
        }
      ],
      "must_not": [
        {
          "match": {
            "step": {
              "query": "chili pepper cheese",
              "operator": "and"
            }
          }
        }
      ],
      "should": [
        {
          "match_phrase": {
            "ingredients.name": "olive oil"
          }
        }
      ], 
      "filter": [
        {
          "range": {
            "preparation_time_minutes": {
              "lte": 30
            }
          }
        },
        {
          "range": {
            "servings.min": {
              "gte": 2,
              "lte": 4
            }
          }
        }
      ]
    }
  }
}
```

### Display matched queries name
``` javascript
// Set name for each query
GET recipes/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "multi_match": {
            "query": "Spaghetti",
            "fields": ["title", "description"],
            "_name": "must_multi_match" // set name
          }
        }
      ],
      "must_not": [
        {
          "match": {
            "step": {
              "query": "chili pepper cheese",
              "operator": "and",
              "_name": "must_not_match" // set name
            }
          }
        }
      ],
      "should": [
        {
          "match_phrase": {
            "ingredients.name": {
              "query": "olive oil",
              "_name": "should_match_phrase" // set name
            }
          }
        }
      ], 
      "filter": [
        {
          "range": {
            "preparation_time_minutes": {
              "lte": 30,
              "_name": "filter_range_preparation_time_minutes" // set name
            }
          }
        },
        {
          "range": {
            "servings.min": {
              "gte": 2,
              "lte": 4,
              "_name": "filter_range_servings_min" // set name
            }
          }
        }
      ]
    }
  },
  "_source": ["title", "ingredients.name"]
}
```

### Aggregation
- https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations.html#search-aggregations

### Metrics aggregations
- https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics.html

### Sum
``` javascript
GET orders/_search
{
  "size": 0,
  "aggs": {
    "sum_of_total_amount": {
      "sum": {
        "field": "total_amount"
      }
    }
  }
}
```

### Avg
``` javascript
GET orders/_search
{
  "size": 0,
  "aggs": {
    "avg_of_total_amount": {
      "avg": {
        "field": "total_amount"
      }
    }
  }
}
```

### Min
``` javascript
GET orders/_search
{
  "size": 0,
  "aggs": {
    "min_of_total_amount": {
      "min": {
        "field": "total_amount"
      }
    }
  }
}
```

### Max
``` javascript
GET orders/_search
{
  "size": 0,
  "aggs": {
    "max_of_total_amount": {
      "max": {
        "field": "total_amount"
      }
    }
  }
}
```

### Cardinality
- https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-cardinality-aggregation.html
``` javascript
GET orders/_search
{
  "size": 0,
  "aggs": {
    "salesman_count": {
      "cardinality": {
        "field": "salesman.id"
      }
    }
  }
}
```

### Value count
- https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-valuecount-aggregation.html
``` javascript
GET orders/_search
{
  "size": 0,
  "aggs": {
    "salesman_count": {
      "value_count": {
        "field": "salesman.id"
      }
    }
  }
}
```

### Stats aggregation
- https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-stats-aggregation.html
``` javascript
GET orders/_search
{
  "size": 0,
  "aggs": {
    "total_amount_stats": {
      "stats": {
        "field": "total_amount"
      }
    }
  }
}
```

### Bucket aggregations
- https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-bucket.html#search-aggregations-bucket

### Terms (Grouping)
``` javascript
GET orders/_search
{
  "size": 0,
  "aggs": {
    "status_terms": {
      "terms": {
        "field": "status.keyword",
        "size": 20 // number of displaying term
      }
    }
  }
}
```

### Search & Grouping stats
``` javascript
GET orders/_search
{
  "size": 0,
  "query": {
    "bool": {
      "filter": [
        {
          "range": {
            "total_amount": {
              "gte": 50,
              "lt": 200
            }
          }
        },
        {
          "range": {
            "purchased_at": {
              "lte": "01/05/2016",
              "format": "dd/MM/yyyy"
            }
          }
        }
      ]
    }
  }, 
  "aggs": {
    "status_terms": { // create term
      "terms": {
        "field": "status.keyword",
        "size": 20
      },
      "aggs": { // in term
        "total_amount_stats": {
          "stats": {
            "field": "total_amount"
          }
        }
      }
    },
    "all_order": { // not in term
      "stats": {
        "field": "total_amount"
      }
    }
  }
}
```

### Filter
``` javascript
GET orders/_search
{
  "size": 0,
  "aggs": {
    "filter_total_amount": {
      "filter": {
        "range": {
          "total_amount": {
            "gte": 200,
            "lte": 400
          }
        }
      },
      "aggs": {
        "total_amount_stats": {
          "stats": {
            "field": "total_amount"
          }
        }
      }
    }
  }
}
```

### Filters (Custom buckets)
``` javascript
GET orders/_search
{
  "size": 0,
  "query": {
    "match": {
      "status": "completed cancelled"
    }
  }, 
  "aggs": {
    "order_statuses": {
      "filters": {
        "other_bucket_key": "other_status", 
        "filters": {
          "processed": {
            "match": {
              "status.keyword": "processed"
            }
          },
          "completed": {
            "match": {
              "status.keyword": "completed"
            }
          }
        }
      },
      "aggs": {
        "total_amount_stats": {
          "stats": {
            "field": "total_amount"
          }
        }
      }
    }
  }
}
```

### Range (Group by range)
``` javascript
GET orders/_search
{
  "size": 0,
  "aggs": {
    "total_amount": {
      "range": {
        "field": "total_amount",
        "ranges": [
          { "to": 100 }, // mean x... to 99 (< 100)
          { "from": 100, "to": 200 }, // mean 100 to 199 (100 >= x < 200)
          { "from": 200 } // mean 200 to ...x (200 >=)
        ]
      },
      "aggs": {
        "total_amount_stats": {
          "stats": {
            "field": "total_amount"
          }
        }
      }
    }
  }
}
```

### Date range (Group by date)
``` javascript
GET orders/_search
{
  "size": 0,
  "aggs": {
    "purchased_at_date_range": {
      "date_range": {
        "field": "purchased_at",
        "format": "dd/MM/yyyy HH:mm:ss",
        "keyed": true, 
        "ranges": [
          {
            "from": "01/01/2016 00:00:00",
            "to": "01/01/2016 00:00:00||+6M",
            "key": "first_half_year"
          },
          {
            "from": "01/01/2016 00:00:00||+6M",
            "to": "01/01/2016 00:00:00||+12M",
            "key": "second_half_year"
          },
          {
            "from": "01/01/2016 00:00:00||+12M",
            "to": "now/d"
          }
        ]
      },
      "aggs": {
        "total_amount_stats": {
          "stats": {
            "field": "total_amount"
          }
        }
      }
    }
  }
}
```

### Histogram
- https://www.elastic.co/guide/en/elasticsearch/reference/8.7/search-aggregations-bucket-histogram-aggregation.html
``` javascript
GET orders/_search
{
  "size": 0,
  "aggs": {
    "total_amount_histogram": {
      "histogram": {
        "field": "total_amount",
        "interval": 25,
        "keyed": true,
        "min_doc_count": 1
      }
    }
  }
}
```

### Date histogram
- https://www.elastic.co/guide/en/elasticsearch/reference/8.7/search-aggregations-bucket-datehistogram-aggregation.html
``` javascript
GET orders/_search
{
  "size": 0,
  "aggs": {
    "purchased_at_date_histogram": {
      "date_histogram": {
        "field": "purchased_at",
        "format": "dd/MM/yyyy HH:ss:mm", 
        "calendar_interval": "quarter",
        //"fixed_interval": "30d",
        "keyed": true,
        "missing": "01/01/2000 00:00:00"
      },
      "aggs": {
        "total_amount_stats": {
          "stats": {
            "field": "total_amount"
          }
        }
      }
    }
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

## Logstash pipeline

### Create new pipeline
1. Add new *logstash.conf in ./logstash/pipeline
2. Add *logstash.conf file to pipeline.yml
```
- pipeline.id: {pipeline name}
  path.config: "/usr/share/logstash/pipeline/*logstash.conf"
```

### Input
- [Input plugins](https://www.elastic.co/guide/en/logstash/8.7/input-plugins.html)
``` javascript
// standard input
input {
  stdin {
    
  }
}

// http
input {
  http {
    port => 8080
  }
}

// file
input {
  http {
    port => 8080
  }
}
```

### Filter
- [Filter plugins](https://www.elastic.co/guide/en/logstash/8.7/filter-plugins.html)
- [Mutate filter plugin](https://www.elastic.co/guide/en/logstash/8.7/plugins-filters-mutate.html)
``` javascript
filter {
  mutate {
    convert => {
      "age" => "integer" # Convert type age to integer
      "[address][street]" => "string" # Convert a nested field name street to string
    }

    join => {
      "hobbies" => "||" # join array
    }

    remove_field => ["@version", "[event][original]"] # [event][original] mean Remove a nested field
  }
}
```

### Grok filter plugin
- [Grok filter plugin](https://www.elastic.co/guide/en/logstash/8.7/plugins-filters-grok.html)
- SYNTAX, Logstash ships with about 120 patterns by default. You can find them here: https://github.com/logstash-plugins/logstash-patterns-core/tree/master/patterns
``` javascript
// pattern %{SYNTAX:SEMANTIC:TYPE}
// TYPE is optional
// e.g. %{IP:ip}
// e.g. %{INT:http_status:int}

filter {
    grok {
        match => {
            "message" => "%{IP:ip} - %{USER:user} \[%{HTTPDATE:date}\] \"%{WORD:method} %{URIPATHPARAM:request}"
        }
    }
}
```

### Output
- [Output plugins](https://www.elastic.co/guide/en/logstash/8.7/output-plugins.html)
- [Elasticsearch output plugin](https://www.elastic.co/guide/en/logstash/8.7/plugins-outputs-elasticsearch.html#plugins-outputs-elasticsearch)
``` javascript
// standard output (output to console)
output {
  stdout {

  }
}

// file
output {
  file {
    path => "/usr/share/logs/output/output-%{type}-%{+ddMMyyyy}.log"
  }
}

// elasticsearch
output {
  elasticsearch {
    hosts => ["http://elasticsearch:9200"]
    index => "log_%{type}_%{+yyyyMMdd}"
  }
}
```

VSCode Extensions
- Elasticsearch for VSCode (Use with .es file for querying Elasticsearch)
- Logstash Editor

## Test

Elasticsearch http://localhost:9200<br>
Kibana http://localhost:5601<br>
Logstash http://localhost:8080 (main pipeline), http://localhost:8081 (apache pipeline)<br>
