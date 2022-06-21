#include <iostream>

class Complex {
private:
  int real;
  int imag;



public:
  Complex(int r = 0, int i = 0) {
    real = r;
    imag = i;
  }


  Complex operator + (Complex const &obj) {
    Complex res;

    res.real = real + obj.real;
    res.imag = imag + obj.imag;



    return res;
  }


  void Print() {
    std::cout << real << " + i" << imag << '\n';
  }
};

int main() {
  Complex c1(10, 5);
  Complex c2(2, 4);
  Complex c3 = c1 + c2;


  c3.Print();



  return 0;
}
