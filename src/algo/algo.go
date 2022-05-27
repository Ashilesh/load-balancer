package algo

type Algo interface {
	Add(string)
	Delete(string)
	GetUrl(string) string
}

// TODO: pass algo type in params and make it thread safe
func GetAlgoFactory() Algo {
	return GetConsistetnHash()
}
