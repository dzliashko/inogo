# inoreader-go — учебный Go backend

Учебный проект: backend RSS/news reader, вдохновлённый Inoreader.

## Что уже настроено

- `AGENTS.md` — правила работы Codex как преподавателя и code reviewer.
- `STUDY_PLAN.md` — учебный маршрут.
- `PROGRESS.md` — постоянный прогресс обучения.
- `.zed/settings.json` — минимальные проектные настройки Zed для Go.
- `.zed/tasks.json` — повторяемые Go-команды.
- `.gitignore` — исключения для build/test artifacts, secrets и временных файлов.
- `go.mod` — минимальный Go module.

Архитектура backend (`cmd/`, `internal/`, БД и т. п.) намеренно НЕ создана заранее.
Она должна появляться постепенно по мере прохождения учебного плана.

## Где хранить проект

Рекомендуемый путь внутри WSL:

```bash
~/projects/inoreader-go
```

Распакуйте содержимое репозитория туда и открывайте корень проекта в Zed через WSL.

## Проверить Go

```bash
go version
which go
go install golang.org/x/tools/gopls@latest
gopls version
```

Если `$HOME/go/bin` отсутствует в `PATH`, добавьте в `~/.bashrc`:

```bash
export PATH="$PATH:$HOME/go/bin"
source ~/.bashrc
```

## Установить Codex в Zed

Через Command Palette:

```text
zed: acp registry
```

или:

```text
agent: open settings
→ External Agents
→ Add Agent
→ Install from Registry
→ Codex
```

После установки создайте новый thread именно с **Codex**.

## Первый запрос Codex

```text
Мы начинаем обучение.

Сначала прочитай:
- AGENTS.md
- STUDY_PLAN.md
- PROGRESS.md

Также осмотри текущее состояние репозитория.

Следуй правилам AGENTS.md:
не реализуй проект за меня и не давай полные решения упражнений без моего явного запроса.

Начни с текущего шага, указанного в PROGRESS.md.
Если требуется первоначальная оценка моего уровня — проведи её.

После оценки выбери первый подходящий учебный блок.
```

## Проверка контекста

```text
Прочитай инструкции проекта и расскажи:
1. какова твоя роль в этом репозитории;
2. какова моя учебная цель;
3. какой проект мы создаём;
4. какой язык используется;
5. должен ли ты выполнять учебные задания вместо меня;
6. какое следующее действие указано в PROGRESS.md.

Ничего в проекте пока не изменяй.
```

## Zed tasks

Откройте:

```text
Ctrl+Shift+P
task: spawn
```

Доступны:

- `Go: Test All`
- `Go: Test Race`
- `Go: Vet`
- `Go: Format`
- `Go: Build All`
- `Go: Mod Tidy`

`Go: Run` намеренно пока не добавлен: в проекте ещё нет `main` package.
Когда в ходе обучения появится реальная точка входа, задачу запуска нужно добавить под фактическую структуру проекта.

## Первый Git commit

```bash
git init
git status
git add .
git commit -m "chore: initialize Go learning repository"
```

## Рабочий сценарий

В новом Codex thread обычно достаточно:

```text
Продолжаем обучение.
```

После выполнения упражнения:

```text
Я закончил упражнение. Проверь моё решение по правилам AGENTS.md.
```
