#include <iostream>
#include <bits/stdc++.h>

int main() {
  std::vector<std::string> vect{"geeksforgeeks practice",
      "geeksforgeeks write",
      "geeksforgeeks ide"};


  for (const auto& x : vect) {
    std::cout << "x = " << x << '\n';
  }

  std::cout << "\nAnother for loop\n";

  for (const std::string& x : vect) {
    std::cout << "x = " << x << '\n';
  }



  return 0;
}
