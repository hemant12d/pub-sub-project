package main

type Message struct {
	topic string
	body  string
}

var Name string = "Message"

func NewMessage(msg string, topic string) *Message {
	// Returns the message object
	return &Message{
		topic: topic,
		body:  msg,
	}
}
func (m *Message) GetTopic() string {
	// returns the topic of the message
	return m.topic
}
func (m *Message) GetMessageBody() string {
	// returns the message body.
	return m.body
}
