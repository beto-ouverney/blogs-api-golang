package blogpostmodel

import (
	"context"
	"time"

	"github.com/beto-ouverney/blogs-api-golang/entities"
	"github.com/beto-ouverney/blogs-api-golang/errors"
)

func (model *modelSqlx) AddBlogPost(ctx context.Context, blogPost *entities.BlogPost) (*entities.BlogPost, *errors.CustomError) {
	tran, err := model.sqlx.BeginTx(ctx, nil)
	if err != nil {
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "blogpostmodel.AddBlogPost", Err: err}
	}
	defer tran.Rollback()

	blogPost.Published = time.Now()
	blogPost.Updated = blogPost.Published
	
	result, err := tran.ExecContext(ctx, "INSERT INTO BlogPosts (title,content,userId,published,updated) VALUES (?,?,?,?,?)", blogPost.Title, blogPost.Content, blogPost.UserID, blogPost.Published, blogPost.Updated)
	if err != nil {
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "blogpostmodel.AddBlogPost", Err: err}
	}

	blogPost.ID, err = result.LastInsertId()
	if err != nil {
		return nil, &errors.CustomError{Code: errors.EINVALID, Op: "blogpostmodel.AddBlogPost", Err: err}
	}

	for _, v := range blogPost.CategoryIDs {
		_, errP := tran.ExecContext(ctx, "INSERT INTO PostCategories (postId,categoryId) VALUES (?,?)", blogPost.ID, v)
		if errP != nil {
			return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "blogpostmodel.AddBlogPost", Err: errP}
		}
	}
	if err = tran.Commit(); err != nil {
		return nil, &errors.CustomError{Code: errors.EINTERNAL, Op: "blogpostmodel.AddBlogPost", Err: err}
	}

	return blogPost, nil
}
