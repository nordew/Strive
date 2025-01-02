package model

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type (
	Goal struct {
		ID          string    `json:"id"`
		UserID      string    `json:"user_id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Chapters    []Chapter `json:"chapters"`
		Progress    int       `json:"progress"` // Progress is a percentage of completed chapters
		IsDone      bool      `json:"is_done"`
		Deadline    time.Time `json:"deadline"`
		Priority    int       `json:"priority"`
		Tags        []string  `json:"tags"`
		Comments    []Comment `json:"comments"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	Chapter struct {
		ID          string    `json:"id"`
		GoalID      string    `json:"goal_id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		IsDone      bool      `json:"is_done"`
		Deadline    time.Time `json:"deadline"`
		Priority    int       `json:"priority"`
		Comments    []Comment `json:"comments"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	Comment struct {
		ID        string    `json:"id"`
		GoalID    string    `json:"goal_id omitempty"`
		ChapterID string    `json:"chapter_id omitempty"`
		Content   string    `json:"content"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func NewGoal(
	id string,
	userID string,
	title string,
	description string,
	chapters []Chapter,
	progress int,
	isDone bool,
	deadline time.Time,
	priority int,
	tags []string,
	comments []Comment,
	createdAt time.Time,
	updatedAt time.Time) (*Goal, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	if userID == "" {
		return nil, errors.New("user_id cannot be empty")
	}

	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	if description == "" {
		return nil, errors.New("description cannot be empty")
	}

	if progress < 0 || progress > 100 {
		return nil, errors.New("progress must be between 0 and 100")
	}

	if priority < 0 {
		return nil, errors.New("priority must be a positive integer")
	}

	if deadline.IsZero() {
		return nil, errors.New("deadline cannot be zero")
	}

	if createdAt.IsZero() {
		return nil, errors.New("created_at cannot be zero")
	}

	if updatedAt.IsZero() {
		return nil, errors.New("updated_at cannot be zero")
	}

	if updatedAt.Before(createdAt) {
		return nil, errors.New("updated_at cannot be before created_at")
	}

	return &Goal{
		ID:          id,
		UserID:      userID,
		Title:       title,
		Description: description,
		Chapters:    chapters,
		Progress:    progress,
		IsDone:      isDone,
		Deadline:    deadline,
		Priority:    priority,
		Tags:        tags,
		Comments:    comments,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}, nil
}

func (g *Goal) SetID(id string) (*Goal, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, errors.New("id must be a valid UUID")
	}

	g.ID = id
	return g, nil
}

func (g *Goal) SetUserID(userID string) (*Goal, error) {
	if userID == "" {
		return nil, errors.New("user_id cannot be empty")
	}

	g.UserID = userID
	return g, nil
}

func (g *Goal) SetTitle(title string) (*Goal, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	g.Title = title
	return g, nil
}

func (g *Goal) SetDescription(description string) (*Goal, error) {
	if description == "" {
		return nil, errors.New("description cannot be empty")
	}

	g.Description = description
	return g, nil
}

func (g *Goal) SetProgress(progress int) (*Goal, error) {
	if progress < 0 || progress > 100 {
		return nil, errors.New("progress must be between 0 and 100")
	}

	g.Progress = progress
	return g, nil
}

func (g *Goal) SetIsDone(isDone bool) (*Goal, error) {
	g.IsDone = isDone
	return g, nil
}

func (g *Goal) SetDeadline(deadline time.Time) (*Goal, error) {
	if deadline.IsZero() {
		return nil, errors.New("deadline cannot be zero")
	} else if deadline.Before(time.Now()) {
		return nil, errors.New("deadline cannot be in the past")
	}

	g.Deadline = deadline
	return g, nil
}

func (g *Goal) SetPriority(priority int) (*Goal, error) {
	if priority < 0 {
		return nil, errors.New("priority must be a positive integer")
	}

	g.Priority = priority
	return g, nil
}

func (g *Goal) SetTags(tags []string) (*Goal, error) {
	g.Tags = tags
	return g, nil
}

func (g *Goal) SetComments(comments []Comment) (*Goal, error) {
	g.Comments = comments
	return g, nil
}

func (g *Goal) AddChapter(chapter Chapter) (*Goal, error) {
	g.Chapters = append(g.Chapters, chapter)
	return g, nil
}

func (g *Goal) RemoveChapter(chapterID string) (*Goal, error) {
	for i, chapter := range g.Chapters {
		if chapter.ID == chapterID {
			g.Chapters = append(g.Chapters[:i], g.Chapters[i+1:]...)
			return g, nil
		}
	}
	return nil, errors.New("chapter not found")
}

func (g *Goal) SetCreatedAt(createdAt time.Time) (*Goal, error) {
	if createdAt.IsZero() {
		return nil, errors.New("created_at cannot be zero")
	} else if createdAt.After(time.Now()) {
		return nil, errors.New("created_at cannot be in the future")
	}

	g.CreatedAt = createdAt
	return g, nil
}

func (g *Goal) SetUpdatedAt(updatedAt time.Time) (*Goal, error) {
	if updatedAt.IsZero() {
		return nil, errors.New("updated_at cannot be zero")
	} else if updatedAt.Before(g.CreatedAt) {
		return nil, errors.New("updated_at cannot be before created_at")
	}

	g.UpdatedAt = updatedAt
	return g, nil
}

func NewChapter(
	id string,
	goalID string,
	title string,
	description string,
	isDone bool,
	deadline time.Time,
	priority int,
	comments []Comment,
	createdAt time.Time,
	updatedAt time.Time) (*Chapter, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, errors.New("id must be a valid UUID")
	}

	if _, err := uuid.Parse(goalID); err != nil {
		return nil, errors.New("goal_id must be a valid UUID")
	}

	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	if description == "" {
		return nil, errors.New("description cannot be empty")
	}

	if priority < 0 {
		return nil, errors.New("priority must be a positive integer")
	}

	if deadline.IsZero() {
		return nil, errors.New("deadline cannot be zero")
	}

	if createdAt.IsZero() {
		return nil, errors.New("created_at cannot be zero")
	}

	if updatedAt.IsZero() {
		return nil, errors.New("updated_at cannot be zero")
	}

	if updatedAt.Before(createdAt) {
		return nil, errors.New("updated_at cannot be before created_at")
	}

	return &Chapter{
		ID:          id,
		GoalID:      goalID,
		Title:       title,
		Description: description,
		IsDone:      isDone,
		Deadline:    deadline,
		Priority:    priority,
		Comments:    comments,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}, nil
}

func (c *Chapter) SetID(id string) (*Chapter, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, errors.New("id must be a valid UUID")
	}

	c.ID = id
	return c, nil
}

func (c *Chapter) SetGoalID(goalID string) (*Chapter, error) {
	if _, err := uuid.Parse(goalID); err != nil {
		return nil, errors.New("id must be a valid UUID")
	}

	c.GoalID = goalID
	return c, nil
}

func (c *Chapter) SetTitle(title string) (*Chapter, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	c.Title = title
	return c, nil
}

func (c *Chapter) SetDescription(description string) (*Chapter, error) {
	if description == "" {
		return nil, errors.New("description cannot be empty")
	}

	c.Description = description
	return c, nil
}

func (c *Chapter) SetIsDone(isDone bool) (*Chapter, error) {
	c.IsDone = isDone
	return c, nil
}

func (c *Chapter) SetDeadline(deadline time.Time) (*Chapter, error) {
	if deadline.IsZero() {
		return nil, errors.New("deadline cannot be zero")
	} else if deadline.Before(time.Now()) {
		return nil, errors.New("deadline cannot be in the past")
	}

	c.Deadline = deadline
	return c, nil
}

func (c *Chapter) SetPriority(priority int) (*Chapter, error) {
	if priority < 0 {
		return nil, errors.New("priority must be a positive integer")
	}

	c.Priority = priority
	return c, nil
}

func (c *Chapter) SetComments(comments []Comment) (*Chapter, error) {
	c.Comments = comments
	return c, nil
}

func (c *Chapter) SetCreatedAt(createdAt time.Time) (*Chapter, error) {
	if createdAt.IsZero() {
		return nil, errors.New("created_at cannot be zero")
	} else if createdAt.After(time.Now()) {
		return nil, errors.New("created_at cannot be in the future")
	}

	c.CreatedAt = createdAt
	return c, nil
}

func (c *Chapter) SetUpdatedAt(updatedAt time.Time) (*Chapter, error) {
	if updatedAt.IsZero() {
		return nil, errors.New("updated_at cannot be zero")
	} else if updatedAt.Before(c.CreatedAt) {
		return nil, errors.New("updated_at cannot be before created_at")
	}

	c.UpdatedAt = updatedAt
	return c, nil
}

func NewComment(
	id string,
	goalID string,
	chapterID string,
	content string,
	createdAt time.Time,
	updatedAt time.Time) (*Comment, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, errors.New("id must be a valid UUID")
	}

	if goalID == "" && chapterID == "" {
		return nil, errors.New("goal_id or chapter_id must be set")
	}

	if content == "" {
		return nil, errors.New("content cannot be empty")
	}

	if createdAt.IsZero() {
		return nil, errors.New("created_at cannot be zero")
	}

	if updatedAt.IsZero() {
		return nil, errors.New("updated_at cannot be zero")
	}

	if updatedAt.Before(createdAt) {
		return nil, errors.New("updated_at cannot be before created_at")
	}

	return &Comment{
		ID:        id,
		GoalID:    goalID,
		ChapterID: chapterID,
		Content:   content,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (c *Comment) SetID(id string) (*Comment, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, errors.New("id must be a valid UUID")
	}

	c.ID = id
	return c, nil
}

func (c *Comment) SetGoalID(goalID string) (*Comment, error) {
	if _, err := uuid.Parse(goalID); err != nil {
		return nil, errors.New("id must be a valid UUID")
	}

	c.GoalID = goalID
	return c, nil
}

func (c *Comment) SetChapterID(chapterID string) (*Comment, error) {
	if chapterID == "" {
		return nil, errors.New("chapter_id cannot be empty")
	}

	c.ChapterID = chapterID
	return c, nil
}

func (c *Comment) SetContent(content string) (*Comment, error) {
	if content == "" {
		return nil, errors.New("content cannot be empty")
	}

	c.Content = content
	return c, nil
}

func (c *Comment) SetCreatedAt(createdAt time.Time) (*Comment, error) {
	if createdAt.IsZero() {
		return nil, errors.New("created_at cannot be zero")
	} else if createdAt.After(time.Now()) {
		return nil, errors.New("created_at cannot be in the future")
	}

	c.CreatedAt = createdAt
	return c, nil
}

func (c *Comment) SetUpdatedAt(updatedAt time.Time) (*Comment, error) {
	if updatedAt.IsZero() {
		return nil, errors.New("updated_at cannot be zero")
	} else if updatedAt.Before(c.CreatedAt) {
		return nil, errors.New("updated_at cannot be before created_at")
	}

	c.UpdatedAt = updatedAt
	return c, nil
}
