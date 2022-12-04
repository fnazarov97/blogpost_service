package storage

import (
	"blockpost/genprotos/article"
	"blockpost/genprotos/author"
)

type StorageI interface {
	AddAuthor(*author.CreateAuthorReq) (*author.CreateAuthorRes, error)
	GetAuthorByID(req *author.Id) (*author.GetAuthorByIdRes, error)
	GetArticlesByAuthorID(req *author.Id) (resp *author.GetArticles, err error)
	GetAuthorList(req *author.GetAuthorListReq) (resp *author.GetAuthors, err error)
	UpdateAuthor(req *author.UpdateAuthorReq) error
	DeleteAuthor(id *author.Id) error

	AddArticle(id string, req *article.AddArticleReq) error
	GetArticleByID(id string) (*article.GetArticleByIdRes, error)
	GetArticleList(offset, limit int, search string) (*article.GetArticleListRes, error)
	UpdateArticle(entity *article.UpdateArticleReq) error
	DeleteArticle(id string) error
}
