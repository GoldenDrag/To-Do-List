import {useState, useEffect} from 'react';
import Popup from 'reactjs-popup';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {GetTasks, AddTask, UpdateTask, DeleteTask} from "../wailsjs/go/main/App";
import {GetActiveTasks, GetCompletedTasks, GetTasksByDate, GetTasksByPriority, GetTasksDueToday, GetTasksDueWeek, GetExpiredTasks} from "../wailsjs/go/main/App";

function App() {
    const [darkMode, setDarkMode] = useState(true);   
    const [filter, setFilter] = useState('all');   
    const [alert, setAlert] = useState('Welcome to the best To-Do list!');
    const [tasks, setTasks] = useState([]);
    const [newTask, setNewTask] = useState('');
    const [priority, setPriority] = useState('low');
    const [dueDate, setDueDate] = useState(new Date());

    const updateFilter = (e) => {
        setFilter(e.target.value);
    };
    const updateNewTask = (e) => setNewTask(e.target.value);
    const updatePriority = (e) => setPriority(e.target.value);
    const updateDueDate = (e) => setDueDate(e.target.value);

    const task = {
        text: newTask,
        completed: false,
        due_date: dueDate,
        priority: priority};

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

    //Loop among tasks and display them
    const UsingArrayMap = () => {
        return (
            <div className='tasks-list'>
                {tasks.map((task, index) => (
                    <div className={"task " + (task.completed ? 'completed ' : ' ') + task.priority} key={index}>
                        <input type="checkbox" checked={task.completed} onChange={() => toggleTaskCompletion(index)} />
                        <span className={"property " + (task.completed ? 'completed ' : ' ')} >
                           {task.text}
                        </span>
                        <span className={"property priority"} >
                           {task.priority}
                        </span>
                        <span className={"property"} >
                           {task.due_date}
                        </span>
                        <Popup trigger={<btn className="del-btn" onClick={() => deleteTask(index)}>X</btn>} 
                            onClick={() => deleteTask(index)}
                            modal nested>
                            {
                                close => (
                                    <div className='modal'>
                                        <div className='content'>
                                            Task deleted
                                        </div>
                                        <div>
                                            <button onClick=
                                                {() => {deleteTask(index) ; close()}}>
                                                    Close modal
                                            </button>
                                        </div>
                                    </div>
                                )
                            }
                        </Popup>
                    </div>
                ))}
            </div>
        );
    }   

    // add new task
    async function addTask() {
        // check input for null
        if(newTask.trim().length == 0) {
            setAlert('Task cannot be doing nothing!');
            return;
        };
        console.log(task);

        const id = await AddTask(task);
        const newTaskWithId = {...tasks, task};
        
        useEffect(() => {
            setTasks([...tasks, newTaskWithId]);
        }, []);
    };

    //change status on checkbox
    async function toggleTaskCompletion(index) {
        const updatedTasks = [...tasks];
        updatedTasks[index].completed = !updatedTasks[index].completed;
        await UpdateTask(updatedTasks[index]);
        setTasks(updatedTasks);
    }

    //delete task
    async function deleteTask(index) {
        const id = tasks[index].id;
        await DeleteTask(id);
        const updatedTasks = tasks.filter((_, i) => i !== index);
        setTasks(updatedTasks);
    }
    
    return (
        <div id="App" className={darkMode ? 'dark' : 'light'}>
            <div className='theme-toggle'>
                <button onClick={() => setDarkMode(!darkMode)}>
                    Toggle Dark Mode
                </button>
                <h5>I need this job</h5>
            </div>
            <img src={logo} id="logo" alt="logo"/>
            <div id="result" className="result">{alert}</div>
            <div id="input" className="input-box">
                <input className="input" placeholder="Enter a new task" onChange={updateNewTask} />
                <input type="datetime-local" className="input" onChange={updateDueDate}/>
                <select className="input" onChange={updatePriority}>
                    <option value="low">Low</option>
                    <option value="mid">Medium</option>
                    <option value="high">High</option>
                </select>
                <button className="btn" onClick={addTask}>Save</button>
            </div>
            <div className='Filters'>
                <select onChange={updateFilter}>
                    <option value="all">All</option>
                    <option value="active">Active</option>
                    <option value="completed">Completed</option>
                    <option value="date">By Date</option>
                    <option value="priority">Priority</option>
                    <option value="today">Due Today</option>
                    <option value="week">Due in a week</option>
                    <option value="expired">Expired</option>
                </select>
                <button onClick={fetch}>Refresh</button>
            </div>
            <UsingArrayMap/>
        </div>
    )
}

export default App
