package models

import (
	"github.com/meifamily/ptt-alertor/models/article"
	"github.com/meifamily/ptt-alertor/models/board"
	"github.com/meifamily/ptt-alertor/models/user"
)

var User = user.NewUser(new(user.Redis))
var Article = func() *article.Article {
	return article.NewArticle(new(article.DynamoDB))
}
var Board = func() *board.Board {
	return board.NewBoard(new(board.DynamoDB), new(board.Redis))
}
