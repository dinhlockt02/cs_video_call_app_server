package requeststore

func GetRequestSenderIdFilter(senderId string) map[string]interface{} {
	return map[string]interface{}{"sender.id": senderId}
}

func GetRequestReceiverIdFilter(receiverId string) map[string]interface{} {
	return map[string]interface{}{"receiver.id": receiverId}
}
