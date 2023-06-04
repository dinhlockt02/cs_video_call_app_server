package meetingstore

func GetGroupFilter(groupId string) map[string]interface{} {
	return map[string]interface{}{"group": groupId}
}
