package entities

type Memo struct {
	Total     string
	Free      string
	Used      string
	Available string
}

type Host struct {
	Host string
	OS   string
	Arch string
	Memo Memo
}
