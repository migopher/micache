package go_cache

type CachenIterface interface {
	Set(string,interface{},int64)bool
	Get() string
	//IsExists()
}


type File struct {
	Type string
	Expire string
	Prefix string
	Path string
}

func (f File)Get(name string) string {
	return  name
}

func (f File) Set(key string ,value interface{},time int64) bool{
	return true
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