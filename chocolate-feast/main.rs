#![allow(unstable)]

fn main() {
    let mut stdin = std::io::stdio::stdin();
    let case_count:usize = stdin.read_line().ok().unwrap().trim().parse().unwrap();

    for _ in range(0, case_count) {
        let line = stdin.read_line().ok().unwrap();
        let v:Vec<&str> = line.trim().split(' ').collect();
        let money:i32 = v[0].parse().unwrap();
        let chocolate_price:i32 = v[1].parse().unwrap();
        let wrapper_ratio:i32 = v[2].parse().unwrap();

        let buy_count = money / chocolate_price;
        let mut eat_count = buy_count;
        let mut wrapper_count = buy_count;
        while wrapper_count >= wrapper_ratio {
            let trade_count = wrapper_count / wrapper_ratio;
            wrapper_count = wrapper_count % wrapper_ratio;
            eat_count += trade_count;
            wrapper_count += trade_count;
        }
        println!("{}", eat_count);
    }
}
