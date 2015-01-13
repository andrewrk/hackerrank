#![allow(unstable)]

fn main() {
    let mut stdin = std::io::stdio::stdin();

    // discard the redundant first line
    stdin.read_line();

    loop {
        match stdin.read_line() {
            Ok(line) => {
                let mut sum:isize = 0;
                for x in line.split(' ') {
                    sum += x.trim().parse().unwrap();
                }
                println!("{}", sum);
            }
            Err(e) => break
        }
    }
}
