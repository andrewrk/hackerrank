#![allow(unstable)]

fn main() {
    let mut stdin = std::io::stdio::stdin();
    let line = stdin.read_line().ok().unwrap();
    let v:Vec<&str> = line.trim().split(' ').collect();
    let jar_count:u64 = v[0].parse().unwrap();
    let operation_count:u64 = v[1].parse().unwrap();

    let mut sum = 0;
    for _ in range(0, operation_count as usize) {
        let line = stdin.read_line().ok().unwrap();
        let v:Vec<&str> = line.trim().split(' ').collect();
        let left:u64 = v[0].parse().unwrap();
        let right:u64 = v[1].parse().unwrap();
        let count:u64 = v[2].parse().unwrap();

        sum += (right - left + 1) * count;
    }
    let answer = sum / jar_count;
    println!("{}", answer);
}
