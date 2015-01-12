fn abs(x: i32) -> i32 {
    if x >= 0 {x} else {-x}
}

fn count_palindromify(full_string:&str) -> i32 {
    let len = full_string.len();
    let half_len = len / 2;
    let first_half = full_string.slice(0, half_len);
    let second_half = full_string.slice(len - half_len, len).bytes().rev();
    let mut count = 0;
    for (first, second) in first_half.bytes().zip(second_half) {
        count += abs(second as i32 - first as i32);
    }
    return count;
}

fn main() {
    let mut stdin = std::io::stdio::stdin();

    // discard the first line
    stdin.read_line();

    loop {
        match stdin.read_line() {
            Ok(line) => {
                let s = line.trim();
                println!("{}", count_palindromify(s));
            }
            Err(e) => break
        }
    }
}
