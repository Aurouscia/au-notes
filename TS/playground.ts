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