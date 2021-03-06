package whitenoise

import (
	"bufio"
	"errors"
	core "github.com/libp2p/go-libp2p-core"
	"github.com/libp2p/go-libp2p-core/network"
)

const SESSIONIDNON string = "SessionIDNon"

//TODO:add mutex
type Session struct {
	Id   string
	pair StreamPair
	Out  bool //mark if host the in/exit point of the circuit
}

type StreamPair []Stream

type Stream struct {
	StreamId   string
	RemotePeer core.PeerID
	RW         *bufio.ReadWriter
}

func NewStream(s network.Stream) Stream {
	return Stream{
		StreamId:   s.ID(),
		RW:         bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s)),
		RemotePeer: s.Conn().RemotePeer(),
	}
}

func NewSession() Session {
	return Session{
		Id:   SESSIONIDNON,
		pair: StreamPair{},
		Out:  false,
	}
}

func (s *Session) SetSessionID(sID string) {
	s.Id = sID
}

func (s *Session) AddStream(stream Stream) {
	if s.IsReady() {
		return
	}
	s.pair = append(s.pair, stream)
	for {
		if len(s.pair) > 2 {
			s.pair = s.pair[1:]
		} else {
			break
		}
	}
}

func (s *Session) Has(streamID string) bool {
	for _, s := range s.pair {
		if s.StreamId == streamID {
			return true
		}
	}
	return false
}

func (s *Session) GetPattern(streamID string) (Stream, error) {
	if len(s.pair) != 2 {
		return Stream{}, errors.New("session not ready")
	}
	for i, stream := range s.pair {
		if stream.StreamId == streamID {
			return s.pair[i^1], nil
		}
	}
	return Stream{}, errors.New("no such stream")
}

func (s *Session) GetPair() StreamPair {
	return s.pair
}

func (s *Session) IsReady() bool {
	return len(s.pair) == 2
}

func (s *Session) SetOutPoint() {
	s.Out = true
}

type SessionManager struct {
	StreamMap    map[string]Stream
	SessionmapID map[string]Session
}

func NewSessionMapper() SessionManager {
	return SessionManager{
		StreamMap:    make(map[string]Stream),
		SessionmapID: make(map[string]Session),
	}
}

func (man *SessionManager) AddSessionNonid(s Stream) {
	man.StreamMap[s.StreamId] = s
}

func (man *SessionManager) AddSessionId(id string, s Session) {
	man.SessionmapID[id] = s
}
