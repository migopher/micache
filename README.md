# micache
go mini file cache 
## get micahce
````sh
$ go get github.com/xuzili1994/micache
````
#### Set Cache File Path

````sh
package main

import "github.com/xuzili1994/micache"

func main() {
	// set cache file path cache/
	micache.Dir="cache/"
}

````

#### Set Cache
set cache  and time
````sh
package main

import "github.com/xuzili1994/micache"

func main() {
	// set key expiration time 3600m
	micache.Set("go", "golang", 3600)
}

````

#### Get Cache
get key cache value
````sh
package main

import (
	"fmt"
	"github.com/xuzili1994/micache"
)

func main() {
	// get key cache value
	v := micache.Get("go")
	fmt.Println(v)
}

````

#### Set Struct
set key cache struct
````sh
package main

import (
	"github.com/xuzili1994/micache"
)

type User struct {
	Uid      int
	UserName string
}

func main() {
	setUser := User{9, "micache"}
	micache.SetEncoding("go", setUser, 3600)
}

````

#### Get Struct
get key cache struct value
````sh
package main

import (
	"fmt"
	"github.com/xuzili1994/micache"
)

type User struct {
	Uid      int
	UserName string
}

func main() {
	getUser:=User{}
	micache.GetDecoding("go", &getUser)
	fmt.Println(getUser)
}

````

#### Key Is Exist 

````sh
package main

import (
	"fmt"
	"github.com/xuzili1994/micache"
)

func main() {
	b:=micache.IsExist("go")
	fmt.Println(b)
}

````

#### Delete Key
delete cache key 
````sh
package main

import (
	"fmt"
	"github.com/xuzili1994/micache"
)

func main() {
	b:=micache.Delete("go")
	fmt.Println(b)
}

````



