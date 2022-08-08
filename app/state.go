package app

import "hamster-client/module/state"

type State struct {
	service state.Service
}

func NewStateApp(service state.Service) State {
	return State{
		service: service,
	}
}

func (s *State) GetStateInfo(key string) (state.Info, error) {
	info, err := s.service.GetStateInfo(key)
	if err != nil {
		return state.Info{}, err
	}
	return info, nil
}
