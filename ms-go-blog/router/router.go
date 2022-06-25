package router

import (
	"ms-go-blog/context"
	"ms-go-blog/views"
	"net/http"
)

func Router()  {
	//1. 页面  views 2. api 数据（json） 3. 静态资源
	http.Handle("/",context.Context)
	//http://localhost:8080/c/1  1参数 分类的id
	context.Context.Handler("/c/{id}",views.HTML.CategoryNew)
	context.Context.Handler("/login",views.HTML.LoginNew)
	//http.HandleFunc("/",views.HTML.Index)
	////http://localhost:8080/c/1  1参数 分类的id
	//http.HandleFunc("/c/",views.HTML.Category)
	//http.HandleFunc("/login",views.HTML.Login)
	////http://localhost:8080/p/7.html
	//http.HandleFunc("/p/",views.HTML.Detail)
	//http.HandleFunc("/writing",views.HTML.Writing)
	//http.HandleFunc("/pigeonhole",views.HTML.Pigeonhole)
	//http.HandleFunc("/api/v1/post",api.API.SaveAndUpdatePost)
	//http.HandleFunc("/api/v1/post/",api.API.GetPost)
	//http.HandleFunc("/api/v1/post/search",api.API.SearchPost)
	//http.HandleFunc("/api/v1/qiniu/token",api.API.QiniuToken)
	//http.HandleFunc("/api/v1/login",api.API.Login)

	http.Handle("/resource/",http.StripPrefix("/resource/",http.FileServer(http.Dir("public/resource/"))))
}
