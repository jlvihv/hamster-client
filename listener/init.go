package listener

import "hamster-client/module/pallet"

type Listener struct {
	listener pallet.ChainListener
}

func (l *Listener) init() {
	l.listener.CancelListen()
	l.listener.StartListen()
}

func NewListener() {
	l := &Listener{}
	l.init()
}
