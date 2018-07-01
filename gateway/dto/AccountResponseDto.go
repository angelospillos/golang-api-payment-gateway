package dto

type AccountResponseDto struct {
	AccountId   string `json:"account_id"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

func CreateAccountResponseDto(accountId string, status string, description string) AccountResponseDto {

	var accountResponseDto AccountResponseDto

	accountResponseDto.AccountId = accountId
	accountResponseDto.Status = status
	accountResponseDto.Description = description

	return accountResponseDto
}
