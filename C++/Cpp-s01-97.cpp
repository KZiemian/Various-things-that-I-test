class sprzeg {
  unsigned int a : 5;
  unsigned int b : 1;
  unsigned int c : 7;
  unsigned int d : 12;
  unsigned int e : 2;
};

class wzor {
public:
  unsigned int a : 4;
  unsigned int : 5;
  unsigned int b : 2;
  unsigned int : 0;
  unsigned int c : 3;
  unsigned int : 0;
  unsigned int : 3;
};

void print_range(int start, int end) {
  for (auto i = start; i != end; i++) {
    std::cout << i << "\n";
  }
}

void train_model(const std::vector<int>& data, auto& model) {
  for (const auto& x : data) {
    model.update(x);
  }
}

void know_your_algorithms() {
  const std::vector<int> data = {-1, -3, -5, -8, 15, -1};


  const auto is_positive = [](const auto &x) {
    return x > 0;
  };

  auto first_pos_it = std::find_if(ndata.cbegin(),
				   data.cend(),
				   is_positive);
}

void using_c_array() {
  const int n = 256;
  std::array<int, n> arr{};


  better_f(arr);
}

template<typename T>

void print_bytes(const T& input) {
  auto *bytes = reinterpret_cast<const std::byte *>(&input);
}

template<typename T>

void print_bytes(const T& input) {
  using bytearray = std::array<std::byte, sizeof(T)>;


  const auto& bytes = std::bit_cast<bytearray, T>(input);
}

const std::string& more_frequent(const std::unordered_map<std::string, int>&
				 word_counts,
				 const std::string& word1,
				 const std::string& word2) {
  return word_counts.at(word1) > word_counts.at(word2) ? word1 : word2;
}

void print_vec_one_per_line(const std::vector<int>& arr) {
  for (const auto& x : arr) {
    std::cout << x << '\n';
  }
}

class Employee {
private:
  string name;
  string desig;

  // ...
};

void PrintEmpDetails(Employee emp) {
  std::cout << emp.GetName();
  std::cout << emp.GetDesig();

  // ...
}

void PrintEmpDetails(const Employee& emp) {
  std::cout << emp.GetName();
  std::cout << emp.GetDesig();

  // ...
}

// class person {
//   char name[20];
//   int  id;
// public:
//   void GetDetails() {}
// };

// int main() {
//   person p1;
// }
