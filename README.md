## Какую задачу решает код
Печатае цифры от 1 до 10 (и их тип) в обратном порядке.

## Практическое задание от GeekBrains
Для закрепления усвоенного материала повторите шаги раздела «Практическая часть»
самостоятельно для своего проекта:
1. Создайте новый проект с использованием инструментария go mod.
2. Опубликуйте проект в репозитории, установив номер версии, указывающий на активный этап
разработки библиотеки.
3. Обновите номера версий зависимостей в библиотеке.
4. Сделайте изменения в проекте и запушьте

## Краткое описание решения

1. Создал проект GO_2_HOMEWORK_3. Командой go mod init github.com/alexgo92/GO_2_homework_3 создал go.mod, затем командой go mod tidy актуализировал зависимости.
Затем добавили командой go get -u github.com/davecgh/go-spew/spew данный пакет.
2. Отправили все в гитхаб и командой git tag -a v1.0.0 -m "version 1.0.0"