#ifndef VM_H
#define VM_H

#include <cstdint>
#include <cstdio>
#include <vector>
#include <cstring>
#include <climits>
using namespace std;

#include "stack.hpp"
#include "instructions.hpp"
#include "binary.hpp"

#define DEFAULT_STACK_SIZE 64
#define DEFAULT_CALL_STACK_SIZE 16
#define READ_INT(CODE, INDEX) *(int64_t*) (CODE + INDEX)
#define READ_UINT(CODE, INDEX) *(uint64_t*) (CODE + INDEX)
#define READ_FLT(CODE, INDEX) *(double*) (CODE + INDEX)

union word {
    int64_t integer;
    uint64_t ui;
    double flt;
    word* pointer;
};

class Machine {
private:
    Stack<word> stack;
    Stack<word*> call_stack;

    uint64_t PC;
    uint8_t* code;
    uint64_t ins_count;

    void ExecIns();
public:
    Machine(Binary& b);
    void Dump();
    void Debug();
    void Run();
};

Machine::Machine(Binary& b) : stack(DEFAULT_STACK_SIZE), call_stack(DEFAULT_CALL_STACK_SIZE) {
    PC = 0;
    code = b.getCode();
    ins_count = b.getIns_len();

    uint64_t fz_len = b.getZF_len();
    if (fz_len != 0) {
        uint8_t* image = b.getZeroFrame();
        word* buf = (word*) malloc(fz_len);
        call_stack.Push(buf);

        for (uint64_t i = 0; i < fz_len; i ++) {
            buf[i] = word { ui: *((uint64_t*) image + (i * sizeof(uint64_t))) };
        }
    }

    if (sizeof(uint64_t) != sizeof(int64_t) && sizeof(int64_t) != sizeof(word*) && sizeof(word*) != sizeof(double)) {
        cerr << "Unequal word size" << endl;
        exit(1);
    }
}

