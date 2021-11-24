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
	RequestId         string `json:"requestId"`
	TimeStamp         string `json:"timeStamp"`
	Content           string `json:"content"`
	TransactionType   string `json:"transactionType"`
	MessageType       string `json:"messageType"`
	TransactionStatus string `json:"transactionStatus"`
	ReasonCode        string `json:"reasonCode"`
	LatestPosition    string `json:"latestPosition"`
}

type SettlementLogStruct struct {
	RequestId              string `json:"requestId"`
	Direction              string `json:"direction"`
	SettlementType         string `json:"settlementType"`
	FromId                 string `json:"fromId"`
	ToId                   string `json:"toId"`
	BusinessMessageId      string `json:"businessMessageId"`
	MessageDefinitionId    string `json:"messageDefinitionId"`
	CreationDate           string `json:"creationDate"`
	MessageId              string `json:"messageId"`
	CreationDateTime       string `json:"creationDateTime"`
	OriginalEndToEndId     string `json:"originalEndToEndId"`
	OriginalTransactionId  string `json:"originalTransactionId"`
	TransactionStatus      string `json:"transactionStatus"`
	ReasonCode             string `json:"reasonCode"`
	AdditionalInfo         string `json:"additionalInfo"`
	CreditorName           string `json:"creditorName"`
	CreditorAccountId      string `json:"creditorAccountId"`
	CreditorAccountType    string `json:"creditorAccountType"`
	CreditorType           string `json:"creditorType"`
	CreditorResidentStatus string `json:"creditorResidentStatus"`
	CreditorTownName       string `json:"creditorTownName"`
}

type CreditTransferLogStruct struct {
	RequestId               string `json:"requestId"`
	Direction               string `json:"direction"`
	ReversalMessage         string `json:"reversalMessage"`
	FromId                  string `json:"fromId"`
	ToId                    string `json:"toId"`
	BusinessMessageId       string `json:"businessMessageId"`
	MessageDefinitionId     string `json:"messageDefinitionId"`
	BusinessServiceId       string `json:"businessServiceId"`
	CreationDate            string `json:"creationDate"`
	CopyDuplicate           string `json:"copyDuplicate"`
	PossibleDuplicate       string `json:"possibleDuplicate"`
	MessageId               string `json:"messageId"`
	CreationDateTime        string `json:"creationDateTime"`
	OriginalMessageId       string `json:"originalMessageId"`
	OriginalMessageNameId   string `json:"originalMessageNameId"`
	OriginalEndToEndId      string `json:"originalEndToEndId"`
	OriginalTransactionId   string `json:"originalTransactionId"`
	TransactionStatus       string `json:"transactionStatus"`
	ReasonCode              string `json:"reasonCode"`
	AdditionalInfo          string `json:"additionalInfo"`
	InterBankSettlementDate string `json:"interBankSettlementDate"`
	CreditorName            string `json:"creditorName"`
	CreditorType            string `json:"creditorType"`
	CreditorNationalId      string `json:"creditorNationalId"`
	CreditorResidentStatus  string `json:"creditorResidentStatus"`
	CreditorTownName        string `json:"creditorTownName"`
}
