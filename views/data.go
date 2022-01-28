package views

const (
	AlertLvError   = "danger"
	AlertLvWarning = "warning"
	AlertLvInfo    = "info"
	AlertLvSuccess = "success"
)

// Alert is used to render Bootstrap Alert messages in templates
type Alert struct {
	Level   string
	Message string
}

// Data is the top level structure that views expect data
// to come in.
type Data struct {
	Alert *Alert // can be nil
	Yield interface{}
}
