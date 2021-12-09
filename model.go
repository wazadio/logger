package main

var (
	mongoConfig = MongoConfig{}
	kafkaConfig = KafkaConfig{}
)

type MongoConfig struct {
	MongoServer                 string
	Database                    string
	SystemLogCollection         string
	SettlementLogCollection     string
	CreditTransferLogCollection string
}

type KafkaConfig struct {
	KafkaAddress           string
	SystemLogTopic         string
	SettlementLogTopic     string
	CreditTransferLogTopic string
}

type SystemLoggerStruct struct {
	RequestId         string `json:"requestId" bson:"requestId"`
	TimeStamp         string `json:"timeStamp" bson:"timeStamp"`
	Content           string `json:"content" bson:"content"`
	TransactionType   string `json:"transactionType" bson:"transactionType"`
	MessageType       string `json:"messageType" bson:"messageType"`
	TransactionStatus string `json:"transactionStatus" bson:"transactionStatus"`
	ReasonCode        string `json:"reasonCode" bson:"reasonCode"`
	LatestPosition    string `json:"latestPosition" bson:"latestPosition"`
}

type SettlementLogStruct struct {
	RequestId                string `json:"requestId" bson:"requestId"`
	Direction                string `json:"direction" bson:"direction"`
	SettlementType           string `json:"settlementType" bson:"settlementType"`
	FromId                   string `json:"fromId" bson:"fromId"`
	ToId                     string `json:"toId" bson:"toId"`
	BusinessMessageId        string `json:"businessMessageId" bson:"businessMessageId"`
	MessageDefinitionId      string `json:"messageDefinitionId" bson:"messageDefinitionId"`
	BusinessServiceId        string `json:"businessServiceId,omitempty" bson:"businessServiceId,omitempty"`
	HeaderDateTime           string `json:"headerDateTime" bson:"headerDateTime"`
	CreationDate             string `json:"creationDate" bson:"creationDate"`
	CopyDuplicate            string `json:"copyDuplicate,omitempty" bson:"copyDuplicate,omitempty"`
	PossibleDuplicate        string `json:"possibleDuplicate,omitempty" bson:"possibleDuplicate,omitempty"`
	MessageId                string `json:"messageId" bson:"messageId"`
	DocumentDateTime         string `json:"documentDateTime" bson:"documentDateTime"`
	OriginalMessageId        string `json:"originalMessageId,omitempty" bson:"originalMessageId,omitempty"`
	OriginalMessageNameId    string `json:"originalMessageNameId,omitempty" bson:"originalMessageNameId,omitempty"`
	OriginalCreationDateTime string `json:"originalCreationDateTime,omitempty" bson:"originalCreationDateTime,omitempty"`
	OriginalEndToEndId       string `json:"originalEndToEndId" bson:"originalEndToEndId"`
	OriginalTransactionId    string `json:"originalTransactionId" bson:"originalTransactionId"`
	TransactionStatus        string `json:"transactionStatus" bson:"transactionStatus"`
	ReasonCode               string `json:"reasonCode" bson:"reasonCode"`
	AdditionalInfo           string `json:"additionalInfo" bson:"additionalInfo"`
	InterBankSettlementDate  string ` json:"interBankSettlementDate,omitempty" bson:"interBankSettlementDate,omitempty"`
	DebtorName               string `json:"debtorName,omitempty" bson:"debtorName,omitempty"`
	DebtorAccountId          string `json:"debtorAccountId,omitempty" bson:"debtorAccountId,omitempty"`
	DebtorAccountType        string `json:"debtorAccountType,omitempty" bson:"debtorAccountType,omitempty"`
	OriginalDebtorFiId       string `json:"originalDebtorFiId,omitempty" bson:"originalDebtorFiId,omitempty"`
	CreditorFiId             string `json:"creditorFiId,omitempty" bson:"creditorFiId,omitempty"`
	CreditorName             string `json:"creditorName" bson:"creditorName"`
	CreditorAccountId        string `json:"creditorAccountId" bson:"creditorAccountId"`
	CreditorAccountType      string `json:"creditorAccountType" bson:"creditorAccountType"`
	DebtorType               string `json:"debtorType,omitempty" bson:"debtorType,omitempty"`
	DebtorNationalId         string `json:"debtorNationalId,omitempty" bson:"debtorNationalId,omitempty"`
	DebtorResidentStatus     string `json:"debtorResidentStatus,omitempty" bson:"debtorResidentStatus,omitempty"`
	DebtorTownName           string `json:"debtorTownName,omitempty" bson:"debtorTownName,omitempty"`
	CreditorType             string `json:"creditorType" bson:"creditorType"`
	CreditorResidentStatus   string `json:"creditorResidentStatus" bson:"creditorResidentStatus"`
	CreditorTownName         string `json:"creditorTownName" bson:"creditorTownName"`
}

type CreditTransferLogStruct struct {
	RequestId               string `json:"requestId" bson:"requestId"`
	ChannelSource           string `json:"channelSource" bson:"channelSource"`
	Direction               string `json:"direction" bson:"direction"`
	ReversalMessage         string `json:"reversalMessage" bson:"reversalMessage"`
	FromId                  string `json:"fromId" bson:"fromId"`
	ToId                    string `json:"toId" bson:"toId"`
	BusinessMessageId       string `json:"businessMessageId" bson:"businessMessageId"`
	MessageDefinitionId     string `json:"messageDefinitionId" bson:"messageDefinitionId"`
	BusinessServiceId       string `json:"businessServiceId" bson:"businessServiceId"`
	HeaderDateTime          string `json:"headerDateTime" bson:"headerDateTime"`
	CreationDate            string `json:"creationDate" bson:"creationDate"`
	CopyDuplicate           string `json:"copyDuplicate" bson:"copyDuplicate"`
	PossibleDuplicate       string `json:"possibleDuplicate" bson:"possibleDuplicate"`
	MessageId               string `json:"messageId" bson:"messageId"`
	DocumentDateTime        string `json:"documentDateTime" bson:"documentDateTime"`
	OriginalMessageId       string `json:"originalMessageId" bson:"originalMessageId"`
	OriginalMessageNameId   string `json:"originalMessageNameId" bson:"originalMessageNameId"`
	OriginalEndToEndId      string `json:"originalEndToEndId" bson:"originalEndToEndId"`
	OriginalTransactionId   string `json:"originalTransactionId" bson:"originalTransactionId"`
	TransactionStatus       string `json:"transactionStatus" bson:"transactionStatus"`
	ReasonCode              string `json:"reasonCode" bson:"reasonCode"`
	AdditionalInfo          string `json:"additionalInfo" bson:"additionalInfo"`
	InterBankSettlementDate string `json:"interBankSettlementDate" bson:"interBankSettlementDate"`
	Amount                  string `json:"amount" bson:"amount"`
	DebitorName             string `json:"debitorName" bson:"debitorName"`
	DebitorAccountId        string `json:"debitorAccountId" bson:"debitorAccountId"`
	CreditorName            string `json:"creditorName" bson:"creditorName"`
	CreditorAccountId       string `json:"creditorAccountId" bson:"creditorAccountId"`
	CreditorAccountType     string `json:"creditorAccountType" bson:"creditorAccountType"`
	CreditorType            string `json:"creditorType" bson:"creditorType"`
	CreditorNationalId      string `json:"creditorNationalId" bson:"creditorNationalId"`
	CreditorResidentStatus  string `json:"creditorResidentStatus" bson:"creditorResidentStatus"`
	CreditorTownName        string `json:"creditorTownName" bson:"creditorTownName"`
}
