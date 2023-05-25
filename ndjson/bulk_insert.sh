curl localhost:9200/accounts/_bulk -H "content-type:application/x-ndjson" --data-binary "@accounts.ndjson"
curl localhost:9200/orders/_bulk -H "content-type:application/x-ndjson" --data-binary "@orders.ndjson"
curl localhost:9200/products/_bulk -H "content-type:application/x-ndjson" --data-binary "@products.ndjson"
curl localhost:9200/recipes/_bulk -H "content-type:application/x-ndjson" --data-binary "@recipes.ndjson"