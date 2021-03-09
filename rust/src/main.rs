fn main() {
    for n in 1..101 {
        println!("{}", check_fizz_buzz(n));
    }
}

fn check_fizz_buzz(num: u32) -> String {
    return match num {
        n if n % 3 == 0 && n % 5 == 0 => String::from("fizzbuzz"),
        n if n % 3 == 0 => String::from("fizz"),
        n if n % 5 == 0 => String::from("buzz"),
        _ => num.to_string(),
    };
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn fizz() {
        assert_eq!("fizz", check_fizz_buzz(6));
        assert_eq!("fizz", check_fizz_buzz(3));
        assert_eq!("fizz", check_fizz_buzz(12));
        assert_eq!("fizz", check_fizz_buzz(18));
    }

    #[test]
    fn buzz() {
        assert_eq!("buzz", check_fizz_buzz(5));
        assert_eq!("buzz", check_fizz_buzz(10));
        assert_eq!("buzz", check_fizz_buzz(20));
        assert_eq!("buzz", check_fizz_buzz(35));
    }

    #[test]
    fn fizz_buzz() {
        assert_eq!("fizzbuzz", check_fizz_buzz(15));
        assert_eq!("fizzbuzz", check_fizz_buzz(30));
        assert_eq!("fizzbuzz", check_fizz_buzz(60));
        assert_eq!("fizzbuzz", check_fizz_buzz(90));
    }
    #[test]
    fn no_match() {
        assert_eq!("1", check_fizz_buzz(1));
        assert_eq!("4", check_fizz_buzz(4));
        assert_eq!("22", check_fizz_buzz(22));
        assert_eq!("52", check_fizz_buzz(52));
    }
}
