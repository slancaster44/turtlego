#ifndef BINARY_H
#define BINARY_H

#include <stdio.h>

class Binary {
private:
    uint8_t* code;
    uint64_t ins_len;

    uint8_t* frame_zero_image;
    uint64_t frame_zero_length;

   void __fputn(uint8_t* v, uint64_t n, FILE* f);
   void __fgetn(uint8_t* out_buf, uint64_t n, FILE* f);
   FILE* __fopen(const char* filename, const char* mod);
public:
    Binary(string filename);
    Binary(uint8_t* code, uint64_t code_length);
    Binary(uint8_t*, uint64_t, uint8_t*, uint64_t);

    void Write(string filename);

    uint8_t* getCode();
    uint64_t getIns_len();
    uint8_t* getZeroFrame();
    uint64_t getZF_len();

    ~Binary();
};

Binary::Binary(uint8_t* c, uint64_t l) {
    code = c;
    ins_len = l;
    frame_zero_image = NULL;
    frame_zero_length = 0;
}

Binary::Binary(uint8_t* c, uint64_t il, uint8_t* fz, uint64_t fzl) {
    code = c;
    ins_len = il;
    frame_zero_image = fz;
    frame_zero_length = fzl;
}

Binary::Binary(string filename) {
    FILE* f = __fopen(filename.c_str(), "r");
    
    __fgetn((uint8_t*) &ins_len, sizeof(uint64_t), f);
    __fgetn((uint8_t*) &frame_zero_length, sizeof(uint64_t), f);

    code = NULL;
    frame_zero_image = NULL;

    if (ins_len != 0)
        code = (uint8_t*) malloc(ins_len);
    if (frame_zero_length != 0)
        frame_zero_image = (uint8_t*) malloc(frame_zero_length);
        
    __fgetn(code, ins_len, f);
    __fgetn(frame_zero_image, frame_zero_length, f);

    fclose(f);
}

Binary::~Binary() {
    if (code != NULL)
        free(code);

    if (frame_zero_image != NULL)
        free(frame_zero_image);
}

uint8_t* Binary::getCode() {
    return code;
}
uint64_t Binary::getIns_len() {
    return ins_len;
}
uint8_t* Binary::getZeroFrame() {
    return frame_zero_image;
}
uint64_t Binary::getZF_len() {
    return frame_zero_length;
}


void Binary::__fputn(uint8_t* v, uint64_t n, FILE* f) {
    for (uint64_t i = 0; i < n; i ++) {
        fputc(v[i], f);
    }
}

void Binary::__fgetn(uint8_t* out_buf, uint64_t n, FILE* f) {
    for (uint64_t i = 0; i < n; i++) {
        out_buf[i] = (uint8_t) fgetc(f);
    }
}

FILE* Binary::__fopen(const char* filename, const char* mod) {
    FILE* f = fopen(filename, mod);
    if (f == NULL) {
        cerr << "Invalid file" << endl;
        exit(1);
    }

    return f;
}

void Binary::Write(string filename) {
    FILE* f = __fopen(filename.c_str(), "w");
    
    uint8_t* code_len_str = (uint8_t*) alloca(sizeof(uint64_t));
    uint8_t* fz_len_str = (uint8_t*) alloca(sizeof(uint64_t));

    memcpy(code_len_str, &ins_len, sizeof(uint64_t));
    memcpy(fz_len_str, &frame_zero_length, sizeof(uint64_t));

    __fputn(code_len_str, sizeof(uint64_t), f);
    __fputn(fz_len_str, sizeof(uint64_t), f);

    __fputn(code, ins_len, f);
    __fputn(frame_zero_image, frame_zero_length, f);

    fclose(f);
    
}

#endif