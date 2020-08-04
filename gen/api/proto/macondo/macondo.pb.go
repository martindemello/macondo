// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: api/proto/macondo/macondo.proto

package macondo

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PlayState int32

const (
	PlayState_PLAYING                PlayState = 0
	PlayState_WAITING_FOR_FINAL_PASS PlayState = 1
	PlayState_GAME_OVER              PlayState = 2
)

// Enum value maps for PlayState.
var (
	PlayState_name = map[int32]string{
		0: "PLAYING",
		1: "WAITING_FOR_FINAL_PASS",
		2: "GAME_OVER",
	}
	PlayState_value = map[string]int32{
		"PLAYING":                0,
		"WAITING_FOR_FINAL_PASS": 1,
		"GAME_OVER":              2,
	}
)

func (x PlayState) Enum() *PlayState {
	p := new(PlayState)
	*p = x
	return p
}

func (x PlayState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayState) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_macondo_macondo_proto_enumTypes[0].Descriptor()
}

func (PlayState) Type() protoreflect.EnumType {
	return &file_api_proto_macondo_macondo_proto_enumTypes[0]
}

func (x PlayState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayState.Descriptor instead.
func (PlayState) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_macondo_macondo_proto_rawDescGZIP(), []int{0}
}

type ChallengeRule int32

const (
	ChallengeRule_VOID       ChallengeRule = 0
	ChallengeRule_SINGLE     ChallengeRule = 1
	ChallengeRule_DOUBLE     ChallengeRule = 2
	ChallengeRule_FIVE_POINT ChallengeRule = 3
	ChallengeRule_TEN_POINT  ChallengeRule = 4
)

// Enum value maps for ChallengeRule.
var (
	ChallengeRule_name = map[int32]string{
		0: "VOID",
		1: "SINGLE",
		2: "DOUBLE",
		3: "FIVE_POINT",
		4: "TEN_POINT",
	}
	ChallengeRule_value = map[string]int32{
		"VOID":       0,
		"SINGLE":     1,
		"DOUBLE":     2,
		"FIVE_POINT": 3,
		"TEN_POINT":  4,
	}
)

func (x ChallengeRule) Enum() *ChallengeRule {
	p := new(ChallengeRule)
	*p = x
	return p
}

func (x ChallengeRule) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChallengeRule) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_macondo_macondo_proto_enumTypes[1].Descriptor()
}

func (ChallengeRule) Type() protoreflect.EnumType {
	return &file_api_proto_macondo_macondo_proto_enumTypes[1]
}

func (x ChallengeRule) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChallengeRule.Descriptor instead.
func (ChallengeRule) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_macondo_macondo_proto_rawDescGZIP(), []int{1}
}

type GameEvent_Type int32

const (
	GameEvent_TILE_PLACEMENT_MOVE  GameEvent_Type = 0
	GameEvent_PHONY_TILES_RETURNED GameEvent_Type = 1
	GameEvent_PASS                 GameEvent_Type = 2
	GameEvent_CHALLENGE_BONUS      GameEvent_Type = 3
	GameEvent_EXCHANGE             GameEvent_Type = 4
	GameEvent_END_RACK_PTS         GameEvent_Type = 5
	GameEvent_TIME_PENALTY         GameEvent_Type = 6
	// Only for international rules (or after 6 zeroes)
	GameEvent_END_RACK_PENALTY GameEvent_Type = 7
	// Lose a turn for challenging a word incorrectly (only for double
	// challenge)
	GameEvent_UNSUCCESSFUL_CHALLENGE_TURN_LOSS GameEvent_Type = 8
)

// Enum value maps for GameEvent_Type.
var (
	GameEvent_Type_name = map[int32]string{
		0: "TILE_PLACEMENT_MOVE",
		1: "PHONY_TILES_RETURNED",
		2: "PASS",
		3: "CHALLENGE_BONUS",
		4: "EXCHANGE",
		5: "END_RACK_PTS",
		6: "TIME_PENALTY",
		7: "END_RACK_PENALTY",
		8: "UNSUCCESSFUL_CHALLENGE_TURN_LOSS",
	}
	GameEvent_Type_value = map[string]int32{
		"TILE_PLACEMENT_MOVE":              0,
		"PHONY_TILES_RETURNED":             1,
		"PASS":                             2,
		"CHALLENGE_BONUS":                  3,
		"EXCHANGE":                         4,
		"END_RACK_PTS":                     5,
		"TIME_PENALTY":                     6,
		"END_RACK_PENALTY":                 7,
		"UNSUCCESSFUL_CHALLENGE_TURN_LOSS": 8,
	}
)

