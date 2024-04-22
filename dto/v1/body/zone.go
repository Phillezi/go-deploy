package body

type ZoneRead struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Capabilities []string `json:"capabilities"`
	Interface    *string  `json:"interface"`
	Legacy       bool     `json:"legacy"`

	// Type
	// Deprecated: use capabilities instead
	Type string `json:"type"`
}
