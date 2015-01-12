fn main() {
    let mut stdin = std::io::stdio::stdin();

    let l:u32 = stdin.read_line().ok().unwrap().trim().parse().unwrap();
    let r:u32 = stdin.read_line().ok().unwrap().trim().parse().unwrap();

    let mut best = 0;
    for x in (l..r + 1) {
        for y in (x..r + 1) {
            let val = x ^ y;
            best = if val > best {val} else {best};
        }
    }
    println!("{}", best);
}
