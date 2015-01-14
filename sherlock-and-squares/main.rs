#![allow(unstable)]

use std::num::Float;

fn main() {
    let mut stdin = std::io::stdio::stdin();
    let case_count:usize = stdin.read_line().ok().unwrap().trim().parse().unwrap();

    for _ in range(0, case_count) {
        let line = stdin.read_line().ok().unwrap();
        let v:Vec<&str> = line.trim().split(' ').collect();
        let left:i64 = v[0].parse().unwrap();
        let right:i64 = v[1].parse().unwrap();

        let sqrt_left = (left as f64).sqrt().ceil() as i64;
        let sqrt_right = (right as f64).sqrt().floor() as i64;
        println!("{}", sqrt_right - sqrt_left + 1);
    }
}
