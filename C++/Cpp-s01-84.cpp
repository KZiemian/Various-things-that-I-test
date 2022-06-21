#include <iostream>

class base {
public:
  base() {
    std::cout << "Constructing base\n";
  }

  ~base() {
    std::cout << "Destructing base\n";
  }
};

class derived : public base {
public:
  derived() {
    std::cout << "Constructing derived\n";
  }

  ~derived() {
    std::cout << "Destructing derived\n";
  }
};

// int main() {
//   derived *d = new derived();
//   base *b = d;

//   delete b;
//   getchar();



//   return 0;
// }
