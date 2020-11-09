package configer

type FileType int

const (
	YAML FileType = iota
	JSON
	TOML
	PROPERTIES
	INI
	HCL
)

func NewFileType(f string) FileType {
	switch f {
	case "yaml":
		return YAML
	case "json":
		return JSON
	case "toml":
		return TOML
	case "properties":
		return PROPERTIES
	case "ini":
		return INI
	case "hcml":
		return HCL
	default:
		return YAML
	}
}

func (t FileType) String() string {
	switch t {
	case YAML:
		return "yaml"
	case JSON:
		return "json"
	case TOML:
		return "toml"
	case PROPERTIES:
		return "properties"
	case INI:
		return "ini"
	case HCL:
		return "hcl"
	default:
		return "yaml"
	}
}
