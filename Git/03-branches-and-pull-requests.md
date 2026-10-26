# 3. Ветки и pull request

На работе новую возможность делают отдельно от рабочей версии, показывают коллегам
и только после проверки добавляют. Для этого есть ветки.

**Ветка** (branch) — отдельная линия коммитов. Основная ветка `main` — рабочая версия
программы, её GitHub создал вместе с репозиторием (ты видела её в **Create codespace
on main** и в строке приглашения: `(main)`). Каждая задача делается в своей ветке,
`main` при этом не меняется. При переключении на другую ветку файлы в Explorer
меняются на её версию: правки, сохранённые коммитом, остаются в своей ветке, правки без
коммита переходят с тобой.

**Pull request** — запрос на GitHub влить ветку в `main`. В нём видны изменения,
к строкам оставляют комментарии — так проходит разбор. **Слияние** (merge) —
перенос коммитов ветки в `main` после одобрения.

## Путь каждой задачи

| № | Что сделать | Команда или кнопка |
|---|---|---|
| 1 | Перейти в `main`, забрать новое, создать ветку | `git switch main`, `git pull`, `git switch -c имя-ветки` |
| 2 | Сделать задачу, сохранить коммитом | `git add .`, `git commit -m "..."` |
| 3 | Отправить ветку в GitHub | первый раз `git push --set-upstream origin имя-ветки`, дальше `git push` |
| 4 | Открыть pull request | **Compare & pull request** на github.com |
| 5 | Ответить на замечания, исправить | коммит и `git push` в ту же ветку |
| 6 | После одобрения влить и вернуться в `main` | **Merge pull request** → **Confirm merge** → **Delete branch**, `git switch main`, `git pull` |

Ниже — пункты подробно, на примере задачи 1 шага 3 и ветки `introduce-frog`.

### 1. Создать ветку

```
$ git switch main
Already on 'main'
Your branch is up to date with 'origin/main'.
$ git pull
Already up to date.
$ git switch -c introduce-frog
Switched to a new branch 'introduce-frog'
```

`-c` — create, создать. `Already on`, `Already up to date` — ты уже в `main`
и нового нет, это не ошибки. Имя ветки — коротко о задаче, латиницей, через дефис.
Текущая ветка видна в строке приглашения и в Status Bar:

![Текущая ветка в Status Bar и в строке приглашения](images/03-branch.png)

*1 — ветка в Status Bar, 2 — в строке приглашения, 3 — облачко: ветки ещё нет
на GitHub, не трогай. На кадре ветка `greet-sister`. Звёздочка после имени — есть
изменения без коммита.*

### 2. Сделать задачу

Код, проверка по критериям, `git add .`, `git commit -m "..."`. Коммитов может быть
несколько.

### 3. Отправить ветку

```
$ git push --set-upstream origin introduce-frog
...
remote: Create a pull request for 'introduce-frog' on GitHub by visiting:
remote:      https://github.com/твой-логин/pond/pull/new/introduce-frog
remote: 
To https://github.com/твой-логин/pond
 * [new branch]      introduce-frog -> introduce-frog
branch 'introduce-frog' set up to track 'origin/introduce-frog'.
```

Команда создаёт ветку на GitHub (`[new branch]`) и связывает её с твоей (`set up
to track`); дальше в этой ветке хватает `git push`. Первые строки ответа (на месте `...`)
такие же, как у `git push` в части 2. В строках `remote:` — ссылка на форму pull
request, запасной путь к пункту 4. Если набрать просто `git push`, git ответит
`fatal: The current branch introduce-frog has no upstream branch.` и напишет нужную
команду; строки про `push.autoSetupRemote` пропусти.

### 4. Открыть pull request

1. Открой `pond` на github.com и нажми **Compare & pull request** в жёлтой плашке.

   ![Плашка Compare & pull request](images/04-compare.png)

2. Заполни форму:

   ![Форма нового pull request](images/05-pr-form.png)

   *1 — **Add a title**: что делает задача (если коммит один, GitHub подставит его
   сообщение). 2 — **Add a description**: ответ на «Вопрос к разбору». 3 — шестерёнка
   **Reviewers**. 4 — **Create pull request**. `base: main ← compare: introduce-frog` —
   ветку вливают в `main`, `Able to merge` — влить можно. Остальное не трогай.*

