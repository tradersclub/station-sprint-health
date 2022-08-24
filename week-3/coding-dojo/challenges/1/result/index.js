class Node {
    next;
    value;

    constructor(value) {
        this.value = value;
    }

    exec() {
        this.value()
    }
}

class Queue {
    root;
    tail;

    push(value) {
        const newNode = new Node(value)

        if (!this.root) {
            this.root = newNode;
        }
        
        if (this.tail) {
            this.tail.next = newNode;
        }

        this.tail = newNode;
    }
    
    pop() {

        if (this.root) {
            const lastRoot = this.root
            this.root = lastRoot.next

            if (this.root === undefined) {
                this.tail = undefined
            }

            return lastRoot
        }

        return 
    }
}



const queue = new Queue();

queue.push(() => console.log("1"))
queue.push(() => console.log("20"))
queue.push(() => console.log("10"))

queue.pop()?.exec()
queue.pop()?.exec()

queue.push(() => console.log("39"))

queue.pop()?.exec()
queue.pop()?.exec()
queue.pop()?.exec()




