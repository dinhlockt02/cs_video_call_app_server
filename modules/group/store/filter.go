package groupstore

import "github.com/dinhlockt02/cs_video_call_app_server/common"

func GetGroupIdInIdListFilter(ids ...string) map[string]interface{} {
	mongoIds := make([]interface{}, 0, len(ids))
	for i := range ids {
		mongoId, err := common.ToObjectId(ids[i])
		if err == nil {
			mongoIds = append(mongoIds, mongoId)
		}
	}
	return common.GetInFilter("_id", mongoIds...)
}

func GetUserIdInIdListFilter(ids ...string) map[string]interface{} {
	mongoIds := make([]interface{}, 0, len(ids))
	for i := range ids {
		mongoId, err := common.ToObjectId(ids[i])
		if err == nil {
			mongoIds = append(mongoIds, mongoId)
		}
	}
	return common.GetInFilter("_id", mongoIds...)
}
