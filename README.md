<!-- @format -->

## GO Kafka Example

This is repository how to use kafka in golang, with segmentio/kafka-go

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
