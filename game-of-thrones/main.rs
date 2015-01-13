#![allow(unstable)]

fn main() {
    let mut stdin = std::io::stdio::stdin();
    let line = stdin.read_line().ok().unwrap();
    let string = line.trim();
    let mut counts = [0; 256];
    let mut odd_count = 0;
    for c in string.bytes() {
        counts[c as usize] += 1;
        odd_count += match counts[c as usize] % 2 {
            0 => -1,
            _ =>  1,
        };
    }
    let ok = string.len() % 2 == odd_count;
    println!("{}", if ok {"YES"} else {"NO"});
}
