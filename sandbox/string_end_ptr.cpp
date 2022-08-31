
#include <iostream>
#include <string>
#include <algorithm>

int main()
{
  std::string s = "hello";
  // s.endはヌル文字になる
  std::cout << "*(s.end()): " << (size_t)*(s.end()) << std::endl;
}
