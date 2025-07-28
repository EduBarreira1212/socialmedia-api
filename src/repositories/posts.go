package repositories

import (
	"database/sql"
	"socialmedia-api/src/models"
)

type Posts struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

func (repository Posts) Create(post models.Post) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO posts (title, content, author_id) VALUES (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

func (repository Posts) GetByID(id uint64) (models.Post, error) {
	row, err := repository.db.Query(`
		SELECT p.*, u.Nickname FROM posts p INNER JOIN users u ON u.id = p.author_id WHERE p.id = ?`,
		id)
	if err != nil {
		return models.Post{}, err
	}
	defer row.Close()

	var post models.Post

	if row.Next() {
		if err := row.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return models.Post{}, err
		}
	}
	return post, nil
}

func (repository Posts) GetByUserID(userID uint64) ([]models.Post, error) {
	rows, err := repository.db.Query(`
		SELECT DISTINCT p.*, u.nickName FROM posts p
		INNER JOIN users u ON u.id = p.author_id
		INNER JOIN followers f ON p.author_id = f.user_id
		WHERE u.id = ? OR f.follower_id = ?
		ORDER BY 1 DESC`,
		userID, userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post

		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository Posts) Update(postID uint64, post models.Post) error {
	statement, err := repository.db.Prepare(
		"UPDATE posts SET title = ?, content = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Content, postID); err != nil {
		return err
	}

	return nil
}

func (repository Posts) Delete(postID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}
