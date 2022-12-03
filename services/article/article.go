//O'zgartiriladi author => articlega
package article

import (
	"blockpost/genprotos/article"
	"blockpost/storage"
	// "context"
)

// ArticleService is a struct that implements the server interface
type ArticleService struct {
	article.UnimplementedArticleServicesServer
	Stg storage.StorageI
}

// ArticleService ...
// func (a *ArticleService) AddArticle(ctx context.Context, req *article.AddArticleReq) (*article.Article, error) {
// 	res, err := a.Stg.AddArticle(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return res, nil
// }

// func (a *ArticleService) GetArticlesByArticleID(ctx context.Context, req *article.Id) (*article.GetArticles, error) {
// 	res, err := a.Stg.GetArticlesByArticleID(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return res, nil
// }

// func (a *ArticleService) GetArticleList(ctx context.Context, req *article.GetArticleListReq) (*article.GetArticles, error) {
// 	res, err := a.Stg.GetArticleList(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return res, nil
// }

// func (a *ArticleService) GetArticleByID(ctx context.Context, req *article.Id) (*article.GetArticleByIdRes, error) {
// 	res, err := a.Stg.GetArticleByID(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return res, nil
// }

// func (a *ArticleService) UpdateArticle(ctx context.Context, req *article.UpdateArticleReq) (*article.CreateArticleRes, error) {
// 	err := a.Stg.UpdateArticle(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return &author.CreateArticleRes{}, nil
// }

// func (a *ArticleService) DeleteArticle(ctx context.Context, req *article.Id) (*article.CreateArticleRes, error) {
// 	err := a.Stg.DeleteArticle(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return &author.CreateArticleRes{}, nil
// }
