export {}

type MyPromiseResultType = 'resolved' | 'rejected' | 'pending'

class MyPromiseResult<T> {
    status: MyPromiseResultType
    value?: T
    reason?: any
    constructor(type: MyPromiseResultType) {
        this.status = type
    }
}

type MyPromiseExecutor<T> = (resolve: (arg: T) => void, reject: (arg: any) => void) => void

class MyPromise<T> {
    task: MyPromiseExecutor<T>
    result: MyPromiseResult<T>
    // ❌ 错误：缺少回调队列。Promise 是异步的，then 注册时可能还处于 pending，
    // 需要把回调存起来，等 resolve/reject 时异步执行
    private callbacks: Array<{ onFulfilled?: Function, onRejected?: Function, nextResolve: Function, nextReject: Function }> = []

    // ❌ 错误：缺少状态锁。Promise 一旦 settled 就不能再改变状态
    private settled: boolean = false

    constructor(task: MyPromiseExecutor<T>) {
        this.task = task
        this.result = {
            status: 'pending'
        }
        // 构造函数里立即执行 task，传入 resolve/reject 的绑定版本
        const resolveBinded = this.resolve.bind(this)
        const rejectBinded = this.reject.bind(this)
        // ❌ 错误：应该在 try/catch 中执行，捕获异常并 reject
        try {
            this.task(resolveBinded, rejectBinded)
        } catch (err) {
            rejectBinded(err)
        }
    }

    then<T2>(onFulfilled?: (arg: T) => T2, onRejected?: (arg: any) => any): MyPromise<T2> {
        // ❌ 错误：在 then 里重新调用 this.task 是错误的。Promise 的执行器只执行一次，
        // then 返回的新 Promise 应该基于当前 Promise 的状态和回调结果来决定 resolve/reject
        return new MyPromise<T2>((resolve, reject) => {
            const handler = () => {
                try {
                    if (this.result.status === 'resolved') {
                        const ret = onFulfilled ? onFulfilled(this.result.value!) : this.result.value as any
                        resolve(ret)
                    } else {
                        const ret = onRejected ? onRejected(this.result.reason) : this.result.reason
                        reject(ret)
                    }
                } catch (err) {
                    reject(err)
                }
            }

            if (this.result.status === 'pending') {
                // 状态未确定，先存起来
                this.callbacks.push({
                    onFulfilled,
                    onRejected,
                    nextResolve: resolve,
                    nextReject: reject
                })
            } else {
                // ❌ 错误：即使状态已确定，回调也应该异步执行（微任务），这里直接同步执行了
                handler()
            }
        })
    }

    private resolve(arg: T) {
        // ❌ 错误：没有状态锁，可以多次 resolve/reject
        if (this.settled) return
        this.settled = true

        this.result.value = arg
        this.result.status = 'resolved'

        // ❌ 错误：没有处理 thenable（即 arg 本身是一个 Promise 的情况）
        // 根据 Promise/A+ 规范，如果 value 是 thenable，需要递归解析

        // 异步执行回调（简化版用 setTimeout 模拟微任务）
        setTimeout(() => {
            this.callbacks.forEach(cb => {
                try {
                    const ret = cb.onFulfilled ? cb.onFulfilled(arg) : arg
                    cb.nextResolve(ret)
                } catch (err) {
                    cb.nextReject(err)
                }
            })
            this.callbacks = []
        }, 0)
    }

    private reject(reason: any) {
        if (this.settled) return
        this.settled = true

        this.result.reason = reason
        this.result.status = 'rejected'

        setTimeout(() => {
            this.callbacks.forEach(cb => {
                try {
                    if (cb.onRejected) {
                        const ret = cb.onRejected(reason)
                        cb.nextResolve(ret)
                    } else {
                        cb.nextReject(reason)
                    }
                } catch (err) {
                    cb.nextReject(err)
                }
            })
            this.callbacks = []
        }, 0)
    }

    static all<T>(...myProms: MyPromise<T>[]): MyPromise<MyPromiseResult<T>[]> {
        // ❌ 错误：do...while(true) 无限循环会导致死锁，无法等待异步完成
        return new MyPromise((resolve, reject) => {
            if (myProms.length === 0) {
                resolve([])
                return
            }

            let completed = 0
            const results: T[] = new Array(myProms.length)

            myProms.forEach((p, i) => {
                p.then(val => {
                    results[i] = val
                    completed++
                    if (completed === myProms.length) {
                        resolve(results)
                    }
                }, err => {
                    reject(err)
                })
            })
        })
    }
}

const p0 = new MyPromise<string>((r) => {
    setTimeout(() => r('Hello'), 1000)
}).then(x => {
    console.log(x)
})
