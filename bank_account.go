package smarthr

type BankAccount struct {
	BankCode          string `json:"bank_code"`
	BankBranchCode    string `json:"bank_branch_code"`
	AccountType       string `json:"account_type"`
	AccountNumber     string `json:"account_number"`
	AccountHolderName string `json:"account_holder_name"`
}
