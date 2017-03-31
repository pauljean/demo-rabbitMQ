package rabbitMqTest;

import java.io.IOException;
import java.io.UnsupportedEncodingException;
import java.util.concurrent.TimeoutException;

import com.rabbitmq.client.AMQP;
import com.rabbitmq.client.Channel;
import com.rabbitmq.client.Connection;
import com.rabbitmq.client.ConnectionFactory;
import com.rabbitmq.client.Consumer;
import com.rabbitmq.client.DefaultConsumer;
import com.rabbitmq.client.Envelope;

public class Receiver {

    public static void main(String[] args) throws IOException, TimeoutException {


	ConnectionFactory connectionFactory = new ConnectionFactory();
	connectionFactory.setHost("10.226.159.191");
	connectionFactory.setUsername("pi");
	connectionFactory.setPassword("pi");
	connectionFactory.setVirtualHost("/pi");
	Connection connection = connectionFactory.newConnection();
	Channel channel = connection.createChannel();
	
	channel.exchangeDeclare("test", "topic");
	String queueName = channel.queueDeclare().getQueue();
	String routingKey = "greeting";
	channel.queueBind(queueName, "test", routingKey);
	System.out.println("Waiting for message");
	 
	
	Consumer consumer = new DefaultConsumer(channel){
	    
	    public void handleDelivery(String consuemrTag, Envelope envelope, AMQP.BasicProperties properties, byte[] body) throws UnsupportedEncodingException{
		
		
		String message = new String(body, "UTF-8");
		System.out.println("Message received: " + message + " On :" + envelope.getRoutingKey());
	    }
	};
	
	channel.basicConsume(queueName, true, consumer);
    }
}
