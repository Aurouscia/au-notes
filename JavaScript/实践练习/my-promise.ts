export {}

type MyPromiseResultType = 'resolved'|'rejected'|'pending'

class MyPromiseResult<T>{
    status: MyPromiseResultType
    value?: T
    reason?: any
    constructor(type: MyPromiseResultType){
        this.status = type
    }
}

type MyPromiseExecutor<T> = (resolve:(arg:T)=>void, reject:(arg:any)=>void)=>void

class MyPromise<T> {
    task: MyPromiseExecutor<T>
    result: MyPromiseResult<T>
    constructor(task: MyPromiseExecutor<T>){
        this.task = task
        this.result = {
            status: 'pending'
        }
    }
    then<T2>(anotherTask:(arg:T)=>T2): MyPromise<T2>{
        // 我不确定这个行为对不对
        return new MyPromise<T2>((resolve, reject)=>{
            const resolveBinded = this.resolve.bind(this)
            const rejectBinded = this.reject.bind(this)
            this.task(resolveBinded, rejectBinded)
            if(this.result.status == 'resolved')
                resolve(anotherTask(this.result.value!))
            if(this.result.status == 'rejected')
                reject(this.result.reason)
            return
        })
    }
    private resolve(arg:T){
        this.result.value = arg
        this.result.status = 'resolved'
    }
    private reject(reason:any){
        this.result.reason = reason
        this.result.status = 'rejected'
    }
    static all<T>(...myProms:MyPromise<T>[]): MyPromise<MyPromiseResult<T>[]>{
        // 我想不明白怎么搞
        return new MyPromise((resolve, reject)=>{
            let result = new MyPromiseResult<MyPromiseResult<T>[]>('pending')
            let values:(T|undefined)[] = []
            do {
                values = myProms.map(x=>x.result.value)
            } while(true)
        })
    }
}

const p0 = new MyPromise<string>((r)=>{
    setTimeout(()=>r('Hello'), 1000)
}).then(x=>{
    console.log(x)
})