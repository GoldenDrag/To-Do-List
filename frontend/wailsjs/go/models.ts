export namespace repository {
	
	export class Task {
	    id: number;
	    text: string;
	    completed: boolean;
	    priority: string;
	    due_date: string;
	
	    static createFrom(source: any = {}) {
	        return new Task(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.text = source["text"];
	        this.completed = source["completed"];
	        this.priority = source["priority"];
	        this.due_date = source["due_date"];
	    }
	}

}

