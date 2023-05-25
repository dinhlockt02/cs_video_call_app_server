package friendrepo

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	requestmdl "github.com/dinhlockt02/cs_video_call_app_server/modules/request/model"
	requeststore "github.com/dinhlockt02/cs_video_call_app_server/modules/request/store"
)

// FindRequest returns the friend request between sender and receiver
// If the request does not exist, it returns nil, nil
func (repo *friendRepository) FindRequest(
	ctx context.Context,
	sender string,
	receiver string,
) (*requestmdl.Request, error) {

	senderFilter := requeststore.GetRequestSenderIdFilter(sender)
	receiverFilter := requeststore.GetRequestReceiverIdFilter(receiver)
	filter := common.GetAndFilter(senderFilter, receiverFilter)

	existedRequest, err := repo.requestStore.FindRequest(ctx, filter)
	if err != nil {
		return nil, err
	}
	return existedRequest, nil
}
