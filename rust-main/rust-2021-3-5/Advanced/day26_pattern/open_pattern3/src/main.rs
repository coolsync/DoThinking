//解构并分解值
//解构元祖、结构体、枚举、引用
//
//解构结构体
struct Point {
    x: i32,
    y: i32,
}

fn main() {
    // let p = Point{x: 1, y: 2};
    // let Point{x: a, y: b} = p;  // variable a,b 分别匹配 x,y
    // assert_eq!(1, a);
    // assert_eq!(2, b);

    // let Point{x, y} = p; // let Point{x:x, y:y} = p;
    // assert_eq!(1, x);
    // assert_eq!(2, y);

    // 一部分 match
    let p = Point{x: 1, y: 0};
    match p {
        Point{x, y:0} => println!("x axis"),
        Point{x:0, y} => println!("y axis"),
        _ => println!("other"),
    };
    
    println!("Hello, world!");
}
