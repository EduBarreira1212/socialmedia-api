insert into users (name, nickName, email, password)
values
("User 1", "user_1", "user1@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"), -- user1
("User 2", "user_2", "user2@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"), -- user2
("User 3", "user_3", "user3@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"); -- user3

insert into followers(user_id, follower_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into posts (title, content, author_id)
values
("Post 1", "Content 1", 1),
("Post 2", "Content 2", 2),
("Post 3", "Content 3", 3);