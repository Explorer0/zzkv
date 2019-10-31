package zzkv

type Zzkv struct {
	Storager
	Compression
}

func New(s Storager, c Compression) Zzkv {
	result := Zzkv{
		Storager:s,
		Compression: c,
	}

	if c == nil {
		result.Compression = NewDefaultCompression()
	}
	return result
}

func NewDefault() Zzkv {
	return Zzkv{
		Storager:Storager{
			outStorage:NewDefaultPstStorager(),
			insideStorage:NewDefaultCacheStorager(),
			storageMap:make(map[string]bool),
		},
		Compression:NewDefaultCompression(),
	}
}




