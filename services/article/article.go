package article

import (
	"blockpost/genprotos/article"
	"blockpost/storage"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ArticleService is a struct that implements the server interface
type ArticleService struct {
	Stg storage.StorageI
	article.UnimplementedArticleServicesServer
}

// ArticleService ...
func (a *ArticleService) AddArticle(c context.Context, req *article.AddArticleReq) (*article.AddArticleRes, error) {
	id := uuid.New()
	err := a.Stg.AddArticle(id.String(), req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.AddArticle: %s", err.Error())
	}
	art, err := a.Stg.GetArticleByID(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.Stg.GetArticleByID: %s", err.Error())
	}
	return &article.AddArticleRes{
		Id:        art.Id,
		Content:   (*article.AddArticleRes_Post)(art.Content),
		AuthorId:  art.Authori.Id,
		CreatedAt: art.CreatedAt,
		UpdatedAt: art.UpdatedAt,
	}, nil
}

// ArticleService ...
func (a *ArticleService) GetArticleByID(c context.Context, req *article.GetArticleByIdReq) (*article.GetArticleByIdRes, error) {
	res, err := a.Stg.GetArticleByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "a.Stg.GetArticleByID: %s", err.Error())
	}
	return res, nil
}
