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
			pstStorager:NewDefaultPstStorager(),
			cacheStorager:NewDefaultCacheStorager(),
			storageMap:make(map[string]bool),
		},
		Compression:NewDefaultCompression(),
	}
}

func (z *Zzkv) Set(key string, val interface{}, sync bool) error {
	data, err := Serialize(val)
	if err != nil {
		return err
	}
	data = z.Compress(data)

	setErr := z.Storager.Set(key, data, sync)
	if setErr != nil {
		return setErr
	}

	return nil
}

func (z *Zzkv) Get(key string, val interface{}) error {
	data := z.Storager.Get(key)

	data  = z.Decompress(data)

	err := Deserialize(data, val)

	if err != nil {
		return err
	}

	return nil
}






