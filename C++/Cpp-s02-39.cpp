// Modules
module;

#include <vector>

export module my.first_module;

export import other.module;

#include "mod.h"

constexpr int someInt = 55;

export std::vector<int>  frob(S)



// mod.h
#pragma once;

  struct S {
    int value = 1;
  }


// Module implementation unit
module;

#include <vector>

module my.first_module;

std::vector<int> frob(S s) {
  return (s.valuse + someInt);
}



module;

#include <vector>

export module SimplexModule;

struct detail {
  int someInt = 55;
};

export namespace SimpleModule {
	 void f(const detail& D);

	 std::vector<detail> g();
}



module SimpleModule;

namespace SimpleModule {
  void f(const detail& D) {
    // do something with D
  }
}



module;

#inculde <vector>

module SimpleModule;

namespace SimpleModule {
  std::vector<detail> g() {
    return { { 41 }, { 43 } };
  }
}



export module LargeModule;

export import : iface.f;
export import : iface.g;



module LargeModule : detail;

struct detali {
  int someInt = 41;
};



export
module LargeModule : iface.f;

import : detail;

namespace LargeModule {
  export
  void f(const detail& D);
}



module;

#include <vector>

export
module LargeModule : iface.g;

import : detail;

namespace LargeModule {
  export
  std::vector<detail> g();
}
