fn main() {
    let mut sum = 0;
    for maybe_line in std::io::stdin().lock().lines() {
        let line = maybe_line.unwrap();
        let x = line.trim().parse().unwrap();
        sum += x;
    }
    println!("{}", sum);
}
