curl -XPUT 'http://localhost:9200/ire' -H 'Content-Type: application/json' \
  -d '{ "settings" : { "number_of_shards" : 3, "number_of_replicas" : 2 }}'