3. В **Reviewers** нажми шестерёнку, выбери `d-listopad`, нажми на пустое место формы.

   ![Список Reviewers](images/06-reviewers.png)

   *1 — шестерёнка, 2 — поиск. У тебя здесь будет `d-listopad`; «Nothing to show» —
   кадр снят из аккаунта преподавателя.*

4. Нажми **Create pull request**. Откроется pull request с номером; найти его потом —
   вкладка **Pull requests** репозитория.

   ![Страница pull request](images/07-pr-page.png)

   *1 — номер, 2 — Conversation (обсуждение), 3 — Files changed (изменённые строки).*

   Если Codespace спросит «Would you like Visual Studio Code to periodically run
   "git fetch"?» — **No**. Иконку GitHub Pull Requests в Activity Bar не используем.

### 5. Разбор

Преподаватель комментирует строки на вкладке **Files changed**:

![Комментарий к строке на вкладке Files changed](images/08-line-comment.png)

*1 — строка, 2 — Comment, 3 — Start a review. Тебе они нужны, только если сама
захочешь прокомментировать строку.*

Придёт уведомление, комментарии видны на **Conversation**. Отвечай в поле **Reply...**
под комментарием и нажимай **Comment**.

![Комментарий на вкладке Conversation](images/09-conversation.png)

*1 — комментарий, 2 — поле Reply.... **Resolve conversation** не нажимай: закрывает
обсуждение преподаватель.*

![Ответ на комментарий](images/10-reply.png)

*1 — набранный ответ, 2 — Comment. Cancel стирает набранное.*

Исправления — коммит и `git push` в ту же ветку: pull request обновится сам.
Комментарий к изменённой строке GitHub пометит **Outdated**. Когда всё в порядке,
преподаватель одобряет (**Approve**) — появится запись с зелёной галочкой.

### 6. Влить ветку

1. Нажми **Merge pull request** (стрелку рядом не трогай).

   ![Кнопка Merge pull request](images/11-merge.png)

2. Поля сообщения не меняй, нажми **Confirm merge**.

   ![Окно Confirm merge](images/12-confirm-merge.png)

   *1, 2 — сообщение слияния, не меняй; 3 — Confirm merge. Серым закрыта почта.*

3. Нажми **Delete branch**: ветка удалится на GitHub, её коммиты уже в `main`
   (в Codespace она останется, не мешает). **Revert** и **Restore branch** не нажимай.

   ![Кнопка Delete branch](images/13-delete-branch.png)

   > **Осторожно.** Рядом появится **Delete codespace** — не нажимай: удалит Codespace
   > вместе с неотправленными файлами.

4. В Codespace:

   ```
   $ git switch main
   Switched to branch 'main'
   Your branch is up to date with 'origin/main'.
   $ git pull
   remote: Enumerating objects: 1, done.
   remote: Counting objects: 100% (1/1), done.
   remote: Total 1 (delta 0), reused 0 (delta 0), pack-reused 0 (from 0)
   Unpacking objects: 100% (1/1), 966 bytes | 966.00 KiB/s, done.
   From https://github.com/твой-логин/pond
      22d87f2..54638e1  main       -> origin/main
   Updating 22d87f2..54638e1
   Fast-forward
    main.go | 1 +
    1 file changed, 1 insertion(+)
   ```

   `up to date` сравнивает с тем, что git видел на GitHub в прошлый раз, поэтому
   `git pull` нужен. `Fast-forward` — слитые изменения пришли. Коммит «Merge pull
   request #1 …» в `git log` создал GitHub.

Следующая задача начинается с пункта 1.

## Если что-то пошло не так

| Ситуация | Что делать |
|---|---|
| Ветку не создала, коммита ещё нет | `git switch -c имя-ветки`: изменения перейдут в неё |
| Коммит сделан в `main`, push не было | Не делай `git push`. `git switch -c имя-ветки` — коммит окажется в ветке; напиши преподавателю, как вернуть `main` |
| Нет плашки **Compare & pull request** | Ссылка из строк `remote:` или **Pull requests → New pull request**, в **compare** своя ветка |
| `git switch main`: `Your local changes ... would be overwritten` | Сначала коммит в текущей ветке. Совет про `stash` пропусти |
| После merge в Codespace нет изменений | Не выполнены `git switch main` и `git pull` |
