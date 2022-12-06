package article

import (
	"blockpost/genprotos/article"
	"blockpost/storage"
	"context"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ArticleService is a struct that implements the server interface
type ArticleService struct {
	Stg storage.StorageI
	article.UnimplementedArticleServicesServer
}

// AddArticle ...
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

// GetArticleByID ...
func (a *ArticleService) GetArticleByID(c context.Context, req *article.GetArticleByIdReq) (*article.GetArticleByIdRes, error) {
	res, err := a.Stg.GetArticleByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "a.Stg.GetArticleByID: %s", err.Error())
	}
	return res, nil
}

// GetArticleList ...
func (a *ArticleService) GetArticleList(c context.Context, req *article.GetArticleListReq) (*article.GetArticleListRes, error) {
	res, err := a.Stg.GetArticleList(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "a.Stg.GetArticleList: %s", err.Error())
	}
	return res, nil
}

// UpdateArticle ...
func (a *ArticleService) UpdateArticle(c context.Context, req *article.UpdateArticleReq) (*article.UpdateArticleRes, error) {
	err := a.Stg.UpdateArticle(req)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "a.Stg.UpdateArticle: %s", err.Error())
	}
	res, e := a.Stg.GetArticleByID(req.Id)
	if e != nil {
		return nil, status.Errorf(codes.NotFound, "a.Stg.UpdateArticle: %s", e.Error())
	}
	return &article.UpdateArticleRes{
		Id:        res.Id,
		Content:   (*article.UpdateArticleRes_Post)(req.Content),
		Authori:   (*article.UpdateArticleRes_Author)(res.Authori),
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		DeletedAt: res.DeletedAt,
	}, nil
}

// DeleteArticle ...
func (a *ArticleService) DeleteArticle(c context.Context, req *article.DeleteArticleReq) (*article.DeleteArticleRes, error) {
	res, e := a.Stg.GetArticleByID(req.Id)
	if e != nil {
		return nil, status.Errorf(codes.NotFound, "a.Stg.UpdateArticle: %s", e.Error())
	}
	err := a.Stg.DeleteArticle(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "a.Stg.DeleteArticle: %s", err.Error())
	}

	return &article.DeleteArticleRes{
		Id:        res.Id,
		Content:   (*article.DeleteArticleRes_Post)(res.Content),
		AuthorId:  res.Authori.Id,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		DeletedAt: time.Now().Format(time.UnixDate),
	}, nil
}
