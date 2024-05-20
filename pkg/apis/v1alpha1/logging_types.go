package v1alpha1

type ContainerLoggingSpec struct {
	// +kubebuilder:validation:Optional
	SparkHistory *LoggingConfigSpec `json:"sparkHistory,omitempty"`
}

type LoggingConfigSpec struct {
	// +kubebuilder:validation:Optional
	Loggers map[string]*LogLevelSpec `json:"loggers,omitempty"`

	// +kubebuilder:validation:Optional
	Console *LogLevelSpec `json:"console,omitempty"`

	// +kubebuilder:validation:Optional
	File *LogLevelSpec `json:"file,omitempty"`
}

// LogLevelSpec
// level mapping example
// |---------------------|-----------------|
// |  App log level      |  kds log level  |
// |---------------------|-----------------|
// |  CRITICAL           |  FATAL          |
// |  ERROR              |  ERROR          |
// |  WARNING            |  WARN           |
// |  INFO               |  INFO           |
// |  DEBUG              |  DEBUG          |
// |  DEBUG              |  TRACE          |
// |---------------------|-----------------|
type LogLevelSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:="INFO"
	// +kubebuilder:validation:Enum=FATAL;ERROR;WARN;INFO;DEBUG;TRACE
	Level string `json:"level,omitempty"`
}
