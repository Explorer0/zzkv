package zzkv

type Zzkv struct {
	Storager
	Compression
	Clear
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

func NewDefault() *Zzkv {
	z := &Zzkv{
		Storager:Storager{
			pstStorager:NewDefaultPstStorager(),
			cacheStorager:NewDefaultCacheStorager(),
			storageMap:make(map[string]bool),
		},
		Compression:NewDefaultCompression(),
	}

	z.Run(&z.Storager)
	return z

}

func (z *Zzkv) Set(key string, val interface{}, sync bool) error {
	// 序列化对象
	data, err := Serialize(val)
	if err != nil {
		return err
	}
	// 压缩数据
	data = z.Compress(data)

	//存储数据
	setErr := z.Storager.Set(key, data, sync)
	if setErr != nil {
		return setErr
	}

	return nil
}

func (z *Zzkv) Get(key string, val interface{}) error {
	// 获取数据
	data := z.Storager.Get(key)

	// 解压数据
	data  = z.Decompress(data)

	// 反序列对象
	err := Deserialize(data, val)

	if err != nil {
		return err
	}

	return nil
}






