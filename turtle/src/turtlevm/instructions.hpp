#ifndef INSTRUCTIONS_HPP
#define INSTRUCTIONS_HPP

typedef unsigned char op;

/* S/26 Base Instruction Set */
const op PUSHU = 0; 
const op POP = 1;
const op DUP = 2;

const op ALLOC = 3; 
const op DEALLOC = 4; 
const op DEREF = 5; 
const op WRITE = 6; 

const op PUSH_FRAME = 7;
const op POP_FRAME = 8;
const op LOAD = 9;
const op STORE = 10; 

const op JC = 11; 
const op CALL = 12; 
const op RET = 13; 

const op OR = 14; 
const op AND = 15;
const op EQ = 16; 
const op INV = 17; 
const op XOR = 18;
const op TZ = 19; 

const op ADDU = 20;
const op SUBU = 21;
const op GTU = 22;
const op LTU = 23;
const op GEU = 24;
const op LEU = 25;

/* S/35 With Signed Integers */

const op PUSH = 26; 
const op ADD = 27;
const op SUB = 28;
const op GT = 29;
const op LT = 30;
const op GE = 31;
const op LE = 32;

const op CUI = 33; 
const op CIU = 34; 

/* S/37 With Multiplicationn */

const op MUL = 35;
const op MULU = 36;

/* S/39 With Division */
const op DIV = 37;
const op DIVU = 38;

/* S/52 The Complete Instruction Set */

const op PUSHF = 39;
const op ADDF = 40;
const op SUBF = 41;
const op GTF = 42;
const op LTF = 43;
const op GEF = 44;
const op LEF = 45;

const op DIVF = 46;
const op MULF = 47;

const op CFI = 48;
const op CIF = 49;
const op CUF = 50;
const op CFU = 51;


#endif