inline void Machine::ExecIns() {
    switch (code[PC]) {
    case PUSH:
        stack.Push(word {integer: READ_INT(code, PC+1)});
        PC += sizeof(int64_t);
        break;
    case PUSHF:
        stack.Push(word {flt: READ_FLT(code, PC+1)});
        PC += sizeof(double);
        break;
    case PUSHU:
        stack.Push(word {ui: READ_UINT(code, PC+1)});
        PC += sizeof(uint64_t);
        break;
    case DUP:
        stack.Push(stack.contents[stack.pointer-1]);
        break;
    case POP:
        stack.pointer--;
        break;
    case ADD:
        stack.contents[stack.pointer-1].integer += stack.contents[--stack.pointer].integer;
        break;
    case SUB:
        stack.contents[stack.pointer-1].integer -= stack.contents[--stack.pointer].integer;
        break;
    case MUL:
        stack.contents[stack.pointer-1].integer *= stack.contents[--stack.pointer].integer;
        break;
    case DIV:
        stack.contents[stack.pointer-1].integer /= stack.contents[--stack.pointer].integer;
        break;
    case ADDU:
        stack.contents[stack.pointer-1].ui += stack.contents[--stack.pointer].ui;
        break;
    case SUBU:
        stack.contents[stack.pointer-1].ui -= stack.contents[--stack.pointer].ui;
        break;
    case MULU:
        stack.contents[stack.pointer-1].ui *= stack.contents[--stack.pointer].ui;
        break;
    case DIVU:
        stack.contents[stack.pointer-1].ui /= stack.contents[--stack.pointer].ui;
        break;
    case ADDF:
        stack.contents[stack.pointer-1].flt += stack.contents[--stack.pointer].flt;
        break;
    case SUBF:
        stack.contents[stack.pointer-1].flt -= stack.contents[--stack.pointer].flt;
        break;
    case MULF:
        stack.contents[stack.pointer-1].flt *= stack.contents[--stack.pointer].flt;
        break;
    case DIVF:
        stack.contents[stack.pointer-1].flt /= stack.contents[--stack.pointer].flt;
        break;
    case GT:
        stack.contents[(--stack.pointer)-1].integer = stack.contents[stack.pointer-1].integer > stack.contents[stack.pointer-2].integer;
        break;
    case LT:
        stack.contents[(--stack.pointer)-1].integer = stack.contents[stack.pointer-1].integer < stack.contents[stack.pointer-2].integer;
        break;
    case GE:
        stack.contents[(--stack.pointer)-1].integer = stack.contents[stack.pointer-1].integer >= stack.contents[stack.pointer-2].integer;
        break;
    case LE:
        stack.contents[(--stack.pointer)-1].integer = stack.contents[stack.pointer-1].integer <= stack.contents[stack.pointer-2].integer;
        break;
    case GTU:
        stack.contents[(--stack.pointer)-1].ui = stack.contents[stack.pointer-1].ui > stack.contents[stack.pointer-2].ui;
        break;
    case LTU:
        stack.contents[(--stack.pointer)-1].ui = stack.contents[stack.pointer-1].ui < stack.contents[stack.pointer-2].ui;
        break;
    case GEU:
        stack.contents[(--stack.pointer)-1].ui = stack.contents[stack.pointer-1].ui >= stack.contents[stack.pointer-2].ui;
        break;
    case LEU:
        stack.contents[(--stack.pointer)-1].ui = stack.contents[stack.pointer-1].ui <= stack.contents[stack.pointer-2].ui;
        break;
    case GTF:
        stack.contents[(--stack.pointer)-1].flt = stack.contents[stack.pointer-1].flt > stack.contents[stack.pointer-2].flt;
        break;
    case LTF:
        stack.contents[(--stack.pointer)-1].flt = stack.contents[stack.pointer-1].flt < stack.contents[stack.pointer-2].flt;
        break;
    case GEF:
        stack.contents[(--stack.pointer)-1].flt = stack.contents[stack.pointer-1].flt >= stack.contents[stack.pointer-2].flt;
        break;
    case LEF:
        stack.contents[(--stack.pointer)-1].flt = stack.contents[stack.pointer-1].flt <= stack.contents[stack.pointer-2].flt;
        break;
    case OR:
        stack.contents[(--stack.pointer)-1].ui = stack.contents[stack.pointer-1].ui | stack.contents[stack.pointer-2].ui;
        break;
    case AND:
        stack.contents[(--stack.pointer)-1].ui = stack.contents[stack.pointer-1].ui & stack.contents[stack.pointer-2].ui;
        break;
    case EQ:
        stack.contents[(--stack.pointer)-1].ui = stack.contents[stack.pointer-1].ui == stack.contents[stack.pointer-2].ui;
        break;
    case XOR:
        stack.contents[(--stack.pointer)-1].ui = stack.contents[stack.pointer-1].ui ^ stack.contents[stack.pointer-2].ui;
        break;
    case INV:
        stack.contents[stack.pointer-1].ui = stack.contents[stack.pointer-1].ui ^ UINT_MAX;
        break;
    case TZ:
        stack.contents[stack.pointer-1].ui = stack.contents[stack.pointer-1].ui == 0;
        break;
    case ALLOC:
        stack.Push(word {pointer: (word*) malloc(stack.contents[--stack.pointer].ui)});
        break;
    case DEALLOC:
        free(stack.contents[--stack.pointer].pointer);
        break;
    case WRITE:
        *(stack.contents[--stack.pointer].pointer) = stack.contents[stack.pointer-2];
        stack.contents[stack.pointer-1] = stack.contents[stack.pointer];
        break;
    case DEREF:
        stack.contents[stack.pointer-1] = *(stack.contents[stack.pointer-1].pointer);
        break;
    case CIU:
        stack.contents[stack.pointer-1].ui = (uint64_t) stack.contents[stack.pointer-1].integer;
        break;
    case CUI:
        stack.contents[stack.pointer-1].integer = (int64_t) stack.contents[stack.pointer-1].ui;
        break;
    case CFI:
        stack.contents[stack.pointer-1].integer = (int64_t) stack.contents[stack.pointer-1].flt;
        break;
    case CIF:
        stack.contents[stack.pointer-1].flt = (double) stack.contents[stack.pointer-1].integer;
        break;
    case CUF:
        stack.contents[stack.pointer-1].flt = (double) stack.contents[stack.pointer-1].ui;
        break;
    case CFU:
        stack.contents[stack.pointer-1].ui = (uint64_t) stack.contents[stack.pointer-1].flt;
        break;
    case PUSH_FRAME:
        call_stack.Push((word*) malloc(stack.contents[--stack.pointer].ui));
        break;
    case POP_FRAME:
        free(call_stack.contents[--call_stack.pointer]);
        break;
    case LOAD: //1: frame number; 2: location in frame
        stack.Push(call_stack.contents[stack.contents[--stack.pointer].ui][stack.contents[--stack.pointer].ui]);
        break;
    case STORE: //1: frame number; 2: location in frame; 3: Value to store
        call_stack.contents[stack.contents[--stack.pointer].ui][stack.contents[--stack.pointer].ui] = stack.contents[(--stack.pointer)-2];
        break;
    case JC:
        if (stack.contents[--stack.pointer].ui != 0) { PC = stack.contents[--stack.pointer].ui; }
        return;
    case CALL:
        call_stack.contents[0][0] = word {ui: PC };
        PC = stack.contents[--stack.pointer].ui;
        return;
    case RET:
        PC = call_stack.contents[0][0].ui;
        break;
    default:
        cerr << "Illegal Instruction" << endl;
        exit(1);
    }

    PC++;
}

void Machine::Run() {
    while (PC < ins_count) {
        ExecIns();
    }
}

void Machine::Debug() {
    int shouldExit = 0;

    while (!shouldExit) {
        string input;
        cout << ">> ";
        cin >> input;
        
        if (input == "ds") {
            stack.Dump();
        } else if (input == "dc") {
            call_stack.Dump();
        } else if (input == "peek" || input == "p") {
            cout << "0x" << hex << (*stack.contents[stack.pointer-1].pointer).ui << endl;
        } else if (input == "exit" || input == "q") {
            shouldExit = 1;
        } else if (input == "step" || input == "s") {
            ExecIns();
        } else if (input == "list" || input == "l") {
            for (int i = 0; i < ins_count; i++) {
                if (i == PC) {
                    cout << "-> ";
                } else {
                    cout << "   ";
                }

                cout << "0x" << hex << (int) code[i] << endl;
            }
        }
    }

}

void Machine::Dump() {
    stack.Dump();
}

#endif