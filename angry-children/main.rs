#![allow(unstable)]

use std::mem;

fn main() {
    let mut stdin = std::io::stdio::stdin();

    // discard the first line
    stdin.read_line();

    let candy_count:usize = stdin.read_line().ok().unwrap().trim().parse().unwrap();
    let mut list = vec![];

    loop {
        match stdin.read_line() {
            Ok(line) => {
                let x:i32 = line.trim().parse().unwrap();
                list.push(x);
            }
            Err(e) => break
        }
    }

    list.sort();

    let mut heap = MinMaxHeap::new(candy_count);
    let mut least_unfairness = std::i32::MAX;
    for item in list.iter() {
        heap.push(*item);
        if heap.len() == candy_count {
            let unfairness:i32 = *heap.peek_max().unwrap() - *heap.peek_min().unwrap();
            least_unfairness = if unfairness < least_unfairness {unfairness} else {least_unfairness};
        }
    }
    println!("{}", least_unfairness);
}


pub struct MinMaxHeap<T> {
    dat: Vec<T>,
    cap: usize
}

impl<T:Ord+Clone> MinMaxHeap<T> {

    pub fn new(cap: usize) -> MinMaxHeap<T> {
        MinMaxHeap { dat: vec![], cap: cap }
    }

    pub fn with_capacity(cap: usize) -> MinMaxHeap<T> {
        MinMaxHeap { dat: Vec::with_capacity(cap), cap: cap }
    }

    pub fn from_vec(v: Vec<T>) -> MinMaxHeap<T> {
        let len = v.len();
        let mut out = MinMaxHeap { dat: v, cap: len };
        for i in range(0, out.len()).rev() {
            out.trickle_down(i);
        }
        out
    }

    pub fn from_vec_growable(v: Vec<T>) -> MinMaxHeap<T> {
        let mut out = MinMaxHeap { dat: v, cap: 0 };
        for i in range(0, out.len()).rev() {
            out.trickle_down(i);
        }
        out
    }

    pub fn len(&self) -> usize {
        self.dat.len()
    }

    pub fn is_empty(&self) -> bool {
        self.len() == 0
    }

    pub fn is_capped(&self) -> bool {
        self.cap != 0
    }

    pub fn peek_min(&self) -> Option<&T> {
        if self.is_empty() { None } else { Some(&self.dat[0]) }
    }

    pub fn peek_max(&self) -> Option<&T> {
        match self.max_idx() {
            None => None,
            Some(i) => Some(&self.dat[i])
        }
    }

    fn max_idx(&self) -> Option<usize> {
        match self.len() {
            0 => None,
            1 => Some(0),
            2 => Some(1),
            _ => Some(if self.dat[1] >= self.dat[2] { 1 } else { 2 })
        }
    }

    pub fn pop_min(&mut self) -> Option<T> {
        match self.len() {
            0 => None,
            1 => self.dat.pop(),
            _ => {
                let out = self.dat.swap_remove(0);
                self.trickle_down(0);
                Some(out)
            }
        }
    }

    pub fn pop_max(&mut self) -> Option<T> {
        match self.len() {
            0 => None,
            1|2 => self.dat.pop(),
            3 => {
                if self.dat[1] >= self.dat[2] {
                    Some(self.dat.swap_remove(1))
                } else {
                    self.dat.pop()
                }
            },
            _ => {
                let idx = if self.dat[1] >= self.dat[2] { 1 } else { 2 };
                let out = self.dat.swap_remove(idx);
                self.trickle_down(idx);
                Some(out)
            }
        }
    }

    pub fn push(&mut self, item: T) -> Option<T> {
        self.push_min(item)
    }

    pub fn push_all(&mut self, items: &[T]) {
        self.push_all_min(items)
    }

    pub fn push_min(&mut self, item: T) -> Option<T> {
        if self.cap == 0 || self.len() < self.cap {
            self.push_grow(item);
            None
        } else if *self.peek_min().unwrap() < item {
            let out = mem::replace(&mut self.dat[0], item);
            self.trickle_down(0);
            Some(out)
        } else {
            Some(item)
        }
    }

    pub fn push_all_min(&mut self, items: &[T]) {
        for i in items.iter() {
            self.push_min((*i).clone());
        }
    }

