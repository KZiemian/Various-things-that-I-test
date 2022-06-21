#include <iostream>

class Point {
private:
  int x;
  int y;



public:
  Point(int i = 0, int j = 0) {
    x = i;
    y = j;
  }

  void Print() {
    std::cout << "x = " << x << ", y = " << y << '\n';
  }
};

int main() {
  Point t(20, 20);


  t.Print();
  t = 30;
  t.Print();



  return 0;
}
