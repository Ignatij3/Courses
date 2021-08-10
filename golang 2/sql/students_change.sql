INSERT INTO students -- 1
(FIRSTNAME, NAME, GRADE, ABC)
VALUES
('Ignatij', 'Krasovskij', 10, 'b');
COMMIT;

INSERT INTO exam -- 2
(FIRSTNAME, NAME, LESSON, GRADE)
VALUES
('Ignatij', 'Krasovskij', 'Chemistry', 6),
('Ignatij', 'Krasovskij', 'English', 10),
('Ignatij', 'Krasovskij', 'History', NULL),
('Ignatij', 'Krasovskij', 'Math', 8),
('Ignatij', 'Krasovskij', 'Physics', 9);
COMMIT;

INSERT INTO lessons -- 3
(LESSON, GRADE, ABC, EDATE)
VALUES
('Archeology', 12, 'D', NULL),
('Arts', 12, 'D', NULL),
('Astronomy', 12, 'D', NULL),
('Biology', 12, 'D', NULL),
('Chemistry', 12, 'D', '2019-05-26'),
('Economics', 12, 'D', '2020-05-18'),
('English', 12, 'D', '2020-06-30'),
('French', 12, 'D', NULL),
('Geography', 12, 'D', NULL),
('Health', 12, 'D', NULL),
('History', 12, 'D', '2020-06-05'),
('Informatics', 12, 'D', '2019-06-01'),
('Literature', 12, 'D', '2019-06-10'),
('Math', 12, 'D', '2019-05-16'),
('Music', 12, 'D', NULL),
('Normative vocabulary', 12, 'D', NULL),
('Physical Education', 12, 'D', NULL),
('Physics', 12, 'D', '2020-05-22'),
('Politics', 12, 'D', '2020-06-13'),
('Science', 12, 'D', '2019-05-29'),
('Spanish', 12, 'D', NULL),
('Statistics', 12, 'D', NULL),
('Technical Education', 12, 'D', NULL);
COMMIT;
-- Записать изменения в отдельный файл, а этот оставить как базу данных

INSERT INTO students -- 4
(FIRSTNAME, NAME, GRADE, ABC)
VALUES
('Andrew', 'Dolginov', 12, 'd'),
('Alina', 'Volkova', 12, 'd'),
('Jeugene', 'Abramov', 12, 'd'),
('Ignatij', 'Krasovskij', 12, 'd');
('Daniel', 'Borovikov', 12, 'd'),
('Karina', 'Celnikova', 12, 'd'),
('Anna', 'Denisova', 12, 'd'),
('Vladimir', 'Melnikov', 12, 'd'),
('Eleonora', 'Ternikova', 12, 'd'),
('Vadim', 'Garikov', 12, 'd'),
('Sintija', 'Jalko', 12, 'd'),
('Alexandra', 'Matkina', 12, 'd'),
('Ilja', 'Polkin', 12, 'd'),
('Sophija', 'Filinova', 12, 'd'),
('Dmitrij', 'Kazanov', 12, 'd');
COMMIT;