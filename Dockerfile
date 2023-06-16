# Use the official RabbitMQ image as the base image
FROM rabbitmq:latest

# Expose the RabbitMQ management port (optional)
EXPOSE 15672

# Expose the RabbitMQ AMQP port (optional)
EXPOSE 5672
