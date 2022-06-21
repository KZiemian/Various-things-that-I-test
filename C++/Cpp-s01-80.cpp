#include <iostream>

class Complex {
public:
  Complex(int realPar = 0, int imagPar = 0) {
    real = realPar;
    imag = imagPar;
  }

  void Print() {
    std::cout << real << " + i" << imag << '\n';
  }

  Complex operator + (Complex const &obj) {
    Complex res;

    res.real = real + obj.real;
    res.imag = imag + obj.imag;

    return res;
  }



private:
  int real;
  int imag;
};

int main() {
  Complex c1(10, 5), c2(2, 4);


  Complex c3 = c1 + c2;
  c3.Print();



  return 0;
}
