package util


func AnyMatchSlices(a, b []string) bool {
		set := make(map[string]struct{}, len(a))
		for _, v := range a {
				set[v] = struct{}{}
		}
		for _, v := range b {
					if _, ok := set[v]; ok {
						return true
				}
		}
		return false
}