    pub fn push_max(&mut self, item: T) -> Option<T> {
        if self.cap == 0 || self.len() < self.cap {
            self.push_grow(item);
            None
        } else if *self.peek_max().unwrap() > item {
            let idx = self.max_idx().unwrap();
            let out = mem::replace(&mut self.dat[idx], item);
            self.trickle_down(idx);
            Some(out)
        } else {
            Some(item)
        }
    }

    pub fn push_all_max(&mut self, items: &[T]) {
        for i in items.iter() {
            self.push_max((*i).clone());
        }
    }

    fn push_grow(&mut self, item: T) {
        self.dat.push(item);
        let last = self.len() - 1;
        self.bubble_up(last);
    }

    fn bubble_up(&mut self, i: usize) {
        if i == 0 {
            return;
        }
        let o = match i.level_type() {
            LevelType::Min => std::cmp::Ordering::Greater,
            LevelType::Max => std::cmp::Ordering::Less
        };
        if self.dat[i].cmp(&self.dat[i.parent()]) == o {
            self.dat.as_mut_slice().swap(i, i.parent());
            self._bubble_up(i.parent(), o);
        } else {
            self._bubble_up(i, rev_order(o));
        }
    }

    fn _bubble_up(&mut self, i: usize, o: std::cmp::Ordering) {
        if i < 3 { // no grandparent
            return;
        }
        if self.dat[i].cmp(&self.dat[i.grandparent()]) == o {
            self.dat.as_mut_slice().swap(i, i.grandparent());
            self._bubble_up(i.grandparent(), o);
        }
    }

    fn trickle_down(&mut self, i: usize) {
        match i.level_type() {
            LevelType::Min => self._trickle_down(i, std::cmp::Ordering::Less),
            LevelType::Max => self._trickle_down(i, std::cmp::Ordering::Greater)
        }
    }

    fn _trickle_down(&mut self, i: usize, o: std::cmp::Ordering) {
        let m = self.child_or_grandchild(i, o);
        if m == 0 {
            return;
        }
        if self.dat[m].cmp(&self.dat[i]) == o {
            self.dat.as_mut_slice().swap(m, i);
        }
        if !m.is_child_of(i) { //m is a grandchild
            if self.dat[m.parent()].cmp(&self.dat[m]) == o {
                self.dat.as_mut_slice().swap(m, m.parent());
            }
            self._trickle_down(m, o);
        }
    }

    fn child_or_grandchild(&self, i: usize, o: std::cmp::Ordering) -> usize {
        let l = i.left();
        if l < self.dat.len() {
            let mut out = l;
            let r = i.right();
            for idx in [r, l.left(), l.right(), r.left(), r.right()].iter() {
                if *idx >= self.dat.len() {
                    break;
                }
                if self.dat[*idx].cmp(&self.dat[out]) == o {
                    out = *idx;
                }
            }
            out
        } else {
            0
        }
    }

}

trait HeapIdx {
    fn left(self) -> Self;
    fn right(self) -> Self;
    fn parent(self) -> Self;
    fn grandparent(self) -> Self;
    fn is_child_of(self, i: Self) -> bool;
    fn level(self) -> usize;
    fn level_type(self) -> LevelType;
}

impl HeapIdx for usize {
    fn left(self) -> usize {
        (self * 2) + 1
    }

    fn right(self) -> usize {
        (self * 2) + 2
    }

    fn parent(self) -> usize {
        if self == 0 { 0 } else { (self - 1) / 2 }
    }

    fn grandparent(self) -> usize {
        self.parent().parent()
    }

    fn is_child_of(self, parent: usize) -> bool {
        self == parent.left() || self == parent.right()
    }

    fn level(self) -> usize {
        let mut c = self;
        let mut out = 0;
        while c != 0 {
            c = c.parent();
            out += 1;
        }
        out
    }

    fn level_type(self) -> LevelType {
        if self.level() % 2 == 0 { LevelType::Min } else { LevelType::Max }
    }

}


enum LevelType {
    Min,
    Max
}

fn rev_order(o: std::cmp::Ordering) -> std::cmp::Ordering {
    match o {
        std::cmp::Ordering::Less    => std::cmp::Ordering::Greater,
        std::cmp::Ordering::Equal   => std::cmp::Ordering::Equal,
        std::cmp::Ordering::Greater => std::cmp::Ordering::Less
    }
}
