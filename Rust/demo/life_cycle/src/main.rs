fn main() {
    let full_str = String::from("abcdef"); // 所有权始终在main
    let sliced = get_slice_str(&full_str); // 借给函数使用
    println!("sliced: {}", sliced)
}

fn get_slice_str<'a>(initial: &'a str) -> &'a str{ // 返回值的生命周期与输入值相同
    let holder = Holder { // holder的所有权在本函数内，本函数结束即销毁
        text: initial // 这里相当于把'a赋给了Holder中的<'b>
    };
    let sliced = holder.slice();
    sliced
}

struct Holder<'b> {
    text: &'b str, // 内部引用使用的生命周期标注必须在struct的泛型参数定义
}

impl<'b> Holder<'b> {
    fn slice(&self) -> &'b str { // 明确slice返回的&str的生命周期与text相同
        &self.text[0..3]
    }
}