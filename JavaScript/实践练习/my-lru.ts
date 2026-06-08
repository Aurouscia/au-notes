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
            // ❌ 错误：当 key 已存在时，只移动了位置，没有更新 value
            node.value = value  // 应添加此行
            this.putToNewest(node)
            return
        }
        node = { key, value }
        node.next = this.linkedListHead
        // ❌ 错误：双向链表只设置了 node.next，没有设置 head.prev
        if (this.linkedListHead) {
            this.linkedListHead.prev = node
        }
        // ❌ 错误：当链表为空时，tail 也应该指向这个 node
        if (!this.linkedListTail) {
            this.linkedListTail = node
        }
        this.linkedListHead = node  
        this.map.set(key, node)
        // 如果容量超了，则移除最后一个元素
        if(this.map.size > this.capacity){
            // ❌ 错误：如果 capacity 为 1，last 可能就是 head，last.prev 为 undefined，这里会报错
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
            // ❌ 错误：node.prev.next 没有断开，链表仍然是连着的
            node.prev.next = undefined  // 应添加此行
            this.linkedListTail = node.prev
            // ❌ 错误：node.prev 置空应该在插入到头部之前，否则 head.prev 指向会出错
            node.prev = undefined
        }
        else if(!node.prev && node.next){
            // ❌ 错误：缺少 node 是 head 的情况。此时 node 已在头部，无需移动
            return
        }
        else{
            // 其他情况：仅node自己，或node无prev，则不处理
            return
        }
        // ❌ 错误：node 被插入到头部后，它的 prev 应该是 undefined，但这里没有处理 node 在中间的情况
        node.prev = undefined
        node.next = this.linkedListHead
        // ❌ 错误：没有设置原 head 的 prev 指向 node
        if (this.linkedListHead) {
            this.linkedListHead.prev = node
        }
        this.linkedListHead = node
    }
}

/* ===== 修正版参考 ===== */
class MyLRUCacheFixed<T> {
    private map: Map<MyKey, MyNode<T>>
    private linkedListHead?: MyNode<T>
    private linkedListTail?: MyNode<T>
    private capacity: number

    constructor(capacity: number) {
        this.map = new Map<MyKey, MyNode<T>>()
        this.linkedListHead = undefined
        this.linkedListTail = undefined
        this.capacity = capacity
    }

    get(key: MyKey): T | undefined {
        const node = this.map.get(key)
        if (!node) return undefined
        this.moveToHead(node)
        return node.value
    }

    put(key: MyKey, value: T) {
        let node = this.map.get(key)
        if (node) {
            node.value = value
            this.moveToHead(node)
            return
        }
        node = { key, value }
        this.map.set(key, node)
        this.addToHead(node)
        if (this.map.size > this.capacity) {
            this.removeTail()
        }
    }

    private addToHead(node: MyNode<T>) {
        node.next = this.linkedListHead
        node.prev = undefined
        if (this.linkedListHead) {
            this.linkedListHead.prev = node
        }
        this.linkedListHead = node
        if (!this.linkedListTail) {
            this.linkedListTail = node
        }
    }

    private removeNode(node: MyNode<T>) {
        if (node.prev) {
            node.prev.next = node.next
        } else {
            this.linkedListHead = node.next
        }
        if (node.next) {
            node.next.prev = node.prev
        } else {
            this.linkedListTail = node.prev
        }
    }

    private moveToHead(node: MyNode<T>) {
        this.removeNode(node)
        this.addToHead(node)
    }

    private removeTail() {
        if (!this.linkedListTail) return
        const node = this.linkedListTail
        this.map.delete(node.key)
        if (node.prev) {
            node.prev.next = undefined
            this.linkedListTail = node.prev
        } else {
            // 只有一个节点
            this.linkedListHead = undefined
            this.linkedListTail = undefined
        }
    }
}
