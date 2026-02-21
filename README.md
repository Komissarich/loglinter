# LogLinter — линтер для проверки логов в Go

**LogLinter** — это кастомный линтер, который проверяет сообщения логов на соответствие хорошим практикам:

- Сообщения логов должны быть на английском языке (только буквы a–z A–Z)
- Сообщение должно начинаться с маленькой буквы
- Запрещены специальные символы (кроме пробела)
- Запрещено включать в логи критическую информацию (пароли, токены, api_key и т.д.). Поиск идет по ключевым словам.

## Настройка

Настройка возможна через файл config.yam в папке config. Там можно добавить новые слова для поиска критической информации, новые методы для распознавания вывода логов, а также включить или выключить какие то проверки

Поддерживает `log/slog` и `go.uber.org/zap`.

## Примеры работы

Тесты можно найти в папке /testdata.

# Проверка маленькой буквы:

```go
log.Println("Uppercase log")                  // want `log message 'Uppercase log' should be named 'uppercase log'`
slog.Info("Uppercase log")                    // want `log message 'Uppercase log' should be named 'uppercase log'`

logger, _ := zap.NewProduction()
logger.Info("Uppercase log")                  // want `log message 'Uppercase log' should be named 'uppercase log'`
```

# Проверка кириллицы

```go
log.Println("сtart server")                   // want `log message 'сtart server' should not use cyrillic characters`
slog.Info("старт сервера")                    // want `log message 'старт сервера' should not use cyrillic characters`

logger, _ := zap.NewProduction()
logger.Info("start sеrvеr")                   // want `log message 'start sеrvеr' should not use cyrillic characters`
```

# Проверка критической информации

```go
password := "qwerty"
log.Println("user password" + password)       // want `log message should not contain critical information like password`
slog.Info("user password" + password)         // want `log message should not contain critical information like password`

token := "api_token"
log.Println("api token" + token)              // want `log message should not contain critical information like token`
slog.Info("api token" + token)                // want `log message should not contain critical information like token`

logger, _ := zap.NewProduction()
logger.Info("api token" + token)              // want `log message should not contain critical information like token`
```

# Составные сложные случаи с несколькими ошибками

```go
password := "qwerty"
token := "token"

log.Println("User инфо")                      // want `log message 'User инфо' should be named 'user инфо'; log message 'User инфо' should not use cyrillic characters`

log.Println("привет" + password)              // want `log message should not contain critical information like password; log message 'привет' should not use cyrillic characters`

slog.Info("complex" + token + password)       // want `log message should not contain critical information like password; log message should not contain critical information like token`

logger, _ := zap.NewProduction()
logger.Info("Very" + "сложный" + token + "log!") // want `log message should not contain critical information like token; log message should not use special symbols; log message 'сложный' should not use cyrillic characters; log message 'Very' should be named 'very'`
```

При этом линтер также указывает место, где произошла ошибка, например:

```go
 analysistest.go:713: /home/komissarich/Documents/go_linter_task/testdata/russian_letters.go:13: no diagnostic was reported matching `lыog message 'старт сервера' should not use cyrillic characters`
```

## Установка и использование

1. Склонировать репозиторий:

   ```bash
   git clone https://github.com/Komissarich/loglinter.git
   cd loglinter
   ```

2. Собрать бинарник
   ```bash
   go build -o loglinter ./cmd/main.go
   ```
3. Запустить на своем проекте
   ```bash
   ./loglinter ./...
   ```

К большому сожалению, у меня не получилось интегрировать его в golangci-lint, я убил на это почти целый день и у меня так ничего и не вышло.
Я пытался настроить автоматический вариант с настройкой .custom-gcl.yml и этот файл можно у меня увидеть. Однако какие бы я не перебирал версии, пути к module или import, убирал import, убирал path, все это приводило лишь к тому что golangci-lint custom -v позволял мне действительно собрать бинарник custom-gcl. Вот только ни разу у меня не получилось увидеть в списке линтеров мой loglinter ./custom-gcl linters | grep -i loglinter - всегда пусто. Я следовал статье https://disaev.me/p/writing-useful-go-analysis-linter/ и добавлял обертку, создававшую \*goanalysis.Linter, однако и это не помогало. Я попробовал manual way из документации, но тоже не получилось. Нагуглить тоже ничего хорошего не вышло. Грустно сдаваться, но идей у меня нет вообще никаких, что я делаю не так, нейронки уже давным давно загаллюцинировали и также не могают. Буду очень признателен, если поможете разобраться и укажете что я неправильно делал для интеграции с golangci.... Благодарю за внимание!
