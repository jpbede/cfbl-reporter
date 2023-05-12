package cfbl

func NewReport(originalMail *[]byte) *Report {
	report := &Report{
		originalMail:   originalMail,
		sendFullReport: false,
	}
	return report
}
