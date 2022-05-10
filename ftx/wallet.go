package ftx

import (
	"fmt"
	"net/http"
)

const pathBalance = "%s/wallet/balances"
const pathWithdrawal = "%s/wallet/withdrawals"
const pathDeposit = "%s/wallet/deposits"
const pathCoins = "%s/wallet/coins"

type IWalletService interface {
	GetBalance() ([]*BalanceResponse, error)
	GetWithdrawalHistory(opts *GetWithdrawalHistoryOptions) ([]*WithdrawalResponse, error)
	RequestWithdrawal(request *WithdrawalRequest) (*WithdrawalResponse, error)
	GetDepositHistory(opts *GetDepositHistoryOptions) ([]*DepositResponse, error)
	GetCoins() ([]*CoinResponse, error)
}

type WalletService service

type BalanceResponse struct {
	Coin                   string
	Free                   float64
	SpotBorrow             float64
	Total                  float64
	UsdValue               float64
	AvailableWithoutBorrow float64
}

type WithdrawalResponse struct {
	Coin    string
	Address string
	Tag     string
	Fee     float64
	Id      float64
	Size    float64
	Status  string
	Time    string
	Method  string
	Txid    string
}

type WithdrawalRequest struct {
	Coin     string `url:"coin" json:"coin"`
	Size     string `url:"size" json:"size"`
	Address  string `url:"address" json:"address"`
	Tag      string `url:"tag" json:"tag"`
	Method   string `url:"method" json:"method"`
	Password string `url:"password" json:"password"`
	Code     string `url:"code" json:"code"`
}

type DepositResponse struct {
	Coin          string
	Confirmations float64
	ConfirmedTime string
	Fee           float64
	Id            float64
	SentTime      string
	Size          float64
	Status        string
	Time          string
	Txid          string
	Notes         string
}

type CoinResponse struct {
	CanDeposit       bool
	CanWithdraw      bool
	HasTag           bool
	Id               string
	Name             string
	Bep2Asset        string
	CanConvert       bool
	Collateral       bool
	CollateralWeight float64
	CreditTo         string
	Erc20Contract    string
	Fiat             bool
	IsToken          bool
	Methods          []string
	SplMint          string
	Trc20Contract    string
	UsdFungible      bool
}

type GetDepositHistoryOptions struct {
	StartTime int64 `url:"start_time"`
	EndTime   int64 `url:"end_time"`
}

type GetWithdrawalHistoryOptions struct {
	StartTime int64 `url:"start_time"`
	EndTime   int64 `url:"end_time"`
}

func (s *WalletService) GetBalance() ([]*BalanceResponse, error) {
	u := fmt.Sprintf(pathBalance, s.client.baseURL)

	var out []*BalanceResponse

	err := s.client.DoPrivate(u, http.MethodGet, nil, &out)

	return out, err

}

func (s *WalletService) GetWithdrawalHistory(opts *GetWithdrawalHistoryOptions) ([]*WithdrawalResponse, error) {
	u := fmt.Sprintf(pathWithdrawal, s.client.baseURL)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}

	var out []*WithdrawalResponse

	err = s.client.DoPrivate(u, http.MethodGet, nil, &out)

	return out, err

}

func (s *WalletService) RequestWithdrawal(request *WithdrawalRequest) (*WithdrawalResponse, error) {
	u := fmt.Sprintf(pathWithdrawal, s.client.baseURL)

	var out *WithdrawalResponse

	err := s.client.DoPrivate(u, http.MethodPost, &request, &out)

	return out, err

}

func (s *WalletService) GetDepositHistory(opts *GetDepositHistoryOptions) ([]*DepositResponse, error) {
	u := fmt.Sprintf(pathDeposit, s.client.baseURL)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}

	var out []*DepositResponse

	err = s.client.DoPrivate(u, http.MethodGet, nil, &out)

	return out, err

}

func (s *WalletService) GetCoins() ([]*CoinResponse, error) {
	u := fmt.Sprintf(pathCoins, s.client.baseURL)

	var out []*CoinResponse

	err := s.client.DoPrivate(u, http.MethodGet, nil, &out)

	return out, err
}
