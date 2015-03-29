package controllers

import (
	"github.com/astaxie/beego"
//   "github.com/astaxie/beego/httplib"
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
  var response []models.Coordinate
  var coordinates =  "[{ \"latitude\" : 37.4419, \"longitude\" : -122.1419},{ \"latitude\" : 7.4419, \"longitude\" : 122.1419},{ \"latitude\" : 37.4219, \"longitude\" : -122.1219},{ \"latitude\" : 37.119, \"longitude\" : -122.1419},{ \"latitude\" : 37.4019, \"longitude\" : -122.1419},{ \"latitude\" : 37.4490, \"longitude\" : -122.1419}]";
  if err := json.Unmarshal([]byte(coordinates), &response); err != nil {
    fmt.Println(err)
  }
  fmt.Println("Response")
  fmt.Println(response)
  c.Data["Coordinate"] = response
  c.TplNames = "maps.tpl" 
}

func (c *MainController) Hello() {
  c.Data["Hello"] = "Hello World"
  c.TplNames = "hello.tpl"
}
