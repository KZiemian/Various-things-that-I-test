#include <iostream>

#include <algorithm>
#include <vector>

int solution(std::vector<int> &A) {
    int N = A.size();
    int result = 0;


    for (int i = 0; i < N; i++)
        for (int j = i; j < N; j++)
            if (A[i] != A[j])
	      result = std::max(result, j - i);
    return result;
}

int main() {
  std::vector<int> v = {1, 1, 1, 1, 0, 1, 1, 1};


  std::cout << "solution(v): " << solution(v) << '\n';



  return 0;
}