func (x GameEvent_Type) Enum() *GameEvent_Type {
	p := new(GameEvent_Type)
	*p = x
	return p
}

func (x GameEvent_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameEvent_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_macondo_macondo_proto_enumTypes[2].Descriptor()
}

func (GameEvent_Type) Type() protoreflect.EnumType {
	return &file_api_proto_macondo_macondo_proto_enumTypes[2]
}

func (x GameEvent_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameEvent_Type.Descriptor instead.
func (GameEvent_Type) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_macondo_macondo_proto_rawDescGZIP(), []int{1, 0}
}

type GameEvent_Direction int32

const (
	GameEvent_HORIZONTAL GameEvent_Direction = 0
	GameEvent_VERTICAL   GameEvent_Direction = 1
)

// Enum value maps for GameEvent_Direction.
var (
	GameEvent_Direction_name = map[int32]string{
		0: "HORIZONTAL",
		1: "VERTICAL",
	}
	GameEvent_Direction_value = map[string]int32{
		"HORIZONTAL": 0,
		"VERTICAL":   1,
	}
)

func (x GameEvent_Direction) Enum() *GameEvent_Direction {
	p := new(GameEvent_Direction)
	*p = x
	return p
}

func (x GameEvent_Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameEvent_Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_macondo_macondo_proto_enumTypes[3].Descriptor()
}

func (GameEvent_Direction) Type() protoreflect.EnumType {
	return &file_api_proto_macondo_macondo_proto_enumTypes[3]
}

func (x GameEvent_Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameEvent_Direction.Descriptor instead.
func (GameEvent_Direction) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_macondo_macondo_proto_rawDescGZIP(), []int{1, 1}
}

// GameHistory encodes a whole history of a game, and it should also encode
// the initial board and tile configuration, etc. It can be considered
// to be an instantiation of a GCG file.
type GameHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events      []*GameEvent  `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	Players     []*PlayerInfo `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
	Version     int32         `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	OriginalGcg string        `protobuf:"bytes,4,opt,name=original_gcg,json=originalGcg,proto3" json:"original_gcg,omitempty"`
	Lexicon     string        `protobuf:"bytes,5,opt,name=lexicon,proto3" json:"lexicon,omitempty"`
	IdAuth      string        `protobuf:"bytes,6,opt,name=id_auth,json=idAuth,proto3" json:"id_auth,omitempty"`
	Uid         string        `protobuf:"bytes,7,opt,name=uid,proto3" json:"uid,omitempty"`
	Title       string        `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`
	Description string        `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	// last_known_racks should only be set in an incomplete / in-progress game.
	// if set, player racks should be set to these values.
	LastKnownRacks []string `protobuf:"bytes,10,rep,name=last_known_racks,json=lastKnownRacks,proto3" json:"last_known_racks,omitempty"`
	// If second_went_first is set, the second player in `players` actually
	// went first. not that this does NOT change the order of `last_known_racks`,
	// which is always in the order of the listed players!
	SecondWentFirst bool          `protobuf:"varint,11,opt,name=second_went_first,json=secondWentFirst,proto3" json:"second_went_first,omitempty"`
	ChallengeRule   ChallengeRule `protobuf:"varint,12,opt,name=challenge_rule,json=challengeRule,proto3,enum=macondo.ChallengeRule" json:"challenge_rule,omitempty"`
	PlayState       PlayState     `protobuf:"varint,13,opt,name=play_state,json=playState,proto3,enum=macondo.PlayState" json:"play_state,omitempty"`
	// Final scores of the game; the order is in the order of the listed players!
	FinalScores []int32 `protobuf:"varint,14,rep,packed,name=final_scores,json=finalScores,proto3" json:"final_scores,omitempty"`
	// The variant, together with the lexicon, would encode the initial board
	// and tile configuration.
	Variant string `protobuf:"bytes,15,opt,name=variant,proto3" json:"variant,omitempty"`
	// The index of the player who won. It's not always the person with the
	// highest score, because there can be timeouts, etc. If it's a tie,
	// it will be a -1.
	Winner int32 `protobuf:"varint,16,opt,name=winner,proto3" json:"winner,omitempty"`
}

