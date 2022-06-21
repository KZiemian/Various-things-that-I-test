#include <iostream>

void solution(int N) {
    int enable_print = N % 10;


    while (N > 0) {
        if (enable_print == 0 && N % 10 != 0) {
            enable_print = 1;
	    std::cout << N % 10;
        }
        else if (enable_print >= 1) {
	  std::cout << N % 10;
        }
        N = N / 10;
    }
    std::cout << '\n';
}

int main() {
  solution(00001);
  solution(10000);
  solution(320002);
  solution(100002);



  return 0;
}
