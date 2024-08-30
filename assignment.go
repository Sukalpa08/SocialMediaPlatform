package main

import (
	"fmt"
	"sync"
	"time"
)

type PostContent struct {
	Content   string
	CreatedAt time.Time
}

type Comment struct {
	ID        int
	PostID    int
	Content   string
	CreatedAt time.Time
}

type LikeDislike struct {
	Likes    int
	Dislikes int
}

type Share struct {
	SharedLink string
}

type Post struct {
	ID          int
	Content     PostContent
	Comments    []Comment
	LikeDislike LikeDislike
	Share       Share
}

type SocialMediaPlatform struct {
	Posts map[int]*Post
	mu    sync.RWMutex
}

func (s *SocialMediaPlatform) CreatePost(content string) *Post {
	s.mu.Lock()
	defer s.mu.Unlock()

	postID := len(s.Posts) + 1
	post := &Post{
		ID: postID,
		Content: PostContent{
			Content:   content,
			CreatedAt: time.Now(),
		},
		LikeDislike: LikeDislike{
			Likes:    0,
			Dislikes: 0,
		},
	}

	s.Posts[postID] = post
	return post
}
func (s *SocialMediaPlatform) AddComment(postID int, commentContent string) *Comment {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, exists := s.Posts[postID]
	if !exists {
		fmt.Println("Post not found!")
		return nil
	}

	commentID := len(post.Comments) + 1
	comment := Comment{
		ID:        commentID,
		PostID:    postID,
		Content:   commentContent,
		CreatedAt: time.Now(),
	}

	post.Comments = append(post.Comments, comment)
	return &comment
}
func (s *SocialMediaPlatform) LikePost(postID int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if post, exists := s.Posts[postID]; exists {
		post.LikeDislike.Likes++
	} else {
		fmt.Println("Post not found!")
	}
}

func (s *SocialMediaPlatform) DislikePost(postID int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if post, exists := s.Posts[postID]; exists {
		post.LikeDislike.Dislikes++
	} else {
		fmt.Println("Post not found!")
	}
}
func (s *SocialMediaPlatform) SharePost(postID int) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, exists := s.Posts[postID]
	if !exists {
		fmt.Println("Post not found!")
		return ""
	}

	post.Share.SharedLink = fmt.Sprintf("https://socialmedia.com/post/%d", post.ID)
	return post.Share.SharedLink
}
func main() {
	platform := &SocialMediaPlatform{
		Posts: make(map[int]*Post),
	}

	post := platform.CreatePost("Hello, this is my first post!")

	platform.AddComment(post.ID, "Nice one!")

	platform.LikePost(post.ID)

	sharedLink := platform.SharePost(post.ID)
	fmt.Printf("Post shared at: %s\n", sharedLink)

	fmt.Printf("Post ID: %d\nContent: %s\nLikes: %d\nDislikes: %d\nComments: %d\nShared Link: %s\n",
		post.ID, post.Content.Content, post.LikeDislike.Likes, post.LikeDislike.Dislikes, len(post.Comments), post.Share.SharedLink)
}
