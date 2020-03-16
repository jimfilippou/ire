<div align="center">
    <h1>IRE</h1>
    <img src="ire.png" alt="ire"  style="margin-bottom: 20px;"/>
</div>

IRE stands for "Information Retrieval" which is an AUEB course. This implementation delivers a tool written
entirely in Go, which handles an "Elastic Search" instance, regarding data & indexes. 

# WIP & Notice

Chances are, this will not be useful to you, since it covers assignment needs, however you can use parts of the code or 
get inspired to build something similar. Work is in progress, no stable build will ever be released because "who cares".

# Usage

Thanks to Go's nature, IRE can be compiled to a single executable "ire.exe" or "ire" for *nix systems.
To compile this use the following command

`go build`

### Converting AUEB provided data to JSON

`ire generate json`

### Inserting data to a cluster

`ire feed`

# Docker setup 🐋

Docker compose is not implemented in this project, so you have to manually start the containers using the
following instructions.

### Pulling the image

Obtaining ElasticSearch for Docker is as simple as issuing a docker pull command against the Elastic Docker registry.

`docker pull docker.elastic.co/elasticsearch/elasticsearch:7.6.1`

### Starting a single node cluster with Docker

`docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.6.1`

### Running kibana for GUI searching

`docker pull docker.elastic.co/kibana/kibana:7.6.1`

`docker run --link YOUR_ELASTICSEARCH_CONTAINER_NAME_OR_ID:elasticsearch -p 5601:5601 {docker.elastic.co/kibana/kibana:7.6.1`

#### Useful references
- https://www.youtube.com/watch?v=jDYeqCkAN_Y