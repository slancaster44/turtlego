

if [[ "$1" == "--run_test" ]]; then
    g++ -g -O3 -Wformat=0 test.cpp
    time ./a.out
    rm ./a.out
elif [[ "$1" == "--run_main" ]]; then
    g++ -g -O3 -Wformat=0 main.cpp
    time ./a.out
    rm ./a.out
fi