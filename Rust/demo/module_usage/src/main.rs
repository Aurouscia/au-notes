mod utils;

use utils::math::add;
use utils::str_op::str_slice;

fn main() {
    println!("2+3={}", add(2, 3));
    println!("abcdefg sliced: {}", str_slice("abcdefg"));
}