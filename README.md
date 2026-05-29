# GoArchiver

CLI-архиватор на Go с двумя алгоритмами сжатия: **VLC (Variable Length Coding)** и **Shannon-Fano**.

Проект сделан как pet-проект для портфолио: здесь акцент на архитектуру, работу с бинарными данными и реализацию алгоритмов кодирования без внешних библиотек для компрессии.

## Что умеет

- Упаковывать файл командой `pack`.
- Распаковывать файл командой `unpack`.
- Переключать метод сжатия через флаг `-m` / `--method`:
  - `vlc`
  - `shennon_fano`
- Прогонять тесты и собирать бинарник через `Makefile`.

## Стек

- **Go 1.25.6**
- **Cobra** для CLI
- Стандартная библиотека (`encoding/gob`, `encoding/binary`, `bytes`, `io`, и т.д.)

## Быстрый старт

### 1) Клонирование и запуск тестов/сборки

```bash
make all
```

Команда делает:
- `go test ./...`
- `go build -o moonArchiver ./src`

### 2) Только сборка

```bash
make build
```

После сборки появится бинарник:

```bash
./moonArchiver
```

### 3) Упаковка файла

```bash
./moonArchiver pack path/to/file.txt -m vlc
```

или

```bash
./moonArchiver pack path/to/file.txt -m shennon_fano
```

### 4) Распаковка файла

```bash
./moonArchiver unpack file.txt.vlc -m vlc
```

или

```bash
./moonArchiver unpack file.txt.vlc -m shennon_fano
```

## Как устроен проект

```text
src/
  main.go                          # точка входа
  cmd/                             # CLI-команды (pack/unpack)
  lib/compression/
    compression.go                 # общие интерфейсы и encode в битовую строку
    chunks/                        # работа с битовыми чанками и байтами
    algorithms/
      vlc/                         # алгоритм VLC
      shennon_fano/                # алгоритм Shannon-Fano
    table/                         # таблицы кодирования + дерево декодирования
```

## Архитектурные решения

- Выделены интерфейсы `Encoder` / `Decoder`, чтобы алгоритмы были взаимозаменяемыми.
- Общая логика преобразования в битовый поток вынесена отдельно, алгоритмы переиспользуют инфраструктуру.
- Для Shannon-Fano таблица кодов сериализуется в файл (`gob`) вместе с данными, поэтому декодирование автономно.
- В `table` реализовано дерево декодирования, которое подходит для обоих алгоритмов.