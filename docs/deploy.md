# Deploy
## 开发环境

- 创建网络：
`docker network create flickFrame --driver bridge`
- 创建Zookeeper:
`docker run -d --name zookeeper --network flickFrame -p 2181:2181 -t zookeeper`
- 创建kafka:
```shell
docker run -d --name kafka --network flickFrame -p 9092:9092 \\
      -e KAFKA_BROKER_ID=0 -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \\ 
      -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092 \\
      -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092 wurstmeister/kafka
```
- 手动传家Topic
```shell
docker exec -it kafka /bin/bash
cd /opt/kafka/bin
./kafka-topics.sh --create --topic topic-flickFrame-like --bootstrap-server localhost:9092
./kafka-topics.sh --describe --topic topic-flickFrame-like --bootstrap-server localhost:9092
./kafka-console-producer.sh --topic topic-flickFrame-like --bootstrap-server localhost:9092
./kafka-console-consumer.sh --topic topic-flickFrame-like --from-beginning --bootstrap-server localhost:9092
```