func (x *GameHistory) Reset() {
	*x = GameHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_macondo_macondo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameHistory) ProtoMessage() {}

func (x *GameHistory) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_macondo_macondo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameHistory.ProtoReflect.Descriptor instead.
func (*GameHistory) Descriptor() ([]byte, []int) {
	return file_api_proto_macondo_macondo_proto_rawDescGZIP(), []int{0}
}

func (x *GameHistory) GetEvents() []*GameEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *GameHistory) GetPlayers() []*PlayerInfo {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *GameHistory) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *GameHistory) GetOriginalGcg() string {
	if x != nil {
		return x.OriginalGcg
	}
	return ""
}

func (x *GameHistory) GetLexicon() string {
	if x != nil {
		return x.Lexicon
	}
	return ""
}

func (x *GameHistory) GetIdAuth() string {
	if x != nil {
		return x.IdAuth
	}
	return ""
}

func (x *GameHistory) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *GameHistory) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GameHistory) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GameHistory) GetLastKnownRacks() []string {
	if x != nil {
		return x.LastKnownRacks
	}
	return nil
}

func (x *GameHistory) GetSecondWentFirst() bool {
	if x != nil {
		return x.SecondWentFirst
	}
	return false
}

func (x *GameHistory) GetChallengeRule() ChallengeRule {
	if x != nil {
		return x.ChallengeRule
	}
	return ChallengeRule_VOID
}

func (x *GameHistory) GetPlayState() PlayState {
	if x != nil {
		return x.PlayState
	}
	return PlayState_PLAYING
}

func (x *GameHistory) GetFinalScores() []int32 {
	if x != nil {
		return x.FinalScores
	}
	return nil
}

func (x *GameHistory) GetVariant() string {
	if x != nil {
		return x.Variant
	}
	return ""
}

func (x *GameHistory) GetWinner() int32 {
	if x != nil {
		return x.Winner
	}
	return 0
}

