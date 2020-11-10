package configer

import "strings"

func GetConfigPath(key string, path ...string) string {
	return key + "." + strings.Join(path, ".")
}
