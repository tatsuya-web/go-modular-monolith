// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package post_service

import (
	"context"
	"github.com/tatuya-web/go-modular-monolith/modules/post_module/post_model"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_model"
	"github.com/tatuya-web/go-modular-monolith/repository"
	"sync"
)

// Ensure, that PostAdderMock does implement PostAdder.
// If this is not the case, regenerate this file with moq.
var _ PostAdder = &PostAdderMock{}

// PostAdderMock is a mock implementation of PostAdder.
//
//	func TestSomethingThatUsesPostAdder(t *testing.T) {
//
//		// make and configure a mocked PostAdder
//		mockedPostAdder := &PostAdderMock{
//			AddPostFunc: func(ctx context.Context, db repository.Execer, p *post_model.Post) error {
//				panic("mock out the AddPost method")
//			},
//		}
//
//		// use mockedPostAdder in code that requires PostAdder
//		// and then make assertions.
//
//	}
type PostAdderMock struct {
	// AddPostFunc mocks the AddPost method.
	AddPostFunc func(ctx context.Context, db repository.Execer, p *post_model.Post) error

	// calls tracks calls to the methods.
	calls struct {
		// AddPost holds details about calls to the AddPost method.
		AddPost []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db repository.Execer
			// P is the p argument value.
			P *post_model.Post
		}
	}
	lockAddPost sync.RWMutex
}

// AddPost calls AddPostFunc.
func (mock *PostAdderMock) AddPost(ctx context.Context, db repository.Execer, p *post_model.Post) error {
	if mock.AddPostFunc == nil {
		panic("PostAdderMock.AddPostFunc: method is nil but PostAdder.AddPost was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Db  repository.Execer
		P   *post_model.Post
	}{
		Ctx: ctx,
		Db:  db,
		P:   p,
	}
	mock.lockAddPost.Lock()
	mock.calls.AddPost = append(mock.calls.AddPost, callInfo)
	mock.lockAddPost.Unlock()
	return mock.AddPostFunc(ctx, db, p)
}

// AddPostCalls gets all the calls that were made to AddPost.
// Check the length with:
//
//	len(mockedPostAdder.AddPostCalls())
func (mock *PostAdderMock) AddPostCalls() []struct {
	Ctx context.Context
	Db  repository.Execer
	P   *post_model.Post
} {
	var calls []struct {
		Ctx context.Context
		Db  repository.Execer
		P   *post_model.Post
	}
	mock.lockAddPost.RLock()
	calls = mock.calls.AddPost
	mock.lockAddPost.RUnlock()
	return calls
}

// Ensure, that PostUpdaterMock does implement PostUpdater.
// If this is not the case, regenerate this file with moq.
var _ PostUpdater = &PostUpdaterMock{}

// PostUpdaterMock is a mock implementation of PostUpdater.
//
//	func TestSomethingThatUsesPostUpdater(t *testing.T) {
//
//		// make and configure a mocked PostUpdater
//		mockedPostUpdater := &PostUpdaterMock{
//			IsOwnPostFunc: func(ctx context.Context, db repository.Queryer, id post_model.PostID) bool {
//				panic("mock out the IsOwnPost method")
//			},
//			UpdatePostFunc: func(ctx context.Context, db repository.Execer, p *post_model.Post) error {
//				panic("mock out the UpdatePost method")
//			},
//		}
//
//		// use mockedPostUpdater in code that requires PostUpdater
//		// and then make assertions.
//
//	}
type PostUpdaterMock struct {
	// IsOwnPostFunc mocks the IsOwnPost method.
	IsOwnPostFunc func(ctx context.Context, db repository.Queryer, id post_model.PostID) bool

	// UpdatePostFunc mocks the UpdatePost method.
	UpdatePostFunc func(ctx context.Context, db repository.Execer, p *post_model.Post) error

	// calls tracks calls to the methods.
	calls struct {
		// IsOwnPost holds details about calls to the IsOwnPost method.
		IsOwnPost []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db repository.Queryer
			// ID is the id argument value.
			ID post_model.PostID
		}
		// UpdatePost holds details about calls to the UpdatePost method.
		UpdatePost []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db repository.Execer
			// P is the p argument value.
			P *post_model.Post
		}
	}
	lockIsOwnPost  sync.RWMutex
	lockUpdatePost sync.RWMutex
}

