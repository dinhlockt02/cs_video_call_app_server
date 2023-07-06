package meetingstore

import meetingmodel "github.com/dinhlockt02/cs_video_call_app_server/modules/meeting/model"

func GetGroupFilter(groupId string) map[string]interface{} {
	return map[string]interface{}{"group": groupId}
}

func GetStatusFilter(status meetingmodel.MeetingStatus) map[string]interface{} {
	return map[string]interface{}{"status": status}
}
