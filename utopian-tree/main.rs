struct Simulation {
    heights: Vec<usize>,
}

enum Season {
    Spring,
    Summer,
}

impl Simulation {
    pub fn new() -> Self {
        Simulation { heights: vec![1] }
    }

    fn compute_to(&mut self, cycle: usize) {
        while self.heights.len() <= cycle {
            let last_value = self.heights[self.heights.len() - 1];
            let season = if (self.heights.len() % 2) == 1 {Season::Spring} else {Season::Summer};
            let new_value = match season {
                Season::Spring => last_value * 2,
                Season::Summer => last_value + 1,
            };
            self.heights.push(new_value);
        }
    }

    pub fn height_at(&mut self, cycle: usize) -> usize {
        self.compute_to(cycle);
        return self.heights[cycle];
    }
}

fn main() {
    let mut stdin = std::io::stdio::stdin();

    // discard the first line
    stdin.read_line();

    let mut sim = Simulation::new();

    loop {
        let maybe_line = stdin.read_line();
        match maybe_line {
            Ok(line) => {
                let cycle:usize = line.trim().parse().unwrap();
                println!("{}", sim.height_at(cycle));
            }
            Err(e) => break
        }
    }
}

