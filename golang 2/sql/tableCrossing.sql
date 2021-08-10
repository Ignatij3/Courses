/*                                                    -- Удаляйте комментарии блок за блоком (так удобнее, чем сразу всё смотреть) + там, где 0.x, там задания из тетради, а остальные с листка "Первое задание по базам данных"
SELECT s.GRADE, s.ABC, e.LESSON, AVG(e.ITEM) -- 0.1
FROM exam e, students s
WHERE (s.FIRSTNAME = e.FIRSTNAME) AND (s.NAME = e.NAME)
  AND (s.GRADE = 10) AND (s.ABC = "a")
  AND (e.LESSON = "Math")
GROUP BY s.GRADE, s.ABC;

SELECT s.GRADE, s.ABC, AVG(e.ITEM) -- 0.2
FROM exam e, students s
WHERE (s.FIRSTNAME = e.FIRSTNAME) AND (s.NAME = e.NAME)
GROUP BY s.GRADE, s.ABC
ORDER BY s.GRADE, s.ABC;

SELECT l.GRADE, l.ABC, e.LESSON, COUNT(e.ITEM) -- 0.3
FROM lessons l
INNER JOIN students s ON (l.GRADE = s.GRADE) AND (l.ABC = s.ABC)
INNER JOIN exam e ON (s.FIRSTNAME = e.FIRSTNAME) AND (s.NAME = e.NAME) AND (e.LESSON = l.LESSON)
WHERE (item IS NULL OR item < 4)
GROUP BY l.GRADE, l.ABC, e.LESSON
HAVING COUNT(item) >= 10
ORDER BY l.GRADE, l.ABC, l.LESSON;

SELECT DISTINCT s.FIRSTNAME, s.NAME -- 0.4
FROM students s
WHERE 1 < (SELECT COUNT(*)
           FROM exam e
			  WHERE e.firstname = s.firstname
			    AND e.NAME = s.NAME
				 AND e.ITEM < 4)
ORDER BY s.FIRSTNAME, s.NAME;

SELECT l.GRADE, l.ABC, l.LESSON -- 0.5
FROM lessons l
WHERE 2 <= (SELECT COUNT(l.LESSON)
            FROM lessons l
			   WHERE l.EDATE IS NOT NULL)
GROUP BY l.GRADE, l.ABC, l.LESSON;

SELECT s.FIRSTNAME, s.NAME, s.GRADE, s.ABC, AVG(e.ITEM) -- 0.6
FROM students s, exam e
WHERE (s.FIRSTNAME = e.FIRSTNAME) AND (s.NAME = e.NAME)
GROUP BY s.GRADE, s.ABC
HAVING (SELECT COUNT(s.FIRSTNAME AND s.NAME)
        FROM students s
		  WHERE AVG(e.ITEM)) > AVG(e.ITEM);

SELECT s.FIRSTNAME, s.NAME -- 1
FROM students s
RIGHT JOIN exam e ON e.FIRSTNAME = s.FIRSTNAME AND e.NAME = s.NAME
GROUP BY s.FIRSTNAME, s.NAME
HAVING COUNT(e.LESSON) = COUNT(e.ITEM);

SELECT s.FIRSTNAME, s.NAME -- 2
FROM students s
INNER JOIN lessons l ON l.GRADE = s.GRADE AND l.ABC = s.ABC AND (l.EDATE IS NOT NULL)
GROUP BY s.FIRSTNAME, s.NAME;

SELECT l.GRADE, l.ABC -- 3
FROM lessons l
INNER JOIN students s ON (l.GRADE = s.GRADE) AND (l.ABC = s.ABC)
INNER JOIN exam e ON (s.FIRSTNAME = e.FIRSTNAME) AND (s.NAME = e.NAME) AND (e.LESSON = l.LESSON) AND (e.ITEM IS NULL OR e.ITEM <= 4)
WHERE l.LESSON = "Chemistry"
GROUP BY l.GRADE, l.ABC
ORDER BY l.GRADE, l.ABC;

SELECT s.FIRSTNAME, s.NAME -- 4 - 5 - 6 - 12
FROM students s
INNER JOIN exam e ON e.FIRSTNAME = s.FIRSTNAME AND e.NAME = s.NAME
WHERE (e.ITEM >= 5 OR e.ITEM IS NULL)
GROUP BY s.FIRSTNAME, s.NAME
HAVING COUNT(e.ITEM) = 0;                             -- Не совсем понятно, что значит "пока" в 5-м + утверждения про время не имеют смысла

SELECT DISTINCT e.FIRSTNAME, e.NAME, e.LESSON, e.ITEM, l.EDATE -- 7
FROM exam e, lessons l, students s
WHERE (l.GRADE = s.GRADE)
  AND (l.ABC = s.ABC)
  AND (e.FIRSTNAME = s.FIRSTNAME)
  AND (e.NAME = s.NAME)
  AND l.EDATE IS NOT NULL
  AND ((e.ITEM <= 4)
  OR (l.EDATE < '2019-12-17' AND e.ITEM IS NULL)
  OR (l.EDATE > '2019-12-17' AND e.ITEM IS NULL))
GROUP BY s.FIRSTNAME, s.NAME, e.ITEM, l.EDATE;        -- Здесь непонятно, какая "сегодня" дата, и к тому же везде уже поставлены оценки

SELECT s.FIRSTNAME, s.NAME -- 8
FROM students s
INNER JOIN exam e ON e.FIRSTNAME = s.FIRSTNAME AND e.NAME = s.NAME
WHERE (e.ITEM < 4 OR e.ITEM IS NULL)
GROUP BY s.FIRSTNAME, s.NAME
HAVING COUNT(e.ITEM) = 0;

SELECT s.FIRSTNAME, s.NAME -- 9 - аналогично 5-му
FROM students s
INNER JOIN exam e ON e.FIRSTNAME = s.FIRSTNAME AND e.NAME = s.NAME
WHERE (e.ITEM >= 7 OR e.ITEM IS NULL)
GROUP BY s.FIRSTNAME, s.NAME
HAVING COUNT(e.ITEM) = 0;

SELECT s.FIRSTNAME, s.NAME -- 10
FROM students s
INNER JOIN exam e ON e.FIRSTNAME = s.FIRSTNAME AND e.NAME = s.NAME AND (e.ITEM < 5 OR e.ITEM IS NULL)
GROUP BY s.FIRSTNAME, s.NAME
HAVING COUNT(e.ITEM) = 0;                             -- Здесь если написать "IS NOT NULL", то никого не покажет, значит никто не сдал сессию на хорошую оценку без неявки на хотя бы один предмет

SELECT e.FIRSTNAME, e.NAME, e.LESSON, e.ITEM, l.EDATE -- 11
FROM exam e
INNER JOIN lessons l ON (l.LESSON = e.LESSON)
INNER JOIN students s ON (s.FIRSTNAME = e.FIRSTNAME) AND (s.NAME = e.NAME) AND (l.ABC = s.ABC)
  WHERE (l.EDATE > '2019-12-17')
  AND e.ITEM IS NULL                                  -- Тут я написал это, что бы внести видимость того, что экзамена ещё не было
GROUP BY e.FIRSTNAME, e.NAME, e.LESSON, e.ITEM, l.EDATE;

SELECT l.LESSON -- 13
FROM lessons l
WHERE l.EDATE IS NOT NULL
GROUP BY l.LESSON
HAVING COUNT(l.LESSON) = 0;                           -- Если вместо нуля написать другую цифру, например 3, то покажет список, в которых дата есть 3 раза, а почему не показывает так?

SELECT l.GRADE, l.ABC -- 14
FROM exam e, lessons l, students s
WHERE (l.GRADE = s.GRADE)
  AND (l.ABC = s.ABC)
  AND (e.FIRSTNAME = s.FIRSTNAME)
  AND (e.NAME = s.NAME)
  AND e.ITEM = 10
GROUP BY l.GRADE, l.ABC
ORDER BY l.GRADE, l.ABC;

SELECT l.GRADE, l.ABC -- 15
FROM lessons l
INNER JOIN exam e ON  (e.LESSON = l.LESSON)
INNER JOIN students s ON (l.GRADE = s.GRADE) AND (l.ABC = s.ABC) AND (s.FIRSTNAME = e.FIRSTNAME) AND (s.NAME = e.NAME)
WHERE (e.ITEM < 4 OR e.ITEM IS NULL)
GROUP BY l.GRADE, l.ABC
ORDER BY l.GRADE, l.ABC;

SELECT e.FIRSTNAME, e.NAME -- 16
FROM exam e
INNER JOIN students s ON (e.FIRSTNAME = s.FIRSTNAME) AND (e.NAME = s.NAME)
INNER JOIN lessons l ON (l.GRADE = s.GRADE) AND (l.ABC = s.ABC) AND (e.LESSON = l.LESSON)
WHERE (e.FIRSTNAME <> s.FIRSTNAME AND e.NAME <> s.NAME)
  AND (s.GRADE <> l.GRADE AND s.ABC <> l.ABC)
GROUP BY e.FIRSTNAME, e.NAME;

SELECT e.FIRSTNAME, e.NAME -- 17
FROM exam e
INNER JOIN students s ON (e.FIRSTNAME = s.FIRSTNAME) AND (e.NAME = s.NAME)
INNER JOIN lessons l ON (l.GRADE = s.GRADE) AND (l.ABC = s.ABC) AND (e.LESSON = l.LESSON)
WHERE (e.ITEM IS NOT NULL OR e.ITEM = 0)
  AND l.EDATE IS NULL
GROUP BY e.FIRSTNAME, e.NAME
HAVING COUNT(e.LESSON) = 0;

SELECT s.FIRSTNAME, s.NAME, s.GRADE, s.ABC, l1.lesson -- 18
FROM students s
INNER JOIN lessons l1 ON (s.grade = l1.grade) AND (s.ABC = l1.abc)
WHERE l1.edate IS NOT NULL
  AND NOT EXISTS (
    SELECT *
	 FROM exam e
	 WHERE (e.FIRSTNAME = s.FIRSTNAME) AND (e.NAME = s.NAME))

SELECT e.LESSON -- 19
FROM exam e
WHERE (e.ITEM IS NULL OR e.ITEM = 1 OR e.ITEM = 2 OR e.ITEM = 3)
GROUP BY e.LESSON;

SELECT s.GRADE, s.ABC -- 20
FROM students s
INNER JOIN lessons l ON (l.GRADE = s.GRADE) AND (l.ABC = s.ABC)
WHERE l.LESSON IS NULL
GROUP BY s.GRADE, s.ABC;

SELECT l.GRADE, l.ABC -- 21
FROM lessons l
INNER JOIN students s ON (s.GRADE = l.GRADE) AND (s.ABC = l.ABC)
WHERE (s.FIRSTNAME IS NULL AND s.NAME IS NULL)
GROUP BY l.GRADE, l.ABC;

SELECT s.FIRSTNAME, s.NAME, s.GRADE, s.ABC -- 22
FROM students s
WHERE s.GRADE = (SELECT s.GRADE
					  FROM students s
					  WHERE (s.FIRSTNAME = "Theo" AND s.NAME = "Allen"))
  AND s.ABC = (SELECT s.ABC
					FROM students s
					WHERE (s.FIRSTNAME = "Theo" AND s.NAME = "Allen"))
GROUP BY s.FIRSTNAME, s.NAME;

SELECT e.LESSON, l.EDATE -- 23
FROM exam e
INNER JOIN lessons l ON (e.LESSON = l.LESSON)
  WHERE (l.EDATE > '2019-12-17')
  AND e.ITEM IS NULL                                  -- Аналогичный 11-му бред с датами
GROUP BY e.LESSON, l.EDATE;

SELECT COUNT(l.LESSON), l.GRADE, l.ABC -- 24
FROM lessons l
GROUP BY l.GRADE, l.ABC
ORDER BY COUNT(l.LESSON) DESC;

SELECT COUNT(l.LESSON) -- 25
FROM lessons l
GROUP BY l.GRADE, l.ABC
LIMIT 1;

SELECT COUNT(s.FIRSTNAME AND s.NAME), s.GRADE, s.ABC -- 26
FROM students s
GROUP BY s.GRADE, s.ABC
ORDER BY COUNT(s.FIRSTNAME AND s.NAME) DESC;
*/