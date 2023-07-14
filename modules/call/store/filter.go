package callstore

import callmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/call/model"

func GetCallerIdFilter(senderId string) map[string]interface{} {
	return map[string]interface{}{"caller.id": senderId}
}

func GetCalleeIdFilter(receiverId string) map[string]interface{} {
	return map[string]interface{}{"callee.id": receiverId}
}

func GetCallRoomFilter(groupId string) map[string]interface{} {
	return map[string]interface{}{"call_room": groupId}
}

func GetCallStatusFilter(status callmdl.Status) map[string]interface{} {
	return map[string]interface{}{"status": status}
}

func GetCallOwnerFilter(owner string) map[string]interface{} {
	return map[string]interface{}{"owner": owner}
}
