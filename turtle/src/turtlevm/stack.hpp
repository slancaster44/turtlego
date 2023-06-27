#ifndef STACK_HPP
#define STACK_HPP

#include <iostream>
using namespace std;


template <typename T>
class Stack {
public:
    T* contents;
    uint64_t size;
    uint64_t pointer;
    uint64_t low_water_mark;

    Stack(uint64_t s);
    ~Stack();

    void Resize(uint64_t new_size);
    void Push(T);
    void Dump();
};

template <typename T>
Stack<T>::Stack(uint64_t s) {
    low_water_mark = s;
    size = s;
    pointer = 0;
    contents = (T*) malloc(s * sizeof(T));
    memset(contents, 0, s);
}

template <typename T>
Stack<T>::~Stack() {
    free(contents);
}

template <typename T>
void Stack<T>::Push(T val) {
    contents[pointer] = val;
    pointer++;

    if (size == pointer) {
        Resize(size << 1);
    } else if (size >> 4 > pointer && size >> 2 > low_water_mark) {
        Resize(size >> 2);
    }
}

//TODO: realloc?
template <typename T>
void Stack<T>::Resize(uint64_t new_size) {
    T* oldArr = contents;
    contents = new T[new_size];
    memcpy(contents, oldArr, pointer);
    size = new_size;
    delete oldArr;
}

template <typename T>
void Stack<T>::Dump() {
    for (int i = size-1; i > -1; i--) {
        if (i == pointer) {
            cout << "SP-> ";
        } else {
            cout << "     ";
        }
        printf("%i: 0x%X\n", i, contents[i]);

    }
}

#endif

