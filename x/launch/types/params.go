package types

import (
	"errors"
	"fmt"
	"time"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// TODO: Determine default values
var (
	DefaultMinLaunchTime = uint64(time.Hour.Seconds() * 24)
	DefaultMaxLaunchTime = uint64(time.Hour.Seconds() * 24 * 7)

	MaxParametrableLaunchTime = uint64(time.Hour.Seconds() * 24 * 31)

	KeyMinLaunchTime = []byte("MinLaunchTime")
	KeyMaxLaunchTime = []byte("MaxLaunchTime")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamTable for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(minLaunchTime, maxLaunchTime uint64) Params {
	return Params{
		MinLaunchTime: minLaunchTime,
		MaxLaunchTime: maxLaunchTime,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMinLaunchTime,
		DefaultMaxLaunchTime,
	)
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMinLaunchTime, &p.MinLaunchTime, validateLaunchTime),
		paramtypes.NewParamSetPair(KeyMaxLaunchTime, &p.MaxLaunchTime, validateLaunchTime),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if p.MinLaunchTime > p.MaxLaunchTime {
		return errors.New("MinLaunchTime can't be higher than MaxLaunchTime")
	}
	if err := validateLaunchTime(p.MinLaunchTime); err != nil {
		return err
	}
	return validateLaunchTime(p.MaxLaunchTime)
}

func validateLaunchTime(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v > MaxParametrableLaunchTime {
		return errors.New("max parametrable launch time reached")
	}
	return nil
}
