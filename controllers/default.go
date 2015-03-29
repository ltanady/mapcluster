package controllers

import (
	"github.com/astaxie/beego"
  "github.com/astaxie/beego/httplib"
  "mapcluster/models"
  "fmt"
  "encoding/json"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "Hello How are you?"
  c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}

func (c *MainController) Maps() {
  var response models.Response
  
  reqUrl := "http://gmaps-utility-library-dev.googlecode.com/svn/tags/markerclusterer/1.0/examples/data.json"
  req := httplib.Get(reqUrl)
  result, err := req.Bytes()
  if err != nil {
		fmt.Println(err)
	} else {
    fmt.Println(result)
  }
  if err := json.Unmarshal(result, &response); err != nil {
    fmt.Println(err)
  }
  fmt.Println("Response")
  fmt.Println(response)
  c.TplNames = "maps.tpl" 
}

func (c *MainController) Hello() {
  c.Data["Hello"] = "Hello World"
  c.TplNames = "hello.tpl"
}
