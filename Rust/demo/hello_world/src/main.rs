fn main() {
    println!("Hello, world!");

    let r0: Rectangle = Rectangle {
        width: 10.,
        height: 5.
    };
    println!("r0: {}", r0.calc_area());
    println!("r0: {}", r0.calc_area()); // 声明时，第一参数必须用&self，不能用self，否则会导致consume（没法再次使用）

    let c0: Circle = Circle {
        radius: 32.
    };
    println!("c0: {}", c0.calc_area());
    println!("c0: {}", c0.calc_area());
    println!("c0: {}", Shape2D::calc_area(&c0)); // 另一种调用方式：使用UCS
}

trait Shape2D{
    fn calc_area(&self) -> f32;
}

struct Rectangle{
    width: f32,
    height: f32
}

impl Shape2D for Rectangle {
    fn calc_area(&self) -> f32{
        self.width * self.height
    }
}

struct Circle{
    radius: f32
}

impl Shape2D for Circle {
    fn calc_area(&self) -> f32 {
        std::f32::consts::PI * self.radius * self.radius
    }
}