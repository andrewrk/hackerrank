#![allow(unstable)]
use std::vec::Vec;

fn main() {
    let mut stdin = std::io::stdio::stdin();
    let line = stdin.read_line().ok().unwrap();
    let v: Vec<&str> = line.trim().split(' ').collect();
    let people_count = v[0].parse().unwrap();
    let topic_count = v[1].parse().unwrap();
    let mut people = Vec::with_capacity(people_count);
    for _ in range(0, people_count as usize) {
        let line = stdin.read_line().ok().unwrap();
        let mut topics = Vec::with_capacity(topic_count);
        for c in line.trim().chars() {
            topics.push(c == '1');
        }
        people.push(topics);
    }
    let mut max_known_count = 0;
    let mut knowledgable_team_count = 0;
    for (p1_index, p1) in people.iter().enumerate() {
        for p2 in people.slice_from(p1_index + 1).iter() {
            let mut known_count = 0;
            for (known1, known2) in p1.iter().zip(p2.iter()) {
                known_count = if *known1 || *known2 {known_count + 1} else {known_count};
            }
            if known_count > max_known_count {
                max_known_count = known_count;
                knowledgable_team_count = 1;
            } else if known_count == max_known_count {
                knowledgable_team_count += 1;
            }
        }
    }
    print!("{}\n{}\n", max_known_count, knowledgable_team_count);
}
