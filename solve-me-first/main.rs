fn main() {
    let mut stdin = std::io::stdio::stdin();
    let mut sum = 0;

    loop {
        match stdin.read_line() {
            Ok(line) => {
                sum += line.trim().parse().unwrap();
            }
            Err(e) => break
        }
    }
    println!("{}", sum);
}
