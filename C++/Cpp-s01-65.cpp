#include <iostream>
#include <bits/stdc++.h>

int main() {
  std::vector<int> vect{10, 20, 30, 40};


  for (int& x : vect) {
    x += 5;
  }

  for (int x : vect) {
    std::cout << "x = " << x << '\n';
  }



  return 0;
}
