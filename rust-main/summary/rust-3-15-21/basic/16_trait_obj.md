# day25

## trait obj, generic

`./Cargo.toml`

```rust
[workspace]
members = [
    "gui",
    "main",
]
```



`./gui/src/lib.rc`

```rust
pub trait Draw {
    fn draw(&self);
}

pub struct Screen {
    pub components: Vec<Box<dyn Draw>>,   // trait obj, use dyn keywords， 动态分发
}
impl Screen {
    pub fn run(&self) {
        for comp in self.components.iter() {
            comp.draw();
        }
    }
}

// Use generic
// pub struct Screen<T: Draw> {
//     pub components: Vec<T>,   // generic type 一旦确定, 不能更改， 静态分发
// }

// impl<T> Screen<T> {
//     where T: Draw {
//         pub fn run(&self) {
//             for comp in self.components.iter() {
//                 comp.draw();
//             }
//         }
//     }
// }

pub struct Button {
    pub width: u32,
    pub height: u32,
    pub label: String,
}

impl Draw for Button {
    fn draw(&self) {
        println!("draw button, width: {}, height: {}, label: {}", self.width, self.height, self.label);
    }
}

pub struct SelectBox {
    pub width: u32,
    pub height: u32,
    pub option: Vec<String>,
}

impl Draw for SelectBox {
    fn draw(&self) {
        println!("draw selectbox, width: {}, height: {}, option: {:?}", self.width, self.height, self.option);
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}
```



`./main/src/main.rc`

```rust
use gui::{Screen, Button, SelectBox};

fn main() {
    let s = Screen{
        components: vec![
            Box::new(Button{
                width: 50,
                height: 10,
                label: String::from("Yes"),
            }),
            Box::new(SelectBox{
                width: 70,
                height: 20,
                option: vec![
                    String::from("Yes"),
                    String::from("No"),
                    String::from("MayBe"),
                ],
            }),
        ],
    };

    // Use generic // mismatched types expected struct `gui::Button`, found struct `gui::SelectBox`
    
    // let s = Screen{
    //     components: vec![
    //         Button{
    //             width: 50,
    //             height: 10,
    //             label: String::from("Yes"),
    //         },
    //         SelectBox{  
    //             width: 70,
    //             height: 20,
    //             option: vec![
    //                 String::from("Yes"),
    //                 String::from("No"),
    //                 String::from("MayBe"),
    //             ],
    //         },
    //     ],
    // };

    s.run();  
    println!("Hello, world!");
}
```



res:

```shell
draw button, width: 50, height: 10, label: Yes
draw selectbox, width: 70, height: 20, option: ["Yes", "No", "MayBe"]
Hello, world!
```





## trait Clone

```
//1、trait对象动态分发
//（1）在上述例子中，对泛型类型使用trait bound编译器进行的方式是单态化处理，单态化的代码进行的是静态分发（就是说编译器在编译的时候就知道调用了什么方法）。
//（2）使用 trait 对象时，Rust 必须使用动态分发。编译器无法知晓所有可能用于 trait 对象代码的类型，所以它也不知道应该调用哪个类型的哪个方法实现。为此，Rust 在运行时使用 trait 对象中的指针来知晓需要调用哪个方法。
//
//2、trait对象要求对象安全
//只有 对象安全（object safe）的 trait 才可以组成 trait 对象。trait的方法满足以下两条要求才是对象安全的：
//(1)返回值类型不为 Self
//(2)方法没有任何泛型类型参数

// pub trait Clone {    // trait Clone obj not safely
//     fn clone(&self) -> Self;
// }

// pub struct Screen {
//     pub compnents: Vec<Box<dyn Clone>>, // `Clone` cannot be made into an object
// }

fn main() {
    println!("Hello, world!");
}





//作业
//将《Rust程序设计语言》中17.3节中的例子敲一遍
//实现一个发布博客的工作流，如下：
//博文从空白的草案开始。
//一旦草案完成，请求审核博文。
//一旦博文过审，它将被发表。
//只有被发表的博文的内容会被打印，这样就不会意外打印出没有被审核的博文的文本。
```



