package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["beego-api-demo/controllers:HeroController"] = append(beego.GlobalControllerRouter["beego-api-demo/controllers:HeroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-api-demo/controllers:HeroController"] = append(beego.GlobalControllerRouter["beego-api-demo/controllers:HeroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-api-demo/controllers:HeroController"] = append(beego.GlobalControllerRouter["beego-api-demo/controllers:HeroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-api-demo/controllers:HeroController"] = append(beego.GlobalControllerRouter["beego-api-demo/controllers:HeroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-api-demo/controllers:HeroController"] = append(beego.GlobalControllerRouter["beego-api-demo/controllers:HeroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-api-demo/controllers:JwtController"] = append(beego.GlobalControllerRouter["beego-api-demo/controllers:JwtController"],
        beego.ControllerComments{
            Method: "GetJwtInfo",
            Router: `/get`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-api-demo/controllers:JwtController"] = append(beego.GlobalControllerRouter["beego-api-demo/controllers:JwtController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
