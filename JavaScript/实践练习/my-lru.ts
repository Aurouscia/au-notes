type MyKey = string | number

type MyNode<T> = {
    key: MyKey
    value: T,
    prev?:MyNode<T>
    next?:MyNode<T>
}

class MyLRUCache<T> {
    private map: Map<MyKey, MyNode<T>>
    private linkedListHead?: MyNode<T>
    private linkedListTail?: MyNode<T>
    private capacity: number
    constructor(capacity:number){
        this.map = new Map<MyKey, MyNode<T>>()
        this.linkedListHead = undefined
        this.linkedListTail = undefined
        this.capacity = capacity
    }
    get(key:MyKey):T|undefined{
        const node = this.map.get(key)
        if(!node) return undefined
        this.putToNewest(node)
        return node.value
    }
    put(key:MyKey, value: T){
        let node = this.map.get(key)
        if(node) {
            this.putToNewest(node)
            return
        }
        node = { key, value }
        node.next = this.linkedListHead
        this.linkedListHead = node  
        this.map.set(key, node)
        // 如果容量超了，则移除最后一个元素
        if(this.map.size > this.capacity){
            const last = this.linkedListTail!
            const lastButOne = last.prev!
            lastButOne.next = undefined
            this.linkedListTail = lastButOne
            this.map.delete(last.key)
        }
    }
    /** 把 node 的上下家连接，自己从中脱离 */
    private putToNewest(node:MyNode<T>){
        if(node.prev && node.next){
            node.prev.next = node.next
            node.next.prev = node.prev
        }
        else if(node.prev && !node.next){
            this.linkedListTail = node.prev
            node.prev = undefined
        }
        else{
            // 其他情况：仅node自己，或node无prev，则不处理
            return
        }
        node.next = this.linkedListHead
        this.linkedListHead = node
    }
}