//O'zgartiriladi author => articlega
package article

import (
	"blockpost/genprotos/article"
	"blockpost/storage"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	// "context"
)

// ArticleService is a struct that implements the server interface
type ArticleService struct {
	article.UnimplementedArticleServicesServer
	Stg storage.StorageI
}

// ArticleService ...
func (a *ArticleService) AddArticle(c context.Context, req *article.AddArticleReq) (*article.Article, error) {
	id := uuid.New()
	err := a.Stg.AddArticle(id.String(), req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.AddArticle: %s", err.Error())
	}
	art, err := a.Stg.GetArticleByID(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.Stg.GetArticleByID: %s", err.Error())
	}
	return &article.Article{
		Id:        art.Id,
		Content:   art.Content,
		AuthorId:  art.Authori.Id,
		CreatedAt: art.CreatedAt,
		UpdatedAt: art.UpdatedAt,
	}, nil
}
