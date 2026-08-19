# hexlet-path-size

Консольная утилита, которая выводит размер файла или директории

[![asciicast](https://asciinema.org/a/cFpoHrJsHgoxqCEl.svg)](https://asciinema.org/a/cFpoHrJsHgoxqCEl)

## Установка

```bash
make build
```

## Использование

```
hexlet-path-size [глобальные опции] <путь>
```

`--recursive` `-r` Заходить во вложенные директории и добавлять их содержимое к сумме. Без этой опции вложенные директории не учитываются
`--human` `-H` Выводить размер в читаемом виде (`7.8KB`) вместо количества байт (`8000B`)
`--all` `-a` Учитывать скрытые элементы

### Примеры

Размер одного файла:

```bash
$ hexlet-path-size demo/notes.txt
4096B	demo/notes.txt
```

Размер директории — без `-r` считается только то, что лежит в ней непосредственно:

```bash
$ hexlet-path-size demo
4096B	demo
```

Всё дерево целиком:

```bash
$ hexlet-path-size --recursive demo
471040B	demo
```

То же самое в читаемом виде:

```bash
$ hexlet-path-size -r -H demo
460.0KB	demo
```

Вместе со скрытыми файлами:

```bash
$ hexlet-path-size -r -a -H demo
972.0KB	demo
```

Несуществующий путь:

```bash
$ hexlet-path-size demo/nope.txt
error processing demo/nope.txt: lstat demo/nope.txt: no such file or directory
$ echo $?
1
```

## Разработка

```bash
make lint       # golangci-lint run
make lint-fix   # golangci-lint run --fix
```
