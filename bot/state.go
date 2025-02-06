package bot

import "sync"

type UserState struct {
	kind string
	data interface{}
}

type State struct {
	usersMap map[int64]*UserState
	mutex    sync.RWMutex
}

func NewState() *State {
	return &State{
		usersMap: make(map[int64]*UserState),
	}
}

func (s *State) GetState(userId int64) (string, interface{}) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	state, ok := s.usersMap[userId]

	if !ok {
		return "", nil
	}

	return state.kind, state.data
}

func (s *State) SetState(userId int64, kind string, data interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.usersMap[userId] = &UserState{
		kind: kind,
		data: data,
	}
}

func (s *State) ClearState(userId int64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.usersMap, userId)
}
