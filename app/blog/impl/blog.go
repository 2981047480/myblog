package impl

import (
	"context"
	"fmt"
	"log"
	"time"
	"vblog/app/blog"
	"vblog/common"
	"vblog/common/config"
	"vblog/exception"

	"dario.cat/mergo"
)

// 创建博客
func (i *Blogimpl) CreateBlog(ctx context.Context, in *blog.CreateBlogRequest) (*blog.Blog, error) {
	// 1、创建博客得先有一个创建博客的对象吧，有了得检查下吧
	if err := in.Validate(); err != nil {
		return nil, err
	}

	// 2、创建Blog对象
	b := blog.NewBlog()
	// 直接把这个传进来的请求给blog给赋进去
	b.CreateBlogRequest = in

	// 3、初始化db对象
	i.Init()

	// 4、入库
	if err := i.WithContext(ctx).Create(b).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	// 5、最后没有问题返回
	return b, nil
}

// 删除博客
func (i *Blogimpl) DeleteBlog(ctx context.Context, in *blog.DeleteBlogRequest) (*blog.Blog, error) {
	// 1、首先要弄清楚要删除必须得要有一个标识吧，这里用BlogId作为标识，那么要删除得先有吧，所以得先能查到
	// 先得获取下db对象
	i.Init()
	var count int64
	if err := i.WithContext(ctx).Table(blog.TB_NAME).Where("id = ?", in.BlogId).Count(&count).Error; err != nil {
		return nil, err
	}
	if count == 0 {
		exception.ErrNotFound("要查询的Blog未找到")
	}

	// 2、开始删
	var b = blog.NewBlog()
	if err := i.WithContext(ctx).Table(blog.TB_NAME).Where("id = ?", in.BlogId).Delete(b).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return b, nil
}

func (i *Blogimpl) UpdateBlog(ctx context.Context, in *blog.UpdateBlogRequest) (*blog.Blog, error) {
	/*
		更新有两种模式：
			全量更新（POST请求）
				直接不用管原本的东西是啥 全部都替换了

			增量更新（PATCH请求）
				传什么更新什么
	*/
	// 1、都一样 先获取db对象
	i.Init()
	// 查到需要修改的博客
	b, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(in.BlogId))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	b.UpdateTime = time.Now().Unix()
	if b.PublishAt != 0 {
		b.PublishAt = b.UpdateTime
	}

	// 2、判断更新方式
	switch in.UpdateMode {
	case common.UPDATE_MODE_PATCH:
		if err := b.Validate(); err != nil {
			fmt.Println(err)
			return nil, err
		}
		// 合并查出来的数据和要修改的数据
		if err := mergo.MergeWithOverwrite(b.CreateBlogRequest, in.CreateBlogRequest); err != nil {
			fmt.Println(err)
			return nil, err
		}
	case common.UPDATE_MODE_POST:
		// 直接全量替换了
		b.CreateBlogRequest = in.CreateBlogRequest
	}

	// 3、开始替换
	if err := i.WithContext(ctx).Table(blog.TB_NAME).Save(b).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return b, nil
}

func (i *Blogimpl) UpdateBlogStatus(ctx context.Context, in *blog.UpdateBlogStatusRequest) (*blog.Blog, error) {
	// 1、查不多 也是先查后改
	// 先查
	i.Init()
	b, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(in.BlogId))
	if err != nil {
		log.Println(err)
		return nil, err

	}
	b.UpdateTime = time.Now().Unix()

	// 2、有了一个blog对象之后，然后对status字段进行修改
	b.ChangeBlogStatusRequest = in.ChangeBlogStatusRequest
	if err = i.WithContext(ctx).Table(blog.TB_NAME).Where("id = ?", in.BlogId).Save(&b).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return b, nil
}

// 查询博客
func (i *Blogimpl) QueryBlog(ctx context.Context, in *blog.QueryBlogRequest) (*blog.BlogSet, error) {
	// 1、首先得有db对象
	var err error
	i.DB, err = config.C().GetConn()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// 2、查询意味着需要有个标识，这里采用的是匹配标题的方法
	query := i.DB.WithContext(ctx).Table(blog.TB_NAME)
	if in.Keywords != "" {
		query = query.Where("title like ?", "%"+in.Keywords+"%")
	}

	// 先得查count，这里需要有一个BlogSet对象
	set := blog.NewBlogSet()
	query.Count(&set.Total)

	// 查完了count再查一下blogset
	if err := query.Limit(in.PageSize).Offset(in.OffSet()).Find(&set.Item).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	// 3、没有错误就返回set
	return set, nil
}

func (i *Blogimpl) DescribeBlog(ctx context.Context, in *blog.DescribeBlogRequest) (*blog.Blog, error) {
	// 先得获取下db对象
	i.Init()
	b := blog.NewBlog()

	if err := i.WithContext(ctx).Where("id = ?", in.BlogId).First(b).Error; err != nil {
		return nil, err
	}

	return b, nil
}