// This should be merged into Move.
type GameEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname    string              `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Note        string              `protobuf:"bytes,2,opt,name=note,proto3" json:"note,omitempty"`
	Rack        string              `protobuf:"bytes,3,opt,name=rack,proto3" json:"rack,omitempty"`
	Type        GameEvent_Type      `protobuf:"varint,4,opt,name=type,proto3,enum=macondo.GameEvent_Type" json:"type,omitempty"`
	Cumulative  int32               `protobuf:"varint,5,opt,name=cumulative,proto3" json:"cumulative,omitempty"`
	Row         int32               `protobuf:"varint,6,opt,name=row,proto3" json:"row,omitempty"`
	Column      int32               `protobuf:"varint,7,opt,name=column,proto3" json:"column,omitempty"`
	Direction   GameEvent_Direction `protobuf:"varint,8,opt,name=direction,proto3,enum=macondo.GameEvent_Direction" json:"direction,omitempty"`
	Position    string              `protobuf:"bytes,9,opt,name=position,proto3" json:"position,omitempty"`
	PlayedTiles string              `protobuf:"bytes,10,opt,name=played_tiles,json=playedTiles,proto3" json:"played_tiles,omitempty"`
	// An event will not have all of these; it depends on the type of the event.
	Exchanged     string `protobuf:"bytes,11,opt,name=exchanged,proto3" json:"exchanged,omitempty"`
	Score         int32  `protobuf:"varint,12,opt,name=score,proto3" json:"score,omitempty"`
	Bonus         int32  `protobuf:"varint,13,opt,name=bonus,proto3" json:"bonus,omitempty"`
	EndRackPoints int32  `protobuf:"varint,14,opt,name=end_rack_points,json=endRackPoints,proto3" json:"end_rack_points,omitempty"`
	LostScore     int32  `protobuf:"varint,15,opt,name=lost_score,json=lostScore,proto3" json:"lost_score,omitempty"`
	IsBingo       bool   `protobuf:"varint,16,opt,name=is_bingo,json=isBingo,proto3" json:"is_bingo,omitempty"`
	// words_formed is a list of all words made by this play, in user-visible
	// pretty form. The first word is the "main" word, anything after it are
	// cross-words.
	WordsFormed     []string `protobuf:"bytes,17,rep,name=words_formed,json=wordsFormed,proto3" json:"words_formed,omitempty"`
	MillisRemaining int32    `protobuf:"varint,18,opt,name=millis_remaining,json=millisRemaining,proto3" json:"millis_remaining,omitempty"`
}

func (x *GameEvent) Reset() {
	*x = GameEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_macondo_macondo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameEvent) ProtoMessage() {}

func (x *GameEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_macondo_macondo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameEvent.ProtoReflect.Descriptor instead.
func (*GameEvent) Descriptor() ([]byte, []int) {
	return file_api_proto_macondo_macondo_proto_rawDescGZIP(), []int{1}
}

func (x *GameEvent) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *GameEvent) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *GameEvent) GetRack() string {
	if x != nil {
		return x.Rack
	}
	return ""
}

func (x *GameEvent) GetType() GameEvent_Type {
	if x != nil {
		return x.Type
	}
	return GameEvent_TILE_PLACEMENT_MOVE
}

func (x *GameEvent) GetCumulative() int32 {
	if x != nil {
		return x.Cumulative
	}
	return 0
}

func (x *GameEvent) GetRow() int32 {
	if x != nil {
		return x.Row
	}
	return 0
}

func (x *GameEvent) GetColumn() int32 {
	if x != nil {
		return x.Column
	}
	return 0
}

func (x *GameEvent) GetDirection() GameEvent_Direction {
	if x != nil {
		return x.Direction
	}
	return GameEvent_HORIZONTAL
}

func (x *GameEvent) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *GameEvent) GetPlayedTiles() string {
	if x != nil {
		return x.PlayedTiles
	}
	return ""
}

func (x *GameEvent) GetExchanged() string {
	if x != nil {
		return x.Exchanged
	}
	return ""
}

func (x *GameEvent) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *GameEvent) GetBonus() int32 {
	if x != nil {
		return x.Bonus
	}
	return 0
}

func (x *GameEvent) GetEndRackPoints() int32 {
	if x != nil {
		return x.EndRackPoints
	}
	return 0
}

func (x *GameEvent) GetLostScore() int32 {
	if x != nil {
		return x.LostScore
	}
	return 0
}

func (x *GameEvent) GetIsBingo() bool {
	if x != nil {
		return x.IsBingo
	}
	return false
}

func (x *GameEvent) GetWordsFormed() []string {
	if x != nil {
		return x.WordsFormed
	}
	return nil
}

func (x *GameEvent) GetMillisRemaining() int32 {
	if x != nil {
		return x.MillisRemaining
	}
	return 0
}

type PlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	RealName string `protobuf:"bytes,2,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	// user_id is an internal, unchangeable user ID, whereas the other two user
	// identifiers might possibly be mutable.
	UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *PlayerInfo) Reset() {
	*x = PlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_macondo_macondo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInfo) ProtoMessage() {}

func (x *PlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_macondo_macondo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInfo.ProtoReflect.Descriptor instead.
func (*PlayerInfo) Descriptor() ([]byte, []int) {
	return file_api_proto_macondo_macondo_proto_rawDescGZIP(), []int{2}
}

func (x *PlayerInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *PlayerInfo) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *PlayerInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type BotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameHistory *GameHistory `protobuf:"bytes,1,opt,name=game_history,json=gameHistory,proto3" json:"game_history,omitempty"`
}

func (x *BotRequest) Reset() {
	*x = BotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_macondo_macondo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotRequest) ProtoMessage() {}

func (x *BotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_macondo_macondo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotRequest.ProtoReflect.Descriptor instead.
func (*BotRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_macondo_macondo_proto_rawDescGZIP(), []int{3}
}

func (x *BotRequest) GetGameHistory() *GameHistory {
	if x != nil {
		return x.GameHistory
	}
	return nil
}

type BotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*BotResponse_Move
	//	*BotResponse_Error
	Response isBotResponse_Response `protobuf_oneof:"response"`
}

func (x *BotResponse) Reset() {
	*x = BotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_macondo_macondo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotResponse) ProtoMessage() {}

func (x *BotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_macondo_macondo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotResponse.ProtoReflect.Descriptor instead.
func (*BotResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_macondo_macondo_proto_rawDescGZIP(), []int{4}
}

func (m *BotResponse) GetResponse() isBotResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *BotResponse) GetMove() *GameEvent {
	if x, ok := x.GetResponse().(*BotResponse_Move); ok {
		return x.Move
	}
	return nil
}

