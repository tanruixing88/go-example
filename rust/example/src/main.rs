fn main() {
    let num = 5 / 3;
    //rust char 大小是4字节
    let c = 'z';
    println!("num:{num} c:{c}");

    let tuple = (500, 6.4, 1);
    let (x, y, z) = tuple;
    let (t1, t2, t3) = (tuple.0, tuple.1, tuple.2);
    println!("x:{x}, y:{y}, z:{z} t1:{t1} t2:{t2} t3:{t3}");

    let a = [3; 5];
    println!("a:{:?}", a);
}
