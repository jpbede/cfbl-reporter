package cfbl

func NewReport(originalMail *[]byte) *Report {
	report := &Report{
		originalMail: originalMail,
	}
	return report
}