func (x *BotResponse) GetError() string {
	if x, ok := x.GetResponse().(*BotResponse_Error); ok {
		return x.Error
	}
	return ""
}

type isBotResponse_Response interface {
	isBotResponse_Response()
}

type BotResponse_Move struct {
	Move *GameEvent `protobuf:"bytes,1,opt,name=move,proto3,oneof"`
}

type BotResponse_Error struct {
	Error string `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*BotResponse_Move) isBotResponse_Response() {}

func (*BotResponse_Error) isBotResponse_Response() {}

var File_api_proto_macondo_macondo_proto protoreflect.FileDescriptor

var file_api_proto_macondo_macondo_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x63, 0x6f,
	0x6e, 0x64, 0x6f, 0x2f, 0x6d, 0x61, 0x63, 0x6f, 0x6e, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x6d, 0x61, 0x63, 0x6f, 0x6e, 0x64, 0x6f, 0x22, 0xbf, 0x04, 0x0a, 0x0b, 0x47,
	0x61, 0x6d, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x61, 0x63,
	0x6f, 0x6e, 0x64, 0x6f, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x61, 0x63, 0x6f, 0x6e, 0x64,
	0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x67, 0x63, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x47,
	0x63, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x65, 0x78, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x65, 0x78, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07,
	0x69, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69,
	0x64, 0x41, 0x75, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x28, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x72, 0x61,
	0x63, 0x6b, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x4b,
	0x6e, 0x6f, 0x77, 0x6e, 0x52, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x5f, 0x77, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x57, 0x65, 0x6e, 0x74,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x6d, 0x61, 0x63, 0x6f, 0x6e, 0x64, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x52, 0x75, 0x6c, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6d, 0x61, 0x63, 0x6f, 0x6e,
	0x64, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x09, 0x70, 0x6c,
	0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x6e, 0x61, 0x6c,
	0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x66,
	0x69, 0x6e, 0x61, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0xaf, 0x06, 0x0a,
	0x09, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61,
	0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x2b,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d,
	0x61, 0x63, 0x6f, 0x6e, 0x64, 0x6f, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72,
	0x6f, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x6f, 0x77, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x3a, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6d, 0x61, 0x63, 0x6f, 0x6e,
	0x64, 0x6f, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x54, 0x69, 0x6c, 0x65, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x6e,
	0x64, 0x5f, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x65, 0x6e, 0x64, 0x52, 0x61, 0x63, 0x6b, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x6f, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x62, 0x69, 0x6e, 0x67, 0x6f, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x42, 0x69, 0x6e, 0x67, 0x6f, 0x12, 0x21, 0x0a, 0x0c,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x11, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x46, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x12,
	0x29, 0x0a, 0x10, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6d, 0x69, 0x6c, 0x6c, 0x69,
	0x73, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0xc6, 0x01, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x49, 0x4c, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x43,
	0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14,
	0x50, 0x48, 0x4f, 0x4e, 0x59, 0x5f, 0x54, 0x49, 0x4c, 0x45, 0x53, 0x5f, 0x52, 0x45, 0x54, 0x55,
	0x52, 0x4e, 0x45, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x41, 0x53, 0x53, 0x10, 0x02,
	0x12, 0x13, 0x0a, 0x0f, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45, 0x5f, 0x42, 0x4f,
	0x4e, 0x55, 0x53, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47,
	0x45, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x4e, 0x44, 0x5f, 0x52, 0x41, 0x43, 0x4b, 0x5f,
	0x50, 0x54, 0x53, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x50, 0x45,
	0x4e, 0x41, 0x4c, 0x54, 0x59, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x44, 0x5f, 0x52,
	0x41, 0x43, 0x4b, 0x5f, 0x50, 0x45, 0x4e, 0x41, 0x4c, 0x54, 0x59, 0x10, 0x07, 0x12, 0x24, 0x0a,
	0x20, 0x55, 0x4e, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x46, 0x55, 0x4c, 0x5f, 0x43, 0x48,
	0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45, 0x5f, 0x54, 0x55, 0x52, 0x4e, 0x5f, 0x4c, 0x4f, 0x53,
	0x53, 0x10, 0x08, 0x22, 0x29, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x0a, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x4f, 0x4e, 0x54, 0x41, 0x4c, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x56, 0x45, 0x52, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x01, 0x22, 0x5e,
	0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x6c,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x45,
	0x0a, 0x0a, 0x42, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0c,
	0x67, 0x61, 0x6d, 0x65, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x61, 0x63, 0x6f, 0x6e, 0x64, 0x6f, 0x2e, 0x47, 0x61, 0x6d,
	0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x0b, 0x67, 0x61, 0x6d, 0x65, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x5b, 0x0a, 0x0b, 0x42, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x61, 0x63, 0x6f, 0x6e, 0x64, 0x6f, 0x2e, 0x47, 0x61, 0x6d,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x04, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x16,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2a, 0x43, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x50, 0x4c, 0x41, 0x59, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16,
	0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x46, 0x49, 0x4e, 0x41,
	0x4c, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x41, 0x4d, 0x45,
	0x5f, 0x4f, 0x56, 0x45, 0x52, 0x10, 0x02, 0x2a, 0x50, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x4f, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x49, 0x4e, 0x47, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x49,
	0x56, 0x45, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x45,
	0x4e, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x10, 0x04, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x6d, 0x69, 0x6e, 0x6f, 0x31, 0x34,
	0x2f, 0x6d, 0x61, 0x63, 0x6f, 0x6e, 0x64, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x63, 0x6f, 0x6e, 0x64, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_macondo_macondo_proto_rawDescOnce sync.Once
	file_api_proto_macondo_macondo_proto_rawDescData = file_api_proto_macondo_macondo_proto_rawDesc
)

func file_api_proto_macondo_macondo_proto_rawDescGZIP() []byte {
	file_api_proto_macondo_macondo_proto_rawDescOnce.Do(func() {
		file_api_proto_macondo_macondo_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_macondo_macondo_proto_rawDescData)
	})
	return file_api_proto_macondo_macondo_proto_rawDescData
}

var file_api_proto_macondo_macondo_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_api_proto_macondo_macondo_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_proto_macondo_macondo_proto_goTypes = []interface{}{
	(PlayState)(0),           // 0: macondo.PlayState
	(ChallengeRule)(0),       // 1: macondo.ChallengeRule
	(GameEvent_Type)(0),      // 2: macondo.GameEvent.Type
	(GameEvent_Direction)(0), // 3: macondo.GameEvent.Direction
	(*GameHistory)(nil),      // 4: macondo.GameHistory
	(*GameEvent)(nil),        // 5: macondo.GameEvent
	(*PlayerInfo)(nil),       // 6: macondo.PlayerInfo
	(*BotRequest)(nil),       // 7: macondo.BotRequest
	(*BotResponse)(nil),      // 8: macondo.BotResponse
}
var file_api_proto_macondo_macondo_proto_depIdxs = []int32{
	5, // 0: macondo.GameHistory.events:type_name -> macondo.GameEvent
	6, // 1: macondo.GameHistory.players:type_name -> macondo.PlayerInfo
	1, // 2: macondo.GameHistory.challenge_rule:type_name -> macondo.ChallengeRule
	0, // 3: macondo.GameHistory.play_state:type_name -> macondo.PlayState
	2, // 4: macondo.GameEvent.type:type_name -> macondo.GameEvent.Type
	3, // 5: macondo.GameEvent.direction:type_name -> macondo.GameEvent.Direction
	4, // 6: macondo.BotRequest.game_history:type_name -> macondo.GameHistory
	5, // 7: macondo.BotResponse.move:type_name -> macondo.GameEvent
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_api_proto_macondo_macondo_proto_init() }
func file_api_proto_macondo_macondo_proto_init() {
	if File_api_proto_macondo_macondo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_macondo_macondo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameHistory); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_macondo_macondo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_macondo_macondo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_macondo_macondo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BotRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_macondo_macondo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BotResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_api_proto_macondo_macondo_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*BotResponse_Move)(nil),
		(*BotResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_macondo_macondo_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_macondo_macondo_proto_goTypes,
		DependencyIndexes: file_api_proto_macondo_macondo_proto_depIdxs,
		EnumInfos:         file_api_proto_macondo_macondo_proto_enumTypes,
		MessageInfos:      file_api_proto_macondo_macondo_proto_msgTypes,
	}.Build()
	File_api_proto_macondo_macondo_proto = out.File
	file_api_proto_macondo_macondo_proto_rawDesc = nil
	file_api_proto_macondo_macondo_proto_goTypes = nil
	file_api_proto_macondo_macondo_proto_depIdxs = nil
}
