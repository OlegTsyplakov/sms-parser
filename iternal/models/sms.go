package models

type sms struct {
	from      string
	from_toa  string
	from_smsc string
	sent      string
	received  string
	subject   string
	modem     string
	imsi      int
	report    string
	alphabet  string
	length    int
	content   string
}
