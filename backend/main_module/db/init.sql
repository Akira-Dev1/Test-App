-- Создание курсов
CREATE TABLE IF NOT EXISTS courses (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT
);

-- Создание тестов
CREATE TABLE IF NOT EXISTS tests (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    course_id INT REFERENCES courses(id) ON DELETE CASCADE
);

-- Стартовые данные (для проверки)
INSERT INTO courses (id, name, description) VALUES (1, 'Программирование', 'Базовый курс') ON CONFLICT DO NOTHING;
INSERT INTO tests (name, description, course_id) VALUES 
('Тест 1', 'Описание первого теста', 1),
('Тест 2', 'Описание второго теста', 1) ON CONFLICT DO NOTHING;

-- Сбрасываем счетчик ID
SELECT setval('courses_id_seq', (SELECT MAX(id) FROM courses));
