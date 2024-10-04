package biz

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"my-template-with-go/internal/model"
	"my-template-with-go/internal/repo"
	"my-template-with-go/request"
	"my-template-with-go/response"
	"sync"
	"time"
)

const workerCount = 100
const userCount = 1000

type IArticleUC interface {
	Sync(ctx context.Context) error

	List(ctx echo.Context) (interface{}, error)
	Detail(id uint) (interface{}, error)
	Create(jBody *request.ArticleCreateReq) error
	Edit(id uint, jBody *request.ArticleUpdateReq) error
	Delete(jBody *request.ArticleDeleteReq) error
}

type articleUC struct {
	articleRepo repo.IArticleRepo
	userRepo    repo.IUserRepo
}

func (b *articleUC) Sync(ctx context.Context) error {
	jobs := make(chan int, userCount)
	var wg sync.WaitGroup

	// Start worker pool
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go b.worker(ctx, &wg, jobs)
	}

	for i := 0; i < userCount; i++ {
		select {
		case jobs <- i: // Enqueue user job
		case <-ctx.Done():
			// Stop sending jobs if context is canceled
			fmt.Println("Job submission canceled:", ctx.Err())
			close(jobs)
			wg.Wait()
			return ctx.Err()
		}
	}

	close(jobs)
	wg.Wait()

	return nil
	//users := 1000
	//var wg sync.WaitGroup
	//
	//for i := 0; i < users; i++ {
	//	wg.Add(1)
	//	go func(user int) {
	//		defer wg.Done()
	//
	//		// Context-aware function call
	//		select {
	//		case <-ctx.Done():
	//			// Exit if the context is canceled
	//			fmt.Printf("User %d job canceled: %v\n", user, ctx.Err())
	//			return
	//		default:
	//			// Do the actual work
	//			b.doSomeThing2(user)
	//		}
	//	}(i)
	//}
	//
	//done := make(chan struct{})
	//go func() {
	//	wg.Wait()
	//	close(done)
	//}()
	//
	//select {
	//case <-done:
	//	// All goroutines finished
	//	fmt.Println("All jobs completed successfully")
	//	return nil
	//case <-ctx.Done():
	//	// Context timeout or cancellation
	//	fmt.Println("Job canceled:", ctx.Err())
	//	return ctx.Err()
	//}

	//go func() {
	//	for errorChan := range chanError {
	//		fmt.Println("error: ", errorChan)
	//	}
	//	close(chanError)
	//}()
	//
	//go func() {
	//	for resultChan := range chanResult {
	//		fmt.Println("success: ", resultChan)
	//	}
	//	close(chanResult)
	//}()
}

func (b *articleUC) worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan int) {
	defer wg.Done()
	for user := range jobs {
		if err := b.doSomeThing2(ctx, user); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func (b *articleUC) doSomeThing(user int, chanError chan<- error, chanResult chan<- string) {
	if user%2 == 0 {
		chanResult <- fmt.Sprintf("user %d ok", user)
	} else {
		chanError <- fmt.Errorf("user %d failed", user)
	}
}

func (b *articleUC) doSomeThing2(ctx context.Context, user int) error {
	//time.Sleep(9 * time.Second)
	//if user%2 == 0 {
	//	fmt.Println(fmt.Sprintf("user %d ok", user))
	//	//chanResult <- fmt.Sprintf("user %d ok", user)
	//} else {
	//	fmt.Println(fmt.Sprintf("user %d failed", user))
	//	//chanError <- fmt.Errorf("user %d failed", user)
	//}

	//return nil

	select {
	//case <-time.After(10 * time.Second): // Simulate work
	//	fmt.Printf("Processed user: %d\n", user)
	//	return nil
	case <-ctx.Done():
		// Handle context cancellation
		fmt.Printf("Processing canceled for user %d: %v\n", user, ctx.Err())
		return ctx.Err()
	default:
		if user%2 == 0 {
			time.Sleep(20 * time.Second)
			fmt.Println(fmt.Sprintf("user %d ok", user))
			//chanResult <- fmt.Sprintf("user %d ok", user)
		} else {
			fmt.Println(fmt.Sprintf("user %d failed", user))
			//chanError <- fmt.Errorf("user %d failed", user)
		}
		return nil
	}
}

func (b *articleUC) List(ctx echo.Context) (interface{}, error) {
	var (
		res []*response.ArticleListRes
	)

	exist, err := b.userRepo.CheckUserExist(1)
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, fmt.Errorf("user not exist")
	}

	// do something with login business
	articles, err := b.articleRepo.List()
	if err != nil {
		return nil, err
	}

	if len(articles) > 0 {
		res = make([]*response.ArticleListRes, 0, len(articles))
		for _, a := range articles {
			temp := &response.ArticleListRes{}
			temp.SetAttributes(a)
			res = append(res, temp)
		}
	}

	return res, nil
}

func (b *articleUC) Detail(id uint) (interface{}, error) {
	var (
		res *response.ArticleDetailRes
	)

	// do something with login business
	article, err := b.articleRepo.Detail(id)
	if err != nil {
		return nil, err
	}

	res = &response.ArticleDetailRes{}
	res.SetAttributes(article)

	return res, nil
}

func (b *articleUC) Create(jBody *request.ArticleCreateReq) error {
	// do something with login business
	articleEntity := model.ToArticleEntity(jBody.Author, jBody.Title)

	if err := b.articleRepo.Create(articleEntity); err != nil {
		return err
	}
	return nil
}

func (b *articleUC) Edit(id uint, jBody *request.ArticleUpdateReq) error {
	var (
		updateItems = map[string]interface{}{}
	)

	// do something with login business
	if jBody.Author != nil {
		updateItems["author"] = jBody.Author
	}

	if jBody.Title != nil {
		updateItems["title"] = jBody.Title
	}

	if err := b.articleRepo.Update(id, updateItems); err != nil {
		return err
	}

	return nil
}

func (b *articleUC) Delete(jBody *request.ArticleDeleteReq) error {
	// do something with login business
	return b.articleRepo.Delete(jBody.IDs)
}

func NewArticleUseCase(articleRepo repo.IArticleRepo, userRepo repo.IUserRepo) IArticleUC {
	return &articleUC{
		articleRepo: articleRepo,
		userRepo:    userRepo,
	}
}
