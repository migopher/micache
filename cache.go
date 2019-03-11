package go_cache

type IAnimal interface {
	Set()
	Get()
	IsExists()
}

type File struct {
	Type string
	Expire string
	Prefix string
	Path string
}
type Redis struct {

}

//func Connect() interface{} {
//
//}

//func NewCache() *Cache {
//	//return
//}

func (c *File)Set(key string, value string,time int64)  {
	
}

func (c *File )IsExists() bool {

}