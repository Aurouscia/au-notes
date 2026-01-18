interface School{
    studentIds:number[]
    teacherIds:Array<number>
}

const NUAA:School = {
    studentIds:[1,7,6,8],
    teacherIds:[9,16,53]
}

const names = ["Alice", "Bob", "Eve"];
 
// Contextual typing for function - parameter s inferred to have type string
names.forEach(function (s) {
  console.log(s.toUpperCase());
})


function isString(x: unknown): x is string {
    return typeof x === 'number'; // 不会报错，因为ts无条件信任is类型谓词
}


const obj = { a: 1, b: 2 };
function isKeyOf<T extends object>(
  k: PropertyKey,
  o: T
): k is keyof T {
  return k in o;
}
for (const key in obj) {
    // console.log(obj[key]); // 直接写会出问题
    if (isKeyOf(key, obj)) {
        console.log(obj[key]); // OK，key 被收窄成 'a' | 'b'
    }
}