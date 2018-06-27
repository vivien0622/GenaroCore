package vm

import (
	"github.com/GenaroNetwork/Genaro-Core/core/types"
	"math/big"
	"time"
	"github.com/GenaroNetwork/Genaro-Core/common"
	"errors"
)

func CheckSpecialTxTypeSyncSidechainStatusParameter( s types.SpecialTxInput) error {
	if 64 != len(s.SpecialTxTypeMortgageInit.Dataversion) {
		return errors.New("Parameter Dataversion  error")
	}

	if 64 != len(s.SpecialTxTypeMortgageInit.FileID) {
		return errors.New("Parameter fileID  error")
	}
	if 20 != len(s.SpecialTxTypeMortgageInit.FromAccount) {
		return errors.New("Parameter fromAccount  error")
	}
	if 1 < len(s.SpecialTxTypeMortgageInit.Sidechain) {
		for k,v := range s.SpecialTxTypeMortgageInit.Sidechain{
			if 20 != len(k) {
				return errors.New("Parameter mortgage account  error")
			}
			if v.ToInt().Cmp(big.NewInt(0)) < 0 {
				return errors.New("Parameter Sidechain")
			}
		}
	} else {
		return errors.New("Parameter side chain length less than zero")
	}
	return nil
}


func CheckspecialTxTypeMortgageInitParameter( s types.SpecialTxInput,caller common.Address) error {
	var tmp  big.Int
	timeLimit := s.SpecialTxTypeMortgageInit.TimeLimit.ToInt()
	tmp.Mul(timeLimit,big.NewInt(86400))
	endTime :=  tmp.Add(&tmp,big.NewInt(s.SpecialTxTypeMortgageInit.CreateTime)).Int64()
	if s.SpecialTxTypeMortgageInit.CreateTime > s.SpecialTxTypeMortgageInit.EndTime ||
		s.SpecialTxTypeMortgageInit.CreateTime > time.Now().Unix() ||
		s.SpecialTxTypeMortgageInit.EndTime != endTime {
		return errors.New("Parameter CreateTime or EndTime  error")
	}
	if caller != s.SpecialTxTypeMortgageInit.FromAccount {
		return errors.New("Parameter FromAccount  error")
	}
	if len(s.SpecialTxTypeMortgageInit.FileID) != 64 {
		return errors.New("Parameter FileID  error")
	}
	mortgageTable := s.SpecialTxTypeMortgageInit.MortgageTable
	authorityTable := s.SpecialTxTypeMortgageInit.AuthorityTable
	if len(authorityTable) != len(mortgageTable) {
		return errors.New("Parameter authorityTable != mortgageTable  error")
	}
	for k,v := range authorityTable {
		if v < 0 || v > 3 {
			return errors.New("Parameter authority type  error")
		}
		if mortgageTable[k].ToInt().Cmp(big.NewInt(0)) < 0 {
			return errors.New("Parameter mortgage amount is less than zero")
		}
	}
	return nil
}

func CheckSynchronizeShareKeyParameter( s types.SpecialTxInput) error {
	if len(s.SynchronizeShareKey.ShareKeyId) != 64 {
		return errors.New("Parameter ShareKeyId  error")
	}
	if len(s.SynchronizeShareKey.ShareKey) > 0 {
		return errors.New("Parameter ShareKey  error")
	}
	if s.SynchronizeShareKey.Shareprice.ToInt().Cmp(big.NewInt(0)) < 0 {
		return errors.New("Parameter Shareprice  is less than zero")
	}
	return nil
}

func CheckUnlockSharedKeyParameter( s types.SpecialTxInput) error {
	if len(s.SynchronizeShareKey.ShareKeyId) != 64 {
		return errors.New("Parameter ShareKeyId  error")
	}
	return nil
}