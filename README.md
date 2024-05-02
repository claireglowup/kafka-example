<!-- @format -->

## GO Kafka Example

This is repository how to use kafka in golang, with segmentio/kafka-go

## Motivations

We rely on both Go and Kafka a lot at Segment. Unfortunately, the state of the Go
client libraries for Kafka at the time of this writing was not ideal. The available
options were:

- [sarama](https://github.com/Shopify/sarama), which is by far the most popular
  but is quite difficult to work with. It is poorly documented, the API exposes
  low level concepts of the Kafka protocol, and it doesn't support recent Go features
  like [contexts](https://golang.org/pkg/context/). It also passes all values as
  pointers which causes large numbers of dynamic memory allocations, more frequent
  garbage collections, and higher memory usage.

- [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go) is a
  cgo based wrapper around [librdkafka](https://github.com/edenhill/librdkafka),
  which means it introduces a dependency to a C library on all Go code that uses
  the package. It has much better documentation than sarama but still lacks support
  for Go contexts.

- [goka](https://github.com/lovoo/goka) is a more recent Kafka client for Go
  which focuses on a specific usage pattern. It provides abstractions for using Kafka
  as a message passing bus between services rather than an ordered log of events, but
  this is not the typical use case of Kafka for us at Segment. The package also
  depends on sarama for all interactions with Kafka.

This is where `kafka-go` comes into play. It provides both low and high level
APIs for interacting with Kafka, mirroring concepts and implementing interfaces of
the Go standard library to make it easy to use and integrate with existing
software.

#### Note:

In order to better align with our newly adopted Code of Conduct, the kafka-go
project has renamed our default branch to `main`. For the full details of our
Code Of Conduct see [this](./CODE_OF_CONDUCT.md) document.

## Kafka versions

`kafka-go` is currently tested with Kafka versions 0.10.1.0 to 2.7.1.
While it should also be compatible with later versions, newer features available
in the Kafka API may not yet be implemented in the client.

## Go versions

`kafka-go` requires Go version 1.15 or later.

## How to run

```bash
# clone this repo
git clone git@github.com:claireglowup/kafka-example.git
# or
git clone https://github.com/claireglowup/kafka-example.git


# Add kafka to your hosts
sudo vim /etc/hosts
127.0.0.1 kafka

cd kafka-example

# running docker compose
docker-compose up

# run consumer with command
go run cmd/consumer/main.go

# run publisher with command
go run cmd/producer/main.go
```
