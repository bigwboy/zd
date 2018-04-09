package controllers

import (
	"miaopost/frontend/models"
)

type ArticleController struct {
	BaseController
}

func (this *ArticleController) View() {
	id, _ := this.GetInt("id")
	article, err := models.GetArticleById(int(id))
	if err != nil {
		this.Tips(err.Error())
	}

	this.Data["article"] = article
	this.Layout = "layout/main.tpl"
	this.TplName = "article/view.tpl"
}
