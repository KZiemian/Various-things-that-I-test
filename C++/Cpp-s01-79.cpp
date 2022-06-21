#include <iostream>

class Complex {
public:
  Complex(int r = 0, int i = 0) {
    real = r;
    imag = i;
  }

  Complex operator + (Complex const &obj) {
    Complex res;


    std::cout << "res.real: " << res.real << '\n'
	      << "res.imag: " << res.imag << '\n';

    res.real = real + obj.real;
    res.imag = imag + obj.imag;

    std::cout << "res.real: " << res.real << '\n'
	      << "res.imag: " << res.imag << '\n';

    return res;
  }

  void Print() {
    std::cout << real << " + i" << imag << '\n';
  }



private:
  int real;
  int imag;
};

int main() {
  Complex c1(10, 5), c2(2, 4);
  // Complex c1(10, 5);
  // Complex c2(2, 4);


  c1.Print();
  c2.Print();

  Complex c3 = c1 + c2;
  c3.Print();



  return 0;
}
