CREATE TABLE users (
                       id UUID PRIMARY KEY,
                       telegram_id BIGINT UNIQUE NOT NULL,
                       first_name VARCHAR(255),
                       last_name VARCHAR(255),
                       role INTEGER NOT NULL,
                       is_authorized BOOLEAN NOT NULL DEFAULT FALSE,
                       created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE goals (
                       id UUID PRIMARY KEY,
                       user_id UUID NOT NULL,
                       title VARCHAR(255) NOT NULL,
                       description TEXT,
                       progress INT,
                       is_done BOOLEAN DEFAULT FALSE,
                       deadline TIMESTAMP,
                       priority INT,
                       tags TEXT[],
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE chapters (
                          id UUID PRIMARY KEY,
                          goal_id UUID NOT NULL REFERENCES goals(id) ON DELETE CASCADE,
                          title VARCHAR(255) NOT NULL,
                          description TEXT,
                          is_done BOOLEAN DEFAULT FALSE,
                          deadline TIMESTAMP,
                          priority INT,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE comments (
                          id UUID PRIMARY KEY,
                          goal_id UUID REFERENCES goals(id) ON DELETE CASCADE,
                          chapter_id UUID REFERENCES chapters(id) ON DELETE CASCADE,
                          content TEXT NOT NULL,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_goals_user_id ON goals(user_id);
CREATE INDEX idx_chapters_goal_id ON chapters(goal_id);
CREATE INDEX idx_comments_goal_id ON comments(goal_id);
CREATE INDEX idx_comments_chapter_id ON comments(chapter_id);
CREATE INDEX idx_users_telegram_id ON users (telegram_id);