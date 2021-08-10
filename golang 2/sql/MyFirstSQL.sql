SELECT AVG(item), firstname, name -- 0.1
FROM exam
GROUP BY firstname, name;

SELECT COUNT(distinct firstname, name) -- 0.2
FROM exam
WHERE (item IS NULL) OR (item < 4);

SELECT MIN(firstname), MAX(firstname), MIN(name), MAX(name), grade, abc -- 0.3
FROM students
GROUP BY grade, abc
ORDER BY grade, abc;

SELECT MAX(edate), lesson -- 0.4
FROM lessons
GROUP BY lesson;

SELECT COUNT(item), lesson -- 0.5
FROM exam
WHERE (item IS NULL)
GROUP BY lesson;

SELECT AVG(item), lesson -- 0.6
FROM exam
WHERE NOT (item IS NULL)
GROUP BY lesson;

SELECT COUNT(edate), lesson -- 0.7 remake?
FROM lessons
WHERE NOT (edate IS NULL)
GROUP BY lesson;

SELECT COUNT(firstname), grade, abc -- 26
FROM students
GROUP BY grade, abc
ORDER BY COUNT(firstname) DESC;


SELECT COUNT(lesson), grade, abc -- 25
FROM lessons
GROUP BY grade, abc
ORDER BY COUNT(lesson) DESC
LIMIT 1;

