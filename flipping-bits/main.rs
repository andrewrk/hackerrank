fn main() {
    let mut stdin = std::io::stdio::stdin();

    // discard the first line
    stdin.read_line();

    loop {
        let maybe_line = stdin.read_line();
        match maybe_line {
            Ok(line) => {
                let x:u32 = line.trim().parse().unwrap();
                let flipped:u32 = x ^ 0b1111111111111111111111111111111111111111;
                println!("{}", flipped);
            }
            Err(e) => break
        }
    }
}
