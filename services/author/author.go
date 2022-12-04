package author

import (
	"blockpost/genprotos/author"
	"blockpost/storage"
	"context"
	"fmt"
)

// AuthorService is a struct that implements the server interface
type AuthorService struct {
	author.UnimplementedAuthorServicesServer
	Stg storage.StorageI
}

// AuthorService ...
func (a *AuthorService) AddAuthor(ctx context.Context, req *author.CreateAuthorReq) (*author.CreateAuthorRes, error) {
	res, err := a.Stg.AddAuthor(req)
	if err != nil {
		fmt.Println("myerr---", err)
	}
	return res, nil
}

func (a *AuthorService) GetArticlesByAuthorID(ctx context.Context, req *author.Id) (*author.GetArticles, error) {
	res, err := a.Stg.GetArticlesByAuthorID(req)
	if err != nil {
		panic(err)
	}
	return res, nil
}

func (a *AuthorService) GetAuthorList(ctx context.Context, req *author.GetAuthorListReq) (*author.GetAuthors, error) {
	res, err := a.Stg.GetAuthorList(req)
	if err != nil {
		panic(err)
	}
	return res, nil
}

func (a *AuthorService) GetAuthorByID(ctx context.Context, req *author.Id) (*author.GetAuthorByIdRes, error) {
	res, err := a.Stg.GetAuthorByID(req)
	if err != nil {
		panic(err)
	}
	return res, nil
}

func (a *AuthorService) UpdateAuthor(ctx context.Context, req *author.UpdateAuthorReq) (*author.CreateAuthorRes, error) {
	err := a.Stg.UpdateAuthor(req)
	if err != nil {
		panic(err)
	}
	return &author.CreateAuthorRes{}, nil
}

func (a *AuthorService) DeleteAuthor(ctx context.Context, req *author.Id) (*author.CreateAuthorRes, error) {
	err := a.Stg.DeleteAuthor(req)
	if err != nil {
		panic(err)
	}
	return &author.CreateAuthorRes{}, nil
}
