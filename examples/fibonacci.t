extern libc {
  fn printf(format *u8, ...) i32;
}

fn fib(n int) int {
  if n <= 1 {
    return n;
  }
  return fib(n - 1) + fib(n - 2);
}

fn main() i32 {
  for(i := 0; i < 10; i = i + 1) {
    result := fib(i);
    libc.printf("%d ", result);
  }
  return 0;
}
