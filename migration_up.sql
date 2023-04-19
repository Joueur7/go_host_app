CREATE DATABASE IF NOT EXISTS sticker_db;

USE sticker_db;

CREATE TABLE IF NOT EXISTS trending_stickers (
  id         INT UNSIGNED NOT NULL AUTO_INCREMENT,
  name       VARCHAR(255) NOT NULL,
  image_url  VARCHAR(255) NOT NULL,
  priority   INT          NOT NULL,
  start_time TIME         NOT NULL,
  end_time   TIME         NOT NULL,
  PRIMARY KEY (id)
);
INSERT INTO trending_stickers (name, image_url, priority, start_time, end_time)
VALUES
    ('Sticker 1', 'https://example.com/sticker1.png', 1, '10:00:00', '12:00:00'),
    ('Sticker 2', 'https://example.com/sticker2.png', 2, '12:00:00', '14:00:00'),
    ('Sticker 3', 'https://example.com/sticker3.png', 3, '14:00:00', '16:00:00'),
    ('Garfield (Night)', 'https://garfield.com/', 5, '00:00:00', '05:59:59'),
    ('Garfield (Day)', 'https://garfield.com/', 5, '06:00:00', '11:59:59'),
    ('Hello Kitty (Morning)', 'https://www.sanrio.com/characters/hello-kitty', 9, '06:00:00', '10:59:59'),
    ('Hello Kitty (Afternoon)', 'https://www.sanrio.com/characters/hello-kitty', 9, '11:00:00', '15:59:59'),
    ('Hello Kitty (Evening)', 'https://www.sanrio.com/characters/hello-kitty', 9, '16:00:00', '20:59:59'),
    ('Hello Kitty (Night)', 'https://www.sanrio.com/characters/hello-kitty', 9, '21:00:00', '23:59:59'),
    ('Snoopy (Morning)', 'https://www.snoopy.com/', 6, '06:00:00', '10:59:59'),
    ('Snoopy (Afternoon)', 'https://www.snoopy.com/', 6, '11:00:00', '15:59:59'),
    ('Snoopy (Evening)', 'https://www.snoopy.com/', 6, '16:00:00', '20:59:59'),
    ('Snoopy (Night)', 'https://www.snoopy.com/', 6, '21:00:00', '23:59:59'),
    ('Doraemon (Morning)', 'https://en.wikipedia.org/wiki/Doraemon', 7, '06:00:00', '10:59:59'),
    ('Doraemon (Afternoon)', 'https://en.wikipedia.org/wiki/Doraemon', 7, '11:00:00', '15:59:59'),
    ('Doraemon (Evening)', 'https://en.wikipedia.org/wiki/Doraemon', 7, '16:00:00', '20:59:59'),
    ('Doraemon (Night)', 'https://en.wikipedia.org/wiki/Doraemon', 7, '21:00:00', '23:59:59'),
    ('SpongeBob SquarePants (Morning)', 'https://www.spongebob.com/', 7, '06:00:00', '10:59:59'),
    ('SpongeBob SquarePants (Afternoon)', 'https://www.spongebob.com/', 7, '11:00:00', '15:59:59'),
    ('SpongeBob SquarePants (Evening)', 'https://www.spongebob.com/', 7, '16:00:00', '20:59:59'),
    ('SpongeBob SquarePants (Night)', 'https://www.spongebob.com/', 7, '21:00:00', '23:59:59'),
    ('Tweety Bird (Morning)', 'https://www.warnerbros.com/tweety-bird/', 5, '06:00:00', '10:59:59');

