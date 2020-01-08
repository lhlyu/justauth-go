package utils

func GetRealState(state string) string {
	if state == "" {
		return GetUUID()
	}
	return state
}
