package requeststore

func GetRequestSenderIdFilter(senderId string) map[string]interface{} {
	return map[string]interface{}{"sender.id": senderId}
}

func GetRequestReceiverIdFilter(receiverId string) map[string]interface{} {
	return map[string]interface{}{"receiver.id": receiverId}
}

func GetTypeFilterFilter(group bool) map[string]interface{} {
	return map[string]interface{}{
		"group": map[string]interface{}{
			"$exists": group,
		},
	}
}
