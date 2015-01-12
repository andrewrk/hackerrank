enum Expect {
    A,
    B,
    Any,
}

fn count_deletions(s:&str) -> i32 {
    let mut del_count = 0;
    let mut expect = Expect::Any;
    for c in s.chars() {
        match (&expect, c) {
            (&Expect::B, 'A') | (&Expect::A, 'B') => del_count += 1,
            (&Expect::A, 'A') => expect = Expect::B,
            (&Expect::B, 'B') => expect = Expect::A,
            (&Expect::Any, 'A') => expect = Expect::B,
            (&Expect::Any, 'B') => expect = Expect::A,
            _ => panic!("invalid character"),
        }
    }
    return del_count;
}

fn main() {
    let mut stdin = std::io::stdio::stdin();

    // discard the first line
    stdin.read_line();

    loop {
        match stdin.read_line() {
            Ok(line) => {
                let s = line.trim();
                println!("{}", count_deletions(s));
            }
            Err(e) => break
        }
    }
}
