
label:
    addi t1, zero, 100
    addi t1, t1, 200

    addi t2, zero, 300
    beq t1, t2, .match

.no_match:
    and a0, t1, t2
    ebreak

.match:
    xor a0, t1, t2
    ebreak