# README

## About

This is the official Wails React template.

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: https://wails.io/docs/reference/project-config

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.


# What has been done

### 1. Интерфейс пользователя 

Основная часть:
 Создать интерфейс с текстовым полем для ввода новой задачи.<img width="1009" height="845" alt="image" src="https://github.com/user-attachments/assets/4bb1a4aa-b140-4c3b-8008-1fa8f5b6da10" />

 Добавить кнопку для добавления задачи в список. <img width="84" height="69" alt="image" src="https://github.com/user-attachments/assets/516ae580-80d3-45c7-a700-7111bed7a0b6" />

 Отображать список всех задач на экране.<img width="772" height="377" alt="image" src="https://github.com/user-attachments/assets/c0c797b2-e0c3-4d33-ae7a-3f4a7bae0387" />

 Использовать CSS для стилизации интерфейса.
 --    **Repository includes App.css** 
     
 Использовать значки или цвета для обозначения выполненных и невыполненных задач.<img width="763" height="93" alt="image" src="https://github.com/user-attachments/assets/29334966-c506-4d33-a0de-97e4c53284d7" />

Бонусная часть:
 Адаптивная верстка (чтобы корректно смотрелось при изменении размера окна).
 <img width="765" height="627" alt="image" src="https://github.com/user-attachments/assets/5c83304b-a378-40cd-a69e-34657e96a373" />

 Возможность переключения светлой/тёмной темы.<img width="744" height="590" alt="image" src="https://github.com/user-attachments/assets/e9124fa1-2161-4af5-9dda-1b1e8ecb19bb" />


### 2. Добавление задач (20 баллов)

Основная часть :
 Реализовать функционал добавления новой задачи в список.<img width="568" height="122" alt="image" src="https://github.com/user-attachments/assets/ee24c2e2-9b2f-428d-90e6-2ef1644d4a25" />

 Валидация ввода (проверка на пустой ввод).<img width="543" height="89" alt="image" src="https://github.com/user-attachments/assets/1f911397-ba3f-462f-a4ed-0f6e0115d872" />
 ```
        // check input for null
        if(newTask.trim().length == 0) {
            setAlert('Task cannot be doing nothing!');
            return;
        };
```
Бонусная часть :
 Возможность добавлять задачи с датой и временем выполнения.<img width="338" height="317" alt="image" src="https://github.com/user-attachments/assets/cc6fbe52-65d6-41c5-8ddf-935eb07f2b70" />

 Установка приоритета задачи (низкий, средний, высокий).<img width="90" height="110" alt="image" src="https://github.com/user-attachments/assets/8bd1bed2-f752-4c21-bd2e-fa57ecd378a4" />


### 3. Удаление задач

Основная часть :
 Реализовать возможность удаления задач из списка.<img width="93" height="43" alt="image" src="https://github.com/user-attachments/assets/8fc42d4e-4266-487f-8453-896fb0165bce" />
```
    //delete task
    async function deleteTask(index) {
        const id = tasks[index].id;
        await DeleteTask(id);
        const updatedTasks = tasks.filter((_, i) => i !== index);
        setTasks(updatedTasks);
    }
```

Бонусная часть :
 Добавить подтверждение удаления задачи (модальное окно).<img width="139" height="56" alt="image" src="https://github.com/user-attachments/assets/a5049efa-6fb7-4458-b106-80aa72a50626" />


### 4. Управление выполнением задач 

Основная часть :
 Реализовать возможность отметки задачи как выполненной.
 
 Зачеркивание текста выполненных задач.<img width="144" height="89" alt="image" src="https://github.com/user-attachments/assets/f913f574-d62b-489a-a2c8-5d2a28ba3548" />

### 5. Сохранение состояния

Основная часть :
 Сохранение состояния задач при закрытии приложения.
   -- Usage of Database allows saving of tasks and their status after shutdown
   
 Загрузка состояния задач при запуске приложения.
 --- Code below ensures loading of all tasks when startup
```
    // fetch tasks from database depending on filter
    async function fetch() {
        switch (filter) {
            case "active":
                const fetchedActiveTasks = await GetActiveTasks();
                setTasks(fetchedActiveTasks || []);
                break;
            case "completed":
                const fetchedCompletedTasks = await GetCompletedTasks();
                setTasks(fetchedCompletedTasks || []);
                break;
            case "date":
                const fetchedTasksByDate = await GetTasksByDate();
                setTasks(fetchedTasksByDate || []);
                break;
            case "priority":
                const fetchedTasksByPriority = await GetTasksByPriority();
                setTasks(fetchedTasksByPriority || []);
                break;
            case "today":
                const fetchedTasksDueToday = await GetTasksDueToday();
                setTasks(fetchedTasksDueToday || []);
                break;
            case "week":
                const fetchedTasksDueWeek = await GetTasksDueWeek();
                setTasks(fetchedTasksDueWeek || []);
                break;
            case "expired":
                const fetchedExpiredTasks = await GetExpiredTasks();
                setTasks(fetchedExpiredTasks || []);
                break;
            default:
                const fetchedTasks = await GetTasks();
                setTasks(fetchedTasks || []);
        }
    };

    //initial fetch on startup to have tasks already represented
    useEffect(() => {
        fetch();
    }, []);
```
Бонусная часть :
 Использование PostgreSQL для хранения задач.
 ```
func InitDB() *sql.DB {
	// Define the connection string with PostgreSQL credentials
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"

	// Open a database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Ping to confirm connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to PostgreSQL successfully!")

	createEnum(db)
	createTable(db)

	return db
}
```

### 6. Фильтрация и сортировка задач (20 баллов)

Основная часть:

 Фильтрация задач по статусу (все / активные / выполненные).<img width="764" height="176" alt="image" src="https://github.com/user-attachments/assets/b3c2b672-821b-4be0-812c-b8b8219b9bc6" />

 
 Сортировка по дате добавления.<img width="753" height="218" alt="image" src="https://github.com/user-attachments/assets/519478c0-443a-4ab8-8a2d-95f774e63c54" />

 
Бонусная часть:

 Сортировка по приоритету.<img width="764" height="226" alt="image" src="https://github.com/user-attachments/assets/23fd2fe4-9744-4e93-9610-107c2dbc696b" />

 
 Фильтрация по дате (сегодня / на неделю / просроченные).
<img width="105" height="218" alt="image" src="https://github.com/user-attachments/assets/49bd6079-3b9c-46e7-a230-f0e454a5b686" />


 

 
