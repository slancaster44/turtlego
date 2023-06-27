#include <stdio.h>
#include "vm.hpp"

int main() {
    Binary b("../turtlego/output.tbin");
    Machine vm(b);
    vm.Debug();
}