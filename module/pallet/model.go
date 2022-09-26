package pallet

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type ComputingResource struct {
	Index     types.U64       `json:"index"`
	AccountId types.AccountID `json:"accountId"`
	PeerId    types.Text      `json:"peerId"`
	Config    struct {
		Cpu      types.U64  `json:"cpu"`
		Memory   types.U64  `json:"memory"`
		System   types.Text `json:"system"`
		CpuModel types.Text `json:"cpuModel"`
	} `json:"config"`
	RentalStatistics struct {
		RentalCount    types.U32 `json:"rentalCount"`
		RentalDuration types.U32 `json:"rentalDuration"`
		FaultCount     types.U32 `json:"faultCount"`
		FaultDuration  types.U32 `json:"faultDuration"`
	} `json:"rentalStatistics"`
	RentalInfo struct {
		RentUnitPrice types.U128 `json:"rentUnitPrice"`
		RentDuration  types.U32  `json:"rentDuration"`
		EndOfRent     types.U32  `json:"endOfRent"`
	} `json:"rentalInfo"`
	Status Status `json:"status"`
}

type Status struct {
	IsInuse   bool `json:"isInuse"`
	IsLocked  bool `json:"isLocked"`
	IsUnused  bool `json:"isUnused"`
	IsOffline bool `json:"isOffline"`
}

func (s *Status) toString() string {
	if s.IsInuse {
		return "Inuse"
	} else if s.IsLocked {
		return "Locked"
	} else if s.IsUnused {
		return "Unuse"
	} else if s.IsOffline {
		return "Offline"
	} else {
		return ""
	}
}

func (m *Status) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	if b == 0 {
		m.IsInuse = true
	} else if b == 1 {
		m.IsLocked = true
	} else if b == 2 {
		m.IsUnused = true
	} else if b == 3 {
		m.IsOffline = true
	}

	if err != nil {
		return err
	}

	return nil
}

func (m *Status) Encode(encoder scale.Encoder) error {
	var err1 error
	if m.IsInuse {
		err1 = encoder.PushByte(0)
	} else if m.IsLocked {
		err1 = encoder.PushByte(1)
	} else if m.IsUnused {
		err1 = encoder.PushByte(2)
	} else if m.IsOffline {
		err1 = encoder.PushByte(3)
	}
	if err1 != nil {
		return err1
	}
	return nil
}
