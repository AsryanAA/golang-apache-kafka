# golang-apache-kafka
payment service with apache-kafka &amp; golang

# steps by install apache kafka
    ## Start the ZooKeeper service
bin/zookeeper-server-start.sh config/zookeeper.properties

    ## Start the Kafka broker service
bin/kafka-server-start.sh config/server.properties

    ## Create topic by name payment_service_topic
bin/kafka-topics.sh --create --topic payment_service_topic --bootstrap-server localhost:9092

    ## Write
bin/kafka-console-producer.sh --topic payment_service_topic --bootstrap-server localhost:9092

    ## Read
bin/kafka-console-consumer.sh --topic payment_service_topic --from-beginning --bootstrap-server localhost:9092
