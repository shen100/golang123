package route

import (
	"gopkg.in/kataras/iris.v6"
	"golang123/config"
	"golang123/controller/common"
	"golang123/controller/auth"
	"golang123/controller/category"
	"golang123/controller/article"
	"golang123/controller/collect"
	"golang123/controller/comment"
	"golang123/controller/user"
	"golang123/controller/message"
)

// Route 路由
func Route(app *iris.Framework) {
	apiPrefix   := config.APIConfig.Prefix

	router := app.Party(apiPrefix) 
	{	
		router.Post("/signin",                  user.Signin)
		router.Post("/signup",                  user.Signup)
		router.Get("/active/:id/:secret",       user.ActiveAccount)
		router.Post("/reset",                   user.ResetPasswordMail)
		router.Post("/reset/:id/:secret",       user.ResetPassword)
		router.Get("/reset/verify/:id/:secret", user.VerifyResetPasswordLink)

		router.Get("/user/info",             user.Info)
		router.Post("/user/update",          auth.SigninRequired,       
										     user.UpdateInfo)
		router.Post("/user/password/update", auth.SigninRequired,       
											 user.UpdatePassword)
		router.Get("/user/score/top10",      user.Top10)
		router.Get("/user/score/top100",     user.Top100)
		router.Post("/upload",               auth.SigninRequired,          
											 common.Upload)
		router.Get("/message/unread",        message.Unread)
		router.Get("/message/unread/count",  message.UnreadCount)

		router.Get("/categories",          category.List)

		router.Get("/articles",             article.List)
		router.Get("/articles/recent",      auth.SigninRequired,  
										    article.RecentList)
		router.Get("/articles/maxcomment",  article.ListMaxComment)
		router.Get("/articles/maxbrowse",   article.ListMaxBrowse)
		router.Get("/article/:id",          article.Info)
		router.Post("/article/create",      auth.SigninRequired, 
										    article.Create)
		router.Post("/article/update",      auth.SigninRequired,    
										    article.Update)
		router.Post("/collect/create",      auth.SigninRequired,
											collect.Collect)
		router.Post("/collect/delete",      auth.SigninRequired,
										    collect.DeleteCollect)
		router.Get("/collects",             auth.SigninRequired,
											collect.List)
		router.Post("/comment/create",      auth.SigninRequired,
										    comment.Create)
    }

	adminRouter := app.Party(apiPrefix + "/admin", auth.AdminRequired)
	{
		adminRouter.Get("/categories",               category.AllList)
		adminRouter.Post("/category/create",         category.Create)
		adminRouter.Post("/category/update",         category.Update)
		adminRouter.Post("/category/status/update",  category.UpdateStatus)

		adminRouter.Get("/articles",                 article.AllList)
		adminRouter.Post("/article/status/update",   article.UpdateStatus)
    }
}