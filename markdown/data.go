package markdown

import "time"

type SubmitWindowedPoStType struct {
	SectorBatchStatInfoType
}

type ProveCommitSectorType struct {
	SectorBatchStatInfoType
}

type PreCommitSectorType struct {
	SectorBatchStatInfoType
}

type SectorBatchStatInfoType struct {
	SuccessNum     uint64 //PreCommitSector扇区成功
	FailureNum     uint64 //PreCommitSector扇区失败
	TotalMinerTips uint64 //总矿工手续费及
	TotalGasFee    uint64 //总gas费用
}

type BillInfoDo struct {
	Id              int64     `json:"id"`
	Miner           string    `json:"miner"`
	Type            string    `json:"type"`
	BlockReward     float64   `json:"block_reward"`
	BlockNumber     int64     `json:"block_number"`
	DailyRelease    float64   `json:"daily_release"`
	VestingFunds    float64   `json:"vesting_funds"`
	IndirectRelease float64   `json:"indirect_release"`
	BusRelease      float64   `json:"bus_release"`
	ActualRelease   float64   `json:"actual_release"`
	StatisticsDate  string    `json:"statistics_date"`
	CreateTime      time.Time `json:"create_time"`
}

type CostInfoDo struct {
	Id                             int64     `json:"id"`
	Miner                          string    `json:"miner"`
	WindowedPostMinerFee           float64   `json:"windowed_post_miner_fee"`
	WindowedPostBurnFee            float64   `json:"windowed_post_burn_fee"`
	WindowedPostTotal              int64     `json:"windowed_post_total"`
	WindowedPostSuccess            int64     `json:"windowed_post_success"`
	WindowedPostFail               int64     `json:"windowed_post_fail"`
	ProvecommitMinerFee            float64   `json:"provecommit_miner_fee"`
	ProvecommitBurnFee             float64   `json:"provecommit_burn_fee"`
	ProvecommitTotal               int64     `json:"provecommit_total"`
	ProvecommitSuccess             int64     `json:"provecommit_success"`
	ProvecommitFail                int64     `json:"provecommit_fail"`
	PrecommitMinerFee              float64   `json:"precommit_miner_fee"`
	PrecommitBurnFee               float64   `json:"precommit_burn_fee"`
	PrecommitTotal                 int64     `json:"precommit_total"`
	PrecommitSuccess               int64     `json:"precommit_success"`
	PrecommitFail                  int64     `json:"precommit_fail"`
	DeclareFaultsMinerFee          float64   `json:"declare_faults_miner_fee"`
	DeclareFaultsBurnFee           float64   `json:"declare_faults_burn_fee"`
	DeclareFaultsRecoveredMinerFee float64   `json:"declare_faults_recovered_miner_fee"`
	DeclareFaultsRecoveredBurnFee  float64   `json:"declare_faults_recovered_burn_fee"`
	WithdrawBalanceMinerMinerFee   float64   `json:"withdraw_balance_miner_miner_fee"`
	WithdrawBalanceMinerBurnFee    float64   `json:"withdraw_balance_miner_burn_fee"`
	WithdrawBalanceMinerTransfer   float64   `json:"withdraw_balance_miner_transfer"`
	TerminateMinerFee              float64   `json:"terminate_miner_fee"`
	TerminateBurnFee               float64   `json:"terminate_burn_fee"`
	SendMinerFee                   float64   `json:"send_miner_fee"`
	SendBurnFee                    float64   `json:"send_burn_fee"`
	SendTransfer                   float64   `json:"send_transfer"`
	BurnFee                        float64   `json:"burn_fee"`
	PenaltyFee                     float64   `json:"penalty_fee"`
	RewardFee                      float64   `json:"reward_fee"`
	OtherFee                       float64   `json:"other_fee"`
	PledgeFee                      float64   `json:"pledge_fee"`
	StatisticsDate                 string    `json:"statistics_date"`
	CreateTime                     time.Time `json:"create_time"`
}

//python转 golang

var Costinfo = CostInfoDo{
	Id:                             11,
	Miner:                          "f02301",
	WindowedPostMinerFee:           0.12,
	WindowedPostBurnFee:            0.54,
	WindowedPostTotal:              12,
	WindowedPostSuccess:            11,
	WindowedPostFail:               12,
	ProvecommitMinerFee:            0.45,
	ProvecommitBurnFee:             1,
	ProvecommitTotal:               10,
	ProvecommitSuccess:             10,
	ProvecommitFail:                0,
	PrecommitMinerFee:              1,
	PrecommitBurnFee:               1,
	PrecommitTotal:                 10,
	PrecommitSuccess:               10,
	PrecommitFail:                  10,
	DeclareFaultsMinerFee:          1,
	DeclareFaultsBurnFee:           1,
	DeclareFaultsRecoveredMinerFee: 1,
	DeclareFaultsRecoveredBurnFee:  1,
	WithdrawBalanceMinerMinerFee:   1,
	WithdrawBalanceMinerBurnFee:    1,
	WithdrawBalanceMinerTransfer:   1,
	TerminateMinerFee:              1,
	TerminateBurnFee:               1,
	SendMinerFee:                   1,
	SendBurnFee:                    1,
	SendTransfer:                   1,
	BurnFee:                        1,
	PenaltyFee:                     1,
	RewardFee:                      1,
	OtherFee:                       1,
	PledgeFee:                      1,
	StatisticsDate:                 "",
	CreateTime:                     time.Now(),
}

var TestBillInfo = BillInfoDo{
	Id:              1,
	Miner:           "f02301",
	Type:            "",
	BlockReward:     1,
	BlockNumber:     1,
	DailyRelease:    1,
	VestingFunds:    1,
	IndirectRelease: 1,
	BusRelease:      1,
	ActualRelease:   1,
	StatisticsDate:  "",
	CreateTime:      time.Now(),
}

var TemplateData = `{
	"测试代码": "这是什么",
	"测试代码": "这是什么",
	"测试代码": "这是什么",
	"测试代码": "这是什么",
}`
