#include <iostream>

class Fraction {
public:
  Fraction(int n, int d) {
    num = n;
    den = d;
  }

  operator float() const {
    return float(num) / float(den);
  }



private:
  int num;
  int den;
};

int main() {
  Fraction f(2, 5);
  float val = f;


  std::cout << "val: " << val << '\n';



  return 0;
}
