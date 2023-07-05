# Notification documentation
## Notification structure
### Channel
There are 2 types of channel at the moment:
1. basic_channel
2. call_channel
### Notification Payload
- Notification `payload` is a Json-Serialized of the Notification. Those key below are required in payload
```json
{
  "id": "the notification id",
  "created_at": "the time notification was created",
  "owner": "the id of user who will receive notification",
  "action": "the type of action, this is unique between each type of notification",
  // ... other optional keys ...
}
```
- Depends on which type of notification, payload also may contains `subject`, `direct`, `indirect` and `prep` key. In each key, e.g. `subject`, the data is
```json
{
  "id": "string",
  "name": "string",
  "image": "nullable string",
  "type": "NotificationObjectType"
}
```
- The `NotificationObjectType` is defined as
  - `User`: "user"
  - `CallRoom`: "call-room"
  - `Request`: "request"
  - `Group`: "group"
- The `NotificationActionType` is defined in each type of notification's details
## Types of notification
1. [AcceptRequest](#acceptrequest)
2. [ReceiveFriendRequest](#receivefriendrequest)
3. [ReceiveGroupRequest](#receivegrouprequest)
4. [InComingCall](#incomingcall)
5. [RejectCall](#rejectcall)
6. [AbandonCall](#abandoncall)
## Notification Details
### AcceptRequest
- NotificationId: 1
- Description: the subject accept the indirect (aka owner)'s friend request
- Channel: "basic_channel"
- Action Button: []
- Notification Action Type: "accept-request"
- Object: 
  - Subject: User
  - Indirect: User
### ReceiveFriendRequest
- NotificationId: 2
- Description: the Subject (aka owner) received the friend request from Prep's
- Channel: "basic_channel"
- Action Button: ["Accept", "Deny"]
- Notification Action Type: "receive-friend-request"
- Object:
    - Subject: User
    - Prep: User
### ReceiveGroupRequest
- NotificationId: 3
- Description: the Subject (aka owner) received the group request (Direct) to Group (Indirect) from Prep's
- Channel: "basic_channel"
- Action Button: ["Accept", "Deny"]
- Notification Action Type: "receive-group-request"
- Object:
    - Subject: User
    - Direct: Request
    - Indirect: Group
    - Prep: User
### InComingCall
- NotificationId: 4
- Description: the Subject call the Direct (aka owner) in a room (Prep)
- Channel: "call_channel"
- Action Button: ["Accept", "Deny"]
- Notification Action Type: "incoming-call"
- Object:
    - Subject: User
    - Direct: Request
    - Prep: CallRoom
### RejectCall
- NotificationId: 5
- Description: the Subject reject the Direct (aka owner) in a room (Prep)
- Channel: "call_channel"
- Action Button: []
- Notification Action Type: "reject-call"
- Object:
    - Subject: User
    - Direct: User
    - Prep: CallRoom
### AbandonCall
- NotificationId: 6
- Description: when the Subject abandon call before the Direct (aka owner) answer in a room (Prep)
- Channel: "call_channel"
- Action Button: []
- Notification Action Type: "abandon-call"
- Object:
    - Subject: User
    - Direct: User
    - Prep: CallRoom

