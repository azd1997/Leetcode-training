package ltmap

type IMap interface {
	Put(k, v interface{})	// k必须可比较，目前只考虑整数和字符串
	Get(k interface{}) (interface{}, error)
	Remove(k interface{})
	Len() int
}
