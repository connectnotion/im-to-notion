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

func (l *LogOptions) Validate() bool {
	if l.Output == "" || l.Level == "" {
		return false
	}

	return true
}

// Dingtalk config
type Dingtalk struct {
	AppKey    string `json:"appKey" yaml:"appKey"`
	AppSecret string `json:"appSecret" yaml:"appSecret"`
}

func (d *Dingtalk) Validate() bool {
	if d.AppKey == "" || d.AppSecret == "" {
		return false
	}
	return true
}

// Source from im
type Source struct {
	Dingtalk *Dingtalk `json:"dingtalk,omitempty" yaml:"dingtalk,omitempty"`
}

func (s *Source) Validate() bool {
	if s.Dingtalk != nil && !s.Dingtalk.Validate() {
		return false
	}

	return true
}

// Notion config
type Notion struct {
	Secret     string `json:"secret" yaml:"secret"`
	DatabaseId string `json:"databaseId" yaml:"databaseId"`
}

func (n *Notion) Validate() bool {
	if n.Secret == "" || n.DatabaseId == "" {
		return false
	}

	return true
}

// Destination send to
type Destination struct {
	Notion Notion `json:"notion" yaml:"notion"`
}

func (d *Destination) Validate() bool {
	return d.Notion.Validate()
}

// Config all configuration
type Config struct {
	Log         LogOptions  `json:"log" yaml:"log"`
	Source      Source      `json:"source" yaml:"source"`
	Destination Destination `json:"destination" yaml:"destination"`
}

func (c *Config) Validate() bool {
	if !c.Log.Validate() || !c.Source.Validate() || !c.Destination.Validate() {
		return false
	}

	return true
}
