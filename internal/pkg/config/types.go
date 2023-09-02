package config

// LogOptions contains several options related to logging.
type LogOptions struct {
	// Level is the minimum logging level that a logging message should have
	// to output itself.
	Level string `json:"level" yaml:"level"`
	// Output defines the destination file path to output logging messages.
	// Two keywords "stderr" and "stdout" can be specified so that message will
	// be written to stderr or stdout.
	Output string `json:"output" yaml:"output"`
}

// Dingtalk config
type Dingtalk struct {
	AppKey    string `json:"appKey" yaml:"appKey"`
	AppSecret string `json:"appSecret" yaml:"appSecret"`
}

// Source from im
type Source struct {
	Dingtalk Dingtalk `json:"dingtalk" yaml:"dingtalk"`
}

// Notion config
type Notion struct {
	Secret string `json:"secret" yaml:"secret"`
}

// Destination send to
type Destination struct {
	Notion Notion `json:"notion" yaml:"notion"`
}

// Config all configuration
type Config struct {
	Log         LogOptions  `json:"log" yaml:"log"`
	Source      Source      `json:"source" yaml:"source"`
	Destination Destination `json:"destination" yaml:"destination"`
}
