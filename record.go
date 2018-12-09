package csvkit

type Record map[string]string

func (r Record) Keys() []string {
	keys := make([]string, 0, len(r))
	for k := range r {
		keys = append(keys, k)
	}
	return keys
}

func (r Record) Values() []string {
	values := make([]string, 0, len(r))
	for _, v := range r {
		values = append(values, v)
	}
	return values
}

func (r Record) Get(key string) string {
	if r == nil {
		return ""
	}
	if v, ok := r[key]; ok {
		return v
	}
	return ""
}
