package cmd

type Options map[string]string

func (o Options) Value(key string) string {
	return o[key]
}

func (o Options) Key(value string) string {
	for k, v := range o {
		if v == value {
			return k
		}
	}
	return ""
}

func (o Options) Values() []string {
	var values []string
	for _, v := range o {
		values = append(values, v)
	}
	return values
}