// IsOwnPost calls IsOwnPostFunc.
func (mock *PostUpdaterMock) IsOwnPost(ctx context.Context, db repository.Queryer, id post_model.PostID) bool {
	if mock.IsOwnPostFunc == nil {
		panic("PostUpdaterMock.IsOwnPostFunc: method is nil but PostUpdater.IsOwnPost was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Db  repository.Queryer
		ID  post_model.PostID
	}{
		Ctx: ctx,
		Db:  db,
		ID:  id,
	}
	mock.lockIsOwnPost.Lock()
	mock.calls.IsOwnPost = append(mock.calls.IsOwnPost, callInfo)
	mock.lockIsOwnPost.Unlock()
	return mock.IsOwnPostFunc(ctx, db, id)
}

// IsOwnPostCalls gets all the calls that were made to IsOwnPost.
// Check the length with:
//
//	len(mockedPostUpdater.IsOwnPostCalls())
func (mock *PostUpdaterMock) IsOwnPostCalls() []struct {
	Ctx context.Context
	Db  repository.Queryer
	ID  post_model.PostID
} {
	var calls []struct {
		Ctx context.Context
		Db  repository.Queryer
		ID  post_model.PostID
	}
	mock.lockIsOwnPost.RLock()
	calls = mock.calls.IsOwnPost
	mock.lockIsOwnPost.RUnlock()
	return calls
}

// UpdatePost calls UpdatePostFunc.
func (mock *PostUpdaterMock) UpdatePost(ctx context.Context, db repository.Execer, p *post_model.Post) error {
	if mock.UpdatePostFunc == nil {
		panic("PostUpdaterMock.UpdatePostFunc: method is nil but PostUpdater.UpdatePost was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Db  repository.Execer
		P   *post_model.Post
	}{
		Ctx: ctx,
		Db:  db,
		P:   p,
	}
	mock.lockUpdatePost.Lock()
	mock.calls.UpdatePost = append(mock.calls.UpdatePost, callInfo)
	mock.lockUpdatePost.Unlock()
	return mock.UpdatePostFunc(ctx, db, p)
}

// UpdatePostCalls gets all the calls that were made to UpdatePost.
// Check the length with:
//
//	len(mockedPostUpdater.UpdatePostCalls())
func (mock *PostUpdaterMock) UpdatePostCalls() []struct {
	Ctx context.Context
	Db  repository.Execer
	P   *post_model.Post
} {
	var calls []struct {
		Ctx context.Context
		Db  repository.Execer
		P   *post_model.Post
	}
	mock.lockUpdatePost.RLock()
	calls = mock.calls.UpdatePost
	mock.lockUpdatePost.RUnlock()
	return calls
}

// Ensure, that PostListerMock does implement PostLister.
// If this is not the case, regenerate this file with moq.
var _ PostLister = &PostListerMock{}

// PostListerMock is a mock implementation of PostLister.
//
//	func TestSomethingThatUsesPostLister(t *testing.T) {
//
//		// make and configure a mocked PostLister
//		mockedPostLister := &PostListerMock{
//			ListPostsFunc: func(ctx context.Context, db repository.Queryer, id user_model.UserID) (post_model.Posts, error) {
//				panic("mock out the ListPosts method")
//			},
//		}
//
//		// use mockedPostLister in code that requires PostLister
//		// and then make assertions.
//
//	}
type PostListerMock struct {
	// ListPostsFunc mocks the ListPosts method.
	ListPostsFunc func(ctx context.Context, db repository.Queryer, id user_model.UserID) (post_model.Posts, error)

	// calls tracks calls to the methods.
	calls struct {
		// ListPosts holds details about calls to the ListPosts method.
		ListPosts []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db repository.Queryer
			// ID is the id argument value.
			ID user_model.UserID
		}
	}
	lockListPosts sync.RWMutex
}

// ListPosts calls ListPostsFunc.
func (mock *PostListerMock) ListPosts(ctx context.Context, db repository.Queryer, id user_model.UserID) (post_model.Posts, error) {
	if mock.ListPostsFunc == nil {
		panic("PostListerMock.ListPostsFunc: method is nil but PostLister.ListPosts was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Db  repository.Queryer
		ID  user_model.UserID
	}{
		Ctx: ctx,
		Db:  db,
		ID:  id,
	}
	mock.lockListPosts.Lock()
	mock.calls.ListPosts = append(mock.calls.ListPosts, callInfo)
	mock.lockListPosts.Unlock()
	return mock.ListPostsFunc(ctx, db, id)
}

// ListPostsCalls gets all the calls that were made to ListPosts.
// Check the length with:
//
//	len(mockedPostLister.ListPostsCalls())
func (mock *PostListerMock) ListPostsCalls() []struct {
	Ctx context.Context
	Db  repository.Queryer
	ID  user_model.UserID
} {
	var calls []struct {
		Ctx context.Context
		Db  repository.Queryer
		ID  user_model.UserID
	}
	mock.lockListPosts.RLock()
	calls = mock.calls.ListPosts
	mock.lockListPosts.RUnlock()
	return calls
}
