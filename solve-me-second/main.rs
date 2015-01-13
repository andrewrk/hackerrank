fn add(x: usize, y: usize) -> usize {
    x + y
}

fn main() {
    let mut stdin = std::io::stdio::stdin();
    let line = stdin.read_line().ok().unwrap();
    let count = line.trim().parse().unwrap();
    for _ in range(0, count) {
        let line = stdin.read_line().ok().unwrap();
        let v: Vec<&str> = line.trim().split(' ').collect();
        let x = v[0].parse().unwrap();
        let y = v[1].parse().unwrap();
        println!("{}", add(x, y));
    }
}
