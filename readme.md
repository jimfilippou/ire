### Pulling the image
Obtaining Elasticsearch for Docker is as simple as issuing a docker pull command against the Elastic Docker registry.

`docker pull docker.elastic.co/elasticsearch/elasticsearch:7.6.1`

### Starting a single node cluster with Docker

`docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.6.1`

### Running kibana for GUI searching
`docker pull docker.elastic.co/kibana/kibana:7.6.1`
`docker run --link YOUR_ELASTICSEARCH_CONTAINER_NAME_OR_ID:elasticsearch -p 5601:5601 {docker.elastic.co/kibana/kibana:7.6.1`

#### Useful references
- https://www.youtube.com/watch?v=jDYeqCkAN_Y