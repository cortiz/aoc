use regex::Regex;
use std::fs::read_to_string;

fn main() {
    let mut total: i32 = 0;
    let reg = Regex::new(r"\d").unwrap();
    for line in read_to_string("./input.txt").unwrap().lines() {
        let digits: Vec<&str> = reg.find_iter(line).map(|m| m.as_str()).collect();
        let num = format!("{}{}", digits[0], digits[digits.len() - 1]);
        total = total + num.parse::<i32>().unwrap();
    }
    println!("Your total is {}", total);
}
