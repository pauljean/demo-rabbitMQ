package rabbitMqTest;

import java.io.IOException;
import java.util.concurrent.TimeoutException;

import com.rabbitmq.client.Channel;
import com.rabbitmq.client.Connection;
import com.rabbitmq.client.ConnectionFactory;

public class Publisher {

    public static void main(String[] args) throws IOException, TimeoutException {

	ConnectionFactory connectionFactory = new ConnectionFactory();
	connectionFactory.setHost("10.226.159.191");
	connectionFactory.setUsername("pi");
	connectionFactory.setPassword("pi");
	connectionFactory.setVirtualHost("/pi");
	Connection connection = connectionFactory.newConnection();
	Channel channel = connection.createChannel();
	
	channel.exchangeDeclare("test", "topic");
	String message = "salut";
	String routingKey = "greeting";
	channel.basicPublish("test", routingKey, null, message.getBytes());
	System.out.println("Message has been sent;");
	channel.close();
	connection.close();
	
    }

}
