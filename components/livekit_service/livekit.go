package lksv

import (
	"context"
	lkauth "github.com/livekit/protocol/auth"
	livekit "github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go"
	"time"
)

type LiveKitService interface {
	CreateRoom(ctx context.Context, roomName string) (*livekit.Room, error)
	CreateJoinToken(room, identity string) (string, error)
	AuthProvider() lkauth.KeyProvider
}

type liveKitService struct {
	roomClient         *lksdk.RoomServiceClient
	rooms              map[string]*livekit.Room
	apiKey             string
	apiSecret          string
	host               string
	timeout            uint32
	maximumParticipant uint32
}

func NewLiveKitService(
	apiKey string,
	apiSecret string,
	host string,
	timeout uint32,
	maximumParticipant uint32,
) LiveKitService {
	return &liveKitService{
		roomClient:         lksdk.NewRoomServiceClient(host, apiKey, apiSecret),
		rooms:              map[string]*livekit.Room{},
		apiSecret:          apiSecret,
		apiKey:             apiKey,
		host:               host,
		timeout:            timeout,
		maximumParticipant: maximumParticipant,
	}
}

func (s *liveKitService) CreateRoom(ctx context.Context, roomName string) (*livekit.Room, error) {
	if room, ok := s.rooms[roomName]; ok {
		return room, nil
	}
	room, err := s.roomClient.CreateRoom(ctx, &livekit.CreateRoomRequest{
		Name:            roomName,
		EmptyTimeout:    s.timeout,
		MaxParticipants: s.maximumParticipant,
	})

	if err != nil {
		return nil, err
	}
	s.rooms[roomName] = room
	return s.rooms[roomName], nil
}

func (s *liveKitService) CreateJoinToken(room, identity string) (string, error) {
	at := lkauth.NewAccessToken(s.apiKey, s.apiSecret)
	t := true
	videoGrant := &lkauth.VideoGrant{
		Room:         room,
		RoomJoin:     true,
		CanPublish:   &t,
		CanSubscribe: &t,
	}

	at.AddGrant(videoGrant).
		SetIdentity(identity).
		SetValidFor(time.Hour)
	return at.ToJWT()
}

func (s *liveKitService) AuthProvider() lkauth.KeyProvider {
	return lkauth.NewSimpleKeyProvider(
		s.apiKey, s.apiSecret,
	)
}
