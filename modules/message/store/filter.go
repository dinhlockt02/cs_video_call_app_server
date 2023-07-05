package messagestore

func GetGroupIdFilter(groupId string) map[string]interface{} {
	return map[string]interface{}{
		"group_id": groupId,
	}
}
