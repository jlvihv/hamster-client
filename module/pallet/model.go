package pallet

import (
	"fmt"
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

type RentalAgreement struct {
	Index      types.U64
	Provider   types.AccountID
	TenantInfo struct {
		AccountId types.AccountID
		PublicKey string
	}
	PeerId        string
	ResourceIndex types.U64
	Config        struct {
		Cpu      types.U64
		Memory   types.U64
		System   string
		CpuModel string
	}
	RentalInfo struct {
		RentUnitPrice types.U128
		RentDuration  types.U32
		EndOfRent     types.U32
	}
	PenaltyAmount types.U128
	ReceiveAmount types.U128
	Start         types.U32
	End           types.U32
	Calculation   types.U32
	Time          types.U64
	Status        RentalAgreementStatus
}

type ResourceOrder struct {
	Index      types.U64
	TenantInfo struct {
		AccountId types.AccountID
		PublicKey types.Text
	}
	ResourceIndex  types.U64
	Create         types.U32
	RentDuration   types.U32
	Time           Duration
	Status         OrderStatus
	AgreementIndex types.OptionU64
}

type Duration struct {
	Secs  types.U64
	Nanos types.U32
}

type OrderStatus struct {
	IsPending  bool
	IsFinished bool
	IsCanceled bool
}

func (m *OrderStatus) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	fmt.Println(b)

	if err != nil {
		return err
	}

	if b == 0 {
		m.IsPending = true
	} else if b == 1 {
		m.IsFinished = true
	} else if b == 2 {
		m.IsCanceled = true
	}

	if err != nil {
		return err
	}

	return nil
}

func (m *OrderStatus) Encode(encoder scale.Encoder) error {
	var err1 error
	if m.IsPending {
		err1 = encoder.PushByte(0)
	} else if m.IsFinished {
		err1 = encoder.PushByte(1)
	} else if m.IsCanceled {
		err1 = encoder.PushByte(2)
	}
	if err1 != nil {
		return err1
	}
	return nil
}

type AccountInfo struct {
	Address string
	Amount  types.U128
}

type RentalAgreementStatus struct {
	// 使用
	IsUsing bool
	// 完成
	IsFinished bool
	// 被惩罚
	IsPunished bool
}

func (m *RentalAgreementStatus) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	fmt.Println(b)

	if err != nil {
		return err
	}

	if b == 0 {
		m.IsUsing = true
	} else if b == 1 {
		m.IsFinished = true
	} else if b == 2 {
		m.IsPunished = true
	}

	if err != nil {
		return err
	}

	return nil
}

func (m *RentalAgreementStatus) Encode(encoder scale.Encoder) error {
	var err1 error
	if m.IsUsing {
		err1 = encoder.PushByte(0)
	} else if m.IsFinished {
		err1 = encoder.PushByte(1)
	} else if m.IsPunished {
		err1 = encoder.PushByte(2)
	}
	if err1 != nil {
		return err1
	}
	return nil
}
